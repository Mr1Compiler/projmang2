package services

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/mustafa91ameen/prjalgo/backend/internal/dtos"
	"github.com/mustafa91ameen/prjalgo/backend/internal/models"
	"github.com/mustafa91ameen/prjalgo/backend/internal/repository"
)

var (
	ErrDebtorNotFound      = errors.New("debtor not found")
	ErrDebtorStatsNotFound = errors.New("debtor stats not found")
)

type DebtorService struct {
	debtorRepo repository.DebtorRepositoryInterface
}

func NewDebtorService(debtorRepo repository.DebtorRepositoryInterface) *DebtorService {
	return &DebtorService{
		debtorRepo: debtorRepo,
	}
}

func (s *DebtorService) GetAll(ctx context.Context, limit, offset int) (dtos.PaginatedResponse[dtos.DebtorSummary], error) {
	debtors, total, err := s.debtorRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return dtos.PaginatedResponse[dtos.DebtorSummary]{}, err
	}

	result := make([]dtos.DebtorSummary, len(debtors))
	for i, d := range debtors {
		result[i] = toDebtorSummaryDTO(d)
	}

	page := (offset / limit) + 1
	return dtos.NewPaginatedResponse(result, total, page, limit), nil
}

func (s *DebtorService) GetByID(ctx context.Context, id int64) (*dtos.Debtor, error) {
	debtor, err := s.debtorRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrDebtorNotFound
		}
		return nil, err
	}
	dto := toDebtorDTO(debtor)
	return &dto, nil
}

func (s *DebtorService) Create(ctx context.Context, req dtos.CreateDebtor) (*dtos.Debtor, error) {
	debtor := &models.Debtor{
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		TotalDebt: req.TotalDebt,
		Currency:  req.Currency,
		DueDate:   req.DueDate,
		Status:    req.Status,
		Notes:     req.Notes,
		CreatedBy: req.CreatedBy,
	}

	created, err := s.debtorRepo.Create(ctx, debtor)
	if err != nil {
		return nil, err
	}

	dto := toDebtorDTO(created)
	return &dto, nil
}

func (s *DebtorService) Update(ctx context.Context, id int64, req dtos.UpdateDebtor) (*dtos.Debtor, error) {
	existing, err := s.debtorRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrDebtorNotFound
		}
		return nil, err
	}

	// Apply partial updates
	if req.Name != nil {
		existing.Name = *req.Name
	}
	if req.Email != nil {
		existing.Email = req.Email
	}
	if req.Phone != nil {
		existing.Phone = req.Phone
	}
	if req.TotalDebt != nil {
		existing.TotalDebt = *req.TotalDebt
	}
	if req.Currency != nil {
		existing.Currency = *req.Currency
	}
	if req.DueDate != nil {
		existing.DueDate = req.DueDate
	}
	if req.Status != nil {
		existing.Status = req.Status
	}
	if req.Notes != nil {
		existing.Notes = req.Notes
	}

	updated, err := s.debtorRepo.Update(ctx, id, existing)
	if err != nil {
		return nil, err
	}

	dto := toDebtorDTO(updated)
	return &dto, nil
}

func (s *DebtorService) Delete(ctx context.Context, id int64) error {
	err := s.debtorRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrDebtorNotFound
		}
		return err
	}
	return nil
}

func (s *DebtorService) GetStats(ctx context.Context, period string) (*dtos.DebtorStatsResponse, error) {
	stats, err := s.debtorRepo.GetStats(ctx, period)
	if err != nil {
		return nil, err
	}

	return &dtos.DebtorStatsResponse{
		Total:       stats.Total,
		Active:      stats.Active,
		Paid:        stats.Paid,
		TotalDebt:   stats.TotalDebt,
		ActiveDebt:  stats.ActiveDebt,
		AverageDebt: stats.AverageDebt,
	}, nil
}

// DTO conversion helpers
func toDebtorSummaryDTO(d models.Debtor) dtos.DebtorSummary {
	return dtos.DebtorSummary{
		ID:        d.ID,
		Name:      d.Name,
		Email:     d.Email,
		Phone:     d.Phone,
		TotalDebt: d.TotalDebt,
		Currency:  d.Currency,
		DueDate:   d.DueDate,
		Status:    d.Status,
	}
}

func toDebtorDTO(d *models.Debtor) dtos.Debtor {
	return dtos.Debtor{
		ID:        d.ID,
		Name:      d.Name,
		Email:     d.Email,
		Phone:     d.Phone,
		TotalDebt: d.TotalDebt,
		Currency:  d.Currency,
		DueDate:   d.DueDate,
		Status:    d.Status,
		Notes:     d.Notes,
		CreatedBy: d.CreatedBy,
		CreatedAt: d.CreatedAt,
	}
}
