package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/auth"
	"github.com/mustafaameen91/project-managment/backend/internal/container"
)

// RegisterWorkCategoryRoutes sets up /work-categories and /work-subcategories routes
func RegisterWorkCategoryRoutes(rg *gin.RouterGroup, c *container.Container) {
	// Work Categories
	categories := rg.Group("/work-categories")
	categoriesAuthz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/work-categories", perm)
	}
	{
		categories.GET("", categoriesAuthz("read"), c.WorkCategoryHandler.GetAll)
		categories.GET("/:id", categoriesAuthz("read"), c.WorkCategoryHandler.GetByID)
		categories.POST("", categoriesAuthz("write"), c.WorkCategoryHandler.Create)
		categories.PUT("/:id", categoriesAuthz("write"), c.WorkCategoryHandler.Update)
		categories.DELETE("/:id", categoriesAuthz("delete"), c.WorkCategoryHandler.Delete)
	}

	// Work Sub-Categories
	subCategories := rg.Group("/work-subcategories")
	subCategoriesAuthz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/work-subcategories", perm)
	}
	{
		subCategories.GET("", subCategoriesAuthz("read"), c.WorkSubCategoryHandler.GetAll)
		subCategories.GET("/:id", subCategoriesAuthz("read"), c.WorkSubCategoryHandler.GetByID)
		subCategories.POST("", subCategoriesAuthz("write"), c.WorkSubCategoryHandler.Create)
		subCategories.PUT("/:id", subCategoriesAuthz("write"), c.WorkSubCategoryHandler.Update)
		subCategories.DELETE("/:id", subCategoriesAuthz("delete"), c.WorkSubCategoryHandler.Delete)
	}
}
