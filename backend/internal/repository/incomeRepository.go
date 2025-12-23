package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafa91ameen/prjalgo/backend/internal/models"
)

type IncomeRepositoryInterface interface {
	GetAll(ctx context.Context, limit, offset int) ([]models.Income, int64, error)
	GetByID(ctx context.Context, id int64) (*models.Income, error)
	Create(ctx context.Context, income *models.Income) (*models.Income, error)
	Update(ctx context.Context, id int64, income *models.Income) (*models.Income, error)
	Delete(ctx context.Context, id int64) error
	GetStats(ctx context.Context, period string) (*IncomeStatsResult, error)
}

// IncomeStatsResult holds aggregated income statistics
type IncomeStatsResult struct {
	Total         int64   `json:"total"`
	TotalAmount   float64 `json:"totalAmount"`
	Pending       int64   `json:"pending"`
	Approved      int64   `json:"approved"`
	Rejected      int64   `json:"rejected"`
	AverageAmount float64 `json:"averageAmount"`
}

type IncomeRepository struct {
	db *pgxpool.Pool
}

func NewIncomeRepository(db *pgxpool.Pool) *IncomeRepository {
	return &IncomeRepository{db: db}
}

func (r *IncomeRepository) GetAll(ctx context.Context, limit, offset int) ([]models.Income, int64, error) {
	// Get total count
	var total int64
	countQuery := `SELECT COUNT(*) FROM income`
	if err := r.db.QueryRow(ctx, countQuery).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT id, name, amount, type, incomeDate, status
		FROM income
		ORDER BY createdAt DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var incomes []models.Income
	for rows.Next() {
		var i models.Income
		err := rows.Scan(
			&i.ID, &i.Name, &i.Amount, &i.Type, &i.IncomeDate, &i.Status,
		)
		if err != nil {
			return nil, 0, err
		}
		incomes = append(incomes, i)
	}

	return incomes, total, rows.Err()
}

func (r *IncomeRepository) GetByID(ctx context.Context, id int64) (*models.Income, error) {
	query := `
		SELECT id, name, amount, type, incomeDate, status, notes, createdBy, createdAt
		FROM income
		WHERE id = $1
	`

	var i models.Income
	err := r.db.QueryRow(ctx, query, id).Scan(
		&i.ID, &i.Name, &i.Amount, &i.Type, &i.IncomeDate,
		&i.Status, &i.Notes, &i.CreatedBy, &i.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &i, nil
}

func (r *IncomeRepository) Create(ctx context.Context, income *models.Income) (*models.Income, error) {
	query := `
		INSERT INTO income (name, amount, type, incomeDate, status, notes, createdBy)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		income.Name, income.Amount, income.Type, income.IncomeDate,
		income.Status, income.Notes, income.CreatedBy,
	).Scan(&income.ID, &income.CreatedAt)

	if err != nil {
		return nil, err
	}

	return income, nil
}

func (r *IncomeRepository) Update(ctx context.Context, id int64, income *models.Income) (*models.Income, error) {
	query := `
		UPDATE income
		SET name = $1, amount = $2, type = $3, incomeDate = $4, status = $5, notes = $6
		WHERE id = $7
		RETURNING id, name, amount, type, incomeDate, status, notes, createdBy, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		income.Name, income.Amount, income.Type, income.IncomeDate,
		income.Status, income.Notes, id,
	).Scan(
		&income.ID, &income.Name, &income.Amount, &income.Type, &income.IncomeDate,
		&income.Status, &income.Notes, &income.CreatedBy, &income.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return income, nil
}

func (r *IncomeRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM income WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *IncomeRepository) GetStats(ctx context.Context, period string) (*IncomeStatsResult, error) {
	whereClause := buildIncomePeriodWhereClause(period, "createdAt")

	query := `
		SELECT
			COUNT(*) as total,
			COALESCE(SUM(amount), 0) as total_amount,
			COUNT(*) FILTER (WHERE status = 'pending') as pending,
			COUNT(*) FILTER (WHERE status = 'approved') as approved,
			COUNT(*) FILTER (WHERE status = 'rejected') as rejected,
			COALESCE(AVG(amount), 0) as average_amount
		FROM income
	` + whereClause

	var stats IncomeStatsResult
	err := r.db.QueryRow(ctx, query).Scan(
		&stats.Total,
		&stats.TotalAmount,
		&stats.Pending,
		&stats.Approved,
		&stats.Rejected,
		&stats.AverageAmount,
	)
	if err != nil {
		return nil, err
	}

	return &stats, nil
}

func buildIncomePeriodWhereClause(period, dateColumn string) string {
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
