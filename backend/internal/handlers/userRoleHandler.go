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

type UserRoleHandler struct {
	userRoleService *services.UserRoleService
	validator       *validator.Validate
}

func NewUserRoleHandler(userRoleService *services.UserRoleService) *UserRoleHandler {
	return &UserRoleHandler{
		userRoleService: userRoleService,
		validator:       validator.New(),
	}
}

// GetAll handles GET /user-roles
func (h *UserRoleHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	userRoles, err := h.userRoleService.GetAll(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to fetch user roles")
		return
	}

	response.Success(w, userRoles)
}

// GetByID handles GET /user-roles/{id}
func (h *UserRoleHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid user role id")
		return
	}

	userRole, err := h.userRoleService.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrUserRoleNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to fetch user role")
		return
	}

	response.Success(w, userRole)
}

// GetByUserID handles GET /users/{userId}/roles
func (h *UserRoleHandler) GetByUserID(w http.ResponseWriter, r *http.Request) {
	userID, err := h.parseID(r, "userId")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid user id")
		return
	}

	userRoles, err := h.userRoleService.GetByUserID(r.Context(), userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to fetch user roles")
		return
	}

	response.Success(w, userRoles)
}

// Create handles POST /user-roles
func (h *UserRoleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateUserRole
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	userRole, err := h.userRoleService.Create(r.Context(), req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to create user role")
		return
	}

	response.Created(w, userRole)
}

// Delete handles DELETE /user-roles/{id}
func (h *UserRoleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid user role id")
		return
	}

	err = h.userRoleService.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrUserRoleNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to delete user role")
		return
	}

	response.Success(w, nil)
}

// parseID extracts an int64 ID from the URL path
func (h *UserRoleHandler) parseID(r *http.Request, param string) (int64, error) {
	idStr := r.PathValue(param)
	return strconv.ParseInt(idStr, 10, 64)
}
