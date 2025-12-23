package services

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/mustafa91ameen/prjalgo/backend/internal/dtos"
	"github.com/mustafa91ameen/prjalgo/backend/internal/models"
	"github.com/mustafa91ameen/prjalgo/backend/internal/repository"
)

var (
	ErrWorkCategoryNotFound      = errors.New("work category not found")
	ErrWorkCategoryStatsNotFound = errors.New("work category stats not found")
)

type WorkCategoryService struct {
	categoryRepo repository.WorkCategoryRepositoryInterface
}

func NewWorkCategoryService(categoryRepo repository.WorkCategoryRepositoryInterface) *WorkCategoryService {
	return &WorkCategoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *WorkCategoryService) GetAll(ctx context.Context, limit, offset int) (dtos.PaginatedResponse[dtos.WorkCategorySummary], error) {
	categories, total, err := s.categoryRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return dtos.PaginatedResponse[dtos.WorkCategorySummary]{}, err
	}

	result := make([]dtos.WorkCategorySummary, len(categories))
	for i, c := range categories {
		result[i] = toWorkCategorySummaryDTO(c)
	}

	page := (offset / limit) + 1
	return dtos.NewPaginatedResponse(result, total, page, limit), nil
}

func (s *WorkCategoryService) GetByID(ctx context.Context, id int64) (*dtos.WorkCategory, error) {
	category, err := s.categoryRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkCategoryNotFound
		}
		return nil, err
	}
	dto := toWorkCategoryDTO(category)
	return &dto, nil
}

func (s *WorkCategoryService) Create(ctx context.Context, req dtos.CreateWorkCategory) (*dtos.WorkCategory, error) {
	category := &models.WorkCategory{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
		CreatedBy:   req.CreatedBy,
	}

	created, err := s.categoryRepo.Create(ctx, category)
	if err != nil {
		return nil, err
	}

	dto := toWorkCategoryDTO(created)
	return &dto, nil
}

func (s *WorkCategoryService) Update(ctx context.Context, id int64, req dtos.UpdateWorkCategory) (*dtos.WorkCategory, error) {
	existing, err := s.categoryRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkCategoryNotFound
		}
		return nil, err
	}

	// Apply partial updates
	if req.Name != nil {
		existing.Name = *req.Name
	}
	if req.Description != nil {
		existing.Description = req.Description
	}
	if req.Status != nil {
		existing.Status = req.Status
	}

	updated, err := s.categoryRepo.Update(ctx, id, existing)
	if err != nil {
		return nil, err
	}

	dto := toWorkCategoryDTO(updated)
	return &dto, nil
}

func (s *WorkCategoryService) Delete(ctx context.Context, id int64) error {
	err := s.categoryRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrWorkCategoryNotFound
		}
		return err
	}
	return nil
}

func (s *WorkCategoryService) GetStats(ctx context.Context, period string) (*dtos.WorkCategoryStatsResponse, error) {
	stats, err := s.categoryRepo.GetStats(ctx, period)
	if err != nil {
		return nil, err
	}

	return &dtos.WorkCategoryStatsResponse{
		Total:            stats.Total,
		Active:           stats.Active,
		Inactive:         stats.Inactive,
		TotalSubcategory: stats.TotalSubcategory,
	}, nil
}

// DTO conversion helpers
func toWorkCategorySummaryDTO(c models.WorkCategory) dtos.WorkCategorySummary {
	return dtos.WorkCategorySummary{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		Status:      c.Status,
	}
}

func toWorkCategoryDTO(c *models.WorkCategory) dtos.WorkCategory {
	return dtos.WorkCategory{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		Status:      c.Status,
		CreatedBy:   c.CreatedBy,
		CreatedAt:   c.CreatedAt,
	}
}
