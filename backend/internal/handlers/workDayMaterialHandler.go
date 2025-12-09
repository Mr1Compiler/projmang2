package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/response"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

type WorkDayMaterialHandler struct {
	materialService *services.WorkDayMaterialService
}

func NewWorkDayMaterialHandler(materialService *services.WorkDayMaterialService) *WorkDayMaterialHandler {
	return &WorkDayMaterialHandler{
		materialService: materialService,
	}
}

// GetAll handles GET /workday-materials
func (h *WorkDayMaterialHandler) GetAll(c *gin.Context) {
	var pagination dtos.PaginationQuery
	if err := c.ShouldBindQuery(&pagination); err != nil {
		response.ValidationError(c, err)
		return
	}
	pagination.Normalize()

	materials, err := h.materialService.GetAll(c.Request.Context(), pagination.Limit, pagination.Offset())
	if err != nil {
		response.InternalError(c, "failed to fetch materials")
		return
	}

	response.Success(c, materials)
}

// GetByWorkDayID handles GET /workdays/:workDayId/materials
func (h *WorkDayMaterialHandler) GetByWorkDayID(c *gin.Context) {
	workDayID, err := h.parseID(c, "workDayId")
	if err != nil {
		response.BadRequest(c, "invalid work day id")
		return
	}

	materials, err := h.materialService.GetByWorkDayID(c.Request.Context(), workDayID)
	if err != nil {
		response.InternalError(c, "failed to fetch materials")
		return
	}

	response.Success(c, materials)
}

// GetByID handles GET /workdays/:workDayId/materials/:id
func (h *WorkDayMaterialHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid material id")
		return
	}

	material, err := h.materialService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayMaterialNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to fetch material")
		return
	}

	response.Success(c, material)
}

// Create handles POST /workday-materials
func (h *WorkDayMaterialHandler) Create(c *gin.Context) {
	var req dtos.CreateWorkDayMaterial
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	material, err := h.materialService.Create(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, "failed to create material")
		return
	}

	response.Created(c, material)
}

// Update handles PUT /workdays/:workDayId/materials/:id
func (h *WorkDayMaterialHandler) Update(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid material id")
		return
	}

	var req dtos.UpdateWorkDayMaterial
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	material, err := h.materialService.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayMaterialNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to update material")
		return
	}

	response.Success(c, material)
}

// Delete handles DELETE /workdays/:workDayId/materials/:id
func (h *WorkDayMaterialHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid material id")
		return
	}

	err = h.materialService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayMaterialNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to delete material")
		return
	}

	response.Success(c, nil)
}

// parseID extracts an int64 ID from the URL path
func (h *WorkDayMaterialHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
