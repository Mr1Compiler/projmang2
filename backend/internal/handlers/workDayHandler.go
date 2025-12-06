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

type WorkDayHandler struct {
	workDayService *services.WorkDayService
	validator      *validator.Validate
}

func NewWorkDayHandler(workDayService *services.WorkDayService) *WorkDayHandler {
	return &WorkDayHandler{
		workDayService: workDayService,
		validator:      validator.New(),
	}
}

// GetAll handles GET /workdays
func (h *WorkDayHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	workDays, err := h.workDayService.GetAll(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to fetch work days")
		return
	}

	response.Success(w, workDays)
}

// GetByID handles GET /workdays/{id}
func (h *WorkDayHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid work day id")
		return
	}

	workDay, err := h.workDayService.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to fetch work day")
		return
	}

	response.Success(w, workDay)
}

// GetByProjectID handles GET /projects/{projectId}/workdays
func (h *WorkDayHandler) GetByProjectID(w http.ResponseWriter, r *http.Request) {
	projectID, err := h.parseID(r, "projectId")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid project id")
		return
	}

	workDays, err := h.workDayService.GetByProjectID(r.Context(), projectID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to fetch work days")
		return
	}

	response.Success(w, workDays)
}

// Create handles POST /workdays
func (h *WorkDayHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateWorkDay
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	workDay, err := h.workDayService.Create(r.Context(), req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to create work day")
		return
	}

	response.Created(w, workDay)
}

// Update handles PUT /workdays/{id}
func (h *WorkDayHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid work day id")
		return
	}

	var req dtos.UpdateWorkDay
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	workDay, err := h.workDayService.Update(r.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to update work day")
		return
	}

	response.Success(w, workDay)
}

// Delete handles DELETE /workdays/{id}
func (h *WorkDayHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid work day id")
		return
	}

	err = h.workDayService.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to delete work day")
		return
	}

	response.Success(w, nil)
}

// parseID extracts an int64 ID from the URL path
func (h *WorkDayHandler) parseID(r *http.Request, param string) (int64, error) {
	idStr := r.PathValue(param)
	return strconv.ParseInt(idStr, 10, 64)
}