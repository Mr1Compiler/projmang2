package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafaameen91/project-managment/backend/internal/auth"
	"github.com/mustafaameen91/project-managment/backend/internal/container"
)

// RegisterFinanceRoutes sets up /expenses and /income routes
func RegisterFinanceRoutes(rg *gin.RouterGroup, c *container.Container) {
	// Expenses
	expenses := rg.Group("/expenses")
	expensesAuthz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/expenses", perm)
	}
	{
		expenses.GET("", expensesAuthz("read"), c.ExpenseHandler.GetAll)
		expenses.GET("/:id", expensesAuthz("read"), c.ExpenseHandler.GetByID)
		expenses.POST("", expensesAuthz("write"), c.ExpenseHandler.Create)
		expenses.PUT("/:id", expensesAuthz("write"), c.ExpenseHandler.Update)
		expenses.DELETE("/:id", expensesAuthz("delete"), c.ExpenseHandler.Delete)
	}

	// Income
	income := rg.Group("/income")
	incomeAuthz := func(perm string) gin.HandlerFunc {
		return auth.AuthorizationMiddleware(c.PermissionChecker, "/income", perm)
	}
	{
		income.GET("", incomeAuthz("read"), c.IncomeHandler.GetAll)
		income.GET("/:id", incomeAuthz("read"), c.IncomeHandler.GetByID)
		income.POST("", incomeAuthz("write"), c.IncomeHandler.Create)
		income.PUT("/:id", incomeAuthz("write"), c.IncomeHandler.Update)
		income.DELETE("/:id", incomeAuthz("delete"), c.IncomeHandler.Delete)
	}
}
