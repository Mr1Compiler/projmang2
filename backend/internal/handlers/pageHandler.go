package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/response"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

type PageHandler struct {
	pageService *services.PageService
}

func NewPageHandler(pageService *services.PageService) *PageHandler {
	return &PageHandler{
		pageService: pageService,
	}
}

// GetAll handles GET /pages
func (h *PageHandler) GetAll(c *gin.Context) {
	pages, err := h.pageService.GetAll(c.Request.Context())
	if err != nil {
		response.InternalError(c, "failed to fetch pages")
		return
	}

	response.Success(c, pages)
}

// GetByID handles GET /pages/:id
func (h *PageHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid page id")
		return
	}

	page, err := h.pageService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrPageNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to fetch page")
		return
	}

	response.Success(c, page)
}

// GetActivePages handles GET /pages/active
func (h *PageHandler) GetActivePages(c *gin.Context) {
	pages, err := h.pageService.GetActivePages(c.Request.Context())
	if err != nil {
		response.InternalError(c, "failed to fetch active pages")
		return
	}

	response.Success(c, pages)
}

// Create handles POST /pages
func (h *PageHandler) Create(c *gin.Context) {
	var req dtos.CreatePage
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	page, err := h.pageService.Create(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, "failed to create page")
		return
	}

	response.Created(c, page)
}

// Update handles PUT /pages/:id
func (h *PageHandler) Update(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid page id")
		return
	}

	var req dtos.UpdatePage
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	page, err := h.pageService.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrPageNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to update page")
		return
	}

	response.Success(c, page)
}

// Delete handles DELETE /pages/:id
func (h *PageHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid page id")
		return
	}

	err = h.pageService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrPageNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to delete page")
		return
	}

	response.Success(c, nil)
}

// parseID extracts an int64 ID from the URL path
func (h *PageHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
