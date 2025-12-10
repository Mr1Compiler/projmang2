package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafa91ameen/prjalgo/backend/internal/models"
)

type TeamMemberRepositoryInterface interface {
	GetAll(ctx context.Context, limit, offset int) ([]models.TeamMember, int64, error)
	GetByID(ctx context.Context, id int64) (*models.TeamMember, error)
	GetByProjectID(ctx context.Context, projectID int64) ([]models.TeamMember, error)
	GetByUserID(ctx context.Context, userID int64) ([]models.TeamMember, error)
	Create(ctx context.Context, teamMember *models.TeamMember) (*models.TeamMember, error)
	Delete(ctx context.Context, id int64) error
}

type TeamMemberRepository struct {
	db *pgxpool.Pool
}

func NewTeamMemberRepository(db *pgxpool.Pool) *TeamMemberRepository {
	return &TeamMemberRepository{db: db}
}

func (r *TeamMemberRepository) GetAll(ctx context.Context, limit, offset int) ([]models.TeamMember, int64, error) {
	// Get total count
	var total int64
	countQuery := `SELECT COUNT(*) FROM teamMembers`
	if err := r.db.QueryRow(ctx, countQuery).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT id, projectId, userId
		FROM teamMembers
		ORDER BY createdAt DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var teamMembers []models.TeamMember
	for rows.Next() {
		var tm models.TeamMember
		err := rows.Scan(&tm.ID, &tm.ProjectID, &tm.UserID)
		if err != nil {
			return nil, 0, err
		}
		teamMembers = append(teamMembers, tm)
	}

	return teamMembers, total, rows.Err()
}

func (r *TeamMemberRepository) GetByID(ctx context.Context, id int64) (*models.TeamMember, error) {
	query := `
		SELECT id, projectId, userId
		FROM teamMembers
		WHERE id = $1
	`

	var tm models.TeamMember
	err := r.db.QueryRow(ctx, query, id).Scan(
		&tm.ID, &tm.ProjectID, &tm.UserID,
	)

	if err != nil {
		return nil, err
	}

	return &tm, nil
}

func (r *TeamMemberRepository) GetByProjectID(ctx context.Context, projectID int64) ([]models.TeamMember, error) {
	query := `
		SELECT id, projectId, userId
		FROM teamMembers
		WHERE projectId = $1
		ORDER BY createdAt DESC
	`

	rows, err := r.db.Query(ctx, query, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teamMembers []models.TeamMember
	for rows.Next() {
		var tm models.TeamMember
		err := rows.Scan(&tm.ID, &tm.ProjectID, &tm.UserID)
		if err != nil {
			return nil, err
		}
		teamMembers = append(teamMembers, tm)
	}

	return teamMembers, rows.Err()
}

func (r *TeamMemberRepository) GetByUserID(ctx context.Context, userID int64) ([]models.TeamMember, error) {
	query := `
		SELECT id, projectId, userId, createdAt, updatedAt
		FROM teamMembers
		WHERE userId = $1
		ORDER BY createdAt DESC
	`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teamMembers []models.TeamMember
	for rows.Next() {
		var tm models.TeamMember
		err := rows.Scan(&tm.ID, &tm.ProjectID, &tm.UserID, &tm.CreatedAt, &tm.UpdatedAt)
		if err != nil {
			return nil, err
		}
		teamMembers = append(teamMembers, tm)
	}

	return teamMembers, rows.Err()
}

func (r *TeamMemberRepository) Create(ctx context.Context, teamMember *models.TeamMember) (*models.TeamMember, error) {
	query := `
		INSERT INTO teamMembers (projectId, userId)
		VALUES ($1, $2)
		RETURNING id, createdAt, updatedAt
	`

	err := r.db.QueryRow(ctx, query,
		teamMember.ProjectID, teamMember.UserID,
	).Scan(&teamMember.ID, &teamMember.CreatedAt, &teamMember.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return teamMember, nil
}

func (r *TeamMemberRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM teamMembers WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
