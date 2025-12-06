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

type ExpenseHandler struct {
	expenseService *services.ExpenseService
	validator      *validator.Validate
}

func NewExpenseHandler(expenseService *services.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{
		expenseService: expenseService,
		validator:      validator.New(),
	}
}

// GetAll handles GET /expenses
func (h *ExpenseHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	expenses, err := h.expenseService.GetAll(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to fetch expenses")
		return
	}

	response.Success(w, expenses)
}

// GetByID handles GET /expenses/{id}
func (h *ExpenseHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid expense id")
		return
	}

	expense, err := h.expenseService.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrExpenseNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to fetch expense")
		return
	}

	response.Success(w, expense)
}

// GetByProjectID handles GET /projects/{projectId}/expenses
func (h *ExpenseHandler) GetByProjectID(w http.ResponseWriter, r *http.Request) {
	projectID, err := h.parseID(r, "projectId")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid project id")
		return
	}

	expenses, err := h.expenseService.GetByProjectID(r.Context(), projectID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to fetch expenses")
		return
	}

	response.Success(w, expenses)
}

// Create handles POST /expenses
func (h *ExpenseHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateExpense
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	expense, err := h.expenseService.Create(r.Context(), req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to create expense")
		return
	}

	response.Created(w, expense)
}

// Update handles PUT /expenses/{id}
func (h *ExpenseHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid expense id")
		return
	}

	var req dtos.UpdateExpense
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	expense, err := h.expenseService.Update(r.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrExpenseNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to update expense")
		return
	}

	response.Success(w, expense)
}

// Delete handles DELETE /expenses/{id}
func (h *ExpenseHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid expense id")
		return
	}

	err = h.expenseService.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrExpenseNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to delete expense")
		return
	}

	response.Success(w, nil)
}

// parseID extracts an int64 ID from the URL path
func (h *ExpenseHandler) parseID(r *http.Request, param string) (int64, error) {
	idStr := r.PathValue(param)
	return strconv.ParseInt(idStr, 10, 64)
}
