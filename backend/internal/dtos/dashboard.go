package dtos

// Response DTOs
type DashboardStats struct {
	TotalProjects     int64 `json:"totalProjects"`
	ActiveProjects    int64 `json:"activeProjects"`
	CompletedProjects int64 `json:"completedProjects"`
	TotalEngineers    int64 `json:"totalEngineers"`
}

type DashboardFinancial struct {
	TotalIncome   float64 `json:"totalIncome"`
	TotalExpenses float64 `json:"totalExpenses"`
}

type ProjectProgressItem struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Progress float64 `json:"progress"`
}

type DashboardProjectProgress struct {
	Projects []ProjectProgressItem `json:"projects"`
}

type DashboardTaskSummary struct {
	Labels    []string `json:"labels"`
	Completed []int    `json:"completed"`
	Pending   []int    `json:"pending"`
}

// Request DTOs
type FinancialQuery struct {
	Period string `form:"period" binding:"omitempty,oneof=all month"`
	Month  string `form:"month" binding:"omitempty"`
}

type ProjectProgressQuery struct {
	Status string `form:"status" binding:"omitempty,oneof=pending in_progress completed all"`
	Limit  int    `form:"limit" binding:"omitempty,min=1,max=20"`
}

type TaskSummaryQuery struct {
	Month string `form:"month" binding:"required"`
}

// Activity DTOs
type ActivityResponse struct {
	ID         int64  `json:"id"`
	ActorName  string `json:"actorName"`
	Action     string `json:"action"`
	TargetType string `json:"targetType"`
	TargetName string `json:"targetName"`
	CreatedAt  string `json:"createdAt"`
}
