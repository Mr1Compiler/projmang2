package models

import "time"

type WorkCategory struct {
	ID          int64     `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description *string   `json:"description" db:"description"`
	Status      *string   `json:"status" db:"status"`
	CreatedBy   *int64    `json:"createdBy" db:"createdby"`
	CreatedAt   time.Time `json:"createdAt" db:"createdat"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updatedat"`
}

type WorkSubCategory struct {
	ID          int64     `json:"id" db:"id"`
	CategoryID  int64     `json:"categoryId" db:"categoryid"`
	Name        string    `json:"name" db:"name"`
	Description *string   `json:"description" db:"description"`
	Percentage  *float64  `json:"percentage" db:"percentage"`
	Status      *string   `json:"status" db:"status"`
	CreatedBy   *int64    `json:"createdBy" db:"createdby"`
	CreatedAt   time.Time `json:"createdAt" db:"createdat"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updatedat"`
}
