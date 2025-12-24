package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafa91ameen/prjalgo/backend/internal/auth"
	"github.com/mustafa91ameen/prjalgo/backend/internal/container"
)

// RegisterFinanceRoutes sets up /expenses and /income routes
func RegisterFinanceRoutes(rg *gin.RouterGroup, c *container.Container) {
	// Expenses
	expenses := rg.Group("/expenses")
	expensesAuthz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/expenses", perm)
	}
	auditExpense := func(action string) gin.HandlerFunc {
		return auth.AuditMiddleware(c.AuditLogService, action, "expense")
	}
	{
		expenses.GET("", expensesAuthz("read"), c.ExpenseHandler.GetAll)
		expenses.GET("/stats", expensesAuthz("read"), c.ExpenseHandler.GetStats)
		expenses.GET("/:id", expensesAuthz("read"), c.ExpenseHandler.GetByID)
		expenses.POST("", expensesAuthz("create"), auditExpense("create"), c.ExpenseHandler.Create)
		expenses.PUT("/:id", expensesAuthz("update"), auditExpense("update"), c.ExpenseHandler.Update)
		expenses.DELETE("/:id", expensesAuthz("delete"), auditExpense("delete"), c.ExpenseHandler.Delete)
	}

	// Income
	income := rg.Group("/income")
	incomeAuthz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/income", perm)
	}
	auditIncome := func(action string) gin.HandlerFunc {
		return auth.AuditMiddleware(c.AuditLogService, action, "income")
	}
	{
		income.GET("", incomeAuthz("read"), c.IncomeHandler.GetAll)
		income.GET("/stats", incomeAuthz("read"), c.IncomeHandler.GetStats)
		income.GET("/:id", incomeAuthz("read"), c.IncomeHandler.GetByID)
		income.POST("", incomeAuthz("create"), auditIncome("create"), c.IncomeHandler.Create)
		income.PUT("/:id", incomeAuthz("update"), auditIncome("update"), c.IncomeHandler.Update)
		income.DELETE("/:id", incomeAuthz("delete"), auditIncome("delete"), c.IncomeHandler.Delete)
	}
}
