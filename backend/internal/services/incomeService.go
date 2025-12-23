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
	ErrIncomeNotFound      = errors.New("income not found")
	ErrIncomeStatsNotFound = errors.New("income stats not found")
)

type IncomeService struct {
	incomeRepo repository.IncomeRepositoryInterface
}

func NewIncomeService(incomeRepo repository.IncomeRepositoryInterface) *IncomeService {
	return &IncomeService{
		incomeRepo: incomeRepo,
	}
}

func (s *IncomeService) GetAll(ctx context.Context, limit, offset int) (dtos.PaginatedResponse[dtos.IncomeSummary], error) {
	incomes, total, err := s.incomeRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return dtos.PaginatedResponse[dtos.IncomeSummary]{}, err
	}

	result := make([]dtos.IncomeSummary, len(incomes))
	for i, inc := range incomes {
		result[i] = toIncomeSummaryDTO(inc)
	}

	page := (offset / limit) + 1
	return dtos.NewPaginatedResponse(result, total, page, limit), nil
}

func (s *IncomeService) GetByID(ctx context.Context, id int64) (*dtos.Income, error) {
	income, err := s.incomeRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrIncomeNotFound
		}
		return nil, err
	}
	dto := toIncomeDTO(income)
	return &dto, nil
}

func (s *IncomeService) Create(ctx context.Context, req dtos.CreateIncome) (*dtos.Income, error) {
	income := &models.Income{
		Name:       req.Name,
		Amount:     req.Amount,
		Type:       req.Type,
		IncomeDate: req.IncomeDate,
		Status:     req.Status,
		Notes:      req.Notes,
		CreatedBy:  req.CreatedBy,
	}

	created, err := s.incomeRepo.Create(ctx, income)
	if err != nil {
		return nil, err
	}

	dto := toIncomeDTO(created)
	return &dto, nil
}

func (s *IncomeService) Update(ctx context.Context, id int64, req dtos.UpdateIncome) (*dtos.Income, error) {
	existing, err := s.incomeRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrIncomeNotFound
		}
		return nil, err
	}

	// Apply partial updates
	if req.Name != nil {
		existing.Name = *req.Name
	}
	if req.Amount != nil {
		existing.Amount = *req.Amount
	}
	if req.Type != nil {
		existing.Type = req.Type
	}
	if req.IncomeDate != nil {
		existing.IncomeDate = *req.IncomeDate
	}
	if req.Status != nil {
		existing.Status = req.Status
	}
	if req.Notes != nil {
		existing.Notes = req.Notes
	}

	updated, err := s.incomeRepo.Update(ctx, id, existing)
	if err != nil {
		return nil, err
	}

	dto := toIncomeDTO(updated)
	return &dto, nil
}

func (s *IncomeService) Delete(ctx context.Context, id int64) error {
	err := s.incomeRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrIncomeNotFound
		}
		return err
	}
	return nil
}

func (s *IncomeService) GetStats(ctx context.Context, period string) (*dtos.IncomeStatsResponse, error) {
	stats, err := s.incomeRepo.GetStats(ctx, period)
	if err != nil {
		return nil, err
	}

	return &dtos.IncomeStatsResponse{
		Total:         stats.Total,
		TotalAmount:   stats.TotalAmount,
		Pending:       stats.Pending,
		Approved:      stats.Approved,
		Rejected:      stats.Rejected,
		AverageAmount: stats.AverageAmount,
	}, nil
}

// DTO conversion helpers
func toIncomeSummaryDTO(i models.Income) dtos.IncomeSummary {
	return dtos.IncomeSummary{
		ID:         i.ID,
		Name:       i.Name,
		Amount:     i.Amount,
		Type:       i.Type,
		IncomeDate: i.IncomeDate,
		Status:     i.Status,
	}
}

func toIncomeDTO(i *models.Income) dtos.Income {
	return dtos.Income{
		ID:         i.ID,
		Name:       i.Name,
		Amount:     i.Amount,
		Type:       i.Type,
		IncomeDate: i.IncomeDate,
		Status:     i.Status,
		Notes:      i.Notes,
		CreatedBy:  i.CreatedBy,
		CreatedAt:  i.CreatedAt,
	}
}
