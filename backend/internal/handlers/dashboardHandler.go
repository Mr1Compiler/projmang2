package handlers

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mustafa91ameen/prjalgo/backend/internal/dtos"
	"github.com/mustafa91ameen/prjalgo/backend/internal/response"
	"github.com/mustafa91ameen/prjalgo/backend/internal/services"
)

type DashboardHandler struct {
	dashboardService *services.DashboardService
}

func NewDashboardHandler(dashboardService *services.DashboardService) *DashboardHandler {
	return &DashboardHandler{
		dashboardService: dashboardService,
	}
}

// GetStats handles GET /dashboard/stats
func (h *DashboardHandler) GetStats(c *gin.Context) {
	stats, err := h.dashboardService.GetStats(c.Request.Context())
	if err != nil {
		if errors.Is(err, services.ErrDashboardStatsNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to fetch dashboard stats")
		return
	}

	response.Success(c, stats)
}

// GetFinancial handles GET /dashboard/financial
func (h *DashboardHandler) GetFinancial(c *gin.Context) {
	var query dtos.FinancialQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.ValidationError(c, err)
		return
	}

	// Default to "all" if not specified
	if query.Period == "" {
		query.Period = "all"
	}

	var month *time.Time
	if query.Period == "month" {
		if query.Month == "" {
			response.BadRequest(c, "month is required when period is 'month'")
			return
		}
		parsedMonth, err := time.Parse("2006-01", query.Month)
		if err != nil {
			response.BadRequest(c, "invalid month format, use YYYY-MM")
			return
		}
		month = &parsedMonth
	}

	financial, err := h.dashboardService.GetFinancial(c.Request.Context(), query.Period, month)
	if err != nil {
		if errors.Is(err, services.ErrFinancialStatsNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to fetch financial stats")
		return
	}

	response.Success(c, financial)
}

// GetProjectProgress handles GET /dashboard/project-progress
func (h *DashboardHandler) GetProjectProgress(c *gin.Context) {
	var query dtos.ProjectProgressQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.ValidationError(c, err)
		return
	}

	// Defaults
	if query.Status == "" {
		query.Status = "in_progress"
	}
	if query.Limit == 0 {
		query.Limit = 5
	}

	progress, err := h.dashboardService.GetProjectProgress(c.Request.Context(), query.Status, query.Limit)
	if err != nil {
		if errors.Is(err, services.ErrProjectProgressNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to fetch project progress")
		return
	}

	response.Success(c, progress)
}

// GetTaskSummary handles GET /dashboard/task-summary
func (h *DashboardHandler) GetTaskSummary(c *gin.Context) {
	var query dtos.TaskSummaryQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.ValidationError(c, err)
		return
	}

	parsedMonth, err := time.Parse("2006-01", query.Month)
	if err != nil {
		response.BadRequest(c, "invalid month format, use YYYY-MM")
		return
	}

	summary, err := h.dashboardService.GetTaskSummary(c.Request.Context(), parsedMonth)
	if err != nil {
		if errors.Is(err, services.ErrTaskSummaryNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to fetch task summary")
		return
	}

	response.Success(c, summary)
}

// GetActivities handles GET /dashboard/activities
func (h *DashboardHandler) GetActivities(c *gin.Context) {
	var pagination dtos.PaginationQuery
	if err := c.ShouldBindQuery(&pagination); err != nil {
		response.ValidationError(c, err)
		return
	}
	pagination.Normalize()

	activities, err := h.dashboardService.GetActivities(c.Request.Context(), pagination.Limit, pagination.Offset())
	if err != nil {
		if errors.Is(err, services.ErrActivitiesNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to fetch activities")
		return
	}

	response.Success(c, activities)
}
