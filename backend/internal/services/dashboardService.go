package services

import (
	"context"
	"errors"
	"time"

	"github.com/mustafa91ameen/prjalgo/backend/internal/dtos"
	"github.com/mustafa91ameen/prjalgo/backend/internal/repository"
)

var (
	ErrDashboardStatsNotFound  = errors.New("dashboard stats not found")
	ErrFinancialStatsNotFound  = errors.New("financial stats not found")
	ErrProjectProgressNotFound = errors.New("project progress not found")
	ErrTaskSummaryNotFound     = errors.New("task summary not found")
	ErrActivitiesNotFound      = errors.New("activities not found")
)

type DashboardService struct {
	dashboardRepo repository.DashboardRepositoryInterface
}

func NewDashboardService(dashboardRepo repository.DashboardRepositoryInterface) *DashboardService {
	return &DashboardService{
		dashboardRepo: dashboardRepo,
	}
}

func (s *DashboardService) GetStats(ctx context.Context) (*dtos.DashboardStats, error) {
	stats, err := s.dashboardRepo.GetProjectStats(ctx)
	if err != nil {
		return nil, err
	}

	return &dtos.DashboardStats{
		TotalProjects:     stats.TotalProjects,
		ActiveProjects:    stats.ActiveProjects,
		CompletedProjects: stats.CompletedProjects,
		TotalEngineers:    stats.TotalEngineers,
	}, nil
}

func (s *DashboardService) GetFinancial(ctx context.Context, period string, month *time.Time) (*dtos.DashboardFinancial, error) {
	stats, err := s.dashboardRepo.GetFinancialStats(ctx, period, month)
	if err != nil {
		return nil, err
	}

	return &dtos.DashboardFinancial{
		TotalIncome:   stats.TotalIncome,
		TotalExpenses: stats.TotalExpenses,
	}, nil
}

func (s *DashboardService) GetProjectProgress(ctx context.Context, status string, limit int) (*dtos.DashboardProjectProgress, error) {
	projects, err := s.dashboardRepo.GetProjectProgress(ctx, status, limit)
	if err != nil {
		return nil, err
	}

	result := make([]dtos.ProjectProgressItem, len(projects))
	for i, p := range projects {
		result[i] = dtos.ProjectProgressItem{
			ID:       p.ID,
			Name:     p.Name,
			Progress: p.Progress,
		}
	}

	return &dtos.DashboardProjectProgress{
		Projects: result,
	}, nil
}

func (s *DashboardService) GetTaskSummary(ctx context.Context, month time.Time) (*dtos.DashboardTaskSummary, error) {
	summary, err := s.dashboardRepo.GetTaskSummary(ctx, month)
	if err != nil {
		return nil, err
	}

	return &dtos.DashboardTaskSummary{
		Labels:    summary.Labels,
		Completed: summary.Completed,
		Pending:   summary.Pending,
	}, nil
}

func (s *DashboardService) GetActivities(ctx context.Context, limit, offset int) (dtos.PaginatedResponse[dtos.ActivityResponse], error) {
	activities, total, err := s.dashboardRepo.GetRecentActivities(ctx, limit, offset)
	if err != nil {
		return dtos.PaginatedResponse[dtos.ActivityResponse]{}, err
	}

	result := make([]dtos.ActivityResponse, len(activities))
	for i, a := range activities {
		result[i] = dtos.ActivityResponse{
			ID:         a.ID,
			ActorName:  a.ActorName,
			Action:     a.Action,
			TargetType: a.TargetType,
			TargetName: a.TargetName,
			CreatedAt:  a.CreatedAt.Format(time.RFC3339),
		}
	}

	page := (offset / limit) + 1
	return dtos.NewPaginatedResponse(result, total, page, limit), nil
}
