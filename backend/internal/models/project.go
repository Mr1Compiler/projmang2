package models

import "time"

type Project struct {
	ID                 int64      `json:"id" db:"id"`
	Name               string     `json:"name" db:"name"`
	Type               *string    `json:"type" db:"type"`
	Description        *string    `json:"description" db:"description"`
	ClientPhone        *string    `json:"clientPhone" db:"clientphone"`
	Location           *string    `json:"location" db:"location"`
	StartDate          *time.Time `json:"startDate" db:"startdate"`
	Duration           *int       `json:"duration" db:"duration"`
	WarningCost        float64    `json:"warningCost" db:"warningcost"`
	TotalCost          float64    `json:"totalCost" db:"totalcost"`
	Status             string     `json:"status" db:"status"`
	ProgressPercentage float64    `json:"progressPercentage" db:"progresspercentage"`
	Notes              *string    `json:"notes" db:"notes"`
	CreatedBy          *int64     `json:"createdBy" db:"createdby"`
	CreatedAt          time.Time  `json:"createdAt" db:"createdat"`
	UpdatedAt          time.Time  `json:"updatedAt" db:"updatedat"`
	IsActive           bool       `json:"isActive" db:"isactive"`
}
