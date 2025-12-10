package dtos

import "time"

// ** TeamMember DTOs **

// Response DTO
type TeamMember struct {
	ID        int64      `json:"id"`
	ProjectID int64      `json:"projectId"`
	UserID    int64      `json:"userId"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// Request DTO
type CreateTeamMember struct {
	ProjectID int64 `json:"projectId" binding:"required"`
	UserID    int64 `json:"userId" binding:"required"`
}
