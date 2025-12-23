package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafa91ameen/prjalgo/backend/internal/models"
)

type WorkCategoryRepositoryInterface interface {
	GetAll(ctx context.Context, limit, offset int) ([]models.WorkCategory, int64, error)
	GetByID(ctx context.Context, id int64) (*models.WorkCategory, error)
	Create(ctx context.Context, category *models.WorkCategory) (*models.WorkCategory, error)
	Update(ctx context.Context, id int64, category *models.WorkCategory) (*models.WorkCategory, error)
	Delete(ctx context.Context, id int64) error
	GetStats(ctx context.Context, period string) (*WorkCategoryStatsResult, error)
}

// WorkCategoryStatsResult holds aggregated work category statistics
type WorkCategoryStatsResult struct {
	Total            int64 `json:"total"`
	Active           int64 `json:"active"`
	Inactive         int64 `json:"inactive"`
	TotalSubcategory int64 `json:"totalSubcategory"`
}

type WorkCategoryRepository struct {
	db *pgxpool.Pool
}

func NewWorkCategoryRepository(db *pgxpool.Pool) *WorkCategoryRepository {
	return &WorkCategoryRepository{db: db}
}

func (r *WorkCategoryRepository) GetAll(ctx context.Context, limit, offset int) ([]models.WorkCategory, int64, error) {
	// Get total count
	var total int64
	countQuery := `SELECT COUNT(*) FROM workCategories`
	if err := r.db.QueryRow(ctx, countQuery).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT id, name, description, status
		FROM workCategories
		ORDER BY createdAt DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var categories []models.WorkCategory
	for rows.Next() {
		var c models.WorkCategory
		err := rows.Scan(&c.ID, &c.Name, &c.Description, &c.Status)
		if err != nil {
			return nil, 0, err
		}
		categories = append(categories, c)
	}

	return categories, total, rows.Err()
}

func (r *WorkCategoryRepository) GetByID(ctx context.Context, id int64) (*models.WorkCategory, error) {
	query := `
		SELECT id, name, description, status, createdBy, createdAt
		FROM workCategories
		WHERE id = $1
	`

	var c models.WorkCategory
	err := r.db.QueryRow(ctx, query, id).Scan(
		&c.ID, &c.Name, &c.Description, &c.Status,
		&c.CreatedBy, &c.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *WorkCategoryRepository) Create(ctx context.Context, category *models.WorkCategory) (*models.WorkCategory, error) {
	query := `
		INSERT INTO workCategories (name, description, status, createdBy)
		VALUES ($1, $2, $3, $4)
		RETURNING id, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		category.Name, category.Description, category.Status, category.CreatedBy,
	).Scan(&category.ID, &category.CreatedAt)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (r *WorkCategoryRepository) Update(ctx context.Context, id int64, category *models.WorkCategory) (*models.WorkCategory, error) {
	query := `
		UPDATE workCategories
		SET name = $1, description = $2, status = $3
		WHERE id = $4
		RETURNING id, name, description, status, createdBy, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		category.Name, category.Description, category.Status, id,
	).Scan(
		&category.ID, &category.Name, &category.Description, &category.Status,
		&category.CreatedBy, &category.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (r *WorkCategoryRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM workCategories WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *WorkCategoryRepository) GetStats(ctx context.Context, period string) (*WorkCategoryStatsResult, error) {
	whereClause := buildWorkCategoryPeriodWhereClause(period, "createdAt")

	query := `
		SELECT
			(SELECT COUNT(*) FROM workCategories` + whereClause + `) as total,
			(SELECT COUNT(*) FROM workCategories WHERE status = 'active'` + buildWorkCategoryAndClause(period, "createdAt") + `) as active,
			(SELECT COUNT(*) FROM workCategories WHERE status = 'inactive'` + buildWorkCategoryAndClause(period, "createdAt") + `) as inactive,
			(SELECT COUNT(*) FROM workSubCategories` + whereClause + `) as total_subcategory
	`

	var stats WorkCategoryStatsResult
	err := r.db.QueryRow(ctx, query).Scan(
		&stats.Total,
		&stats.Active,
		&stats.Inactive,
		&stats.TotalSubcategory,
	)
	if err != nil {
		return nil, err
	}

	return &stats, nil
}

func buildWorkCategoryPeriodWhereClause(period, dateColumn string) string {
	now := time.Now()
	switch period {
	case "month":
		startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		return " WHERE " + dateColumn + " >= '" + startOfMonth.Format("2006-01-02") + "'"
	case "year":
		startOfYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
		return " WHERE " + dateColumn + " >= '" + startOfYear.Format("2006-01-02") + "'"
	default:
		return ""
	}
}

func buildWorkCategoryAndClause(period, dateColumn string) string {
	now := time.Now()
	switch period {
	case "month":
		startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		return " AND " + dateColumn + " >= '" + startOfMonth.Format("2006-01-02") + "'"
	case "year":
		startOfYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
		return " AND " + dateColumn + " >= '" + startOfYear.Format("2006-01-02") + "'"
	default:
		return ""
	}
}
