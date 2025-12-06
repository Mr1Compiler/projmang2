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

type WorkDayEquipmentHandler struct {
	equipmentService *services.WorkDayEquipmentService
	validator        *validator.Validate
}

func NewWorkDayEquipmentHandler(equipmentService *services.WorkDayEquipmentService) *WorkDayEquipmentHandler {
	return &WorkDayEquipmentHandler{
		equipmentService: equipmentService,
		validator:        validator.New(),
	}
}

// GetByWorkDayID handles GET /workdays/{workDayId}/equipments
func (h *WorkDayEquipmentHandler) GetByWorkDayID(w http.ResponseWriter, r *http.Request) {
	workDayID, err := h.parseID(r, "workDayId")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid work day id")
		return
	}

	equipments, err := h.equipmentService.GetByWorkDayID(r.Context(), workDayID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to fetch equipments")
		return
	}

	response.Success(w, equipments)
}

// GetByID handles GET /workdays/{workDayId}/equipments/{id}
func (h *WorkDayEquipmentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid equipment id")
		return
	}

	equipment, err := h.equipmentService.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayEquipmentNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to fetch equipment")
		return
	}

	response.Success(w, equipment)
}

// Create handles POST /workdays/{workDayId}/equipments
func (h *WorkDayEquipmentHandler) Create(w http.ResponseWriter, r *http.Request) {
	workDayID, err := h.parseID(r, "workDayId")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid work day id")
		return
	}

	var req dtos.CreateWorkDayEquipment
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

	equipment, err := h.equipmentService.Create(r.Context(), req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to create equipment")
		return
	}

	response.Created(w, equipment)
}

// Update handles PUT /workdays/{workDayId}/equipments/{id}
func (h *WorkDayEquipmentHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid equipment id")
		return
	}

	var req dtos.UpdateWorkDayEquipment
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	equipment, err := h.equipmentService.Update(r.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayEquipmentNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to update equipment")
		return
	}

	response.Success(w, equipment)
}

// Delete handles DELETE /workdays/{workDayId}/equipments/{id}
func (h *WorkDayEquipmentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid equipment id")
		return
	}

	err = h.equipmentService.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayEquipmentNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to delete equipment")
		return
	}

	response.Success(w, nil)
}

// parseID extracts an int64 ID from the URL path
func (h *WorkDayEquipmentHandler) parseID(r *http.Request, param string) (int64, error) {
	idStr := r.PathValue(param)
	return strconv.ParseInt(idStr, 10, 64)
}