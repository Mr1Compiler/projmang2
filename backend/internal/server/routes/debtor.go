package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafa91ameen/prjalgo/backend/internal/auth"
	"github.com/mustafa91ameen/prjalgo/backend/internal/container"
)

// RegisterDebtorRoutes sets up /debtors routes
func RegisterDebtorRoutes(rg *gin.RouterGroup, c *container.Container) {
	debtors := rg.Group("/debtors")
	authz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/debtors", perm)
	}
	audit := func(action string) gin.HandlerFunc {
		return auth.AuditMiddleware(c.AuditLogService, action, "debtor")
	}
	{
		debtors.GET("", authz("read"), c.DebtorHandler.GetAll)
		debtors.GET("/stats", authz("read"), c.DebtorHandler.GetStats)
		debtors.GET("/:id", authz("read"), c.DebtorHandler.GetByID)
		debtors.POST("", authz("create"), audit("create"), c.DebtorHandler.Create)
		debtors.PUT("/:id", authz("update"), audit("update"), c.DebtorHandler.Update)
		debtors.DELETE("/:id", authz("delete"), audit("delete"), c.DebtorHandler.Delete)
	}
}
