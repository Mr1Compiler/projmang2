package dtos

import "time"

// ** Expense DTOs **

// Response DTOs
type ExpenseSummary struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Amount      float64   `json:"amount"`
	Type        *string   `json:"type"`
	ExpenseDate time.Time `json:"expenseDate"`
	ProjectID   *int64    `json:"projectId"`
	IsDebtor    bool      `json:"isDebtor"`
	DebtorID    *int64    `json:"debtorId"`
	Status      *string   `json:"status"`
}

type Expense struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Amount      float64   `json:"amount"`
	Type        *string   `json:"type"`
	ExpenseDate time.Time `json:"expenseDate"`
	ProjectID   *int64    `json:"projectId"`
	IsDebtor    bool      `json:"isDebtor"`
	DebtorID    *int64    `json:"debtorId"`
	Status      *string   `json:"status"`
	Notes       *string   `json:"notes"`
	CreatedBy   *int64    `json:"createdBy"`
	CreatedAt   time.Time `json:"createdAt"`
}

// Request DTOs
type CreateExpense struct {
	Name        string    `json:"name" binding:"required"`
	Amount      float64   `json:"amount" binding:"required"`
	Type        *string   `json:"type"`
	ExpenseDate time.Time `json:"expenseDate" binding:"required"`
	ProjectID   *int64    `json:"projectId"`
	IsDebtor    bool      `json:"isDebtor"`
	DebtorID    *int64    `json:"debtorId"`
	Status      *string   `json:"status"`
	Notes       *string   `json:"notes"`
	CreatedBy   *int64    `json:"createdBy"`
}

type UpdateExpense struct {
	Name        *string    `json:"name"`
	Amount      *float64   `json:"amount"`
	Type        *string    `json:"type"`
	ExpenseDate *time.Time `json:"expenseDate"`
	ProjectID   *int64     `json:"projectId"`
	IsDebtor    *bool      `json:"isDebtor"`
	DebtorID    *int64     `json:"debtorId"`
	Status      *string    `json:"status"`
	Notes       *string    `json:"notes"`
}

// ** Income DTOs **

// Response DTOs
type IncomeSummary struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Amount     float64   `json:"amount"`
	Type       *string   `json:"type"`
	IncomeDate time.Time `json:"incomeDate"`
	Status     *string   `json:"status"`
}

type Income struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Amount     float64   `json:"amount"`
	Type       *string   `json:"type"`
	IncomeDate time.Time `json:"incomeDate"`
	Status     *string   `json:"status"`
	Notes      *string   `json:"notes"`
	CreatedBy  *int64    `json:"createdBy"`
	CreatedAt  time.Time `json:"createdAt"`
}

// Request DTOs
type CreateIncome struct {
	Name       string    `json:"name" binding:"required"`
	Amount     float64   `json:"amount" binding:"required"`
	Type       *string   `json:"type"`
	IncomeDate time.Time `json:"incomeDate" binding:"required"`
	Status     *string   `json:"status"`
	Notes      *string   `json:"notes"`
	CreatedBy  *int64    `json:"createdBy"`
}

type UpdateIncome struct {
	Name       *string    `json:"name"`
	Amount     *float64   `json:"amount"`
	Type       *string    `json:"type"`
	IncomeDate *time.Time `json:"incomeDate"`
	Status     *string    `json:"status"`
	Notes      *string    `json:"notes"`
}

// ** Stats Response DTOs **

type ExpenseStatsResponse struct {
	Total         int64   `json:"total"`
	TotalAmount   float64 `json:"totalAmount"`
	Pending       int64   `json:"pending"`
	Approved      int64   `json:"approved"`
	Rejected      int64   `json:"rejected"`
	DebtorCount   int64   `json:"debtorCount"`
	AverageAmount float64 `json:"averageAmount"`
}

type IncomeStatsResponse struct {
	Total         int64   `json:"total"`
	TotalAmount   float64 `json:"totalAmount"`
	Pending       int64   `json:"pending"`
	Approved      int64   `json:"approved"`
	Rejected      int64   `json:"rejected"`
	AverageAmount float64 `json:"averageAmount"`
}
