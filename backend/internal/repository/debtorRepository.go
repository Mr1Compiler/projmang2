package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafa91ameen/prjalgo/backend/internal/models"
)

type DebtorRepositoryInterface interface {
	GetAll(ctx context.Context, limit, offset int) ([]models.Debtor, int64, error)
	GetByID(ctx context.Context, id int64) (*models.Debtor, error)
	Create(ctx context.Context, debtor *models.Debtor) (*models.Debtor, error)
	Update(ctx context.Context, id int64, debtor *models.Debtor) (*models.Debtor, error)
	Delete(ctx context.Context, id int64) error
	GetStats(ctx context.Context, period string) (*DebtorStatsResult, error)
}

// DebtorStatsResult holds aggregated debtor statistics
type DebtorStatsResult struct {
	Total         int64   `json:"total"`
	Active        int64   `json:"active"`
	Paid          int64   `json:"paid"`
	TotalDebt     float64 `json:"totalDebt"`
	ActiveDebt    float64 `json:"activeDebt"`
	AverageDebt   float64 `json:"averageDebt"`
}

type DebtorRepository struct {
	db *pgxpool.Pool
}

func NewDebtorRepository(db *pgxpool.Pool) *DebtorRepository {
	return &DebtorRepository{db: db}
}

func (r *DebtorRepository) GetAll(ctx context.Context, limit, offset int) ([]models.Debtor, int64, error) {
	// Get total count
	var total int64
	countQuery := `SELECT COUNT(*) FROM debtors`
	if err := r.db.QueryRow(ctx, countQuery).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT id, name, email, phone, totalDebt, currency, dueDate, status
		FROM debtors
		ORDER BY createdAt DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var debtors []models.Debtor
	for rows.Next() {
		var d models.Debtor
		err := rows.Scan(
			&d.ID, &d.Name, &d.Email, &d.Phone, &d.TotalDebt,
			&d.Currency, &d.DueDate, &d.Status,
		)
		if err != nil {
			return nil, 0, err
		}
		debtors = append(debtors, d)
	}

	return debtors, total, rows.Err()
}

func (r *DebtorRepository) GetByID(ctx context.Context, id int64) (*models.Debtor, error) {
	query := `
		SELECT id, name, email, phone, totalDebt, currency, dueDate, status, notes, createdBy, createdAt
		FROM debtors
		WHERE id = $1
	`

	var d models.Debtor
	err := r.db.QueryRow(ctx, query, id).Scan(
		&d.ID, &d.Name, &d.Email, &d.Phone, &d.TotalDebt, &d.Currency,
		&d.DueDate, &d.Status, &d.Notes, &d.CreatedBy, &d.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &d, nil
}

func (r *DebtorRepository) Create(ctx context.Context, debtor *models.Debtor) (*models.Debtor, error) {
	query := `
		INSERT INTO debtors (name, email, phone, totalDebt, currency, dueDate, status, notes, createdBy)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		debtor.Name, debtor.Email, debtor.Phone, debtor.TotalDebt, debtor.Currency,
		debtor.DueDate, debtor.Status, debtor.Notes, debtor.CreatedBy,
	).Scan(&debtor.ID, &debtor.CreatedAt)

	if err != nil {
		return nil, err
	}

	return debtor, nil
}

func (r *DebtorRepository) Update(ctx context.Context, id int64, debtor *models.Debtor) (*models.Debtor, error) {
	query := `
		UPDATE debtors
		SET name = $1, email = $2, phone = $3, totalDebt = $4, currency = $5,
		    dueDate = $6, status = $7, notes = $8
		WHERE id = $9
		RETURNING id, name, email, phone, totalDebt, currency, dueDate, status, notes, createdBy, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		debtor.Name, debtor.Email, debtor.Phone, debtor.TotalDebt, debtor.Currency,
		debtor.DueDate, debtor.Status, debtor.Notes, id,
	).Scan(
		&debtor.ID, &debtor.Name, &debtor.Email, &debtor.Phone, &debtor.TotalDebt,
		&debtor.Currency, &debtor.DueDate, &debtor.Status, &debtor.Notes,
		&debtor.CreatedBy, &debtor.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return debtor, nil
}

func (r *DebtorRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM debtors WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *DebtorRepository) GetStats(ctx context.Context, period string) (*DebtorStatsResult, error) {
	whereClause := buildDebtorPeriodWhereClause(period, "createdAt")

	query := `
		SELECT
			COUNT(*) as total,
			COUNT(*) FILTER (WHERE status = 'active') as active,
			COUNT(*) FILTER (WHERE status = 'paid') as paid,
			COALESCE(SUM(totalDebt), 0) as total_debt,
			COALESCE(SUM(totalDebt) FILTER (WHERE status = 'active'), 0) as active_debt,
			COALESCE(AVG(totalDebt), 0) as average_debt
		FROM debtors
	` + whereClause

	var stats DebtorStatsResult
	err := r.db.QueryRow(ctx, query).Scan(
		&stats.Total,
		&stats.Active,
		&stats.Paid,
		&stats.TotalDebt,
		&stats.ActiveDebt,
		&stats.AverageDebt,
	)
	if err != nil {
		return nil, err
	}

	return &stats, nil
}

func buildDebtorPeriodWhereClause(period, dateColumn string) string {
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
