package handlers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/auth"
	"github.com/mustafaameen91/project-managment/backend/internal/response"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body")
		return
	}

	result, err := h.authService.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			response.Unauthorized(c, "invalid username or password")
			return
		}
		response.InternalError(c, "login failed")
		return
	}

	response.Success(c, result)
}

// GetMyPages returns the current user's accessible pages with permissions
func (h *AuthHandler) GetMyPages(c *gin.Context) {
	userID, ok := auth.GetUserID(c)
	if !ok {
		response.Unauthorized(c, "user not authenticated")
		return
	}

	pages, err := h.authService.GetUserPages(c.Request.Context(), userID)
	if err != nil {
		response.InternalError(c, "failed to get user pages")
		return
	}

	response.Success(c, pages)
}
