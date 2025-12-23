package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafa91ameen/prjalgo/backend/internal/auth"
	"github.com/mustafa91ameen/prjalgo/backend/internal/container"
)

// RegisterProjectRoutes sets up /projects routes
func RegisterProjectRoutes(rg *gin.RouterGroup, c *container.Container) {
	projects := rg.Group("/projects")
	authz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/projects", perm)
	}
	audit := func(action string) gin.HandlerFunc {
		return auth.AuditMiddleware(c.AuditLogService, action, "project")
	}
	{
		projects.GET("", authz("read"), c.ProjectHandler.GetAll)
		projects.GET("/stats", authz("read"), c.ProjectHandler.GetStats)
		projects.GET("/:id", authz("read"), c.ProjectHandler.GetByID)
		projects.POST("", authz("write"), audit("create"), c.ProjectHandler.Create)
		projects.PUT("/:id", authz("write"), audit("update"), c.ProjectHandler.Update)
		projects.DELETE("/:id", authz("delete"), audit("delete"), c.ProjectHandler.Delete)

		// Nested resources under project
		projects.GET("/:id/workdays", authz("read"), c.WorkDayHandler.GetByProjectID)
		projects.GET("/:id/expenses", authz("read"), c.ExpenseHandler.GetByProjectID)
		projects.GET("/:id/team-members", authz("read"), c.TeamMemberHandler.GetByProjectID)
	}
}
