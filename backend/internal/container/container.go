package container

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafa91ameen/prjalgo/backend/internal/auth"
	"github.com/mustafa91ameen/prjalgo/backend/internal/config"
	"github.com/mustafa91ameen/prjalgo/backend/internal/handlers"
	"github.com/mustafa91ameen/prjalgo/backend/internal/repository"
	"github.com/mustafa91ameen/prjalgo/backend/internal/services"
)

// Container holds all application dependencies
type Container struct {
	// Project
	ProjectHandler *handlers.ProjectHandler

	// WorkDay
	WorkDayHandler          *handlers.WorkDayHandler
	WorkDayMaterialHandler  *handlers.WorkDayMaterialHandler
	WorkDayLaborHandler     *handlers.WorkDayLaborHandler
	WorkDayEquipmentHandler *handlers.WorkDayEquipmentHandler

	// Work Categories
	WorkCategoryHandler    *handlers.WorkCategoryHandler
	WorkSubCategoryHandler *handlers.WorkSubCategoryHandler

	// Finance
	ExpenseHandler *handlers.ExpenseHandler
	IncomeHandler  *handlers.IncomeHandler

	// Debtor
	DebtorHandler *handlers.DebtorHandler

	// User
	UserHandler       *handlers.UserHandler
	UserRoleHandler   *handlers.UserRoleHandler
	TeamMemberHandler *handlers.TeamMemberHandler

	// Role & Permissions
	RoleHandler     *handlers.RoleHandler
	PageHandler     *handlers.PageHandler
	RolePageHandler *handlers.RolePageHandler

	// Auth
	AuthHandler       *handlers.AuthHandler
	JWTManager        *auth.JWTManager
	PermissionChecker auth.PermissionChecker

	// Audit
	AuditLogService *services.AuditLogService
}

// New creates a new container with all dependencies wired up
func New(db *pgxpool.Pool, cfg *config.Config) *Container {
	// Parse JWT expiry
	jwtExpiry, _ := time.ParseDuration(cfg.JWTExpiry)
	if jwtExpiry == 0 {
		jwtExpiry = 15 * time.Minute
	}
	jwtManager := auth.NewJWTManager(cfg.JWTSecret, jwtExpiry)

	// Parse refresh token expiry
	refreshExpiry, _ := time.ParseDuration(cfg.RefreshTokenExpiry)
	if refreshExpiry == 0 {
		refreshExpiry = 7 * 24 * time.Hour
	}

	// Repositories
	projectRepo := repository.NewProjectRepository(db)
	workDayRepo := repository.NewWorkDayRepository(db)
	workDayMaterialRepo := repository.NewWorkDayMaterialRepository(db)
	workDayLaborRepo := repository.NewWorkDayLaborRepository(db)
	workDayEquipmentRepo := repository.NewWorkDayEquipmentRepository(db)
	workCategoryRepo := repository.NewWorkCategoryRepository(db)
	workSubCategoryRepo := repository.NewWorkSubCategoryRepository(db)
	expenseRepo := repository.NewExpenseRepository(db)
	incomeRepo := repository.NewIncomeRepository(db)
	debtorRepo := repository.NewDebtorRepository(db)
	userRepo := repository.NewUserRepository(db)
	userRoleRepo := repository.NewUserRoleRepository(db)
	teamMemberRepo := repository.NewTeamMemberRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	pageRepo := repository.NewPageRepository(db)
	rolePageRepo := repository.NewRolePageRepository(db)
	refreshTokenRepo := repository.NewRefreshTokenRepository(db)
	auditLogRepo := repository.NewAuditLogRepository(db)

	// Services
	projectService := services.NewProjectService(projectRepo)
	workDayService := services.NewWorkDayService(workDayRepo, projectRepo, workSubCategoryRepo)
	workDayMaterialService := services.NewWorkDayMaterialService(workDayMaterialRepo)
	workDayLaborService := services.NewWorkDayLaborService(workDayLaborRepo)
	workDayEquipmentService := services.NewWorkDayEquipmentService(workDayEquipmentRepo)
	workCategoryService := services.NewWorkCategoryService(workCategoryRepo)
	workSubCategoryService := services.NewWorkSubCategoryService(workSubCategoryRepo)
	expenseService := services.NewExpenseService(expenseRepo)
	incomeService := services.NewIncomeService(incomeRepo)
	debtorService := services.NewDebtorService(debtorRepo)
	userService := services.NewUserService(userRepo, userRoleRepo, refreshTokenRepo)
	userRoleService := services.NewUserRoleService(userRoleRepo)
	teamMemberService := services.NewTeamMemberService(teamMemberRepo)
	roleService := services.NewRoleService(roleRepo)
	pageService := services.NewPageService(pageRepo)
	rolePageService := services.NewRolePageService(rolePageRepo)
	authService := services.NewAuthService(userRepo, userRoleRepo, refreshTokenRepo, jwtManager, refreshExpiry)
	auditLogService := services.NewAuditLogService(auditLogRepo)

	// Handlers
	return &Container{
		ProjectHandler:          handlers.NewProjectHandler(projectService),
		WorkDayHandler:          handlers.NewWorkDayHandler(workDayService),
		WorkDayMaterialHandler:  handlers.NewWorkDayMaterialHandler(workDayMaterialService),
		WorkDayLaborHandler:     handlers.NewWorkDayLaborHandler(workDayLaborService),
		WorkDayEquipmentHandler: handlers.NewWorkDayEquipmentHandler(workDayEquipmentService),
		WorkCategoryHandler:     handlers.NewWorkCategoryHandler(workCategoryService),
		WorkSubCategoryHandler:  handlers.NewWorkSubCategoryHandler(workSubCategoryService),
		ExpenseHandler:          handlers.NewExpenseHandler(expenseService),
		IncomeHandler:           handlers.NewIncomeHandler(incomeService),
		DebtorHandler:           handlers.NewDebtorHandler(debtorService),
		UserHandler:             handlers.NewUserHandler(userService),
		UserRoleHandler:         handlers.NewUserRoleHandler(userRoleService),
		TeamMemberHandler:       handlers.NewTeamMemberHandler(teamMemberService),
		RoleHandler:             handlers.NewRoleHandler(roleService),
		PageHandler:             handlers.NewPageHandler(pageService),
		RolePageHandler:         handlers.NewRolePageHandler(rolePageService),
		AuthHandler:             handlers.NewAuthHandler(authService),
		JWTManager:              jwtManager,
		PermissionChecker:       userRoleRepo,
		AuditLogService:         auditLogService,
	}
}
