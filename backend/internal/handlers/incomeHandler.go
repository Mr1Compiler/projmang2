package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafa91ameen/prjalgo/backend/internal/dtos"
	"github.com/mustafa91ameen/prjalgo/backend/internal/response"
	"github.com/mustafa91ameen/prjalgo/backend/internal/services"
)

type IncomeHandler struct {
	incomeService *services.IncomeService
}

func NewIncomeHandler(incomeService *services.IncomeService) *IncomeHandler {
	return &IncomeHandler{
		incomeService: incomeService,
	}
}

// GetAll handles GET /incomes
func (h *IncomeHandler) GetAll(c *gin.Context) {
	var pagination dtos.PaginationQuery
	if err := c.ShouldBindQuery(&pagination); err != nil {
		response.ValidationError(c, err)
		return
	}
	pagination.Normalize()

	incomes, err := h.incomeService.GetAll(c.Request.Context(), pagination.Limit, pagination.Offset())
	if err != nil {
		response.InternalError(c, "failed to fetch incomes")
		return
	}

	response.Success(c, incomes)
}

// GetByID handles GET /incomes/:id
func (h *IncomeHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid income id")
		return
	}

	income, err := h.incomeService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrIncomeNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to fetch income")
		return
	}

	response.Success(c, income)
}

// Create handles POST /incomes
func (h *IncomeHandler) Create(c *gin.Context) {
	var req dtos.CreateIncome
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	income, err := h.incomeService.Create(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, "failed to create income")
		return
	}

	response.Created(c, income)
}

// Update handles PUT /incomes/:id
func (h *IncomeHandler) Update(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid income id")
		return
	}

	var req dtos.UpdateIncome
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	income, err := h.incomeService.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrIncomeNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to update income")
		return
	}

	response.Success(c, income)
}

// Delete handles DELETE /incomes/:id
func (h *IncomeHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid income id")
		return
	}

	err = h.incomeService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrIncomeNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to delete income")
		return
	}

	response.Success(c, nil)
}

// GetStats handles GET /income/stats
func (h *IncomeHandler) GetStats(c *gin.Context) {
	var periodQuery dtos.PeriodQuery
	if err := c.ShouldBindQuery(&periodQuery); err != nil {
		response.ValidationError(c, err)
		return
	}
	periodQuery.Normalize()

	stats, err := h.incomeService.GetStats(c.Request.Context(), periodQuery.Period)
	if err != nil {
		if errors.Is(err, services.ErrIncomeStatsNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to fetch income stats")
		return
	}

	response.Success(c, stats)
}

// parseID extracts an int64 ID from the URL path
func (h *IncomeHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
