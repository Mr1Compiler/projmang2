package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafa91ameen/prjalgo/backend/internal/models"
)

type RolePageRepositoryInterface interface {
	GetAll(ctx context.Context, limit, offset int) ([]models.RolePage, int64, error)
	GetByID(ctx context.Context, id int64) (*models.RolePage, error)
	GetByRoleID(ctx context.Context, roleID int64) ([]models.RolePage, error)
	Create(ctx context.Context, rolePage *models.RolePage) (*models.RolePage, error)
	Update(ctx context.Context, id int64, rolePage *models.RolePage) (*models.RolePage, error)
	Delete(ctx context.Context, id int64) error
	DeleteAllByRoleIDWithTx(ctx context.Context, tx pgx.Tx, roleID int64) error
	CreateWithTx(ctx context.Context, tx pgx.Tx, rolePage *models.RolePage) (*models.RolePage, error)
}

type RolePageRepository struct {
	db *pgxpool.Pool
}

func NewRolePageRepository(db *pgxpool.Pool) *RolePageRepository {
	return &RolePageRepository{db: db}
}

func (r *RolePageRepository) GetAll(ctx context.Context, limit, offset int) ([]models.RolePage, int64, error) {
	// Get total count
	var total int64
	countQuery := `SELECT COUNT(*) FROM rolePages`
	if err := r.db.QueryRow(ctx, countQuery).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT id, roleId, pageId, permissions, createdAt
		FROM rolePages
		ORDER BY createdAt DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var rolePages []models.RolePage
	for rows.Next() {
		var rp models.RolePage
		err := rows.Scan(&rp.ID, &rp.RoleID, &rp.PageID, &rp.Permissions, &rp.CreatedAt)
		if err != nil {
			return nil, 0, err
		}
		rolePages = append(rolePages, rp)
	}

	return rolePages, total, rows.Err()
}

func (r *RolePageRepository) GetByID(ctx context.Context, id int64) (*models.RolePage, error) {
	query := `
		SELECT id, roleId, pageId, permissions, createdAt
		FROM rolePages
		WHERE id = $1
	`

	var rp models.RolePage
	err := r.db.QueryRow(ctx, query, id).Scan(
		&rp.ID, &rp.RoleID, &rp.PageID, &rp.Permissions, &rp.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &rp, nil
}

func (r *RolePageRepository) GetByRoleID(ctx context.Context, roleID int64) ([]models.RolePage, error) {
	query := `
		SELECT id, roleId, pageId, permissions, createdAt
		FROM rolePages
		WHERE roleId = $1
		ORDER BY createdAt DESC
	`

	rows, err := r.db.Query(ctx, query, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rolePages []models.RolePage
	for rows.Next() {
		var rp models.RolePage
		err := rows.Scan(&rp.ID, &rp.RoleID, &rp.PageID, &rp.Permissions, &rp.CreatedAt)
		if err != nil {
			return nil, err
		}
		rolePages = append(rolePages, rp)
	}

	return rolePages, rows.Err()
}

func (r *RolePageRepository) Create(ctx context.Context, rolePage *models.RolePage) (*models.RolePage, error) {
	query := `
		INSERT INTO rolePages (roleId, pageId, permissions)
		VALUES ($1, $2, $3)
		RETURNING id, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		rolePage.RoleID, rolePage.PageID, rolePage.Permissions,
	).Scan(&rolePage.ID, &rolePage.CreatedAt)

	if err != nil {
		return nil, err
	}

	return rolePage, nil
}

func (r *RolePageRepository) Update(ctx context.Context, id int64, rolePage *models.RolePage) (*models.RolePage, error) {
	query := `
		UPDATE rolePages
		SET permissions = $1
		WHERE id = $2
		RETURNING id, roleId, pageId, permissions, createdAt
	`

	err := r.db.QueryRow(ctx, query, rolePage.Permissions, id).Scan(
		&rolePage.ID, &rolePage.RoleID, &rolePage.PageID,
		&rolePage.Permissions, &rolePage.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return rolePage, nil
}

func (r *RolePageRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM rolePages WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *RolePageRepository) DeleteAllByRoleIDWithTx(ctx context.Context, tx pgx.Tx, roleID int64) error {
	query := `DELETE FROM rolePages WHERE roleId = $1`
	_, err := tx.Exec(ctx, query, roleID)
	return err
}

func (r *RolePageRepository) CreateWithTx(ctx context.Context, tx pgx.Tx, rolePage *models.RolePage) (*models.RolePage, error) {
	query := `
		INSERT INTO rolePages (roleId, pageId, permissions)
		VALUES ($1, $2, $3)
		RETURNING id, createdAt
	`

	err := tx.QueryRow(ctx, query,
		rolePage.RoleID, rolePage.PageID, rolePage.Permissions,
	).Scan(&rolePage.ID, &rolePage.CreatedAt)

	if err != nil {
		return nil, err
	}

	return rolePage, nil
}
