package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafa91ameen/prjalgo/backend/internal/models"
)

type UserCredentials struct {
	ID       int64
	Username string
	Password string
	Status   *string
}

type UserRepositoryInterface interface {
	GetAll(ctx context.Context, limit, offset int) ([]models.User, int64, error)
	GetByID(ctx context.Context, id int64) (*models.User, error)
	GetCredentialsByUsername(ctx context.Context, username string) (*UserCredentials, error)
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Update(ctx context.Context, id int64, user *models.User) (*models.User, error)
	UpdatePassword(ctx context.Context, id int64, hashedPassword string) error
	UpdateStatus(ctx context.Context, id int64, status string) error
	UpdateStatusWithTx(ctx context.Context, tx pgx.Tx, id int64, status string) error
	Delete(ctx context.Context, id int64) error
}

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll(ctx context.Context, limit, offset int) ([]models.User, int64, error) {
	// Get total count
	var total int64
	countQuery := `SELECT COUNT(*) FROM users`
	if err := r.db.QueryRow(ctx, countQuery).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT id, username, email, fullName, phone, avatar,
		       jobTitle, status, lastLogin
		FROM users
		ORDER BY createdAt DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(
			&u.ID, &u.Username, &u.Email, &u.FullName,
			&u.Phone, &u.Avatar, &u.JobTitle, &u.Status, &u.LastLogin,
		)
		if err != nil {
			return nil, 0, err
		}
		users = append(users, u)
	}

	return users, total, rows.Err()
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	query := `
		INSERT INTO users (username, email, password, fullName, phone, avatar, jobTitle, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, createdAt, updatedAt
	`

	err := r.db.QueryRow(ctx, query,
		user.Username, user.Email, user.Password, user.FullName,
		user.Phone, user.Avatar, user.JobTitle, user.Status,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, id int64, user *models.User) (*models.User, error) {
	query := `
		UPDATE users
		SET username = $1, email = $2, fullName = $3, phone = $4,
		    avatar = $5, jobTitle = $6, status = $7
		WHERE id = $8
		RETURNING id, username, email, fullName, phone, avatar,
		          jobTitle, status, lastLogin, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		user.Username, user.Email, user.FullName, user.Phone,
		user.Avatar, user.JobTitle, user.Status, id,
	).Scan(
		&user.ID, &user.Username, &user.Email, &user.FullName,
		&user.Phone, &user.Avatar, &user.JobTitle, &user.Status,
		&user.LastLogin, &user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	query := `
		SELECT id, username, email, fullName, phone, avatar,
		       jobTitle, status, lastLogin, createdAt
		FROM users
		WHERE id = $1
	`

	var user models.User
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID, &user.Username, &user.Email, &user.FullName,
		&user.Phone, &user.Avatar, &user.JobTitle, &user.Status,
		&user.LastLogin, &user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetCredentialsByUsername(ctx context.Context, username string) (*UserCredentials, error) {
	query := `SELECT id, username, password, status FROM users WHERE username = $1`

	var creds UserCredentials
	err := r.db.QueryRow(ctx, query, username).Scan(&creds.ID, &creds.Username, &creds.Password, &creds.Status)
	if err != nil {
		return nil, err
	}

	return &creds, nil
}

func (r *UserRepository) UpdatePassword(ctx context.Context, id int64, hashedPassword string) error {
	query := `UPDATE users SET password = $1 WHERE id = $2`
	result, err := r.db.Exec(ctx, query, hashedPassword, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *UserRepository) UpdateStatus(ctx context.Context, id int64, status string) error {
	query := `UPDATE users SET status = $1 WHERE id = $2`
	result, err := r.db.Exec(ctx, query, status, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *UserRepository) UpdateStatusWithTx(ctx context.Context, tx pgx.Tx, id int64, status string) error {
	query := `UPDATE users SET status = $1 WHERE id = $2`
	result, err := tx.Exec(ctx, query, status, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *UserRepository) UpdateLastLogin(ctx context.Context, userID int64) error {
	query := `UPDATE users SET lastlogin = NOW() AT TIME ZONE 'UTC' WHERE id = $1`
	result, err := r.db.Exec(ctx, query, userID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
