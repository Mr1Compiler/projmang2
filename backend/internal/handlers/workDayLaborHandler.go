package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/response"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

type WorkDayLaborHandler struct {
	laborService *services.WorkDayLaborService
}

func NewWorkDayLaborHandler(laborService *services.WorkDayLaborService) *WorkDayLaborHandler {
	return &WorkDayLaborHandler{
		laborService: laborService,
	}
}

// GetAll handles GET /workday-labor
func (h *WorkDayLaborHandler) GetAll(c *gin.Context) {
	var pagination dtos.PaginationQuery
	if err := c.ShouldBindQuery(&pagination); err != nil {
		response.ValidationError(c, err)
		return
	}
	pagination.Normalize()

	labors, err := h.laborService.GetAll(c.Request.Context(), pagination.Limit, pagination.Offset())
	if err != nil {
		response.InternalError(c, "failed to fetch labors")
		return
	}

	response.Success(c, labors)
}

// GetByWorkDayID handles GET /workdays/:workDayId/labors
func (h *WorkDayLaborHandler) GetByWorkDayID(c *gin.Context) {
	workDayID, err := h.parseID(c, "workDayId")
	if err != nil {
		response.BadRequest(c, "invalid work day id")
		return
	}

	labors, err := h.laborService.GetByWorkDayID(c.Request.Context(), workDayID)
	if err != nil {
		response.InternalError(c, "failed to fetch labors")
		return
	}

	response.Success(c, labors)
}

// GetByID handles GET /workdays/:workDayId/labors/:id
func (h *WorkDayLaborHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid labor id")
		return
	}

	labor, err := h.laborService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayLaborNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to fetch labor")
		return
	}

	response.Success(c, labor)
}

// Create handles POST /workday-labor
func (h *WorkDayLaborHandler) Create(c *gin.Context) {
	var req dtos.CreateWorkDayLabor
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	labor, err := h.laborService.Create(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, "failed to create labor")
		return
	}

	response.Created(c, labor)
}

// Update handles PUT /workdays/:workDayId/labors/:id
func (h *WorkDayLaborHandler) Update(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid labor id")
		return
	}

	var req dtos.UpdateWorkDayLabor
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	labor, err := h.laborService.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayLaborNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to update labor")
		return
	}

	response.Success(c, labor)
}

// Delete handles DELETE /workdays/:workDayId/labors/:id
func (h *WorkDayLaborHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid labor id")
		return
	}

	err = h.laborService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayLaborNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to delete labor")
		return
	}

	response.Success(c, nil)
}

// parseID extracts an int64 ID from the URL path
func (h *WorkDayLaborHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
