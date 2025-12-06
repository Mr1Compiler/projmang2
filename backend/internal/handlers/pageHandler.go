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

type PageHandler struct {
	pageService *services.PageService
	validator   *validator.Validate
}

func NewPageHandler(pageService *services.PageService) *PageHandler {
	return &PageHandler{
		pageService: pageService,
		validator:   validator.New(),
	}
}

// GetAll handles GET /pages
func (h *PageHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	pages, err := h.pageService.GetAll(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to fetch pages")
		return
	}

	response.Success(w, pages)
}

// GetByID handles GET /pages/{id}
func (h *PageHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid page id")
		return
	}

	page, err := h.pageService.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrPageNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to fetch page")
		return
	}

	response.Success(w, page)
}

// GetActivePages handles GET /pages/active
func (h *PageHandler) GetActivePages(w http.ResponseWriter, r *http.Request) {
	pages, err := h.pageService.GetActivePages(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to fetch active pages")
		return
	}

	response.Success(w, pages)
}

// Create handles POST /pages
func (h *PageHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreatePage
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	page, err := h.pageService.Create(r.Context(), req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to create page")
		return
	}

	response.Created(w, page)
}

// Update handles PUT /pages/{id}
func (h *PageHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid page id")
		return
	}

	var req dtos.UpdatePage
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	page, err := h.pageService.Update(r.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrPageNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to update page")
		return
	}

	response.Success(w, page)
}

// Delete handles DELETE /pages/{id}
func (h *PageHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseID(r, "id")
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid page id")
		return
	}

	err = h.pageService.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrPageNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to delete page")
		return
	}

	response.Success(w, nil)
}

// parseID extracts an int64 ID from the URL path
func (h *PageHandler) parseID(r *http.Request, param string) (int64, error) {
	idStr := r.PathValue(param)
	return strconv.ParseInt(idStr, 10, 64)
}
