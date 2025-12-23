package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafa91ameen/prjalgo/backend/internal/models"
)

type ExpenseRepositoryInterface interface {
	GetAll(ctx context.Context, limit, offset int) ([]models.Expense, int64, error)
	GetByID(ctx context.Context, id int64) (*models.Expense, error)
	GetByProjectID(ctx context.Context, projectID int64) ([]models.Expense, error)
	Create(ctx context.Context, expense *models.Expense) (*models.Expense, error)
	Update(ctx context.Context, id int64, expense *models.Expense) (*models.Expense, error)
	Delete(ctx context.Context, id int64) error
	GetStats(ctx context.Context, period string) (*ExpenseStatsResult, error)
}

// ExpenseStatsResult holds aggregated expense statistics
type ExpenseStatsResult struct {
	Total         int64   `json:"total"`
	TotalAmount   float64 `json:"totalAmount"`
	Pending       int64   `json:"pending"`
	Approved      int64   `json:"approved"`
	Rejected      int64   `json:"rejected"`
	DebtorCount   int64   `json:"debtorCount"`
	AverageAmount float64 `json:"averageAmount"`
}

type ExpenseRepository struct {
	db *pgxpool.Pool
}

func NewExpenseRepository(db *pgxpool.Pool) *ExpenseRepository {
	return &ExpenseRepository{db: db}
}

func (r *ExpenseRepository) GetAll(ctx context.Context, limit, offset int) ([]models.Expense, int64, error) {
	// Get total count
	var total int64
	countQuery := `SELECT COUNT(*) FROM expenses`
	if err := r.db.QueryRow(ctx, countQuery).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT id, name, amount, type, expenseDate, projectId, isDebtor, debtorId, status
		FROM expenses
		ORDER BY createdAt DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var expenses []models.Expense
	for rows.Next() {
		var e models.Expense
		err := rows.Scan(
			&e.ID, &e.Name, &e.Amount, &e.Type, &e.ExpenseDate,
			&e.ProjectID, &e.IsDebtor, &e.DebtorID, &e.Status,
		)
		if err != nil {
			return nil, 0, err
		}
		expenses = append(expenses, e)
	}

	return expenses, total, rows.Err()
}

func (r *ExpenseRepository) GetByID(ctx context.Context, id int64) (*models.Expense, error) {
	query := `
		SELECT id, name, amount, type, expenseDate, projectId, isDebtor, debtorId, status, notes, createdBy, createdAt
		FROM expenses
		WHERE id = $1
	`

	var e models.Expense
	err := r.db.QueryRow(ctx, query, id).Scan(
		&e.ID, &e.Name, &e.Amount, &e.Type, &e.ExpenseDate, &e.ProjectID,
		&e.IsDebtor, &e.DebtorID, &e.Status, &e.Notes, &e.CreatedBy, &e.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (r *ExpenseRepository) GetByProjectID(ctx context.Context, projectID int64) ([]models.Expense, error) {
	query := `
		SELECT id, name, amount, type, expenseDate, projectId, isDebtor, debtorId, status
		FROM expenses
		WHERE projectId = $1
		ORDER BY expenseDate DESC
	`

	rows, err := r.db.Query(ctx, query, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []models.Expense
	for rows.Next() {
		var e models.Expense
		err := rows.Scan(
			&e.ID, &e.Name, &e.Amount, &e.Type, &e.ExpenseDate,
			&e.ProjectID, &e.IsDebtor, &e.DebtorID, &e.Status,
		)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, e)
	}

	return expenses, rows.Err()
}

func (r *ExpenseRepository) Create(ctx context.Context, expense *models.Expense) (*models.Expense, error) {
	query := `
		INSERT INTO expenses (name, amount, type, expenseDate, projectId, isDebtor, debtorId, status, notes, createdBy)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		expense.Name, expense.Amount, expense.Type, expense.ExpenseDate, expense.ProjectID,
		expense.IsDebtor, expense.DebtorID, expense.Status, expense.Notes, expense.CreatedBy,
	).Scan(&expense.ID, &expense.CreatedAt)

	if err != nil {
		return nil, err
	}

	return expense, nil
}

func (r *ExpenseRepository) Update(ctx context.Context, id int64, expense *models.Expense) (*models.Expense, error) {
	query := `
		UPDATE expenses
		SET name = $1, amount = $2, type = $3, expenseDate = $4, projectId = $5,
		    isDebtor = $6, debtorId = $7, status = $8, notes = $9
		WHERE id = $10
		RETURNING id, name, amount, type, expenseDate, projectId, isDebtor, debtorId, status, notes, createdBy, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		expense.Name, expense.Amount, expense.Type, expense.ExpenseDate, expense.ProjectID,
		expense.IsDebtor, expense.DebtorID, expense.Status, expense.Notes, id,
	).Scan(
		&expense.ID, &expense.Name, &expense.Amount, &expense.Type, &expense.ExpenseDate,
		&expense.ProjectID, &expense.IsDebtor, &expense.DebtorID, &expense.Status, &expense.Notes,
		&expense.CreatedBy, &expense.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return expense, nil
}

func (r *ExpenseRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM expenses WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *ExpenseRepository) GetStats(ctx context.Context, period string) (*ExpenseStatsResult, error) {
	whereClause := buildExpensePeriodWhereClause(period, "createdAt")

	query := `
		SELECT
			COUNT(*) as total,
			COALESCE(SUM(amount), 0) as total_amount,
			COUNT(*) FILTER (WHERE status = 'pending') as pending,
			COUNT(*) FILTER (WHERE status = 'approved') as approved,
			COUNT(*) FILTER (WHERE status = 'rejected') as rejected,
			COUNT(*) FILTER (WHERE isDebtor = true) as debtor_count,
			COALESCE(AVG(amount), 0) as average_amount
		FROM expenses
	` + whereClause

	var stats ExpenseStatsResult
	err := r.db.QueryRow(ctx, query).Scan(
		&stats.Total,
		&stats.TotalAmount,
		&stats.Pending,
		&stats.Approved,
		&stats.Rejected,
		&stats.DebtorCount,
		&stats.AverageAmount,
	)
	if err != nil {
		return nil, err
	}

	return &stats, nil
}

// buildExpensePeriodWhereClause returns a WHERE clause based on the period filter
func buildExpensePeriodWhereClause(period, dateColumn string) string {
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
