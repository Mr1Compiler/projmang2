package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/response"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

type WorkSubCategoryHandler struct {
	subCategoryService *services.WorkSubCategoryService
}

func NewWorkSubCategoryHandler(subCategoryService *services.WorkSubCategoryService) *WorkSubCategoryHandler {
	return &WorkSubCategoryHandler{
		subCategoryService: subCategoryService,
	}
}

// GetAll handles GET /work-subcategories
func (h *WorkSubCategoryHandler) GetAll(c *gin.Context) {
	subCategories, err := h.subCategoryService.GetAll(c.Request.Context())
	if err != nil {
		response.InternalError(c, "failed to fetch work subcategories")
		return
	}

	response.Success(c, subCategories)
}

// GetByID handles GET /work-subcategories/:id
func (h *WorkSubCategoryHandler) GetByID(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid work subcategory id")
		return
	}

	subCategory, err := h.subCategoryService.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkSubCategoryNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to fetch work subcategory")
		return
	}

	response.Success(c, subCategory)
}

// GetByCategoryID handles GET /work-categories/:categoryId/subcategories
func (h *WorkSubCategoryHandler) GetByCategoryID(c *gin.Context) {
	categoryID, err := h.parseID(c, "categoryId")
	if err != nil {
		response.BadRequest(c, "invalid category id")
		return
	}

	subCategories, err := h.subCategoryService.GetByCategoryID(c.Request.Context(), categoryID)
	if err != nil {
		response.InternalError(c, "failed to fetch work subcategories")
		return
	}

	response.Success(c, subCategories)
}

// Create handles POST /work-subcategories
func (h *WorkSubCategoryHandler) Create(c *gin.Context) {
	var req dtos.CreateWorkSubCategory
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	subCategory, err := h.subCategoryService.Create(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, "failed to create work subcategory")
		return
	}

	response.Created(c, subCategory)
}

// Update handles PUT /work-subcategories/:id
func (h *WorkSubCategoryHandler) Update(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid work subcategory id")
		return
	}

	var req dtos.UpdateWorkSubCategory
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	subCategory, err := h.subCategoryService.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, services.ErrWorkSubCategoryNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to update work subcategory")
		return
	}

	response.Success(c, subCategory)
}

// Delete handles DELETE /work-subcategories/:id
func (h *WorkSubCategoryHandler) Delete(c *gin.Context) {
	id, err := h.parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid work subcategory id")
		return
	}

	err = h.subCategoryService.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrWorkSubCategoryNotFound) {
			response.NotFound(c, err.Error())
			return
		}
		response.InternalError(c, "failed to delete work subcategory")
		return
	}

	response.Success(c, nil)
}

// parseID extracts an int64 ID from the URL path
func (h *WorkSubCategoryHandler) parseID(c *gin.Context, param string) (int64, error) {
	idStr := c.Param(param)
	return strconv.ParseInt(idStr, 10, 64)
}
