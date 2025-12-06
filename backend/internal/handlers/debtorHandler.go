package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/response"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

type DebtorHandler struct {
	debtorService *services.DebtorService
}

func NewDebtorHandler(debtorService *services.DebtorService) *DebtorHandler {
	return &DebtorHandler{
		debtorService: debtorService,
	}
}

// GetAll handles GET /debtors
func (h *DebtorHandler) GetAll(c *gin.Context) {
	debtors, err := h.debtorService.GetAll(c.Request.Context())
	if err != nil {
		response.InternalError(c, "failed to fetch debtors")
		return
	}

	response.Success(c, debtors)
}

// GetByID handles GET /debtors/:id
func (h *DebtorHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid debtor id")
		return
	}

	debtor, err := h.debtorService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrDebtorNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to fetch debtor")
		return
	}

	response.Success(c, debtor)
}

// Create handles POST /debtors
func (h *DebtorHandler) Create(c *gin.Context) {
	var req dtos.CreateDebtor
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	debtor, err := h.debtorService.Create(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, "failed to create debtor")
		return
	}

	response.Created(c, debtor)
}

// Update handles PUT /debtors/:id
func (h *DebtorHandler) Update(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid debtor id")
		return
	}

	var req dtos.UpdateDebtor
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	debtor, err := h.debtorService.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrDebtorNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to update debtor")
		return
	}

	response.Success(c, debtor)
}

// Delete handles DELETE /debtors/:id
func (h *DebtorHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid debtor id")
		return
	}

	err = h.debtorService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrDebtorNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to delete debtor")
		return
	}

	response.Success(c, nil)
}

// parseID extracts an int64 ID from the URL path
func (h *DebtorHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
