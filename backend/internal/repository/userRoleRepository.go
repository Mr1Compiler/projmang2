package repository

import (
	"context"
	"encoding/json"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafa91ameen/prjalgo/backend/internal/models"
)

// UserPage represents a page with its permissions for a user
type UserPage struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Route       string   `json:"route"`
	Icon        *string  `json:"icon"`
	Permissions []string `json:"permissions"`
}

type UserRoleRepositoryInterface interface {
	GetAll(ctx context.Context, limit, offset int) ([]models.UserRole, int64, error)
	GetByID(ctx context.Context, id int64) (*models.UserRole, error)
	GetByUserID(ctx context.Context, userID int64) ([]models.UserRole, error)
	GetUserPages(ctx context.Context, userID int64) ([]UserPage, error)
	HasPermission(ctx context.Context, userID int64, route string, permission string) (bool, error)
	Create(ctx context.Context, userRole *models.UserRole) (*models.UserRole, error)
	Delete(ctx context.Context, id int64) error
}

type UserRoleRepository struct {
	db *pgxpool.Pool
}

func NewUserRoleRepository(db *pgxpool.Pool) *UserRoleRepository {
	return &UserRoleRepository{db: db}
}

func (r *UserRoleRepository) GetAll(ctx context.Context, limit, offset int) ([]models.UserRole, int64, error) {
	// Get total count
	var total int64
	countQuery := `SELECT COUNT(*) FROM userRoles`
	if err := r.db.QueryRow(ctx, countQuery).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT id, userId, roleId, createdAt
		FROM userRoles
		ORDER BY createdAt DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var userRoles []models.UserRole
	for rows.Next() {
		var ur models.UserRole
		err := rows.Scan(&ur.ID, &ur.UserID, &ur.RoleID, &ur.CreatedAt)
		if err != nil {
			return nil, 0, err
		}
		userRoles = append(userRoles, ur)
	}

	return userRoles, total, rows.Err()
}

func (r *UserRoleRepository) GetByID(ctx context.Context, id int64) (*models.UserRole, error) {
	query := `
		SELECT id, userId, roleId, createdAt
		FROM userRoles
		WHERE id = $1
	`

	var ur models.UserRole
	err := r.db.QueryRow(ctx, query, id).Scan(
		&ur.ID, &ur.UserID, &ur.RoleID, &ur.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &ur, nil
}

func (r *UserRoleRepository) GetByUserID(ctx context.Context, userID int64) ([]models.UserRole, error) {
	query := `
		SELECT id, userId, roleId, createdAt
		FROM userRoles
		WHERE userId = $1
		ORDER BY createdAt DESC
	`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userRoles []models.UserRole
	for rows.Next() {
		var ur models.UserRole
		err := rows.Scan(&ur.ID, &ur.UserID, &ur.RoleID, &ur.CreatedAt)
		if err != nil {
			return nil, err
		}
		userRoles = append(userRoles, ur)
	}

	return userRoles, rows.Err()
}

func (r *UserRoleRepository) Create(ctx context.Context, userRole *models.UserRole) (*models.UserRole, error) {
	query := `
		INSERT INTO userRoles (userId, roleId)
		VALUES ($1, $2)
		RETURNING id, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		userRole.UserID, userRole.RoleID,
	).Scan(&userRole.ID, &userRole.CreatedAt)

	if err != nil {
		return nil, err
	}

	return userRole, nil
}

func (r *UserRoleRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM userRoles WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

// GetUserPages returns all pages a user has access to with their permissions
// Merges permissions from all roles the user has for each page
func (r *UserRoleRepository) GetUserPages(ctx context.Context, userID int64) ([]UserPage, error) {
	query := `
		SELECT p.id, p.name, p.route, p.icon, rp.permissions
		FROM userRoles ur
		JOIN rolePages rp ON ur.roleId = rp.roleId
		JOIN pages p ON rp.pageId = p.id
		WHERE ur.userId = $1 AND (p.status = 'active' OR p.status IS NULL)
		ORDER BY p.id
	`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Use map to merge permissions from multiple roles
	pageMap := make(map[int64]*UserPage)
	permSet := make(map[int64]map[string]bool)
	var pageOrder []int64 // Track insertion order

	for rows.Next() {
		var pageID int64
		var name, route string
		var icon *string
		var permissionsJSON *string

		err := rows.Scan(&pageID, &name, &route, &icon, &permissionsJSON)
		if err != nil {
			return nil, err
		}

		// Initialize page if not seen before
		if _, exists := pageMap[pageID]; !exists {
			pageMap[pageID] = &UserPage{
				ID:    pageID,
				Name:  name,
				Route: route,
				Icon:  icon,
			}
			permSet[pageID] = make(map[string]bool)
			pageOrder = append(pageOrder, pageID) // Preserve order
		}

		// Parse and merge permissions
		if permissionsJSON != nil && *permissionsJSON != "" {
			var perms []string
			if err := json.Unmarshal([]byte(*permissionsJSON), &perms); err == nil {
				for _, p := range perms {
					permSet[pageID][p] = true
				}
			}
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Convert map to slice preserving order and assign merged permissions
	var pages []UserPage
	for _, pageID := range pageOrder {
		page := pageMap[pageID]
		for perm := range permSet[pageID] {
			page.Permissions = append(page.Permissions, perm)
		}
		pages = append(pages, *page)
	}

	return pages, nil
}

// HasPermission checks if a user has a specific permission for a route
// Checks ALL roles the user has - if ANY role grants the permission, access is allowed
func (r *UserRoleRepository) HasPermission(ctx context.Context, userID int64, route string, permission string) (bool, error) {
	query := `
		SELECT rp.permissions
		FROM userRoles ur
		JOIN rolePages rp ON ur.roleId = rp.roleId
		JOIN pages p ON rp.pageId = p.id
		WHERE ur.userId = $1 AND p.route = $2
	`

	rows, err := r.db.Query(ctx, query, userID, route)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	for rows.Next() {
		var permissionsJSON *string
		if err := rows.Scan(&permissionsJSON); err != nil {
			return false, err
		}

		if permissionsJSON == nil || *permissionsJSON == "" {
			continue
		}

		var permissions []string
		if err := json.Unmarshal([]byte(*permissionsJSON), &permissions); err != nil {
			continue
		}

		for _, p := range permissions {
			if p == permission {
				return true, nil
			}
		}
	}

	return false, rows.Err()
}
