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

type WorkCategoryHandler struct {
	categoryService *services.WorkCategoryService
	validator       *validator.Validate
}

func NewWorkCategoryHandler(categoryService *services.WorkCategoryService) *WorkCategoryHandler {
	return &WorkCategoryHandler{
		categoryService: categoryService,
		validator:       validator.New(),
	}
}

// GetAll handles GET /work-categories
func (h *WorkCategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.categoryService.GetAll(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to fetch work categories")
		return
	}

	response.Success(w, categories)
}

// GetByID handles GET /work-categories/{id}
func (h *WorkCategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid work category id")
		return
	}

	category, err := h.categoryService.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkCategoryNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to fetch work category")
		return
	}

	response.Success(w, category)
}

// Create handles POST /work-categories
func (h *WorkCategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateWorkCategory
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	category, err := h.categoryService.Create(r.Context(), req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to create work category")
		return
	}

	response.Created(w, category)
}

// Update handles PUT /work-categories/{id}
func (h *WorkCategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid work category id")
		return
	}

	var req dtos.UpdateWorkCategory
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	category, err := h.categoryService.Update(r.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrWorkCategoryNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to update work category")
		return
	}

	response.Success(w, category)
}

// Delete handles DELETE /work-categories/{id}
func (h *WorkCategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid work category id")
		return
	}

	err = h.categoryService.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkCategoryNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to delete work category")
		return
	}

	response.Success(w, nil)
}

// parseID extracts an int64 ID from the URL path
func (h *WorkCategoryHandler) parseID(r *http.Request, param string) (int64, error) {
	idStr := r.PathValue(param)
	return strconv.ParseInt(idStr, 10, 64)
}
