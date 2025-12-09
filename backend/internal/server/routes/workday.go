package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/auth"
	"github.com/mustafaameen91/project-managment/backend/internal/container"
)

// RegisterWorkDayRoutes sets up /workdays and related routes
func RegisterWorkDayRoutes(rg *gin.RouterGroup, c *container.Container) {
	// WorkDay
	workDays := rg.Group("/workdays")
	workDaysAuthz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/workdays", perm)
	}
	auditWorkDay := func(action string) gin.HandlerFunc {
		return auth.AuditMiddleware(c.AuditLogService, action, "workday")
	}
	{
		workDays.GET("", workDaysAuthz("read"), c.WorkDayHandler.GetAll)
		workDays.GET("/:id", workDaysAuthz("read"), c.WorkDayHandler.GetByID)
		workDays.POST("", workDaysAuthz("write"), auditWorkDay("create"), c.WorkDayHandler.Create)
		workDays.PUT("/:id", workDaysAuthz("write"), auditWorkDay("update"), c.WorkDayHandler.Update)
		workDays.DELETE("/:id", workDaysAuthz("delete"), auditWorkDay("delete"), c.WorkDayHandler.Delete)
	}

	// WorkDay Materials
	materials := rg.Group("/workday-materials")
	materialsAuthz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/workday-materials", perm)
	}
	auditMaterial := func(action string) gin.HandlerFunc {
		return auth.AuditMiddleware(c.AuditLogService, action, "workday_material")
	}
	{
		materials.GET("", materialsAuthz("read"), c.WorkDayMaterialHandler.GetAll)
		materials.GET("/:id", materialsAuthz("read"), c.WorkDayMaterialHandler.GetByID)
		materials.POST("", materialsAuthz("write"), auditMaterial("create"), c.WorkDayMaterialHandler.Create)
		materials.PUT("/:id", materialsAuthz("write"), auditMaterial("update"), c.WorkDayMaterialHandler.Update)
		materials.DELETE("/:id", materialsAuthz("delete"), auditMaterial("delete"), c.WorkDayMaterialHandler.Delete)
	}

	// WorkDay Labor
	labor := rg.Group("/workday-labor")
	laborAuthz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/workday-labor", perm)
	}
	auditLabor := func(action string) gin.HandlerFunc {
		return auth.AuditMiddleware(c.AuditLogService, action, "workday_labor")
	}
	{
		labor.GET("", laborAuthz("read"), c.WorkDayLaborHandler.GetAll)
		labor.GET("/:id", laborAuthz("read"), c.WorkDayLaborHandler.GetByID)
		labor.POST("", laborAuthz("write"), auditLabor("create"), c.WorkDayLaborHandler.Create)
		labor.PUT("/:id", laborAuthz("write"), auditLabor("update"), c.WorkDayLaborHandler.Update)
		labor.DELETE("/:id", laborAuthz("delete"), auditLabor("delete"), c.WorkDayLaborHandler.Delete)
	}

	// WorkDay Equipment
	equipment := rg.Group("/workday-equipment")
	equipmentAuthz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/workday-equipment", perm)
	}
	auditEquipment := func(action string) gin.HandlerFunc {
		return auth.AuditMiddleware(c.AuditLogService, action, "workday_equipment")
	}
	{
		equipment.GET("", equipmentAuthz("read"), c.WorkDayEquipmentHandler.GetAll)
		equipment.GET("/:id", equipmentAuthz("read"), c.WorkDayEquipmentHandler.GetByID)
		equipment.POST("", equipmentAuthz("write"), auditEquipment("create"), c.WorkDayEquipmentHandler.Create)
		equipment.PUT("/:id", equipmentAuthz("write"), auditEquipment("update"), c.WorkDayEquipmentHandler.Update)
		equipment.DELETE("/:id", equipmentAuthz("delete"), auditEquipment("delete"), c.WorkDayEquipmentHandler.Delete)
	}
}
