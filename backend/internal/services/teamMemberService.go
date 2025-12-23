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
	ErrTeamMemberNotFound      = errors.New("team member not found")
	ErrTeamMemberStatsNotFound = errors.New("team member stats not found")
)

type TeamMemberService struct {
	teamMemberRepo repository.TeamMemberRepositoryInterface
}

func NewTeamMemberService(teamMemberRepo repository.TeamMemberRepositoryInterface) *TeamMemberService {
	return &TeamMemberService{
		teamMemberRepo: teamMemberRepo,
	}
}

func (s *TeamMemberService) GetAll(ctx context.Context, limit, offset int) (dtos.PaginatedResponse[dtos.TeamMember], error) {
	teamMembers, total, err := s.teamMemberRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return dtos.PaginatedResponse[dtos.TeamMember]{}, err
	}

	result := make([]dtos.TeamMember, len(teamMembers))
	for i, tm := range teamMembers {
		result[i] = toTeamMemberDTO(tm)
	}

	page := (offset / limit) + 1
	return dtos.NewPaginatedResponse(result, total, page, limit), nil
}

func (s *TeamMemberService) GetByID(ctx context.Context, id int64) (*dtos.TeamMember, error) {
	teamMember, err := s.teamMemberRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrTeamMemberNotFound
		}
		return nil, err
	}
	dto := toTeamMemberDTO(*teamMember)
	return &dto, nil
}

func (s *TeamMemberService) GetByProjectID(ctx context.Context, projectID int64) ([]dtos.TeamMember, error) {
	teamMembers, err := s.teamMemberRepo.GetByProjectID(ctx, projectID)
	if err != nil {
		return nil, err
	}

	result := make([]dtos.TeamMember, len(teamMembers))
	for i, tm := range teamMembers {
		result[i] = toTeamMemberDTO(tm)
	}
	return result, nil
}

func (s *TeamMemberService) GetByUserID(ctx context.Context, userID int64) ([]dtos.TeamMember, error) {
	teamMembers, err := s.teamMemberRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := make([]dtos.TeamMember, len(teamMembers))
	for i, tm := range teamMembers {
		result[i] = toTeamMemberDTO(tm)
	}
	return result, nil
}

func (s *TeamMemberService) Create(ctx context.Context, req dtos.CreateTeamMember) (*dtos.TeamMember, error) {
	teamMember := &models.TeamMember{
		ProjectID: req.ProjectID,
		UserID:    req.UserID,
	}

	created, err := s.teamMemberRepo.Create(ctx, teamMember)
	if err != nil {
		return nil, err
	}

	dto := toTeamMemberDTO(*created)
	return &dto, nil
}

func (s *TeamMemberService) Delete(ctx context.Context, id int64) error {
	err := s.teamMemberRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrTeamMemberNotFound
		}
		return err
	}
	return nil
}

func (s *TeamMemberService) GetStats(ctx context.Context, period string) (*dtos.TeamMemberStatsResponse, error) {
	stats, err := s.teamMemberRepo.GetStats(ctx, period)
	if err != nil {
		return nil, err
	}

	return &dtos.TeamMemberStatsResponse{
		Total:          stats.Total,
		UniqueUsers:    stats.UniqueUsers,
		UniqueProjects: stats.UniqueProjects,
		AvgPerProject:  stats.AvgPerProject,
	}, nil
}

// DTO conversion helper
func toTeamMemberDTO(tm models.TeamMember) dtos.TeamMember {
	dto := dtos.TeamMember{
		ID:        tm.ID,
		ProjectID: tm.ProjectID,
		UserID:    tm.UserID,
	}

	// Only include timestamps if they are not zero values
	if !tm.CreatedAt.IsZero() {
		dto.CreatedAt = &tm.CreatedAt
	}
	if !tm.UpdatedAt.IsZero() {
		dto.UpdatedAt = &tm.UpdatedAt
	}

	return dto
}
