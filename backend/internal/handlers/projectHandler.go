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

type ProjectHandler struct {
	projectService *services.ProjectService
	validator      *validator.Validate
}

func NewProjectHandler(projectService *services.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		projectService: projectService,
		validator:      validator.New(),
	}
}

// GetAll handles GET /projects
func (h *ProjectHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	projects, err := h.projectService.GetAll(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to fetch projects")
		return
	}

	response.Success(w, projects)
}

// GetByID handles GET /projects/{id}
func (h *ProjectHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid project id")
		return
	}

	project, err := h.projectService.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrProjectNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to fetch project")
		return
	}

	response.Success(w, project)
}

// Create handles POST /projects
func (h *ProjectHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateProject
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	project, err := h.projectService.Create(r.Context(), req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to create project")
		return
	}

	response.Created(w, project)
}

// Update handles PUT /projects/{id}
func (h *ProjectHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid project id")
		return
	}

	var req dtos.UpdateProject
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	project, err := h.projectService.Update(r.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrProjectNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to update project")
		return
	}

	response.Success(w, project)
}

// Delete handles DELETE /projects/{id}
func (h *ProjectHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid project id")
		return
	}

	err = h.projectService.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrProjectNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to delete project")
		return
	}

	response.Success(w, nil)
}

// parseID extracts an int64 ID from the URL path
func (h *ProjectHandler) parseID(r *http.Request, param string) (int64, error) {
	idStr := r.PathValue(param)
	return strconv.ParseInt(idStr, 10, 64)
}