package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/response"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

type RoleHandler struct {
	roleService *services.RoleService
}

func NewRoleHandler(roleService *services.RoleService) *RoleHandler {
	return &RoleHandler{
		roleService: roleService,
	}
}

// GetAll handles GET /roles
func (h *RoleHandler) GetAll(c *gin.Context) {
	roles, err := h.roleService.GetAll(c.Request.Context())
	if err != nil {
		response.InternalError(c, "failed to fetch roles")
		return
	}

	response.Success(c, roles)
}

// GetByID handles GET /roles/:id
func (h *RoleHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid role id")
		return
	}

	role, err := h.roleService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrRoleNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to fetch role")
		return
	}

	response.Success(c, role)
}

// Create handles POST /roles
func (h *RoleHandler) Create(c *gin.Context) {
	var req dtos.CreateRole
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	role, err := h.roleService.Create(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, "failed to create role")
		return
	}

	response.Created(c, role)
}

// Update handles PUT /roles/:id
func (h *RoleHandler) Update(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid role id")
		return
	}

	var req dtos.UpdateRole
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	role, err := h.roleService.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrRoleNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to update role")
		return
	}

	response.Success(c, role)
}

// Delete handles DELETE /roles/:id
func (h *RoleHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid role id")
		return
	}

	err = h.roleService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrRoleNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to delete role")
		return
	}

	response.Success(c, nil)
}

// parseID extracts an int64 ID from the URL path
func (h *RoleHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
