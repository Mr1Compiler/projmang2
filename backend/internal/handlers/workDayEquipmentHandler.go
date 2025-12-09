package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/response"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

type WorkDayEquipmentHandler struct {
	equipmentService *services.WorkDayEquipmentService
}

func NewWorkDayEquipmentHandler(equipmentService *services.WorkDayEquipmentService) *WorkDayEquipmentHandler {
	return &WorkDayEquipmentHandler{
		equipmentService: equipmentService,
	}
}

// GetAll handles GET /workday-equipment
func (h *WorkDayEquipmentHandler) GetAll(c *gin.Context) {
	var pagination dtos.PaginationQuery
	if err := c.ShouldBindQuery(&pagination); err != nil {
		response.ValidationError(c, err)
		return
	}
	pagination.Normalize()

	equipments, err := h.equipmentService.GetAll(c.Request.Context(), pagination.Limit, pagination.Offset())
	if err != nil {
		response.InternalError(c, "failed to fetch equipments")
		return
	}

	response.Success(c, equipments)
}

// GetByWorkDayID handles GET /workdays/:workDayId/equipments
func (h *WorkDayEquipmentHandler) GetByWorkDayID(c *gin.Context) {
	workDayID, err := h.parseID(c, "workDayId")
	if err != nil {
		response.BadRequest(c, "invalid work day id")
		return
	}

	equipments, err := h.equipmentService.GetByWorkDayID(c.Request.Context(), workDayID)
	if err != nil {
		response.InternalError(c, "failed to fetch equipments")
		return
	}

	response.Success(c, equipments)
}

// GetByID handles GET /workdays/:workDayId/equipments/:id
func (h *WorkDayEquipmentHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid equipment id")
		return
	}

	equipment, err := h.equipmentService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayEquipmentNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to fetch equipment")
		return
	}

	response.Success(c, equipment)
}

// Create handles POST /workday-equipment
func (h *WorkDayEquipmentHandler) Create(c *gin.Context) {
	var req dtos.CreateWorkDayEquipment
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	equipment, err := h.equipmentService.Create(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, "failed to create equipment")
		return
	}

	response.Created(c, equipment)
}

// Update handles PUT /workdays/:workDayId/equipments/:id
func (h *WorkDayEquipmentHandler) Update(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid equipment id")
		return
	}

	var req dtos.UpdateWorkDayEquipment
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	equipment, err := h.equipmentService.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayEquipmentNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to update equipment")
		return
	}

	response.Success(c, equipment)
}

// Delete handles DELETE /workdays/:workDayId/equipments/:id
func (h *WorkDayEquipmentHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid equipment id")
		return
	}

	err = h.equipmentService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayEquipmentNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to delete equipment")
		return
	}

	response.Success(c, nil)
}

// parseID extracts an int64 ID from the URL path
func (h *WorkDayEquipmentHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
