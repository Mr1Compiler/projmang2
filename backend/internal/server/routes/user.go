package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/auth"
	"github.com/mustafaameen91/project-managment/backend/internal/container"
)

// RegisterUserRoutes sets up /users and /user-roles routes
func RegisterUserRoutes(rg *gin.RouterGroup, c *container.Container) {
	// Users
	users := rg.Group("/users")
	usersAuthz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/users", perm)
	}
	{
		users.GET("", usersAuthz("read"), c.UserHandler.GetAll)
		users.GET("/:id", usersAuthz("read"), c.UserHandler.GetByID)
		users.POST("", usersAuthz("write"), c.UserHandler.Create)
		users.PUT("/:id", usersAuthz("write"), c.UserHandler.Update)
		users.DELETE("/:id", usersAuthz("delete"), c.UserHandler.Delete)

		// Nested route under user
		users.GET("/:id/roles", usersAuthz("read"), c.UserRoleHandler.GetByUserID)
	}

	// User Roles
	userRoles := rg.Group("/user-roles")
	userRolesAuthz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/user-roles", perm)
	}
	{
		userRoles.GET("", userRolesAuthz("read"), c.UserRoleHandler.GetAll)
		userRoles.GET("/:id", userRolesAuthz("read"), c.UserRoleHandler.GetByID)
		userRoles.POST("", userRolesAuthz("write"), c.UserRoleHandler.Create)
		userRoles.DELETE("/:id", userRolesAuthz("delete"), c.UserRoleHandler.Delete)
	}
}
