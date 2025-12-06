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

type WorkSubCategoryHandler struct {
	subCategoryService *services.WorkSubCategoryService
	validator          *validator.Validate
}

func NewWorkSubCategoryHandler(subCategoryService *services.WorkSubCategoryService) *WorkSubCategoryHandler {
	return &WorkSubCategoryHandler{
		subCategoryService: subCategoryService,
		validator:          validator.New(),
	}
}

// GetAll handles GET /work-subcategories
func (h *WorkSubCategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	subCategories, err := h.subCategoryService.GetAll(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to fetch work subcategories")
		return
	}

	response.Success(w, subCategories)
}

// GetByID handles GET /work-subcategories/{id}
func (h *WorkSubCategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid work subcategory id")
		return
	}

	subCategory, err := h.subCategoryService.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkSubCategoryNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to fetch work subcategory")
		return
	}

	response.Success(w, subCategory)
}

// GetByCategoryID handles GET /work-categories/{categoryId}/subcategories
func (h *WorkSubCategoryHandler) GetByCategoryID(w http.ResponseWriter, r *http.Request) {
	categoryID, err := h.parseID(r, "categoryId")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid category id")
		return
	}

	subCategories, err := h.subCategoryService.GetByCategoryID(r.Context(), categoryID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to fetch work subcategories")
		return
	}

	response.Success(w, subCategories)
}

// Create handles POST /work-subcategories
func (h *WorkSubCategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateWorkSubCategory
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	subCategory, err := h.subCategoryService.Create(r.Context(), req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to create work subcategory")
		return
	}

	response.Created(w, subCategory)
}

// Update handles PUT /work-subcategories/{id}
func (h *WorkSubCategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid work subcategory id")
		return
	}

	var req dtos.UpdateWorkSubCategory
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	subCategory, err := h.subCategoryService.Update(r.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrWorkSubCategoryNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to update work subcategory")
		return
	}

	response.Success(w, subCategory)
}

// Delete handles DELETE /work-subcategories/{id}
func (h *WorkSubCategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid work subcategory id")
		return
	}

	err = h.subCategoryService.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkSubCategoryNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to delete work subcategory")
		return
	}

	response.Success(w, nil)
}

// parseID extracts an int64 ID from the URL path
func (h *WorkSubCategoryHandler) parseID(r *http.Request, param string) (int64, error) {
	idStr := r.PathValue(param)
	return strconv.ParseInt(idStr, 10, 64)
}
