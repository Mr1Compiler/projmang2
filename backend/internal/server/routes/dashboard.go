package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafa91ameen/prjalgo/backend/internal/auth"
	"github.com/mustafa91ameen/prjalgo/backend/internal/container"
)

// RegisterDashboardRoutes sets up /dashboard routes
func RegisterDashboardRoutes(rg *gin.RouterGroup, c *container.Container) {
	dashboard := rg.Group("/dashboard")
	authz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/dashboard", perm)
	}
	{
		dashboard.GET("/stats", authz("read"), c.DashboardHandler.GetStats)
		dashboard.GET("/financial", authz("read"), c.DashboardHandler.GetFinancial)
		dashboard.GET("/project-progress", authz("read"), c.DashboardHandler.GetProjectProgress)
		dashboard.GET("/task-summary", authz("read"), c.DashboardHandler.GetTaskSummary)
		dashboard.GET("/activities", authz("read"), c.DashboardHandler.GetActivities)
	}
}
