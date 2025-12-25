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
	ErrWorkDayNotFound         = errors.New("work day not found")
	ErrWorkDayAlreadyCompleted = errors.New("work day already completed")
	ErrWorkDayNotCompleted     = errors.New("work day is not completed")
	ErrWorkDayNoSubCategory    = errors.New("work day has no subcategory assigned")
)

type WorkDayService struct {
	db                  *pgxpool.Pool
	workDayRepo         repository.WorkDayRepositoryInterface
	projectService      *ProjectService
	workSubCategoryRepo repository.WorkSubCategoryRepositoryInterface
}

func NewWorkDayService(
	db *pgxpool.Pool,
	workDayRepo repository.WorkDayRepositoryInterface,
	projectService *ProjectService,
	workSubCategoryRepo repository.WorkSubCategoryRepositoryInterface,
) *WorkDayService {
	return &WorkDayService{
		db:                  db,
		workDayRepo:         workDayRepo,
		projectService:      projectService,
		workSubCategoryRepo: workSubCategoryRepo,
	}
}

func (s *WorkDayService) GetAll(ctx context.Context, limit, offset int) (dtos.PaginatedResponse[dtos.WorkDaySummary], error) {
	workDays, total, err := s.workDayRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return dtos.PaginatedResponse[dtos.WorkDaySummary]{}, err
	}

	result := make([]dtos.WorkDaySummary, len(workDays))
	for i, w := range workDays {
		result[i] = toWorkDaySummaryDTO(w)
	}

	page := (offset / limit) + 1
	return dtos.NewPaginatedResponse(result, total, page, limit), nil
}

func (s *WorkDayService) GetByID(ctx context.Context, id int64) (*dtos.WorkDay, error) {
	workDay, err := s.workDayRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkDayNotFound
		}
		return nil, err
	}
	dto := toWorkDayDTO(workDay)
	return &dto, nil
}

func (s *WorkDayService) GetByProjectID(ctx context.Context, projectID int64) ([]dtos.WorkDaySummary, error) {
	workDays, err := s.workDayRepo.GetByProjectID(ctx, projectID)
	if err != nil {
		return nil, err
	}

	result := make([]dtos.WorkDaySummary, len(workDays))
	for i, w := range workDays {
		result[i] = toWorkDaySummaryDTO(w)
	}
	return result, nil
}

func (s *WorkDayService) Create(ctx context.Context, req dtos.CreateWorkDay) (*dtos.WorkDay, error) {
	workDay := &models.WorkDay{
		ProjectID:         req.ProjectID,
		WorkSubCategoryID: req.WorkSubCategoryID,
		WorkDate:          req.WorkDate,
		Description:       req.Description,
		Status:            "pending",
		TotalCost:         0,
		Notes:             req.Notes,
		CreatedBy:         req.CreatedBy,
	}

	created, err := s.workDayRepo.Create(ctx, workDay)
	if err != nil {
		return nil, err
	}

	dto := toWorkDayDTO(created)
	return &dto, nil
}

func (s *WorkDayService) Update(ctx context.Context, id int64, req dtos.UpdateWorkDay) (*dtos.WorkDay, error) {
	existing, err := s.workDayRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkDayNotFound
		}
		return nil, err
	}

	// Apply partial updates
	if req.WorkSubCategoryID != nil {
		existing.WorkSubCategoryID = req.WorkSubCategoryID
	}
	if req.WorkDate != nil {
		existing.WorkDate = *req.WorkDate
	}
	if req.Description != nil {
		existing.Description = req.Description
	}
	if req.Status != nil {
		existing.Status = *req.Status
	}
	if req.TotalCost != nil {
		existing.TotalCost = *req.TotalCost
	}
	if req.Notes != nil {
		existing.Notes = req.Notes
	}

	updated, err := s.workDayRepo.Update(ctx, id, existing)
	if err != nil {
		return nil, err
	}

	dto := toWorkDayDTO(updated)
	return &dto, nil
}

func (s *WorkDayService) Delete(ctx context.Context, id int64) error {
	err := s.workDayRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrWorkDayNotFound
		}
		return err
	}
	return nil
}

func (s *WorkDayService) Complete(ctx context.Context, id int64) (*dtos.WorkDay, error) {
	// Start transaction first to avoid race conditions
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// Get the work day with row lock to prevent concurrent updates
	workDay, err := s.workDayRepo.GetByIDWithTx(ctx, tx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkDayNotFound
		}
		return nil, err
	}

	// Check if already completed
	if workDay.Status == "completed" {
		return nil, ErrWorkDayAlreadyCompleted
	}

	// Check if work day has a subcategory assigned
	if workDay.WorkSubCategoryID == nil {
		return nil, ErrWorkDayNoSubCategory
	}

	// Get the subcategory to get its percentage
	subCategory, err := s.workSubCategoryRepo.GetByID(ctx, *workDay.WorkSubCategoryID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("work sub category not found")
		}
		return nil, err
	}

	// Update project progress percentage if subcategory has a percentage
	if subCategory.Percentage != nil && *subCategory.Percentage > 0 {
		err = s.projectService.UpdateProgressPercentageWithTx(ctx, tx, workDay.ProjectID, *subCategory.Percentage)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return nil, errors.New("project not found")
			}
			return nil, err
		}
	}

	// Update work day status to completed
	workDay.Status = "completed"
	updated, err := s.workDayRepo.UpdateWithTx(ctx, tx, id, workDay)
	if err != nil {
		return nil, err
	}

	// Commit transaction
	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	dto := toWorkDayDTO(updated)
	return &dto, nil
}

func (s *WorkDayService) Uncomplete(ctx context.Context, id int64) (*dtos.WorkDay, error) {
	// Start transaction first to avoid race conditions
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// Get the work day with row lock to prevent concurrent updates
	workDay, err := s.workDayRepo.GetByIDWithTx(ctx, tx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkDayNotFound
		}
		return nil, err
	}

	// Check if work day is completed
	if workDay.Status != "completed" {
		return nil, ErrWorkDayNotCompleted
	}

	// Check if work day has a subcategory assigned
	if workDay.WorkSubCategoryID == nil {
		return nil, ErrWorkDayNoSubCategory
	}

	// Get the subcategory to get its percentage
	subCategory, err := s.workSubCategoryRepo.GetByID(ctx, *workDay.WorkSubCategoryID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("work sub category not found")
		}
		return nil, err
	}

	// Decrease project progress percentage if subcategory has a percentage
	if subCategory.Percentage != nil && *subCategory.Percentage > 0 {
		// Pass negative value to decrease the percentage
		err = s.projectService.UpdateProgressPercentageWithTx(ctx, tx, workDay.ProjectID, -*subCategory.Percentage)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return nil, errors.New("project not found")
			}
			return nil, err
		}
	}

	// Update work day status to in_progress
	workDay.Status = "in_progress"
	updated, err := s.workDayRepo.UpdateWithTx(ctx, tx, id, workDay)
	if err != nil {
		return nil, err
	}

	// Commit transaction
	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	dto := toWorkDayDTO(updated)
	return &dto, nil
}

// UpdateTotalCostWithTx updates the work day total cost by adding the specified amount within a transaction
func (s *WorkDayService) UpdateTotalCostWithTx(ctx context.Context, tx pgx.Tx, workDayID int64, amountToAdd float64) error {
	return s.workDayRepo.UpdateTotalCostWithTx(ctx, tx, workDayID, amountToAdd)
}

// DTO conversion helpers
func toWorkDaySummaryDTO(w models.WorkDay) dtos.WorkDaySummary {
	return dtos.WorkDaySummary{
		ID:                w.ID,
		ProjectID:         w.ProjectID,
		WorkSubCategoryID: w.WorkSubCategoryID,
		WorkDate:          w.WorkDate,
		Description:       w.Description,
		Notes:             w.Notes,
		Status:            w.Status,
		TotalCost:         w.TotalCost,
	}
}

func toWorkDayDTO(w *models.WorkDay) dtos.WorkDay {
	return dtos.WorkDay{
		ID:                w.ID,
		ProjectID:         w.ProjectID,
		WorkSubCategoryID: w.WorkSubCategoryID,
		WorkDate:          w.WorkDate,
		Description:       w.Description,
		Status:            w.Status,
		TotalCost:         w.TotalCost,
		Notes:             w.Notes,
		CreatedBy:         w.CreatedBy,
		CreatedAt:         w.CreatedAt,
	}
}
