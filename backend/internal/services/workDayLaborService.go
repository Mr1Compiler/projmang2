package services

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafa91ameen/prjalgo/backend/internal/dtos"
	"github.com/mustafa91ameen/prjalgo/backend/internal/models"
	"github.com/mustafa91ameen/prjalgo/backend/internal/repository"
)

var (
	ErrWorkDayLaborNotFound = errors.New("work day labor not found")
)

type WorkDayLaborService struct {
	db             *pgxpool.Pool
	laborRepo      repository.WorkDayLaborRepositoryInterface
	workDayService *WorkDayService
}

func NewWorkDayLaborService(db *pgxpool.Pool, laborRepo repository.WorkDayLaborRepositoryInterface, workDayService *WorkDayService) *WorkDayLaborService {
	return &WorkDayLaborService{
		db:             db,
		laborRepo:      laborRepo,
		workDayService: workDayService,
	}
}

func (s *WorkDayLaborService) GetAll(ctx context.Context, limit, offset int) (dtos.PaginatedResponse[dtos.WorkDayLabor], error) {
	labors, total, err := s.laborRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return dtos.PaginatedResponse[dtos.WorkDayLabor]{}, err
	}

	result := make([]dtos.WorkDayLabor, len(labors))
	for i, l := range labors {
		result[i] = toWorkDayLaborDTO(l)
	}

	page := (offset / limit) + 1
	return dtos.NewPaginatedResponse(result, total, page, limit), nil
}

func (s *WorkDayLaborService) GetByWorkDayID(ctx context.Context, workDayID int64) ([]dtos.WorkDayLabor, error) {
	labors, err := s.laborRepo.GetByWorkDayID(ctx, workDayID)
	if err != nil {
		return nil, err
	}

	result := make([]dtos.WorkDayLabor, len(labors))
	for i, l := range labors {
		result[i] = toWorkDayLaborDTO(l)
	}
	return result, nil
}

func (s *WorkDayLaborService) GetByID(ctx context.Context, id int64) (*dtos.WorkDayLabor, error) {
	labor, err := s.laborRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkDayLaborNotFound
		}
		return nil, err
	}
	dto := toWorkDayLaborDTO(*labor)
	return &dto, nil
}

func (s *WorkDayLaborService) Create(ctx context.Context, req dtos.CreateWorkDayLabor) (*dtos.WorkDayLabor, error) {
	labor := &models.WorkDayLabor{
		WorkDayID:  req.WorkDayID,
		WorkerName: req.WorkerName,
		JobTitle:   req.JobTitle,
		Phone:      req.Phone,
		Address:    req.Address,
		Quantity:   req.Quantity,
		Cost:       req.Cost,
		Notes:      req.Notes,
	}

	// Calculate amount to add to work day total cost
	amountToAdd := req.Quantity * req.Cost

	// Start transaction
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// Create labor within transaction
	created, err := s.laborRepo.CreateWithTx(ctx, tx, labor)
	if err != nil {
		return nil, err
	}

	// Update work day total cost
	err = s.workDayService.UpdateTotalCostWithTx(ctx, tx, req.WorkDayID, amountToAdd)
	if err != nil {
		return nil, err
	}

	// Commit transaction
	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	dto := toWorkDayLaborDTO(*created)
	return &dto, nil
}

func (s *WorkDayLaborService) Update(ctx context.Context, id int64, req dtos.UpdateWorkDayLabor) (*dtos.WorkDayLabor, error) {
	existing, err := s.laborRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkDayLaborNotFound
		}
		return nil, err
	}

	// Calculate old and new amounts
	oldAmount := existing.Quantity * existing.Cost
	newQuantity := existing.Quantity
	newCost := existing.Cost

	// Apply partial updates
	if req.WorkerName != nil {
		existing.WorkerName = *req.WorkerName
	}
	if req.JobTitle != nil {
		existing.JobTitle = req.JobTitle
	}
	if req.Phone != nil {
		existing.Phone = req.Phone
	}
	if req.Address != nil {
		existing.Address = req.Address
	}
	if req.Quantity != nil {
		newQuantity = *req.Quantity
		existing.Quantity = *req.Quantity
	}
	if req.Cost != nil {
		newCost = *req.Cost
		existing.Cost = *req.Cost
	}
	if req.Notes != nil {
		existing.Notes = req.Notes
	}

	newAmount := newQuantity * newCost
	amountDiff := newAmount - oldAmount

	// Start transaction
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// Update labor within transaction
	updated, err := s.laborRepo.UpdateWithTx(ctx, tx, id, existing)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkDayLaborNotFound
		}
		return nil, err
	}

	// Update work day total cost with difference
	if amountDiff != 0 {
		err = s.workDayService.UpdateTotalCostWithTx(ctx, tx, existing.WorkDayID, amountDiff)
		if err != nil {
			return nil, err
		}
	}

	// Commit transaction
	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	dto := toWorkDayLaborDTO(*updated)
	return &dto, nil
}

func (s *WorkDayLaborService) Delete(ctx context.Context, id int64) error {
	// Get labor to know work day ID and amount to subtract
	labor, err := s.laborRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrWorkDayLaborNotFound
		}
		return err
	}

	// Calculate amount to subtract from work day total cost
	amountToSubtract := labor.Quantity * labor.Cost

	// Start transaction
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Delete labor within transaction
	err = s.laborRepo.DeleteWithTx(ctx, tx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrWorkDayLaborNotFound
		}
		return err
	}

	// Update work day total cost (subtract amount)
	err = s.workDayService.UpdateTotalCostWithTx(ctx, tx, labor.WorkDayID, -amountToSubtract)
	if err != nil {
		return err
	}

	// Commit transaction
	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

// DTO conversion helper
func toWorkDayLaborDTO(l models.WorkDayLabor) dtos.WorkDayLabor {
	return dtos.WorkDayLabor{
		ID:         l.ID,
		WorkDayID:  l.WorkDayID,
		WorkerName: l.WorkerName,
		JobTitle:   l.JobTitle,
		Phone:      l.Phone,
		Address:    l.Address,
		Quantity:   l.Quantity,
		Cost:       l.Cost,
		Notes:      l.Notes,
		CreatedAt:  l.CreatedAt,
	}
}
