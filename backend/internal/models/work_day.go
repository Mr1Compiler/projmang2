package models

import "time"

type WorkDay struct {
	ID                int64     `json:"id" db:"id"`
	ProjectID         int64     `json:"projectId" db:"projectid"`
	WorkSubCategoryID *int64    `json:"workSubCategoryId" db:"worksubcategoryid"`
	WorkDate          time.Time `json:"workDate" db:"workdate"`
	Description       *string   `json:"description" db:"description"`
	Status            string    `json:"status" db:"status"`
	TotalCost         float64   `json:"totalCost" db:"totalcost"`
	Notes             *string   `json:"notes" db:"notes"`
	CreatedBy         *int64    `json:"createdBy" db:"createdby"`
	CreatedAt         time.Time `json:"createdAt" db:"createdat"`
	UpdatedAt         time.Time `json:"updatedAt" db:"updatedat"`
}

type WorkDayMaterial struct {
	ID           int64     `json:"id" db:"id"`
	WorkDayID    int64     `json:"workDayId" db:"workdayid"`
	MaterialName string    `json:"materialName" db:"materialname"`
	Quantity     float64   `json:"quantity" db:"quantity"`
	Cost         float64   `json:"cost" db:"cost"`
	Notes        *string   `json:"notes" db:"notes"`
	CreatedAt    time.Time `json:"createdAt" db:"createdat"`
	UpdatedAt    time.Time `json:"updatedAt" db:"updatedat"`
}

type WorkDayLabor struct {
	ID         int64     `json:"id" db:"id"`
	WorkDayID  int64     `json:"workDayId" db:"workdayid"`
	WorkerName string    `json:"workerName" db:"workername"`
	JobTitle   *string   `json:"jobTitle" db:"jobtitle"`
	Phone      *string   `json:"phone" db:"phone"`
	Address    *string   `json:"address" db:"address"`
	Quantity   float64   `json:"quantity" db:"quantity"`
	Cost       float64   `json:"cost" db:"cost"`
	Notes      *string   `json:"notes" db:"notes"`
	CreatedAt  time.Time `json:"createdAt" db:"createdat"`
	UpdatedAt  time.Time `json:"updatedAt" db:"updatedat"`
}

type WorkDayEquipment struct {
	ID            int64     `json:"id" db:"id"`
	WorkDayID     int64     `json:"workDayId" db:"workdayid"`
	EquipmentName string    `json:"equipmentName" db:"equipmentname"`
	EquipmentType *string   `json:"equipmentType" db:"equipmenttype"`
	Quantity      float64   `json:"quantity" db:"quantity"`
	Cost          float64   `json:"cost" db:"cost"`
	Notes         *string   `json:"notes" db:"notes"`
	CreatedAt     time.Time `json:"createdAt" db:"createdat"`
	UpdatedAt     time.Time `json:"updatedAt" db:"updatedat"`
}
