package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafa91ameen/prjalgo/backend/internal/models"
)

type WorkDayLaborRepositoryInterface interface {
	GetAll(ctx context.Context, limit, offset int) ([]models.WorkDayLabor, int64, error)
	GetByID(ctx context.Context, id int64) (*models.WorkDayLabor, error)
	GetByWorkDayID(ctx context.Context, workDayID int64) ([]models.WorkDayLabor, error)
	Create(ctx context.Context, labor *models.WorkDayLabor) (*models.WorkDayLabor, error)
	CreateWithTx(ctx context.Context, tx pgx.Tx, labor *models.WorkDayLabor) (*models.WorkDayLabor, error)
	Update(ctx context.Context, id int64, labor *models.WorkDayLabor) (*models.WorkDayLabor, error)
	UpdateWithTx(ctx context.Context, tx pgx.Tx, id int64, labor *models.WorkDayLabor) (*models.WorkDayLabor, error)
	Delete(ctx context.Context, id int64) error
	DeleteWithTx(ctx context.Context, tx pgx.Tx, id int64) error
}

type WorkDayLaborRepository struct {
	db *pgxpool.Pool
}

func NewWorkDayLaborRepository(db *pgxpool.Pool) *WorkDayLaborRepository {
	return &WorkDayLaborRepository{db: db}
}

func (r *WorkDayLaborRepository) GetAll(ctx context.Context, limit, offset int) ([]models.WorkDayLabor, int64, error) {
	// Get total count
	var total int64
	countQuery := `SELECT COUNT(*) FROM workDayLabor`
	if err := r.db.QueryRow(ctx, countQuery).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT id, workDayId, workerName, jobTitle, phone, address,
		       quantity, cost, notes, createdAt
		FROM workDayLabor
		ORDER BY createdAt DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var labors []models.WorkDayLabor
	for rows.Next() {
		var l models.WorkDayLabor
		err := rows.Scan(
			&l.ID, &l.WorkDayID, &l.WorkerName, &l.JobTitle, &l.Phone,
			&l.Address, &l.Quantity, &l.Cost, &l.Notes, &l.CreatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		labors = append(labors, l)
	}

	return labors, total, rows.Err()
}

func (r *WorkDayLaborRepository) GetByID(ctx context.Context, id int64) (*models.WorkDayLabor, error) {
	query := `
		SELECT id, workDayId, workerName, jobTitle, phone, address,
		       quantity, cost, notes, createdAt
		FROM workDayLabor
		WHERE id = $1
	`

	var l models.WorkDayLabor
	err := r.db.QueryRow(ctx, query, id).Scan(
		&l.ID, &l.WorkDayID, &l.WorkerName, &l.JobTitle, &l.Phone,
		&l.Address, &l.Quantity, &l.Cost, &l.Notes, &l.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &l, nil
}

func (r *WorkDayLaborRepository) GetByWorkDayID(ctx context.Context, workDayID int64) ([]models.WorkDayLabor, error) {
	query := `
		SELECT id, workDayId, workerName, jobTitle, phone, address,
		       quantity, cost, notes, createdAt
		FROM workDayLabor
		WHERE workDayId = $1
		ORDER BY createdAt DESC
	`

	rows, err := r.db.Query(ctx, query, workDayID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var labors []models.WorkDayLabor
	for rows.Next() {
		var l models.WorkDayLabor
		err := rows.Scan(
			&l.ID, &l.WorkDayID, &l.WorkerName, &l.JobTitle, &l.Phone,
			&l.Address, &l.Quantity, &l.Cost, &l.Notes, &l.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		labors = append(labors, l)
	}

	return labors, rows.Err()
}

func (r *WorkDayLaborRepository) Create(ctx context.Context, labor *models.WorkDayLabor) (*models.WorkDayLabor, error) {
	query := `
		INSERT INTO workDayLabor (workDayId, workerName, jobTitle, phone, address, quantity, cost, notes)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		labor.WorkDayID, labor.WorkerName, labor.JobTitle, labor.Phone,
		labor.Address, labor.Quantity, labor.Cost, labor.Notes,
	).Scan(&labor.ID, &labor.CreatedAt)

	if err != nil {
		return nil, err
	}

	return labor, nil
}

func (r *WorkDayLaborRepository) CreateWithTx(ctx context.Context, tx pgx.Tx, labor *models.WorkDayLabor) (*models.WorkDayLabor, error) {
	query := `
		INSERT INTO workDayLabor (workDayId, workerName, jobTitle, phone, address, quantity, cost, notes)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, createdAt
	`

	err := tx.QueryRow(ctx, query,
		labor.WorkDayID, labor.WorkerName, labor.JobTitle, labor.Phone,
		labor.Address, labor.Quantity, labor.Cost, labor.Notes,
	).Scan(&labor.ID, &labor.CreatedAt)

	if err != nil {
		return nil, err
	}

	return labor, nil
}

func (r *WorkDayLaborRepository) Update(ctx context.Context, id int64, labor *models.WorkDayLabor) (*models.WorkDayLabor, error) {
	query := `
		UPDATE workDayLabor
		SET workerName = $1, jobTitle = $2, phone = $3, address = $4,
		    quantity = $5, cost = $6, notes = $7
		WHERE id = $8
		RETURNING id, workDayId, workerName, jobTitle, phone, address,
		          quantity, cost, notes, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		labor.WorkerName, labor.JobTitle, labor.Phone,
		labor.Address, labor.Quantity, labor.Cost, labor.Notes, id,
	).Scan(
		&labor.ID, &labor.WorkDayID, &labor.WorkerName, &labor.JobTitle, &labor.Phone,
		&labor.Address, &labor.Quantity, &labor.Cost, &labor.Notes, &labor.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return labor, nil
}

func (r *WorkDayLaborRepository) UpdateWithTx(ctx context.Context, tx pgx.Tx, id int64, labor *models.WorkDayLabor) (*models.WorkDayLabor, error) {
	query := `
		UPDATE workDayLabor
		SET workerName = $1, jobTitle = $2, phone = $3, address = $4,
		    quantity = $5, cost = $6, notes = $7
		WHERE id = $8
		RETURNING id, workDayId, workerName, jobTitle, phone, address,
		          quantity, cost, notes, createdAt
	`

	err := tx.QueryRow(ctx, query,
		labor.WorkerName, labor.JobTitle, labor.Phone,
		labor.Address, labor.Quantity, labor.Cost, labor.Notes, id,
	).Scan(
		&labor.ID, &labor.WorkDayID, &labor.WorkerName, &labor.JobTitle, &labor.Phone,
		&labor.Address, &labor.Quantity, &labor.Cost, &labor.Notes, &labor.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return labor, nil
}

func (r *WorkDayLaborRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM workDayLabor WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *WorkDayLaborRepository) DeleteWithTx(ctx context.Context, tx pgx.Tx, id int64) error {
	query := `DELETE FROM workDayLabor WHERE id = $1`
	result, err := tx.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
