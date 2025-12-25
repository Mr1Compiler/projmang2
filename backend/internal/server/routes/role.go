package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafa91ameen/prjalgo/backend/internal/auth"
	"github.com/mustafa91ameen/prjalgo/backend/internal/container"
)

// RegisterRoleRoutes sets up /roles, /pages, and /role-pages routes
func RegisterRoleRoutes(rg *gin.RouterGroup, c *container.Container) {
	// Roles
	roles := rg.Group("/roles")
	rolesAuthz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/roles", perm)
	}
	auditRole := func(action string) gin.HandlerFunc {
		return auth.AuditMiddleware(c.AuditLogService, action, "role")
	}
	{
		roles.GET("", rolesAuthz("read"), c.RoleHandler.GetAll)
		roles.GET("/:id", rolesAuthz("read"), c.RoleHandler.GetByID)
		roles.POST("", rolesAuthz("create"), auditRole("create"), c.RoleHandler.Create)
		roles.PUT("/:id", rolesAuthz("update"), auditRole("update"), c.RoleHandler.Update)
		roles.DELETE("/:id", rolesAuthz("delete"), auditRole("delete"), c.RoleHandler.Delete)
	}

	// Pages
	pages := rg.Group("/pages")
	pagesAuthz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/pages", perm)
	}
	auditPage := func(action string) gin.HandlerFunc {
		return auth.AuditMiddleware(c.AuditLogService, action, "page")
	}
	{
		pages.GET("", pagesAuthz("read"), c.PageHandler.GetAll)
		pages.GET("/:id", pagesAuthz("read"), c.PageHandler.GetByID)
		pages.POST("", pagesAuthz("create"), auditPage("create"), c.PageHandler.Create)
		pages.PUT("/:id", pagesAuthz("update"), auditPage("update"), c.PageHandler.Update)
		pages.DELETE("/:id", pagesAuthz("delete"), auditPage("delete"), c.PageHandler.Delete)
	}

	// Role Pages
	rolePages := rg.Group("/rolePages")
	rolePagesAuthz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/rolePages", perm)
	}
	auditRolePage := func(action string) gin.HandlerFunc {
		return auth.AuditMiddleware(c.AuditLogService, action, "role_page")
	}
	{
		rolePages.GET("", rolePagesAuthz("read"), c.RolePageHandler.GetAll)
		rolePages.GET("/:id", rolePagesAuthz("read"), c.RolePageHandler.GetByID)
		rolePages.POST("", rolePagesAuthz("create"), auditRolePage("create"), c.RolePageHandler.Create)
		rolePages.PUT("/:id", rolePagesAuthz("update"), auditRolePage("update"), c.RolePageHandler.Update)
		rolePages.DELETE("/:id", rolePagesAuthz("delete"), auditRolePage("delete"), c.RolePageHandler.Delete)
	}
}
