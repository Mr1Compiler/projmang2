package services

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafa91ameen/prjalgo/backend/internal/dtos"
	"github.com/mustafa91ameen/prjalgo/backend/internal/models"
	"github.com/mustafa91ameen/prjalgo/backend/internal/repository"
)

var (
	ErrRolePageNotFound = errors.New("role page not found")
)

type RolePageService struct {
	db           *pgxpool.Pool
	rolePageRepo repository.RolePageRepositoryInterface
}

func NewRolePageService(db *pgxpool.Pool, rolePageRepo repository.RolePageRepositoryInterface) *RolePageService {
	return &RolePageService{
		db:           db,
		rolePageRepo: rolePageRepo,
	}
}

func (s *RolePageService) GetAll(ctx context.Context, limit, offset int) (dtos.PaginatedResponse[dtos.RolePage], error) {
	rolePages, total, err := s.rolePageRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return dtos.PaginatedResponse[dtos.RolePage]{}, err
	}

	result := make([]dtos.RolePage, len(rolePages))
	for i, rp := range rolePages {
		result[i] = toRolePageDTO(rp)
	}

	page := (offset / limit) + 1
	return dtos.NewPaginatedResponse(result, total, page, limit), nil
}

func (s *RolePageService) GetByID(ctx context.Context, id int64) (*dtos.RolePage, error) {
	rolePage, err := s.rolePageRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrRolePageNotFound
		}
		return nil, err
	}
	dto := toRolePageDTO(*rolePage)
	return &dto, nil
}

func (s *RolePageService) GetByRoleID(ctx context.Context, roleID int64) ([]dtos.RolePage, error) {
	rolePages, err := s.rolePageRepo.GetByRoleID(ctx, roleID)
	if err != nil {
		return nil, err
	}

	result := make([]dtos.RolePage, len(rolePages))
	for i, rp := range rolePages {
		result[i] = toRolePageDTO(rp)
	}
	return result, nil
}

func (s *RolePageService) Create(ctx context.Context, req dtos.CreateRolePage) (*dtos.RolePage, error) {
	rolePage := &models.RolePage{
		RoleID:      req.RoleID,
		PageID:      req.PageID,
		Permissions: req.Permissions,
	}

	created, err := s.rolePageRepo.Create(ctx, rolePage)
	if err != nil {
		return nil, err
	}

	dto := toRolePageDTO(*created)
	return &dto, nil
}

func (s *RolePageService) Update(ctx context.Context, id int64, req dtos.UpdateRolePage) (*dtos.RolePage, error) {
	existing, err := s.rolePageRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrRolePageNotFound
		}
		return nil, err
	}

	// Apply partial updates (only permissions can be updated)
	if req.Permissions != nil {
		existing.Permissions = req.Permissions
	}

	updated, err := s.rolePageRepo.Update(ctx, id, existing)
	if err != nil {
		return nil, err
	}

	dto := toRolePageDTO(*updated)
	return &dto, nil
}

func (s *RolePageService) Delete(ctx context.Context, id int64) error {
	err := s.rolePageRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrRolePageNotFound
		}
		return err
	}
	return nil
}

func (s *RolePageService) UpdateAllByRoleID(ctx context.Context, roleID int64, req dtos.UpdateRolePagesRequest) ([]dtos.RolePage, error) {
	// Start transaction
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// Delete all existing role pages for this role
	err = s.rolePageRepo.DeleteAllByRoleIDWithTx(ctx, tx, roleID)
	if err != nil {
		return nil, err
	}

	// Create new role pages
	var result []dtos.RolePage
	for _, item := range req.Pages {
		rolePage := &models.RolePage{
			RoleID:      roleID,
			PageID:      item.PageID,
			Permissions: &item.Permissions,
		}

		created, err := s.rolePageRepo.CreateWithTx(ctx, tx, rolePage)
		if err != nil {
			return nil, err
		}

		result = append(result, toRolePageDTO(*created))
	}

	// Commit transaction
	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return result, nil
}

// DTO conversion helper
func toRolePageDTO(rp models.RolePage) dtos.RolePage {
	return dtos.RolePage{
		ID:          rp.ID,
		RoleID:      rp.RoleID,
		PageID:      rp.PageID,
		Permissions: rp.Permissions,
		CreatedAt:   rp.CreatedAt,
	}
}
