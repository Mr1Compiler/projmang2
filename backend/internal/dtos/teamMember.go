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

// Stats Response DTO
type TeamMemberStatsResponse struct {
	Total          int64   `json:"total"`
	UniqueUsers    int64   `json:"uniqueUsers"`
	UniqueProjects int64   `json:"uniqueProjects"`
	AvgPerProject  float64 `json:"avgPerProject"`
}
