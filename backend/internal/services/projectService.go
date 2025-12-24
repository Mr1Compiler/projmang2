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
	ErrProjectNotFound            = errors.New("project not found")
	ErrProjectStatsNotFound       = errors.New("project stats not found")
	ErrWarningCostExceedsTotalCost = errors.New("warning cost must be less than or equal to total cost")
)

type ProjectService struct {
	projectRepo repository.ProjectRepositoryInterface
}

func NewProjectService(projectRepo repository.ProjectRepositoryInterface) *ProjectService {
	return &ProjectService{
		projectRepo: projectRepo,
	}
}

func (s *ProjectService) GetAll(ctx context.Context, limit, offset int) (dtos.PaginatedResponse[dtos.ProjectSummary], error) {
	projects, total, err := s.projectRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return dtos.PaginatedResponse[dtos.ProjectSummary]{}, err
	}

	// Extract project IDs to fetch spending
	projectIDs := make([]int64, len(projects))
	for i, p := range projects {
		projectIDs[i] = p.ID
	}

	// Get spending for all projects at once
	spendingMap, err := s.projectRepo.GetActualSpendingForProjects(ctx, projectIDs)
	if err != nil {
		return dtos.PaginatedResponse[dtos.ProjectSummary]{}, err
	}

	result := make([]dtos.ProjectSummary, len(projects))
	for i, p := range projects {
		dto := toProjectSummaryDTO(p)
		dto.CurrentSpending = spendingMap[p.ID]
		result[i] = dto
	}

	page := (offset / limit) + 1
	return dtos.NewPaginatedResponse(result, total, page, limit), nil
}

func (s *ProjectService) GetByID(ctx context.Context, id int64) (*dtos.Project, error) {
	project, err := s.projectRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrProjectNotFound
		}
		return nil, err
	}

	// Get actual spending for this project
	currentSpending, err := s.projectRepo.GetActualSpending(ctx, id)
	if err != nil {
		return nil, err
	}

	dto := toProjectDTO(project)
	dto.CurrentSpending = currentSpending
	return &dto, nil
}

func (s *ProjectService) Create(ctx context.Context, req dtos.CreateProject) (*dtos.Project, error) {
	// Validate that warning cost is less than or equal to total cost
	if req.WarningCost > req.TotalCost {
		return nil, ErrWarningCostExceedsTotalCost
	}

	project := &models.Project{
		Name:               req.Name,
		Type:               req.Type,
		Description:        req.Description,
		ClientPhone:        req.ClientPhone,
		Location:           req.Location,
		StartDate:          req.StartDate,
		Duration:           req.Duration,
		WarningCost:        req.WarningCost,
		TotalCost:          req.TotalCost,
		Status:             "in_progress",
		ProgressPercentage: 0,
		Notes:              req.Notes,
		CreatedBy:          req.CreatedBy,
	}

	created, err := s.projectRepo.Create(ctx, project)
	if err != nil {
		return nil, err
	}

	dto := toProjectDTO(created)
	return &dto, nil
}

func (s *ProjectService) Update(ctx context.Context, id int64, req dtos.UpdateProject) (*dtos.Project, error) {
	existing, err := s.projectRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrProjectNotFound
		}
		return nil, err
	}

	// Apply partial updates
	if req.Name != nil {
		existing.Name = *req.Name
	}
	if req.Type != nil {
		existing.Type = req.Type
	}
	if req.Description != nil {
		existing.Description = req.Description
	}
	if req.ClientPhone != nil {
		existing.ClientPhone = req.ClientPhone
	}
	if req.Location != nil {
		existing.Location = req.Location
	}
	if req.StartDate != nil {
		existing.StartDate = req.StartDate
	}
	if req.Duration != nil {
		existing.Duration = req.Duration
	}
	if req.WarningCost != nil {
		existing.WarningCost = *req.WarningCost
	}
	if req.TotalCost != nil {
		existing.TotalCost = *req.TotalCost
	}
	if req.Status != nil {
		existing.Status = *req.Status
	}
	if req.ProgressPercentage != nil {
		existing.ProgressPercentage = *req.ProgressPercentage
	}
	if req.Notes != nil {
		existing.Notes = req.Notes
	}

	// Validate that warning cost is less than or equal to total cost
	if existing.WarningCost > existing.TotalCost {
		return nil, ErrWarningCostExceedsTotalCost
	}

	updated, err := s.projectRepo.Update(ctx, id, existing)
	if err != nil {
		return nil, err
	}

	dto := toProjectDTO(updated)
	return &dto, nil
}

func (s *ProjectService) Delete(ctx context.Context, id int64) error {
	err := s.projectRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrProjectNotFound
		}
		return err
	}
	return nil
}

func (s *ProjectService) GetStats(ctx context.Context, period string) (*dtos.ProjectStatsResponse, error) {
	stats, err := s.projectRepo.GetStats(ctx, period)
	if err != nil {
		return nil, err
	}

	return &dtos.ProjectStatsResponse{
		Total:           stats.Total,
		Pending:         stats.Pending,
		InProgress:      stats.InProgress,
		Completed:       stats.Completed,
		TotalBudget:     stats.TotalBudget,
		AverageProgress: stats.AverageProgress,
	}, nil
}

// UpdateProgressPercentageWithTx updates the project progress percentage by adding the specified amount within a transaction
func (s *ProjectService) UpdateProgressPercentageWithTx(ctx context.Context, tx pgx.Tx, projectID int64, percentageToAdd float64) error {
	return s.projectRepo.UpdateProgressPercentageWithTx(ctx, tx, projectID, percentageToAdd)
}

// DTO conversion helpers
func toProjectSummaryDTO(p models.Project) dtos.ProjectSummary {
	return dtos.ProjectSummary{
		ID:                 p.ID,
		Name:               p.Name,
		Type:               p.Type,
		Description:        p.Description,
		ClientPhone:        p.ClientPhone,
		Location:           p.Location,
		StartDate:          p.StartDate,
		Duration:           p.Duration,
		Status:             p.Status,
		ProgressPercentage: p.ProgressPercentage,
		WarningCost:        p.WarningCost,
		TotalCost:          p.TotalCost,
		Notes:              p.Notes,
	}
}

func toProjectDTO(p *models.Project) dtos.Project {
	return dtos.Project{
		ID:                 p.ID,
		Name:               p.Name,
		Type:               p.Type,
		Description:        p.Description,
		ClientPhone:        p.ClientPhone,
		Location:           p.Location,
		StartDate:          p.StartDate,
		Duration:           p.Duration,
		WarningCost:        p.WarningCost,
		TotalCost:          p.TotalCost,
		Status:             p.Status,
		ProgressPercentage: p.ProgressPercentage,
		Notes:              p.Notes,
		CreatedBy:          p.CreatedBy,
		CreatedAt:          p.CreatedAt,
	}
}
