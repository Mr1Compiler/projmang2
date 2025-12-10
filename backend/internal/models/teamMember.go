package models

import "time"

type TeamMember struct {
	ID        int64     `json:"id" db:"id"`
	ProjectID int64     `json:"projectId" db:"projectid"`
	UserID    int64     `json:"userId" db:"userid"`
	CreatedAt time.Time `json:"createdAt" db:"createdat"`
	UpdatedAt time.Time `json:"updatedAt" db:"updatedat"`
}
