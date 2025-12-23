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
	ErrExpenseNotFound      = errors.New("expense not found")
	ErrExpenseStatsNotFound = errors.New("expense stats not found")
)

type ExpenseService struct {
	expenseRepo repository.ExpenseRepositoryInterface
}

func NewExpenseService(expenseRepo repository.ExpenseRepositoryInterface) *ExpenseService {
	return &ExpenseService{
		expenseRepo: expenseRepo,
	}
}

func (s *ExpenseService) GetAll(ctx context.Context, limit, offset int) (dtos.PaginatedResponse[dtos.ExpenseSummary], error) {
	expenses, total, err := s.expenseRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return dtos.PaginatedResponse[dtos.ExpenseSummary]{}, err
	}

	result := make([]dtos.ExpenseSummary, len(expenses))
	for i, e := range expenses {
		result[i] = toExpenseSummaryDTO(e)
	}

	page := (offset / limit) + 1
	return dtos.NewPaginatedResponse(result, total, page, limit), nil
}

func (s *ExpenseService) GetByID(ctx context.Context, id int64) (*dtos.Expense, error) {
	expense, err := s.expenseRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrExpenseNotFound
		}
		return nil, err
	}
	dto := toExpenseDTO(expense)
	return &dto, nil
}

func (s *ExpenseService) GetByProjectID(ctx context.Context, projectID int64) ([]dtos.ExpenseSummary, error) {
	expenses, err := s.expenseRepo.GetByProjectID(ctx, projectID)
	if err != nil {
		return nil, err
	}

	result := make([]dtos.ExpenseSummary, len(expenses))
	for i, e := range expenses {
		result[i] = toExpenseSummaryDTO(e)
	}
	return result, nil
}

func (s *ExpenseService) Create(ctx context.Context, req dtos.CreateExpense) (*dtos.Expense, error) {
	expense := &models.Expense{
		Name:        req.Name,
		Amount:      req.Amount,
		Type:        req.Type,
		ExpenseDate: req.ExpenseDate,
		ProjectID:   req.ProjectID,
		IsDebtor:    req.IsDebtor,
		DebtorID:    req.DebtorID,
		Status:      req.Status,
		Notes:       req.Notes,
		CreatedBy:   req.CreatedBy,
	}

	created, err := s.expenseRepo.Create(ctx, expense)
	if err != nil {
		return nil, err
	}

	dto := toExpenseDTO(created)
	return &dto, nil
}

func (s *ExpenseService) Update(ctx context.Context, id int64, req dtos.UpdateExpense) (*dtos.Expense, error) {
	existing, err := s.expenseRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrExpenseNotFound
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
	if req.ExpenseDate != nil {
		existing.ExpenseDate = *req.ExpenseDate
	}
	if req.ProjectID != nil {
		existing.ProjectID = req.ProjectID
	}
	if req.IsDebtor != nil {
		existing.IsDebtor = *req.IsDebtor
	}
	if req.DebtorID != nil {
		existing.DebtorID = req.DebtorID
	}
	if req.Status != nil {
		existing.Status = req.Status
	}
	if req.Notes != nil {
		existing.Notes = req.Notes
	}

	updated, err := s.expenseRepo.Update(ctx, id, existing)
	if err != nil {
		return nil, err
	}

	dto := toExpenseDTO(updated)
	return &dto, nil
}

func (s *ExpenseService) Delete(ctx context.Context, id int64) error {
	err := s.expenseRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrExpenseNotFound
		}
		return err
	}
	return nil
}

func (s *ExpenseService) GetStats(ctx context.Context, period string) (*dtos.ExpenseStatsResponse, error) {
	stats, err := s.expenseRepo.GetStats(ctx, period)
	if err != nil {
		return nil, err
	}

	return &dtos.ExpenseStatsResponse{
		Total:         stats.Total,
		TotalAmount:   stats.TotalAmount,
		Pending:       stats.Pending,
		Approved:      stats.Approved,
		Rejected:      stats.Rejected,
		DebtorCount:   stats.DebtorCount,
		AverageAmount: stats.AverageAmount,
	}, nil
}

// DTO conversion helpers
func toExpenseSummaryDTO(e models.Expense) dtos.ExpenseSummary {
	return dtos.ExpenseSummary{
		ID:          e.ID,
		Name:        e.Name,
		Amount:      e.Amount,
		Type:        e.Type,
		ExpenseDate: e.ExpenseDate,
		ProjectID:   e.ProjectID,
		IsDebtor:    e.IsDebtor,
		DebtorID:    e.DebtorID,
		Status:      e.Status,
	}
}

func toExpenseDTO(e *models.Expense) dtos.Expense {
	return dtos.Expense{
		ID:          e.ID,
		Name:        e.Name,
		Amount:      e.Amount,
		Type:        e.Type,
		ExpenseDate: e.ExpenseDate,
		ProjectID:   e.ProjectID,
		IsDebtor:    e.IsDebtor,
		DebtorID:    e.DebtorID,
		Status:      e.Status,
		Notes:       e.Notes,
		CreatedBy:   e.CreatedBy,
		CreatedAt:   e.CreatedAt,
	}
}
