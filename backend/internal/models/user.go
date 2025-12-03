package models

import "time"

type User struct {
	ID        int64      `json:"id" db:"id"`
	Username  string     `json:"username" db:"username"`
	Email     string     `json:"email" db:"email"`
	Password  string     `json:"-" db:"password"`
	FullName  *string    `json:"fullName" db:"fullname"`
	Phone     *string    `json:"phone" db:"phone"`
	Avatar    *string    `json:"avatar" db:"avatar"`
	JobTitle  *string    `json:"jobTitle" db:"jobtitle"`
	Status    *string    `json:"status" db:"status"`
	LastLogin *time.Time `json:"lastLogin" db:"lastlogin"`
	CreatedAt time.Time  `json:"createdAt" db:"createdat"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updatedat"`
}

type UserRole struct {
	ID        int64     `json:"id" db:"id"`
	UserID    int64     `json:"userId" db:"userid"`
	RoleID    int64     `json:"roleId" db:"roleid"`
	CreatedAt time.Time `json:"createdAt" db:"createdat"`
}
