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
	ErrWorkDayEquipmentNotFound = errors.New("work day equipment not found")
)

type WorkDayEquipmentService struct {
	db             *pgxpool.Pool
	equipmentRepo  repository.WorkDayEquipmentRepositoryInterface
	workDayService *WorkDayService
}

func NewWorkDayEquipmentService(db *pgxpool.Pool, equipmentRepo repository.WorkDayEquipmentRepositoryInterface, workDayService *WorkDayService) *WorkDayEquipmentService {
	return &WorkDayEquipmentService{
		db:             db,
		equipmentRepo:  equipmentRepo,
		workDayService: workDayService,
	}
}

func (s *WorkDayEquipmentService) GetAll(ctx context.Context, limit, offset int) (dtos.PaginatedResponse[dtos.WorkDayEquipment], error) {
	equipments, total, err := s.equipmentRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return dtos.PaginatedResponse[dtos.WorkDayEquipment]{}, err
	}

	result := make([]dtos.WorkDayEquipment, len(equipments))
	for i, e := range equipments {
		result[i] = toWorkDayEquipmentDTO(e)
	}

	page := (offset / limit) + 1
	return dtos.NewPaginatedResponse(result, total, page, limit), nil
}

func (s *WorkDayEquipmentService) GetByWorkDayID(ctx context.Context, workDayID int64) ([]dtos.WorkDayEquipment, error) {
	equipments, err := s.equipmentRepo.GetByWorkDayID(ctx, workDayID)
	if err != nil {
		return nil, err
	}

	result := make([]dtos.WorkDayEquipment, len(equipments))
	for i, e := range equipments {
		result[i] = toWorkDayEquipmentDTO(e)
	}
	return result, nil
}

func (s *WorkDayEquipmentService) GetByID(ctx context.Context, id int64) (*dtos.WorkDayEquipment, error) {
	equipment, err := s.equipmentRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkDayEquipmentNotFound
		}
		return nil, err
	}
	dto := toWorkDayEquipmentDTO(*equipment)
	return &dto, nil
}

func (s *WorkDayEquipmentService) Create(ctx context.Context, req dtos.CreateWorkDayEquipment) (*dtos.WorkDayEquipment, error) {
	equipment := &models.WorkDayEquipment{
		WorkDayID:     req.WorkDayID,
		EquipmentName: req.EquipmentName,
		Quantity:      req.Quantity,
		Cost:          req.Cost,
		Notes:         req.Notes,
	}

	// Calculate amount to add to work day total cost
	amountToAdd := req.Quantity * req.Cost

	// Start transaction
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// Create equipment within transaction
	created, err := s.equipmentRepo.CreateWithTx(ctx, tx, equipment)
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

	dto := toWorkDayEquipmentDTO(*created)
	return &dto, nil
}

func (s *WorkDayEquipmentService) Update(ctx context.Context, id int64, req dtos.UpdateWorkDayEquipment) (*dtos.WorkDayEquipment, error) {
	existing, err := s.equipmentRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkDayEquipmentNotFound
		}
		return nil, err
	}

	// Calculate old and new amounts
	oldAmount := existing.Quantity * existing.Cost
	newQuantity := existing.Quantity
	newCost := existing.Cost

	// Apply partial updates
	if req.EquipmentName != nil {
		existing.EquipmentName = *req.EquipmentName
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

	// Update equipment within transaction
	updated, err := s.equipmentRepo.UpdateWithTx(ctx, tx, id, existing)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkDayEquipmentNotFound
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

	dto := toWorkDayEquipmentDTO(*updated)
	return &dto, nil
}

func (s *WorkDayEquipmentService) Delete(ctx context.Context, id int64) error {
	// Get equipment to know work day ID and amount to subtract
	equipment, err := s.equipmentRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrWorkDayEquipmentNotFound
		}
		return err
	}

	// Calculate amount to subtract from work day total cost
	amountToSubtract := equipment.Quantity * equipment.Cost

	// Start transaction
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Delete equipment within transaction
	err = s.equipmentRepo.DeleteWithTx(ctx, tx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrWorkDayEquipmentNotFound
		}
		return err
	}

	// Update work day total cost (subtract amount)
	err = s.workDayService.UpdateTotalCostWithTx(ctx, tx, equipment.WorkDayID, -amountToSubtract)
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
func toWorkDayEquipmentDTO(e models.WorkDayEquipment) dtos.WorkDayEquipment {
	return dtos.WorkDayEquipment{
		ID:            e.ID,
		WorkDayID:     e.WorkDayID,
		EquipmentName: e.EquipmentName,
		Quantity:      e.Quantity,
		Cost:          e.Cost,
		Notes:         e.Notes,
		CreatedAt:     e.CreatedAt,
	}
}
