package models

import "time"

type Debtor struct {
	ID        int64      `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Email     *string    `json:"email" db:"email"`
	Phone     *string    `json:"phone" db:"phone"`
	TotalDebt float64    `json:"totalDebt" db:"totaldebt"`
	Currency  string     `json:"currency" db:"currency"`
	DueDate   *time.Time `json:"dueDate" db:"duedate"`
	Status    *string    `json:"status" db:"status"`
	Notes     *string    `json:"notes" db:"notes"`
	CreatedBy *int64     `json:"createdBy" db:"createdby"`
	CreatedAt time.Time  `json:"createdAt" db:"createdat"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updatedat"`
}
