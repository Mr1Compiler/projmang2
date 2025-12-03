package models

import "time"

type Role struct {
	ID          int64     `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description *string   `json:"description" db:"description"`
	CreatedAt   time.Time `json:"createdAt" db:"createdat"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updatedat"`
}

type Page struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Icon      *string   `json:"icon" db:"icon"`
	Route     string    `json:"route" db:"route"`
	Status    *string   `json:"status" db:"status"`
	CreatedAt time.Time `json:"createdAt" db:"createdat"`
	UpdatedAt time.Time `json:"updatedAt" db:"updatedat"`
}

type RolePage struct {
	ID          int64     `json:"id" db:"id"`
	RoleID      int64     `json:"roleId" db:"roleid"`
	PageID      int64     `json:"pageId" db:"pageid"`
	Permissions *string   `json:"permissions" db:"permissions"`
	CreatedAt   time.Time `json:"createdAt" db:"createdat"`
}
