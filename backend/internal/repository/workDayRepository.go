package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafa91ameen/prjalgo/backend/internal/models"
)

type WorkDayRepositoryInterface interface {
	GetAll(ctx context.Context, limit, offset int) ([]models.WorkDay, int64, error)
	GetByID(ctx context.Context, id int64) (*models.WorkDay, error)
	GetByIDWithTx(ctx context.Context, tx pgx.Tx, id int64) (*models.WorkDay, error)
	GetByProjectID(ctx context.Context, projectID int64) ([]models.WorkDay, error)
	Create(ctx context.Context, workDay *models.WorkDay) (*models.WorkDay, error)
	Update(ctx context.Context, id int64, workDay *models.WorkDay) (*models.WorkDay, error)
	UpdateWithTx(ctx context.Context, tx pgx.Tx, id int64, workDay *models.WorkDay) (*models.WorkDay, error)
	UpdateTotalCostWithTx(ctx context.Context, tx pgx.Tx, id int64, amountToAdd float64) error
	Delete(ctx context.Context, id int64) error
}

type WorkDayRepository struct {
	db *pgxpool.Pool
}

func NewWorkDayRepository(db *pgxpool.Pool) *WorkDayRepository {
	return &WorkDayRepository{db: db}
}

func (r *WorkDayRepository) GetAll(ctx context.Context, limit, offset int) ([]models.WorkDay, int64, error) {
	// Get total count
	var total int64
	countQuery := `SELECT COUNT(*) FROM workDays`
	if err := r.db.QueryRow(ctx, countQuery).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT id, projectId, workSubCategoryId, workDate, status, totalCost
		FROM workDays
		ORDER BY createdAt DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var workDays []models.WorkDay
	for rows.Next() {
		var w models.WorkDay
		err := rows.Scan(
			&w.ID, &w.ProjectID, &w.WorkSubCategoryID, &w.WorkDate,
			&w.Status, &w.TotalCost,
		)
		if err != nil {
			return nil, 0, err
		}
		workDays = append(workDays, w)
	}

	return workDays, total, rows.Err()
}

func (r *WorkDayRepository) GetByID(ctx context.Context, id int64) (*models.WorkDay, error) {
	query := `
		SELECT id, projectId, workSubCategoryId, workDate, description,
		       status, totalCost, notes, createdBy, createdAt
		FROM workDays
		WHERE id = $1
	`

	var w models.WorkDay
	err := r.db.QueryRow(ctx, query, id).Scan(
		&w.ID, &w.ProjectID, &w.WorkSubCategoryID, &w.WorkDate, &w.Description,
		&w.Status, &w.TotalCost, &w.Notes, &w.CreatedBy, &w.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &w, nil
}

func (r *WorkDayRepository) GetByIDWithTx(ctx context.Context, tx pgx.Tx, id int64) (*models.WorkDay, error) {
	query := `
		SELECT id, projectId, workSubCategoryId, workDate, description,
		       status, totalCost, notes, createdBy, createdAt
		FROM workDays
		WHERE id = $1
		FOR UPDATE
	`

	var w models.WorkDay
	err := tx.QueryRow(ctx, query, id).Scan(
		&w.ID, &w.ProjectID, &w.WorkSubCategoryID, &w.WorkDate, &w.Description,
		&w.Status, &w.TotalCost, &w.Notes, &w.CreatedBy, &w.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &w, nil
}

func (r *WorkDayRepository) GetByProjectID(ctx context.Context, projectID int64) ([]models.WorkDay, error) {
	query := `
		SELECT id, projectId, workSubCategoryId, workDate, status, totalCost
		FROM workDays
		WHERE projectId = $1
		ORDER BY workDate DESC
	`

	rows, err := r.db.Query(ctx, query, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workDays []models.WorkDay
	for rows.Next() {
		var w models.WorkDay
		err := rows.Scan(
			&w.ID, &w.ProjectID, &w.WorkSubCategoryID, &w.WorkDate,
			&w.Status, &w.TotalCost,
		)
		if err != nil {
			return nil, err
		}
		workDays = append(workDays, w)
	}

	return workDays, rows.Err()
}

func (r *WorkDayRepository) Create(ctx context.Context, workDay *models.WorkDay) (*models.WorkDay, error) {
	query := `
		INSERT INTO workDays (projectId, workSubCategoryId, workDate, description,
		                      status, totalCost, notes, createdBy)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		workDay.ProjectID, workDay.WorkSubCategoryID, workDay.WorkDate, workDay.Description,
		workDay.Status, workDay.TotalCost, workDay.Notes, workDay.CreatedBy,
	).Scan(&workDay.ID, &workDay.CreatedAt)

	if err != nil {
		return nil, err
	}

	return workDay, nil
}

func (r *WorkDayRepository) Update(ctx context.Context, id int64, workDay *models.WorkDay) (*models.WorkDay, error) {
	query := `
		UPDATE workDays
		SET workSubCategoryId = $1, workDate = $2, description = $3,
		    status = $4, totalCost = $5, notes = $6
		WHERE id = $7
		RETURNING id, projectId, workSubCategoryId, workDate, description,
		          status, totalCost, notes, createdBy, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		workDay.WorkSubCategoryID, workDay.WorkDate, workDay.Description,
		workDay.Status, workDay.TotalCost, workDay.Notes, id,
	).Scan(
		&workDay.ID, &workDay.ProjectID, &workDay.WorkSubCategoryID, &workDay.WorkDate,
		&workDay.Description, &workDay.Status, &workDay.TotalCost, &workDay.Notes,
		&workDay.CreatedBy, &workDay.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return workDay, nil
}

func (r *WorkDayRepository) UpdateWithTx(ctx context.Context, tx pgx.Tx, id int64, workDay *models.WorkDay) (*models.WorkDay, error) {
	query := `
		UPDATE workDays
		SET workSubCategoryId = $1, workDate = $2, description = $3,
		    status = $4, totalCost = $5, notes = $6
		WHERE id = $7
		RETURNING id, projectId, workSubCategoryId, workDate, description,
		          status, totalCost, notes, createdBy, createdAt
	`

	err := tx.QueryRow(ctx, query,
		workDay.WorkSubCategoryID, workDay.WorkDate, workDay.Description,
		workDay.Status, workDay.TotalCost, workDay.Notes, id,
	).Scan(
		&workDay.ID, &workDay.ProjectID, &workDay.WorkSubCategoryID, &workDay.WorkDate,
		&workDay.Description, &workDay.Status, &workDay.TotalCost, &workDay.Notes,
		&workDay.CreatedBy, &workDay.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return workDay, nil
}

func (r *WorkDayRepository) UpdateTotalCostWithTx(ctx context.Context, tx pgx.Tx, id int64, amountToAdd float64) error {
	query := `UPDATE workDays SET totalCost = totalCost + $1 WHERE id = $2`
	result, err := tx.Exec(ctx, query, amountToAdd, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *WorkDayRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM workDays WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
