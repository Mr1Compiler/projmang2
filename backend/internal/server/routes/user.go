package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafa91ameen/prjalgo/backend/internal/auth"
	"github.com/mustafa91ameen/prjalgo/backend/internal/container"
)

// RegisterUserRoutes sets up /users and /user-roles routes
func RegisterUserRoutes(rg *gin.RouterGroup, c *container.Container) {
	// Users
	users := rg.Group("/users")
	usersAuthz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/users", perm)
	}
	audit := func(action string) gin.HandlerFunc {
		return auth.AuditMiddleware(c.AuditLogService, action, "user")
	}
	{
		users.GET("", usersAuthz("read"), c.UserHandler.GetAll)
		users.GET("/:id", usersAuthz("read"), c.UserHandler.GetByID)
		users.POST("", usersAuthz("create"), audit("create"), c.UserHandler.Create)
		users.PUT("/:id", usersAuthz("update"), audit("update"), c.UserHandler.Update)
		users.DELETE("/:id", usersAuthz("delete"), audit("delete"), c.UserHandler.Delete)
		users.PUT("/:id/password", usersAuthz("updatePassword"), audit("password_change"), c.UserHandler.UpdatePassword)
		users.PATCH("/:id/status", usersAuthz("updateStatus"), audit("status_change"), c.UserHandler.UpdateStatus)

		// Nested routes under user
		users.GET("/:id/roles", usersAuthz("read"), c.UserRoleHandler.GetByUserID)
		users.GET("/:id/team-members", usersAuthz("read"), c.TeamMemberHandler.GetByUserID)
	}

	// User Roles
	userRoles := rg.Group("/userRoles")
	userRolesAuthz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/userRoles", perm)
	}
	auditUserRole := func(action string) gin.HandlerFunc {
		return auth.AuditMiddleware(c.AuditLogService, action, "user_role")
	}
	{
		userRoles.GET("", userRolesAuthz("read"), c.UserRoleHandler.GetAll)
		userRoles.GET("/:id", userRolesAuthz("read"), c.UserRoleHandler.GetByID)
		userRoles.POST("", userRolesAuthz("create"), auditUserRole("assign_role"), c.UserRoleHandler.Create)
		userRoles.DELETE("/:id", userRolesAuthz("delete"), auditUserRole("remove_role"), c.UserRoleHandler.Delete)
	}
}
