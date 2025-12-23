package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafa91ameen/prjalgo/backend/internal/dtos"
	"github.com/mustafa91ameen/prjalgo/backend/internal/response"
	"github.com/mustafa91ameen/prjalgo/backend/internal/services"
)

type WorkCategoryHandler struct {
	categoryService *services.WorkCategoryService
}

func NewWorkCategoryHandler(categoryService *services.WorkCategoryService) *WorkCategoryHandler {
	return &WorkCategoryHandler{
		categoryService: categoryService,
	}
}

// GetAll handles GET /work-categories
func (h *WorkCategoryHandler) GetAll(c *gin.Context) {
	var pagination dtos.PaginationQuery
	if err := c.ShouldBindQuery(&pagination); err != nil {
		response.ValidationError(c, err)
		return
	}
	pagination.Normalize()

	categories, err := h.categoryService.GetAll(c.Request.Context(), pagination.Limit, pagination.Offset())
	if err != nil {
		response.InternalError(c, "failed to fetch work categories")
		return
	}

	response.Success(c, categories)
}

// GetByID handles GET /work-categories/:id
func (h *WorkCategoryHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid work category id")
		return
	}

	category, err := h.categoryService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkCategoryNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to fetch work category")
		return
	}

	response.Success(c, category)
}

// Create handles POST /work-categories
func (h *WorkCategoryHandler) Create(c *gin.Context) {
	var req dtos.CreateWorkCategory
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	category, err := h.categoryService.Create(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, "failed to create work category")
		return
	}

	response.Created(c, category)
}

// Update handles PUT /work-categories/:id
func (h *WorkCategoryHandler) Update(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid work category id")
		return
	}

	var req dtos.UpdateWorkCategory
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	category, err := h.categoryService.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrWorkCategoryNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to update work category")
		return
	}

	response.Success(c, category)
}

// Delete handles DELETE /work-categories/:id
func (h *WorkCategoryHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid work category id")
		return
	}

	err = h.categoryService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkCategoryNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to delete work category")
		return
	}

	response.Success(c, nil)
}

// GetStats handles GET /work-categories/stats
func (h *WorkCategoryHandler) GetStats(c *gin.Context) {
	var periodQuery dtos.PeriodQuery
	if err := c.ShouldBindQuery(&periodQuery); err != nil {
		response.ValidationError(c, err)
		return
	}
	periodQuery.Normalize()

	stats, err := h.categoryService.GetStats(c.Request.Context(), periodQuery.Period)
	if err != nil {
		if errors.Is(err, services.ErrWorkCategoryStatsNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to fetch work category stats")
		return
	}

	response.Success(c, stats)
}

// parseID extracts an int64 ID from the URL path
func (h *WorkCategoryHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
