package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/auth"
	"github.com/mustafaameen91/project-managment/backend/internal/container"
)

// RegisterRoleRoutes sets up /roles, /pages, and /role-pages routes
func RegisterRoleRoutes(rg *gin.RouterGroup, c *container.Container) {
	// Roles
	roles := rg.Group("/roles")
	rolesAuthz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/roles", perm)
	}
	{
		roles.GET("", rolesAuthz("read"), c.RoleHandler.GetAll)
		roles.GET("/:id", rolesAuthz("read"), c.RoleHandler.GetByID)
		roles.POST("", rolesAuthz("write"), c.RoleHandler.Create)
		roles.PUT("/:id", rolesAuthz("write"), c.RoleHandler.Update)
		roles.DELETE("/:id", rolesAuthz("delete"), c.RoleHandler.Delete)
	}

	// Pages
	pages := rg.Group("/pages")
	pagesAuthz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/pages", perm)
	}
	{
		pages.GET("", pagesAuthz("read"), c.PageHandler.GetAll)
		pages.GET("/:id", pagesAuthz("read"), c.PageHandler.GetByID)
		pages.POST("", pagesAuthz("write"), c.PageHandler.Create)
		pages.PUT("/:id", pagesAuthz("write"), c.PageHandler.Update)
		pages.DELETE("/:id", pagesAuthz("delete"), c.PageHandler.Delete)
	}

	// Role Pages
	rolePages := rg.Group("/role-pages")
	rolePagesAuthz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/role-pages", perm)
	}
	{
		rolePages.GET("", rolePagesAuthz("read"), c.RolePageHandler.GetAll)
		rolePages.GET("/:id", rolePagesAuthz("read"), c.RolePageHandler.GetByID)
		rolePages.POST("", rolePagesAuthz("write"), c.RolePageHandler.Create)
		rolePages.PUT("/:id", rolePagesAuthz("write"), c.RolePageHandler.Update)
		rolePages.DELETE("/:id", rolePagesAuthz("delete"), c.RolePageHandler.Delete)
	}
}
