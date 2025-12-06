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

type IncomeHandler struct {
	incomeService *services.IncomeService
	validator     *validator.Validate
}

func NewIncomeHandler(incomeService *services.IncomeService) *IncomeHandler {
	return &IncomeHandler{
		incomeService: incomeService,
		validator:     validator.New(),
	}
}

// GetAll handles GET /incomes
func (h *IncomeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	incomes, err := h.incomeService.GetAll(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to fetch incomes")
		return
	}

	response.Success(w, incomes)
}

// GetByID handles GET /incomes/{id}
func (h *IncomeHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid income id")
		return
	}

	income, err := h.incomeService.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrIncomeNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to fetch income")
		return
	}

	response.Success(w, income)
}

// Create handles POST /incomes
func (h *IncomeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateIncome
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	income, err := h.incomeService.Create(r.Context(), req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to create income")
		return
	}

	response.Created(w, income)
}

// Update handles PUT /incomes/{id}
func (h *IncomeHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid income id")
		return
	}

	var req dtos.UpdateIncome
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	income, err := h.incomeService.Update(r.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrIncomeNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to update income")
		return
	}

	response.Success(w, income)
}

// Delete handles DELETE /incomes/{id}
func (h *IncomeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid income id")
		return
	}

	err = h.incomeService.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrIncomeNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to delete income")
		return
	}

	response.Success(w, nil)
}

// parseID extracts an int64 ID from the URL path
func (h *IncomeHandler) parseID(r *http.Request, param string) (int64, error) {
	idStr := r.PathValue(param)
	return strconv.ParseInt(idStr, 10, 64)
}
