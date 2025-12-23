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
	ErrWorkDayMaterialNotFound = errors.New("work day material not found")
)

type WorkDayMaterialService struct {
	db             *pgxpool.Pool
	materialRepo   repository.WorkDayMaterialRepositoryInterface
	workDayService *WorkDayService
}

func NewWorkDayMaterialService(db *pgxpool.Pool, materialRepo repository.WorkDayMaterialRepositoryInterface, workDayService *WorkDayService) *WorkDayMaterialService {
	return &WorkDayMaterialService{
		db:             db,
		materialRepo:   materialRepo,
		workDayService: workDayService,
	}
}

func (s *WorkDayMaterialService) GetAll(ctx context.Context, limit, offset int) (dtos.PaginatedResponse[dtos.WorkDayMaterial], error) {
	materials, total, err := s.materialRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return dtos.PaginatedResponse[dtos.WorkDayMaterial]{}, err
	}

	result := make([]dtos.WorkDayMaterial, len(materials))
	for i, m := range materials {
		result[i] = toWorkDayMaterialDTO(m)
	}

	page := (offset / limit) + 1
	return dtos.NewPaginatedResponse(result, total, page, limit), nil
}

func (s *WorkDayMaterialService) GetByWorkDayID(ctx context.Context, workDayID int64) ([]dtos.WorkDayMaterial, error) {
	materials, err := s.materialRepo.GetByWorkDayID(ctx, workDayID)
	if err != nil {
		return nil, err
	}

	result := make([]dtos.WorkDayMaterial, len(materials))
	for i, m := range materials {
		result[i] = toWorkDayMaterialDTO(m)
	}
	return result, nil
}

func (s *WorkDayMaterialService) GetByID(ctx context.Context, id int64) (*dtos.WorkDayMaterial, error) {
	material, err := s.materialRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkDayMaterialNotFound
		}
		return nil, err
	}
	dto := toWorkDayMaterialDTO(*material)
	return &dto, nil
}

func (s *WorkDayMaterialService) Create(ctx context.Context, req dtos.CreateWorkDayMaterial) (*dtos.WorkDayMaterial, error) {
	material := &models.WorkDayMaterial{
		WorkDayID:    req.WorkDayID,
		MaterialName: req.MaterialName,
		Quantity:     req.Quantity,
		Cost:         req.Cost,
		Notes:        req.Notes,
	}

	// Calculate amount to add to work day total cost
	amountToAdd := req.Quantity * req.Cost

	// Start transaction
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// Create material within transaction
	created, err := s.materialRepo.CreateWithTx(ctx, tx, material)
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

	dto := toWorkDayMaterialDTO(*created)
	return &dto, nil
}

func (s *WorkDayMaterialService) Update(ctx context.Context, id int64, req dtos.UpdateWorkDayMaterial) (*dtos.WorkDayMaterial, error) {
	existing, err := s.materialRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkDayMaterialNotFound
		}
		return nil, err
	}

	// Calculate old and new amounts
	oldAmount := existing.Quantity * existing.Cost
	newQuantity := existing.Quantity
	newCost := existing.Cost

	// Apply partial updates
	if req.MaterialName != nil {
		existing.MaterialName = *req.MaterialName
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

	// Update material within transaction
	updated, err := s.materialRepo.UpdateWithTx(ctx, tx, id, existing)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkDayMaterialNotFound
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

	dto := toWorkDayMaterialDTO(*updated)
	return &dto, nil
}

func (s *WorkDayMaterialService) Delete(ctx context.Context, id int64) error {
	// Get material to know work day ID and amount to subtract
	material, err := s.materialRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrWorkDayMaterialNotFound
		}
		return err
	}

	// Calculate amount to subtract from work day total cost
	amountToSubtract := material.Quantity * material.Cost

	// Start transaction
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Delete material within transaction
	err = s.materialRepo.DeleteWithTx(ctx, tx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrWorkDayMaterialNotFound
		}
		return err
	}

	// Update work day total cost (subtract amount)
	err = s.workDayService.UpdateTotalCostWithTx(ctx, tx, material.WorkDayID, -amountToSubtract)
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
func toWorkDayMaterialDTO(m models.WorkDayMaterial) dtos.WorkDayMaterial {
	return dtos.WorkDayMaterial{
		ID:           m.ID,
		WorkDayID:    m.WorkDayID,
		MaterialName: m.MaterialName,
		Quantity:     m.Quantity,
		Cost:         m.Cost,
		Notes:        m.Notes,
		CreatedAt:    m.CreatedAt,
	}
}
