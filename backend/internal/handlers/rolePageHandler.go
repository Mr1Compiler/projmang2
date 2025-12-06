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

type RolePageHandler struct {
	rolePageService *services.RolePageService
	validator       *validator.Validate
}

func NewRolePageHandler(rolePageService *services.RolePageService) *RolePageHandler {
	return &RolePageHandler{
		rolePageService: rolePageService,
		validator:       validator.New(),
	}
}

// GetAll handles GET /role-pages
func (h *RolePageHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	rolePages, err := h.rolePageService.GetAll(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to fetch role pages")
		return
	}

	response.Success(w, rolePages)
}

// GetByID handles GET /role-pages/{id}
func (h *RolePageHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid role page id")
		return
	}

	rolePage, err := h.rolePageService.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrRolePageNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to fetch role page")
		return
	}

	response.Success(w, rolePage)
}

// GetByRoleID handles GET /roles/{roleId}/pages
func (h *RolePageHandler) GetByRoleID(w http.ResponseWriter, r *http.Request) {
	roleID, err := h.parseID(r, "roleId")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid role id")
		return
	}

	rolePages, err := h.rolePageService.GetByRoleID(r.Context(), roleID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to fetch role pages")
		return
	}

	response.Success(w, rolePages)
}

// Create handles POST /role-pages
func (h *RolePageHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateRolePage
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	rolePage, err := h.rolePageService.Create(r.Context(), req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to create role page")
		return
	}

	response.Created(w, rolePage)
}

// Update handles PUT /role-pages/{id}
func (h *RolePageHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid role page id")
		return
	}

	var req dtos.UpdateRolePage
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	rolePage, err := h.rolePageService.Update(r.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrRolePageNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to update role page")
		return
	}

	response.Success(w, rolePage)
}

// Delete handles DELETE /role-pages/{id}
func (h *RolePageHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid role page id")
		return
	}

	err = h.rolePageService.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrRolePageNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to delete role page")
		return
	}

	response.Success(w, nil)
}

// parseID extracts an int64 ID from the URL path
func (h *RolePageHandler) parseID(r *http.Request, param string) (int64, error) {
	idStr := r.PathValue(param)
	return strconv.ParseInt(idStr, 10, 64)
}
