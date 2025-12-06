package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/auth"
	"github.com/mustafaameen91/project-managment/backend/internal/container"
)

// RegisterRoutes sets up all API routes
func RegisterRoutes(router *gin.Engine, c *container.Container) {
	// Health check
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	// Public routes (no auth required)
	RegisterAuthRoutes(router, c)

	// Protected API v1 group
	api := router.Group("/api/v1")
	api.Use(auth.JWTMiddleware(c.JWTManager))

	// Register all domain routes
	RegisterProjectRoutes(api, c)
	RegisterWorkDayRoutes(api, c)
	RegisterWorkCategoryRoutes(api, c)
	RegisterFinanceRoutes(api, c)
	RegisterDebtorRoutes(api, c)
	RegisterUserRoutes(api, c)
	RegisterRoleRoutes(api, c)
}
