package models

import "time"

type AuditLog struct {
	ID         int64     `json:"id" db:"id"`
	ActorID    *int64    `json:"actorId" db:"actorid"`
	Action     string    `json:"action" db:"action"`
	TargetType string    `json:"targetType" db:"targettype"`
	TargetID   *int64    `json:"targetId" db:"targetid"`
	Method     string    `json:"method" db:"method"`
	Status     string    `json:"status" db:"status"`
	IPAddress  string    `json:"ipAddress" db:"ipaddress"`
	CreatedAt  time.Time `json:"createdAt" db:"createdat"`
}