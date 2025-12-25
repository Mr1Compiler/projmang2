package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafa91ameen/prjalgo/backend/internal/auth"
	"github.com/mustafa91ameen/prjalgo/backend/internal/container"
)

// RegisterWorkCategoryRoutes sets up /workCategories and /workSubcategories routes
func RegisterWorkCategoryRoutes(rg *gin.RouterGroup, c *container.Container) {
	// Work Categories - API route is /workCategories but permission check uses /categories (DB page route)
	categories := rg.Group("/workCategories")
	categoriesAuthz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/categories", perm)
	}
	auditCategory := func(action string) gin.HandlerFunc {
		return auth.AuditMiddleware(c.AuditLogService, action, "work_category")
	}
	{
		categories.GET("", categoriesAuthz("read"), c.WorkCategoryHandler.GetAll)
		categories.GET("/stats", categoriesAuthz("read"), c.WorkCategoryHandler.GetStats)
		categories.GET("/:id", categoriesAuthz("read"), c.WorkCategoryHandler.GetByID)
		categories.POST("", categoriesAuthz("create"), auditCategory("create"), c.WorkCategoryHandler.Create)
		categories.PUT("/:id", categoriesAuthz("update"), auditCategory("update"), c.WorkCategoryHandler.Update)
		categories.DELETE("/:id", categoriesAuthz("delete"), auditCategory("delete"), c.WorkCategoryHandler.Delete)
	}

	// Work Sub-Categories
	subCategories := rg.Group("/workSubcategories")
	subCategoriesAuthz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/workSubcategories", perm)
	}
	auditSubCategory := func(action string) gin.HandlerFunc {
		return auth.AuditMiddleware(c.AuditLogService, action, "work_subcategory")
	}
	{
		subCategories.GET("", subCategoriesAuthz("read"), c.WorkSubCategoryHandler.GetAll)
		subCategories.GET("/:id", subCategoriesAuthz("read"), c.WorkSubCategoryHandler.GetByID)
		subCategories.POST("", subCategoriesAuthz("create"), auditSubCategory("create"), c.WorkSubCategoryHandler.Create)
		subCategories.PUT("/:id", subCategoriesAuthz("update"), auditSubCategory("update"), c.WorkSubCategoryHandler.Update)
		subCategories.DELETE("/:id", subCategoriesAuthz("delete"), auditSubCategory("delete"), c.WorkSubCategoryHandler.Delete)
	}
}
