package dtos

import "time"

// ** WorkDay DTOs **

// Response DTOs
type WorkDaySummary struct {
	ID                int64     `json:"id"`
	ProjectID         int64     `json:"projectId"`
	WorkSubCategoryID *int64    `json:"workSubCategoryId"`
	WorkDate          time.Time `json:"workDate"`
	Description       *string   `json:"description"`
	Notes             *string   `json:"notes"`
	Status            string    `json:"status"`
	TotalCost         float64   `json:"totalCost"`
}

type WorkDay struct {
	ID                int64     `json:"id"`
	ProjectID         int64     `json:"projectId"`
	WorkSubCategoryID *int64    `json:"workSubCategoryId"`
	WorkDate          time.Time `json:"workDate"`
	Description       *string   `json:"description"`
	Status            string    `json:"status"`
	TotalCost         float64   `json:"totalCost"`
	Notes             *string   `json:"notes"`
	CreatedBy         *int64    `json:"createdBy"`
	CreatedAt         time.Time `json:"createdAt"`
}

// Request DTOs
type CreateWorkDay struct {
	ProjectID         int64     `json:"projectId" binding:"required"`
	WorkSubCategoryID *int64    `json:"workSubCategoryId"`
	WorkDate          time.Time `json:"workDate" binding:"required"`
	Description       *string   `json:"description"`
	Notes             *string   `json:"notes"`
	CreatedBy         *int64    `json:"createdBy"`
}

type UpdateWorkDay struct {
	WorkSubCategoryID *int64     `json:"workSubCategoryId"`
	WorkDate          *time.Time `json:"workDate"`
	Description       *string    `json:"description"`
	Status            *string    `json:"status"`
	TotalCost         *float64   `json:"totalCost"`
	Notes             *string    `json:"notes"`
}

// ** WorkDayMaterial DTOs **

// Response DTOs
type WorkDayMaterial struct {
	ID           int64     `json:"id"`
	WorkDayID    int64     `json:"workDayId"`
	MaterialName string    `json:"materialName"`
	Quantity     float64   `json:"quantity"`
	Cost         float64   `json:"cost"`
	Notes        *string   `json:"notes"`
	CreatedAt    time.Time `json:"createdAt"`
}

// Request DTOs
type CreateWorkDayMaterial struct {
	WorkDayID    int64   `json:"workDayId" binding:"required"`
	MaterialName string  `json:"materialName" binding:"required"`
	Quantity     float64 `json:"quantity" binding:"required"`
	Cost         float64 `json:"cost" binding:"required"`
	Notes        *string `json:"notes"`
}

type UpdateWorkDayMaterial struct {
	MaterialName *string  `json:"materialName"`
	Quantity     *float64 `json:"quantity"`
	Cost         *float64 `json:"cost"`
	Notes        *string  `json:"notes"`
}

// ** WorkDayLabor DTOs **

// Response DTOs
type WorkDayLabor struct {
	ID         int64     `json:"id"`
	WorkDayID  int64     `json:"workDayId"`
	WorkerName string    `json:"workerName"`
	JobTitle   *string   `json:"jobTitle"`
	Phone      *string   `json:"phone"`
	Address    *string   `json:"address"`
	Quantity   float64   `json:"quantity"`
	Cost       float64   `json:"cost"`
	Notes      *string   `json:"notes"`
	CreatedAt  time.Time `json:"createdAt"`
}

// Request DTOs
type CreateWorkDayLabor struct {
	WorkDayID  int64   `json:"workDayId" binding:"required"`
	WorkerName string  `json:"workerName" binding:"required"`
	JobTitle   *string `json:"jobTitle"`
	Phone      *string `json:"phone"`
	Address    *string `json:"address"`
	Quantity   float64 `json:"quantity" binding:"required"`
	Cost       float64 `json:"cost" binding:"required"`
	Notes      *string `json:"notes"`
}

type UpdateWorkDayLabor struct {
	WorkerName *string  `json:"workerName"`
	JobTitle   *string  `json:"jobTitle"`
	Phone      *string  `json:"phone"`
	Address    *string  `json:"address"`
	Quantity   *float64 `json:"quantity"`
	Cost       *float64 `json:"cost"`
	Notes      *string  `json:"notes"`
}

// ** WorkDayEquipment DTOs **

// Response DTOs
type WorkDayEquipment struct {
	ID            int64     `json:"id"`
	WorkDayID     int64     `json:"workDayId"`
	EquipmentName string    `json:"equipmentName"`
	Quantity      float64   `json:"quantity"`
	Cost          float64   `json:"cost"`
	Notes         *string   `json:"notes"`
	CreatedAt     time.Time `json:"createdAt"`
}

// Request DTOs
type CreateWorkDayEquipment struct {
	WorkDayID     int64   `json:"workDayId" binding:"required"`
	EquipmentName string  `json:"equipmentName" binding:"required"`
	Quantity      float64 `json:"quantity" binding:"required"`
	Cost          float64 `json:"cost" binding:"required"`
	Notes         *string `json:"notes"`
}

type UpdateWorkDayEquipment struct {
	EquipmentName *string  `json:"equipmentName"`
	Quantity      *float64 `json:"quantity"`
	Cost          *float64 `json:"cost"`
	Notes         *string  `json:"notes"`
}
