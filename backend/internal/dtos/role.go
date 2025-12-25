package dtos

import "time"

// ** Role DTOs **

// Response DTO
type Role struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

// Request DTOs
type CreateRole struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
}

type UpdateRole struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

// ** Page DTOs **

// Response DTO
type Page struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Icon      *string   `json:"icon"`
	Route     string    `json:"route"`
	Status    *string   `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

// Request DTOs
type CreatePage struct {
	Name   string  `json:"name" binding:"required"`
	Icon   *string `json:"icon"`
	Route  string  `json:"route" binding:"required"`
	Status *string `json:"status"`
}

type UpdatePage struct {
	Name   *string `json:"name"`
	Icon   *string `json:"icon"`
	Route  *string `json:"route"`
	Status *string `json:"status"`
}

// ** RolePage DTOs **

// Response DTO
type RolePage struct {
	ID          int64     `json:"id"`
	RoleID      int64     `json:"roleId"`
	PageID      int64     `json:"pageId"`
	Permissions *string   `json:"permissions"`
	CreatedAt   time.Time `json:"createdAt"`
}

// Request DTOs
type CreateRolePage struct {
	RoleID      int64   `json:"roleId" binding:"required"`
	PageID      int64   `json:"pageId" binding:"required"`
	Permissions *string `json:"permissions"`
}

type UpdateRolePage struct {
	Permissions *string `json:"permissions"`
}

// RolePageItem represents a single page permission for batch update
type RolePageItem struct {
	PageID      int64  `json:"pageId" binding:"required"`
	Permissions string `json:"permissions" binding:"required"`
}

// UpdateRolePagesRequest for atomic update of all role pages
type UpdateRolePagesRequest struct {
	Pages []RolePageItem `json:"pages" binding:"required"`
}

// ** UserRole DTOs **

// Response DTO
type UserRole struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"userId"`
	RoleID    int64     `json:"roleId"`
	CreatedAt time.Time `json:"createdAt"`
}

// Request DTO
type CreateUserRole struct {
	UserID int64 `json:"userId" binding:"required"`
	RoleID int64 `json:"roleId" binding:"required"`
}
