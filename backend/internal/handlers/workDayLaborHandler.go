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

type WorkDayLaborHandler struct {
	laborService *services.WorkDayLaborService
	validator    *validator.Validate
}

func NewWorkDayLaborHandler(laborService *services.WorkDayLaborService) *WorkDayLaborHandler {
	return &WorkDayLaborHandler{
		laborService: laborService,
		validator:    validator.New(),
	}
}

// GetByWorkDayID handles GET /workdays/{workDayId}/labors
func (h *WorkDayLaborHandler) GetByWorkDayID(w http.ResponseWriter, r *http.Request) {
	workDayID, err := h.parseID(r, "workDayId")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid work day id")
		return
	}

	labors, err := h.laborService.GetByWorkDayID(r.Context(), workDayID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to fetch labors")
		return
	}

	response.Success(w, labors)
}

// GetByID handles GET /workdays/{workDayId}/labors/{id}
func (h *WorkDayLaborHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid labor id")
		return
	}

	labor, err := h.laborService.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayLaborNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to fetch labor")
		return
	}

	response.Success(w, labor)
}

// Create handles POST /workdays/{workDayId}/labors
func (h *WorkDayLaborHandler) Create(w http.ResponseWriter, r *http.Request) {
	workDayID, err := h.parseID(r, "workDayId")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid work day id")
		return
	}

	var req dtos.CreateWorkDayLabor
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// Set workDayId from URL path
	req.WorkDayID = workDayID

	if err := h.validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	labor, err := h.laborService.Create(r.Context(), req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to create labor")
		return
	}

	response.Created(w, labor)
}

// Update handles PUT /workdays/{workDayId}/labors/{id}
func (h *WorkDayLaborHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid labor id")
		return
	}

	var req dtos.UpdateWorkDayLabor
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	labor, err := h.laborService.Update(r.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayLaborNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to update labor")
		return
	}

	response.Success(w, labor)
}

// Delete handles DELETE /workdays/{workDayId}/labors/{id}
func (h *WorkDayLaborHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid labor id")
		return
	}

	err = h.laborService.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayLaborNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to delete labor")
		return
	}

	response.Success(w, nil)
}

// parseID extracts an int64 ID from the URL path
func (h *WorkDayLaborHandler) parseID(r *http.Request, param string) (int64, error) {
	idStr := r.PathValue(param)
	return strconv.ParseInt(idStr, 10, 64)
}