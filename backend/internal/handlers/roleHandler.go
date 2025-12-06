package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/response"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

type RoleHandler struct {
	roleService *services.RoleService
	validator   *validator.Validate
}

func NewRoleHandler(roleService *services.RoleService) *RoleHandler {
	return &RoleHandler{
		roleService: roleService,
		validator:   validator.New(),
	}
}

// GetAll handles GET /roles
func (h *RoleHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	roles, err := h.roleService.GetAll(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to fetch roles")
		return
	}

	response.Success(w, roles)
}

// GetByID handles GET /roles/{id}
func (h *RoleHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid role id")
		return
	}

	role, err := h.roleService.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrRoleNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to fetch role")
		return
	}

	response.Success(w, role)
}

// Create handles POST /roles
func (h *RoleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateRole
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	role, err := h.roleService.Create(r.Context(), req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to create role")
		return
	}

	response.Created(w, role)
}

// Update handles PUT /roles/{id}
func (h *RoleHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid role id")
		return
	}

	var req dtos.UpdateRole
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	role, err := h.roleService.Update(r.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrRoleNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to update role")
		return
	}

	response.Success(w, role)
}

// Delete handles DELETE /roles/{id}
func (h *RoleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid role id")
		return
	}

	err = h.roleService.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrRoleNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to delete role")
		return
	}

	response.Success(w, nil)
}

// parseID extracts an int64 ID from the URL path
func (h *RoleHandler) parseID(r *http.Request, param string) (int64, error) {
	idStr := r.PathValue(param)
	return strconv.ParseInt(idStr, 10, 64)
}
