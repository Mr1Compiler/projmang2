package models

import "time"

type Expense struct {
	ID          int64     `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Amount      float64   `json:"amount" db:"amount"`
	Type        *string   `json:"type" db:"type"`
	ExpenseDate time.Time `json:"expenseDate" db:"expensedate"`
	ProjectID   *int64    `json:"projectId" db:"projectid"`
	IsDebtor    bool      `json:"isDebtor" db:"isdebtor"`
	DebtorID    *int64    `json:"debtorId" db:"debtorid"`
	Status      *string   `json:"status" db:"status"`
	Notes       *string   `json:"notes" db:"notes"`
	CreatedBy   *int64    `json:"createdBy" db:"createdby"`
	CreatedAt   time.Time `json:"createdAt" db:"createdat"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updatedat"`
}

type Income struct {
	ID         int64     `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Amount     float64   `json:"amount" db:"amount"`
	Type       *string   `json:"type" db:"type"`
	IncomeDate time.Time `json:"incomeDate" db:"incomedate"`
	Status     *string   `json:"status" db:"status"`
	Notes      *string   `json:"notes" db:"notes"`
	CreatedBy  *int64    `json:"createdBy" db:"createdby"`
	CreatedAt  time.Time `json:"createdAt" db:"createdat"`
	UpdatedAt  time.Time `json:"updatedAt" db:"updatedat"`
}
