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

type WorkDayMaterialHandler struct {
	materialService *services.WorkDayMaterialService
	validator       *validator.Validate
}

func NewWorkDayMaterialHandler(materialService *services.WorkDayMaterialService) *WorkDayMaterialHandler {
	return &WorkDayMaterialHandler{
		materialService: materialService,
		validator:       validator.New(),
	}
}

// GetByWorkDayID handles GET /workdays/{workDayId}/materials
func (h *WorkDayMaterialHandler) GetByWorkDayID(w http.ResponseWriter, r *http.Request) {
	workDayID, err := h.parseID(r, "workDayId")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid work day id")
		return
	}

	materials, err := h.materialService.GetByWorkDayID(r.Context(), workDayID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to fetch materials")
		return
	}

	response.Success(w, materials)
}

// GetByID handles GET /workdays/{workDayId}/materials/{id}
func (h *WorkDayMaterialHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid material id")
		return
	}

	material, err := h.materialService.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayMaterialNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to fetch material")
		return
	}

	response.Success(w, material)
}

// Create handles POST /workdays/{workDayId}/materials
func (h *WorkDayMaterialHandler) Create(w http.ResponseWriter, r *http.Request) {
	workDayID, err := h.parseID(r, "workDayId")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid work day id")
		return
	}

	var req dtos.CreateWorkDayMaterial
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

	material, err := h.materialService.Create(r.Context(), req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to create material")
		return
	}

	response.Created(w, material)
}

// Update handles PUT /workdays/{workDayId}/materials/{id}
func (h *WorkDayMaterialHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid material id")
		return
	}

	var req dtos.UpdateWorkDayMaterial
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	material, err := h.materialService.Update(r.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayMaterialNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to update material")
		return
	}

	response.Success(w, material)
}

// Delete handles DELETE /workdays/{workDayId}/materials/{id}
func (h *WorkDayMaterialHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid material id")
		return
	}

	err = h.materialService.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkDayMaterialNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to delete material")
		return
	}

	response.Success(w, nil)
}

// parseID extracts an int64 ID from the URL path
func (h *WorkDayMaterialHandler) parseID(r *http.Request, param string) (int64, error) {
	idStr := r.PathValue(param)
	return strconv.ParseInt(idStr, 10, 64)
}