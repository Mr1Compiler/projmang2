package dtos

import "time"

// ** WorkCategory DTOs **

// Response DTOs
type WorkCategorySummary struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
}

type WorkCategory struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Status      *string   `json:"status"`
	CreatedBy   *int64    `json:"createdBy"`
	CreatedAt   time.Time `json:"createdAt"`
}

// Request DTOs
type CreateWorkCategory struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
	CreatedBy   *int64  `json:"createdBy"`
}

type UpdateWorkCategory struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
}

// ** WorkSubCategory DTOs **

// Response DTOs
type WorkSubCategorySummary struct {
	ID          int64    `json:"id"`
	CategoryID  int64    `json:"categoryId"`
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	Percentage  *float64 `json:"percentage"`
	Status      *string  `json:"status"`
}

type WorkSubCategory struct {
	ID          int64     `json:"id"`
	CategoryID  int64     `json:"categoryId"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Percentage  *float64  `json:"percentage"`
	Status      *string   `json:"status"`
	CreatedBy   *int64    `json:"createdBy"`
	CreatedAt   time.Time `json:"createdAt"`
}

// Request DTOs
type CreateWorkSubCategory struct {
	CategoryID  int64    `json:"categoryId" binding:"required"`
	Name        string   `json:"name" binding:"required"`
	Description *string  `json:"description"`
	Percentage  *float64 `json:"percentage"`
	Status      *string  `json:"status"`
	CreatedBy   *int64   `json:"createdBy"`
}

type UpdateWorkSubCategory struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Percentage  *float64 `json:"percentage"`
	Status      *string  `json:"status"`
}

// Stats Response DTO
type WorkCategoryStatsResponse struct {
	Total            int64 `json:"total"`
	Active           int64 `json:"active"`
	Inactive         int64 `json:"inactive"`
	TotalSubcategory int64 `json:"totalSubcategory"`
}
