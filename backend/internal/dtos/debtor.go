package dtos

import "time"

// ** Debtor DTOs **

// Response DTOs
type DebtorSummary struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Email     *string    `json:"email"`
	Phone     *string    `json:"phone"`
	TotalDebt float64    `json:"totalDebt"`
	Currency  string     `json:"currency"`
	DueDate   *time.Time `json:"dueDate"`
	Status    *string    `json:"status"`
}

type Debtor struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Email     *string    `json:"email"`
	Phone     *string    `json:"phone"`
	TotalDebt float64    `json:"totalDebt"`
	Currency  string     `json:"currency"`
	DueDate   *time.Time `json:"dueDate"`
	Status    *string    `json:"status"`
	Notes     *string    `json:"notes"`
	CreatedBy *int64     `json:"createdBy"`
	CreatedAt time.Time  `json:"createdAt"`
}

// Request DTOs
type CreateDebtor struct {
	Name      string     `json:"name" binding:"required"`
	Email     *string    `json:"email" binding:"omitempty,email"`
	Phone     *string    `json:"phone"`
	TotalDebt float64    `json:"totalDebt" binding:"required"`
	Currency  string     `json:"currency" binding:"required"`
	DueDate   *time.Time `json:"dueDate"`
	Status    *string    `json:"status"`
	Notes     *string    `json:"notes"`
	CreatedBy *int64     `json:"createdBy"`
}

type UpdateDebtor struct {
	Name      *string    `json:"name"`
	Email     *string    `json:"email" binding:"omitempty,email"`
	Phone     *string    `json:"phone"`
	TotalDebt *float64   `json:"totalDebt"`
	Currency  *string    `json:"currency"`
	DueDate   *time.Time `json:"dueDate"`
	Status    *string    `json:"status"`
	Notes     *string    `json:"notes"`
}

// Stats Response DTO
type DebtorStatsResponse struct {
	Total       int64   `json:"total"`
	Active      int64   `json:"active"`
	Paid        int64   `json:"paid"`
	TotalDebt   float64 `json:"totalDebt"`
	ActiveDebt  float64 `json:"activeDebt"`
	AverageDebt float64 `json:"averageDebt"`
}
