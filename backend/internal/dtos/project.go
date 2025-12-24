package dtos

import "time"

// Response DTOs
type ProjectSummary struct {
	ID                 int64      `json:"id"`
	Name               string     `json:"name"`
	Type               *string    `json:"type"`
	Description        *string    `json:"description"`
	ClientPhone        *string    `json:"clientPhone"`
	Location           *string    `json:"location"`
	StartDate          *time.Time `json:"startDate"`
	Duration           *int       `json:"duration"`
	Status             string     `json:"status"`
	ProgressPercentage float64    `json:"progressPercentage"`
	WarningCost        float64    `json:"warningCost"`
	TotalCost          float64    `json:"totalCost"`
	Notes              *string    `json:"notes"`
	CurrentSpending    float64    `json:"currentSpending"`
}

type Project struct {
	ID                 int64      `json:"id"`
	Name               string     `json:"name"`
	Type               *string    `json:"type"`
	Description        *string    `json:"description"`
	ClientPhone        *string    `json:"clientPhone"`
	Location           *string    `json:"location"`
	StartDate          *time.Time `json:"startDate"`
	Duration           *int       `json:"duration"`
	WarningCost        float64    `json:"warningCost"`
	TotalCost          float64    `json:"totalCost"`
	Status             string     `json:"status"`
	ProgressPercentage float64    `json:"progressPercentage"`
	Notes              *string    `json:"notes"`
	CreatedBy          *int64     `json:"createdBy"`
	CreatedAt          time.Time  `json:"createdAt"`
	CurrentSpending    float64    `json:"currentSpending"`
}

// Request DTOs
type CreateProject struct {
	Name        string     `json:"name" binding:"required"`
	Type        *string    `json:"type"`
	Description *string    `json:"description"`
	ClientPhone *string    `json:"clientPhone"`
	Location    *string    `json:"location"`
	StartDate   *time.Time `json:"startDate"`
	Duration    *int       `json:"duration"`
	WarningCost float64    `json:"warningCost" binding:"required"`
	TotalCost   float64    `json:"totalCost" binding:"required"`
	Notes       *string    `json:"notes"`
	CreatedBy   *int64     `json:"createdBy"`
}

type UpdateProject struct {
	Name               *string    `json:"name"`
	Type               *string    `json:"type"`
	Description        *string    `json:"description"`
	ClientPhone        *string    `json:"clientPhone"`
	Location           *string    `json:"location"`
	StartDate          *time.Time `json:"startDate"`
	Duration           *int       `json:"duration"`
	WarningCost        *float64   `json:"warningCost"`
	TotalCost          *float64   `json:"totalCost"`
	Status             *string    `json:"status"`
	ProgressPercentage *float64   `json:"progressPercentage"`
	Notes              *string    `json:"notes"`
}

// Stats Response DTO
type ProjectStatsResponse struct {
	Total           int64   `json:"total"`
	Pending         int64   `json:"pending"`
	InProgress      int64   `json:"inProgress"`
	Completed       int64   `json:"completed"`
	TotalBudget     float64 `json:"totalBudget"`
	AverageProgress float64 `json:"averageProgress"`
}
