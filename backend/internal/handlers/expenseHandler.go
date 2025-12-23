package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafa91ameen/prjalgo/backend/internal/dtos"
	"github.com/mustafa91ameen/prjalgo/backend/internal/response"
	"github.com/mustafa91ameen/prjalgo/backend/internal/services"
)

type ExpenseHandler struct {
	expenseService *services.ExpenseService
}

func NewExpenseHandler(expenseService *services.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{
		expenseService: expenseService,
	}
}

// GetAll handles GET /expenses
func (h *ExpenseHandler) GetAll(c *gin.Context) {
	var pagination dtos.PaginationQuery
	if err := c.ShouldBindQuery(&pagination); err != nil {
		response.ValidationError(c, err)
		return
	}
	pagination.Normalize()

	expenses, err := h.expenseService.GetAll(c.Request.Context(), pagination.Limit, pagination.Offset())
	if err != nil {
		response.InternalError(c, "failed to fetch expenses")
		return
	}

	response.Success(c, expenses)
}

// GetByID handles GET /expenses/:id
func (h *ExpenseHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid expense id")
		return
	}

	expense, err := h.expenseService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrExpenseNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to fetch expense")
		return
	}

	response.Success(c, expense)
}

// GetByProjectID handles GET /projects/:id/expenses
func (h *ExpenseHandler) GetByProjectID(c *gin.Context) {
	projectID, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid project id")
		return
	}

	expenses, err := h.expenseService.GetByProjectID(c.Request.Context(), projectID)
	if err != nil {
		response.InternalError(c, "failed to fetch expenses")
		return
	}

	response.Success(c, expenses)
}

// Create handles POST /expenses
func (h *ExpenseHandler) Create(c *gin.Context) {
	var req dtos.CreateExpense
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	expense, err := h.expenseService.Create(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, "failed to create expense")
		return
	}

	response.Created(c, expense)
}

// Update handles PUT /expenses/:id
func (h *ExpenseHandler) Update(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid expense id")
		return
	}

	var req dtos.UpdateExpense
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	expense, err := h.expenseService.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrExpenseNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to update expense")
		return
	}

	response.Success(c, expense)
}

// Delete handles DELETE /expenses/:id
func (h *ExpenseHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid expense id")
		return
	}

	err = h.expenseService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrExpenseNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to delete expense")
		return
	}

	response.Success(c, nil)
}

// GetStats handles GET /expenses/stats
func (h *ExpenseHandler) GetStats(c *gin.Context) {
	var periodQuery dtos.PeriodQuery
	if err := c.ShouldBindQuery(&periodQuery); err != nil {
		response.ValidationError(c, err)
		return
	}
	periodQuery.Normalize()

	stats, err := h.expenseService.GetStats(c.Request.Context(), periodQuery.Period)
	if err != nil {
		if errors.Is(err, services.ErrExpenseStatsNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to fetch expense stats")
		return
	}

	response.Success(c, stats)
}

// parseID extracts an int64 ID from the URL path
func (h *ExpenseHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
