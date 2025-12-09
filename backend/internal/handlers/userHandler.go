package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/response"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetAll handles GET /users
func (h *UserHandler) GetAll(c *gin.Context) {
	var pagination dtos.PaginationQuery
	if err := c.ShouldBindQuery(&pagination); err != nil {
		response.ValidationError(c, err)
		return
	}
	pagination.Normalize()

	users, err := h.userService.GetAll(c.Request.Context(), pagination.Limit, pagination.Offset())
	if err != nil {
		response.InternalError(c, "failed to fetch users")
		return
	}

	response.Success(c, users)
}

// GetByID handles GET /users/:id
func (h *UserHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid user id")
		return
	}

	user, err := h.userService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to fetch user")
		return
	}

	response.Success(c, user)
}

// Create handles POST /users
func (h *UserHandler) Create(c *gin.Context) {
	var req dtos.CreateUser
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	user, err := h.userService.Create(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, services.ErrUserExists) {
			response.Error(c, http.StatusConflict, err.Error())
			return
		}
		response.InternalError(c, "failed to create user")
		return
	}

	response.Created(c, user)
}

// Update handles PUT /users/:id
func (h *UserHandler) Update(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid user id")
		return
	}

	var req dtos.UpdateUser
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	user, err := h.userService.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to update user")
		return
	}

	response.Success(c, user)
}

// Delete handles DELETE /users/:id
func (h *UserHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid user id")
		return
	}

	err = h.userService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to delete user")
		return
	}

	response.Success(c, nil)
}

// UpdatePassword handles PUT /users/:id/password
func (h *UserHandler) UpdatePassword(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid user id")
		return
	}

	var req dtos.UpdatePassword
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	err = h.userService.UpdatePassword(c.Request.Context(), id, req.Password)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to update password")
		return
	}

	response.Success(c, nil)
}

// UpdateStatus handles PATCH /users/:id/status
func (h *UserHandler) UpdateStatus(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid user id")
		return
	}

	var req dtos.UpdateStatus
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	// Normalize to lowercase
	status := strings.ToLower(req.Status)
	if status != "active" && status != "inactive" {
		response.BadRequest(c, "status must be 'active' or 'inactive'")
		return
	}

	err = h.userService.UpdateStatus(c.Request.Context(), id, status)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to update status")
		return
	}

	response.Success(c, gin.H{"status": status})
}

// parseID extracts an int64 ID from the URL path
func (h *UserHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
