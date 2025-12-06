package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/auth"
	"github.com/mustafaameen91/project-managment/backend/internal/container"
)

func RegisterAuthRoutes(router *gin.Engine, c *container.Container) {
	// Public routes
	authPublic := router.Group("/api/v1/auth")
	{
		authPublic.POST("/login", c.AuthHandler.Login)
	}

	// Protected routes (require JWT)
	authProtected := router.Group("/api/v1/auth")
	authProtected.Use(auth.JWTMiddleware(c.JWTManager))
	{
		authProtected.GET("/pages", c.AuthHandler.GetMyPages)
	}
}
