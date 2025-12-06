package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/auth"
	"github.com/mustafaameen91/project-managment/backend/internal/container"
)

// RegisterProjectRoutes sets up /projects routes
func RegisterProjectRoutes(rg *gin.RouterGroup, c *container.Container) {
	projects := rg.Group("/projects")
	authz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/projects", perm)
	}
	{
		projects.GET("", authz("read"), c.ProjectHandler.GetAll)
		projects.GET("/:id", authz("read"), c.ProjectHandler.GetByID)
		projects.POST("", authz("write"), c.ProjectHandler.Create)
		projects.PUT("/:id", authz("write"), c.ProjectHandler.Update)
		projects.DELETE("/:id", authz("delete"), c.ProjectHandler.Delete)

		// Nested resources under project
		projects.GET("/:id/workdays", authz("read"), c.WorkDayHandler.GetByProjectID)
		projects.GET("/:id/expenses", authz("read"), c.ExpenseHandler.GetByProjectID)
	}
}
