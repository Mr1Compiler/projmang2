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

type DebtorHandler struct {
	debtorService *services.DebtorService
	validator     *validator.Validate
}

func NewDebtorHandler(debtorService *services.DebtorService) *DebtorHandler {
	return &DebtorHandler{
		debtorService: debtorService,
		validator:     validator.New(),
	}
}

// GetAll handles GET /debtors
func (h *DebtorHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	debtors, err := h.debtorService.GetAll(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to fetch debtors")
		return
	}

	response.Success(w, debtors)
}

// GetByID handles GET /debtors/{id}
func (h *DebtorHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid debtor id")
		return
	}

	debtor, err := h.debtorService.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrDebtorNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to fetch debtor")
		return
	}

	response.Success(w, debtor)
}

// Create handles POST /debtors
func (h *DebtorHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateDebtor
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	debtor, err := h.debtorService.Create(r.Context(), req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to create debtor")
		return
	}

	response.Created(w, debtor)
}

// Update handles PUT /debtors/{id}
func (h *DebtorHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid debtor id")
		return
	}

	var req dtos.UpdateDebtor
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	debtor, err := h.debtorService.Update(r.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrDebtorNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to update debtor")
		return
	}

	response.Success(w, debtor)
}

// Delete handles DELETE /debtors/{id}
func (h *DebtorHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid debtor id")
		return
	}

	err = h.debtorService.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrDebtorNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to delete debtor")
		return
	}

	response.Success(w, nil)
}

// parseID extracts an int64 ID from the URL path
func (h *DebtorHandler) parseID(r *http.Request, param string) (int64, error) {
	idStr := r.PathValue(param)
	return strconv.ParseInt(idStr, 10, 64)
}
