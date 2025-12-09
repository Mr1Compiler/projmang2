package dtos

import "time"

// Response DTOs
type UserSummary struct {
	ID        int64      `json:"id"`
	FullName  string     `json:"fullName"`
	UserName  string     `json:"userName"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	JobTitle  string     `json:"jobTitle"`
	Status    *string    `json:"status"`
	LastLogin *time.Time `json:"lastLogin"`
}

type User struct {
	ID        int64      `json:"id"`
	Username  string     `json:"username"`
	FullName  string     `json:"fullName"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Avatar    *string    `json:"avatar"`
	JobTitle  string     `json:"jobTitle"`
	Status    *string    `json:"status"`
	LastLogin *time.Time `json:"lastLogin"`
	CreatedAt time.Time  `json:"createdAt"`
}

// Request DTOs
type CreateUser struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	FullName string `json:"fullName" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	JobTitle string `json:"jobTitle" binding:"required"`
}

type UpdateUser struct {
	Username *string `json:"username"`
	Email    *string `json:"email" binding:"omitempty,email"`
	FullName *string `json:"fullName"`
	Phone    *string `json:"phone"`
	Avatar   *string `json:"avatar"`
	JobTitle *string `json:"jobTitle"`
	Status   *string `json:"status"`
}

type UpdatePassword struct {
	Password string `json:"password" binding:"required,min=8"`
}

type UpdateStatus struct {
	Status string `json:"status" binding:"required"`
}
