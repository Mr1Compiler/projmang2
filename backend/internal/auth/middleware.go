package auth

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/response"
)

const (
	AuthorizationHeader = "Authorization"
	BearerPrefix        = "Bearer "
	UserIDKey           = "user_id"
	UsernameKey         = "username"
)

// PermissionChecker interface for checking user permissions
type PermissionChecker interface {
	HasPermission(ctx context.Context, userID int64, route string, permission string) (bool, error)
}

func JWTMiddleware(jwtManager *JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(AuthorizationHeader)
		if authHeader == "" {
			response.Unauthorized(c, "authorization header required")
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, BearerPrefix) {
			response.Unauthorized(c, "invalid authorization format")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, BearerPrefix)
		claims, err := jwtManager.ValidateToken(tokenString)
		if err != nil {
			response.Unauthorized(c, err.Error())
			c.Abort()
			return
		}

		c.Set(UserIDKey, claims.UserID)
		c.Set(UsernameKey, claims.Username)
		c.Next()
	}
}

// AuthorizationMiddleware checks if user has permission to access a route
func AuthorizationMiddleware(checker PermissionChecker, route string, permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get(UserIDKey)
		if !exists {
			response.Unauthorized(c, "user not authenticated")
			c.Abort()
			return
		}

		hasPermission, err := checker.HasPermission(c.Request.Context(), userID.(int64), route, permission)
		if err != nil {
			response.InternalError(c, "failed to check permissions")
			c.Abort()
			return
		}

		if !hasPermission {
			response.Error(c, http.StatusForbidden, "access denied")
			c.Abort()
			return
		}

		c.Next()
	}
}

// GetUserID helper to get user ID from context
func GetUserID(c *gin.Context) (int64, bool) {
	userID, exists := c.Get(UserIDKey)
	if !exists {
		return 0, false
	}
	return userID.(int64), true
}

// AuditLogger interface for audit logging
type AuditLogger interface {
	Log(ctx context.Context, actorID *int64, action, targetType string, targetID *int64, method, status, ipAddress string)
}

// AuditMiddleware logs actions to audit log
func AuditMiddleware(auditService AuditLogger, action, targetType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		status := "success"
		if c.Writer.Status() >= 400 {
			status = "failure"
		}

		var actorID *int64
		if userID, exists := c.Get(UserIDKey); exists {
			id := userID.(int64)
			actorID = &id
		}

		var targetID *int64
		if idStr := c.Param("id"); idStr != "" {
			if id, err := strconv.ParseInt(idStr, 10, 64); err == nil {
				targetID = &id
			}
		}

		auditService.Log(
			c.Request.Context(),
			actorID,
			action,
			targetType,
			targetID,
			c.Request.Method,
			status,
			c.ClientIP(),
		)
	}
}

// AuditAuthMiddleware for login/logout (no authenticated user yet)
func AuditAuthMiddleware(auditService AuditLogger, action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		status := "success"
		if c.Writer.Status() >= 400 {
			status = "failure"
		}

		auditService.Log(
			c.Request.Context(),
			nil,
			action,
			"auth",
			nil,
			c.Request.Method,
			status,
			c.ClientIP(),
		)
	}
}
