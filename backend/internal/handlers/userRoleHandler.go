package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/response"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

type UserRoleHandler struct {
	userRoleService *services.UserRoleService
}

func NewUserRoleHandler(userRoleService *services.UserRoleService) *UserRoleHandler {
	return &UserRoleHandler{
		userRoleService: userRoleService,
	}
}

// GetAll handles GET /user-roles
func (h *UserRoleHandler) GetAll(c *gin.Context) {
	userRoles, err := h.userRoleService.GetAll(c.Request.Context())
	if err != nil {
		response.InternalError(c, "failed to fetch user roles")
		return
	}

	response.Success(c, userRoles)
}

// GetByID handles GET /user-roles/:id
func (h *UserRoleHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid user role id")
		return
	}

	userRole, err := h.userRoleService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrUserRoleNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to fetch user role")
		return
	}

	response.Success(c, userRole)
}

// GetByUserID handles GET /users/:id/roles
func (h *UserRoleHandler) GetByUserID(c *gin.Context) {
	userID, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid user id")
		return
	}

	userRoles, err := h.userRoleService.GetByUserID(c.Request.Context(), userID)
	if err != nil {
		response.InternalError(c, "failed to fetch user roles")
		return
	}

	response.Success(c, userRoles)
}

// Create handles POST /user-roles
func (h *UserRoleHandler) Create(c *gin.Context) {
	var req dtos.CreateUserRole
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	userRole, err := h.userRoleService.Create(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, "failed to create user role")
		return
	}

	response.Created(c, userRole)
}

// Delete handles DELETE /user-roles/:id
func (h *UserRoleHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid user role id")
		return
	}

	err = h.userRoleService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrUserRoleNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to delete user role")
		return
	}

	response.Success(c, nil)
}

// parseID extracts an int64 ID from the URL path
func (h *UserRoleHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
