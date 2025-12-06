package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/response"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

type ProjectHandler struct {
	projectService *services.ProjectService
}

func NewProjectHandler(projectService *services.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		projectService: projectService,
	}
}

// GetAll handles GET /projects
func (h *ProjectHandler) GetAll(c *gin.Context) {
	projects, err := h.projectService.GetAll(c.Request.Context())
	if err != nil {
		response.InternalError(c, "failed to fetch projects")
		return
	}

	response.Success(c, projects)
}

// GetByID handles GET /projects/:id
func (h *ProjectHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid project id")
		return
	}

	project, err := h.projectService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrProjectNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to fetch project")
		return
	}

	response.Success(c, project)
}

// Create handles POST /projects
func (h *ProjectHandler) Create(c *gin.Context) {
	var req dtos.CreateProject
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	project, err := h.projectService.Create(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, "failed to create project")
		return
	}

	response.Created(c, project)
}

// Update handles PUT /projects/:id
func (h *ProjectHandler) Update(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid project id")
		return
	}

	var req dtos.UpdateProject
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	project, err := h.projectService.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrProjectNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to update project")
		return
	}

	response.Success(c, project)
}

// Delete handles DELETE /projects/:id
func (h *ProjectHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid project id")
		return
	}

	err = h.projectService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrProjectNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to delete project")
		return
	}

	response.Success(c, nil)
}

// parseID extracts an int64 ID from the URL path
func (h *ProjectHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
