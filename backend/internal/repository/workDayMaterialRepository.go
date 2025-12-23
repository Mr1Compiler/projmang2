package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafa91ameen/prjalgo/backend/internal/models"
)

type WorkDayMaterialRepositoryInterface interface {
	GetAll(ctx context.Context, limit, offset int) ([]models.WorkDayMaterial, int64, error)
	GetByID(ctx context.Context, id int64) (*models.WorkDayMaterial, error)
	GetByWorkDayID(ctx context.Context, workDayID int64) ([]models.WorkDayMaterial, error)
	Create(ctx context.Context, material *models.WorkDayMaterial) (*models.WorkDayMaterial, error)
	CreateWithTx(ctx context.Context, tx pgx.Tx, material *models.WorkDayMaterial) (*models.WorkDayMaterial, error)
	Update(ctx context.Context, id int64, material *models.WorkDayMaterial) (*models.WorkDayMaterial, error)
	UpdateWithTx(ctx context.Context, tx pgx.Tx, id int64, material *models.WorkDayMaterial) (*models.WorkDayMaterial, error)
	Delete(ctx context.Context, id int64) error
	DeleteWithTx(ctx context.Context, tx pgx.Tx, id int64) error
}

type WorkDayMaterialRepository struct {
	db *pgxpool.Pool
}

func NewWorkDayMaterialRepository(db *pgxpool.Pool) *WorkDayMaterialRepository {
	return &WorkDayMaterialRepository{db: db}
}

func (r *WorkDayMaterialRepository) GetAll(ctx context.Context, limit, offset int) ([]models.WorkDayMaterial, int64, error) {
	// Get total count
	var total int64
	countQuery := `SELECT COUNT(*) FROM workDayMaterials`
	if err := r.db.QueryRow(ctx, countQuery).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT id, workDayId, materialName, quantity, cost, notes, createdAt
		FROM workDayMaterials
		ORDER BY createdAt DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var materials []models.WorkDayMaterial
	for rows.Next() {
		var m models.WorkDayMaterial
		err := rows.Scan(
			&m.ID, &m.WorkDayID, &m.MaterialName, &m.Quantity,
			&m.Cost, &m.Notes, &m.CreatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		materials = append(materials, m)
	}

	return materials, total, rows.Err()
}

func (r *WorkDayMaterialRepository) GetByID(ctx context.Context, id int64) (*models.WorkDayMaterial, error) {
	query := `
		SELECT id, workDayId, materialName, quantity, cost, notes, createdAt
		FROM workDayMaterials
		WHERE id = $1
	`

	var m models.WorkDayMaterial
	err := r.db.QueryRow(ctx, query, id).Scan(
		&m.ID, &m.WorkDayID, &m.MaterialName, &m.Quantity,
		&m.Cost, &m.Notes, &m.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (r *WorkDayMaterialRepository) GetByWorkDayID(ctx context.Context, workDayID int64) ([]models.WorkDayMaterial, error) {
	query := `
		SELECT id, workDayId, materialName, quantity, cost, notes, createdAt
		FROM workDayMaterials
		WHERE workDayId = $1
		ORDER BY createdAt DESC
	`

	rows, err := r.db.Query(ctx, query, workDayID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var materials []models.WorkDayMaterial
	for rows.Next() {
		var m models.WorkDayMaterial
		err := rows.Scan(
			&m.ID, &m.WorkDayID, &m.MaterialName, &m.Quantity,
			&m.Cost, &m.Notes, &m.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		materials = append(materials, m)
	}

	return materials, rows.Err()
}

func (r *WorkDayMaterialRepository) Create(ctx context.Context, material *models.WorkDayMaterial) (*models.WorkDayMaterial, error) {
	query := `
		INSERT INTO workDayMaterials (workDayId, materialName, quantity, cost, notes)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		material.WorkDayID, material.MaterialName, material.Quantity,
		material.Cost, material.Notes,
	).Scan(&material.ID, &material.CreatedAt)

	if err != nil {
		return nil, err
	}

	return material, nil
}

func (r *WorkDayMaterialRepository) CreateWithTx(ctx context.Context, tx pgx.Tx, material *models.WorkDayMaterial) (*models.WorkDayMaterial, error) {
	query := `
		INSERT INTO workDayMaterials (workDayId, materialName, quantity, cost, notes)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, createdAt
	`

	err := tx.QueryRow(ctx, query,
		material.WorkDayID, material.MaterialName, material.Quantity,
		material.Cost, material.Notes,
	).Scan(&material.ID, &material.CreatedAt)

	if err != nil {
		return nil, err
	}

	return material, nil
}

func (r *WorkDayMaterialRepository) Update(ctx context.Context, id int64, material *models.WorkDayMaterial) (*models.WorkDayMaterial, error) {
	query := `
		UPDATE workDayMaterials
		SET materialName = $1, quantity = $2, cost = $3, notes = $4
		WHERE id = $5
		RETURNING id, workDayId, materialName, quantity, cost, notes, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		material.MaterialName, material.Quantity,
		material.Cost, material.Notes, id,
	).Scan(
		&material.ID, &material.WorkDayID, &material.MaterialName, &material.Quantity,
		&material.Cost, &material.Notes, &material.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return material, nil
}

func (r *WorkDayMaterialRepository) UpdateWithTx(ctx context.Context, tx pgx.Tx, id int64, material *models.WorkDayMaterial) (*models.WorkDayMaterial, error) {
	query := `
		UPDATE workDayMaterials
		SET materialName = $1, quantity = $2, cost = $3, notes = $4
		WHERE id = $5
		RETURNING id, workDayId, materialName, quantity, cost, notes, createdAt
	`

	err := tx.QueryRow(ctx, query,
		material.MaterialName, material.Quantity,
		material.Cost, material.Notes, id,
	).Scan(
		&material.ID, &material.WorkDayID, &material.MaterialName, &material.Quantity,
		&material.Cost, &material.Notes, &material.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return material, nil
}

func (r *WorkDayMaterialRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM workDayMaterials WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *WorkDayMaterialRepository) DeleteWithTx(ctx context.Context, tx pgx.Tx, id int64) error {
	query := `DELETE FROM workDayMaterials WHERE id = $1`
	result, err := tx.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
