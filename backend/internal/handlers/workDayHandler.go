package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/response"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

type WorkDayHandler struct {
	workDayService *services.WorkDayService
}

func NewWorkDayHandler(workDayService *services.WorkDayService) *WorkDayHandler {
	return &WorkDayHandler{
		workDayService: workDayService,
	}
}

// GetAll handles GET /workdays
func (h *WorkDayHandler) GetAll(c *gin.Context) {
	var pagination dtos.PaginationQuery
	if err := c.ShouldBindQuery(&pagination); err != nil {
		response.ValidationError(c, err)
		return
	}
	pagination.Normalize()

	workDays, err := h.workDayService.GetAll(c.Request.Context(), pagination.Limit, pagination.Offset())
	if err != nil {
		response.InternalError(c, "failed to fetch work days")
		return
	}

	response.Success(c, workDays)
}

// GetByID handles GET /workdays/:id
func (h *WorkDayHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid work day id")
		return
	}

	workDay, err := h.workDayService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to fetch work day")
		return
	}

	response.Success(c, workDay)
}

// GetByProjectID handles GET /projects/:id/workdays
func (h *WorkDayHandler) GetByProjectID(c *gin.Context) {
	projectID, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid project id")
		return
	}

	workDays, err := h.workDayService.GetByProjectID(c.Request.Context(), projectID)
	if err != nil {
		response.InternalError(c, "failed to fetch work days")
		return
	}

	response.Success(c, workDays)
}

// Create handles POST /workdays
func (h *WorkDayHandler) Create(c *gin.Context) {
	var req dtos.CreateWorkDay
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	workDay, err := h.workDayService.Create(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, "failed to create work day")
		return
	}

	response.Created(c, workDay)
}

// Update handles PUT /workdays/:id
func (h *WorkDayHandler) Update(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid work day id")
		return
	}

	var req dtos.UpdateWorkDay
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	workDay, err := h.workDayService.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to update work day")
		return
	}

	response.Success(c, workDay)
}

// Delete handles DELETE /workdays/:id
func (h *WorkDayHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid work day id")
		return
	}

	err = h.workDayService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to delete work day")
		return
	}

	response.Success(c, nil)
}

// Complete handles PATCH /workdays/:id/complete
func (h *WorkDayHandler) Complete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid work day id")
		return
	}

	workDay, err := h.workDayService.Complete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		if errors.Is(err, services.ErrWorkDayAlreadyCompleted) {
			response.BadRequest(c, err.Error())
			return
		}
		if errors.Is(err, services.ErrWorkDayNoSubCategory) {
			response.BadRequest(c, err.Error())
			return
		}
		response.InternalError(c, "failed to complete work day")
		return
	}

	response.Success(c, workDay)
}

// parseID extracts an int64 ID from the URL path
func (h *WorkDayHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
