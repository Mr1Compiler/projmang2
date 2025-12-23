package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DashboardRepositoryInterface interface {
	GetProjectStats(ctx context.Context) (*ProjectStats, error)
	GetFinancialStats(ctx context.Context, period string, month *time.Time) (*FinancialStats, error)
	GetProjectProgress(ctx context.Context, status string, limit int) ([]ProjectProgress, error)
	GetTaskSummary(ctx context.Context, month time.Time) (*TaskSummary, error)
	GetRecentActivities(ctx context.Context, limit, offset int) ([]ActivityResult, int64, error)
}

// Stats DTOs
type ProjectStats struct {
	TotalProjects     int64 `json:"totalProjects"`
	PendingProjects   int64 `json:"pendingProjects"`
	ActiveProjects    int64 `json:"activeProjects"`
	CompletedProjects int64 `json:"completedProjects"`
	TotalEngineers    int64 `json:"totalEngineers"`
}

type FinancialStats struct {
	TotalIncome   float64 `json:"totalIncome"`
	TotalExpenses float64 `json:"totalExpenses"`
}

type ProjectProgress struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Progress float64 `json:"progress"`
}

type TaskSummary struct {
	Labels    []string `json:"labels"`
	Completed []int    `json:"completed"`
	Pending   []int    `json:"pending"`
}

type ActivityResult struct {
	ID         int64     `json:"id"`
	ActorName  string    `json:"actorName"`
	Action     string    `json:"action"`
	TargetType string    `json:"targetType"`
	TargetName string    `json:"targetName"`
	CreatedAt  time.Time `json:"createdAt"`
}

type DashboardRepository struct {
	db *pgxpool.Pool
}

func NewDashboardRepository(db *pgxpool.Pool) *DashboardRepository {
	return &DashboardRepository{db: db}
}

func (r *DashboardRepository) GetProjectStats(ctx context.Context) (*ProjectStats, error) {
	query := `
		SELECT
			COUNT(*) as total,
			COUNT(*) FILTER (WHERE status = 'pending') as pending,
			COUNT(*) FILTER (WHERE status = 'in_progress') as active,
			COUNT(*) FILTER (WHERE status = 'completed') as completed
		FROM projects
	`

	var stats ProjectStats
	err := r.db.QueryRow(ctx, query).Scan(
		&stats.TotalProjects,
		&stats.PendingProjects,
		&stats.ActiveProjects,
		&stats.CompletedProjects,
	)
	if err != nil {
		return nil, err
	}

	// Get total engineers (team members)
	engineerQuery := `SELECT COUNT(*) FROM teammembers`
	err = r.db.QueryRow(ctx, engineerQuery).Scan(&stats.TotalEngineers)
	if err != nil {
		return nil, err
	}

	return &stats, nil
}

func (r *DashboardRepository) GetFinancialStats(ctx context.Context, period string, month *time.Time) (*FinancialStats, error) {
	var stats FinancialStats

	if period == "month" && month != nil {
		// Get first and last day of the month
		firstDay := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, time.UTC)
		lastDay := firstDay.AddDate(0, 1, -1)

		incomeQuery := `
			SELECT COALESCE(SUM(amount), 0)
			FROM income
			WHERE incomeDate >= $1 AND incomeDate <= $2
		`
		err := r.db.QueryRow(ctx, incomeQuery, firstDay, lastDay).Scan(&stats.TotalIncome)
		if err != nil {
			return nil, err
		}

		expenseQuery := `
			SELECT COALESCE(SUM(amount), 0)
			FROM expenses
			WHERE expenseDate >= $1 AND expenseDate <= $2
		`
		err = r.db.QueryRow(ctx, expenseQuery, firstDay, lastDay).Scan(&stats.TotalExpenses)
		if err != nil {
			return nil, err
		}
	} else {
		// All time
		incomeQuery := `SELECT COALESCE(SUM(amount), 0) FROM income`
		err := r.db.QueryRow(ctx, incomeQuery).Scan(&stats.TotalIncome)
		if err != nil {
			return nil, err
		}

		expenseQuery := `SELECT COALESCE(SUM(amount), 0) FROM expenses`
		err = r.db.QueryRow(ctx, expenseQuery).Scan(&stats.TotalExpenses)
		if err != nil {
			return nil, err
		}
	}

	return &stats, nil
}

func (r *DashboardRepository) GetProjectProgress(ctx context.Context, status string, limit int) ([]ProjectProgress, error) {
	var query string
	var args []interface{}

	if status == "all" {
		query = `
			SELECT id, name, progressPercentage
			FROM projects
			ORDER BY progressPercentage DESC
			LIMIT $1
		`
		args = []interface{}{limit}
	} else {
		query = `
			SELECT id, name, progressPercentage
			FROM projects
			WHERE status = $1
			ORDER BY progressPercentage DESC
			LIMIT $2
		`
		args = []interface{}{status, limit}
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []ProjectProgress
	for rows.Next() {
		var p ProjectProgress
		err := rows.Scan(&p.ID, &p.Name, &p.Progress)
		if err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	return projects, rows.Err()
}

func (r *DashboardRepository) GetTaskSummary(ctx context.Context, month time.Time) (*TaskSummary, error) {
	// Get first and last day of the month
	firstDay := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, -1)
	daysInMonth := lastDay.Day()

	// Initialize arrays
	labels := make([]string, daysInMonth)
	completed := make([]int, daysInMonth)
	pending := make([]int, daysInMonth)

	for i := 0; i < daysInMonth; i++ {
		labels[i] = time.Date(month.Year(), month.Month(), i+1, 0, 0, 0, 0, time.UTC).Format("2")
	}

	// Query completed workdays per day
	completedQuery := `
		SELECT EXTRACT(DAY FROM workDate)::int as day, COUNT(*) as count
		FROM workDays
		WHERE workDate >= $1 AND workDate <= $2 AND status = 'completed'
		GROUP BY EXTRACT(DAY FROM workDate)
	`
	rows, err := r.db.Query(ctx, completedQuery, firstDay, lastDay)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var day, count int
		if err := rows.Scan(&day, &count); err != nil {
			return nil, err
		}
		if day > 0 && day <= daysInMonth {
			completed[day-1] = count
		}
	}

	// Query pending workdays per day
	pendingQuery := `
		SELECT EXTRACT(DAY FROM workDate)::int as day, COUNT(*) as count
		FROM workDays
		WHERE workDate >= $1 AND workDate <= $2 AND status = 'pending'
		GROUP BY EXTRACT(DAY FROM workDate)
	`
	rows2, err := r.db.Query(ctx, pendingQuery, firstDay, lastDay)
	if err != nil {
		return nil, err
	}
	defer rows2.Close()

	for rows2.Next() {
		var day, count int
		if err := rows2.Scan(&day, &count); err != nil {
			return nil, err
		}
		if day > 0 && day <= daysInMonth {
			pending[day-1] = count
		}
	}

	return &TaskSummary{
		Labels:    labels,
		Completed: completed,
		Pending:   pending,
	}, nil
}

func (r *DashboardRepository) GetRecentActivities(ctx context.Context, limit, offset int) ([]ActivityResult, int64, error) {
	// Get total count
	var total int64
	countQuery := `SELECT COUNT(*) FROM auditLogs WHERE status = 'success'`
	if err := r.db.QueryRow(ctx, countQuery).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT
			a.id,
			COALESCE(u.username, 'System') as actorName,
			a.action,
			a.targetType,
			COALESCE(
				CASE a.targetType
					WHEN 'project' THEN (SELECT name FROM projects WHERE id = a.targetId)
					WHEN 'expense' THEN (SELECT name FROM expenses WHERE id = a.targetId)
					WHEN 'income' THEN (SELECT name FROM income WHERE id = a.targetId)
					WHEN 'debtor' THEN (SELECT name FROM debtors WHERE id = a.targetId)
					WHEN 'work_category' THEN (SELECT name FROM workCategories WHERE id = a.targetId)
					WHEN 'work_subcategory' THEN (SELECT name FROM workSubCategories WHERE id = a.targetId)
					WHEN 'user' THEN (SELECT username FROM users WHERE id = a.targetId)
					WHEN 'team_member' THEN (SELECT u2.username FROM teamMembers tm JOIN users u2 ON tm.userId = u2.id WHERE tm.id = a.targetId)
					ELSE NULL
				END,
				'Unknown'
			) as targetName,
			a.createdAt
		FROM auditLogs a
		LEFT JOIN users u ON a.actorId = u.id
		WHERE a.status = 'success'
		ORDER BY a.createdAt DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var activities []ActivityResult
	for rows.Next() {
		var act ActivityResult
		err := rows.Scan(
			&act.ID,
			&act.ActorName,
			&act.Action,
			&act.TargetType,
			&act.TargetName,
			&act.CreatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		activities = append(activities, act)
	}

	return activities, total, rows.Err()
}
