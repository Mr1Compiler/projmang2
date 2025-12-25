package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafa91ameen/prjalgo/backend/internal/auth"
	"github.com/mustafa91ameen/prjalgo/backend/internal/container"
)

// RegisterTeamMemberRoutes sets up /teamMembers routes
func RegisterTeamMemberRoutes(rg *gin.RouterGroup, c *container.Container) {
	teamMembers := rg.Group("/teamMembers")
	authz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/teamMembers", perm)
	}
	audit := func(action string) gin.HandlerFunc {
		return auth.AuditMiddleware(c.AuditLogService, action, "team_member")
	}
	{
		teamMembers.GET("", authz("read"), c.TeamMemberHandler.GetAll)
		teamMembers.GET("/stats", authz("read"), c.TeamMemberHandler.GetStats)
		teamMembers.GET("/:id", authz("read"), c.TeamMemberHandler.GetByID)
		teamMembers.POST("", authz("create"), audit("create"), c.TeamMemberHandler.Create)
		teamMembers.DELETE("/:id", authz("delete"), audit("delete"), c.TeamMemberHandler.Delete)
	}
}
