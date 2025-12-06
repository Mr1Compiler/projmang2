package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/auth"
	"github.com/mustafaameen91/project-managment/backend/internal/container"
)

// RegisterDebtorRoutes sets up /debtors routes
func RegisterDebtorRoutes(rg *gin.RouterGroup, c *container.Container) {
	debtors := rg.Group("/debtors")
	authz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/debtors", perm)
	}
	{
		debtors.GET("", authz("read"), c.DebtorHandler.GetAll)
		debtors.GET("/:id", authz("read"), c.DebtorHandler.GetByID)
		debtors.POST("", authz("write"), c.DebtorHandler.Create)
		debtors.PUT("/:id", authz("write"), c.DebtorHandler.Update)
		debtors.DELETE("/:id", authz("delete"), c.DebtorHandler.Delete)
	}
}
