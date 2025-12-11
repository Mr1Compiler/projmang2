package services

import (
	"context"

	"github.com/mustafa91ameen/prjalgo/backend/internal/dtos"
	"github.com/mustafa91ameen/prjalgo/backend/internal/logger"
	"github.com/mustafa91ameen/prjalgo/backend/internal/models"
	"github.com/mustafa91ameen/prjalgo/backend/internal/repository"
)

type AuditLogService struct {
	repo repository.AuditLogRepositoryInterface
}

func NewAuditLogService(repo repository.AuditLogRepositoryInterface) *AuditLogService {
	return &AuditLogService{repo: repo}
}

func (s *AuditLogService) Log(ctx context.Context, actorID *int64, action, targetType string, targetID *int64, method, status, ipAddress string) {
	log := &models.AuditLog{
		ActorID:    actorID,
		Action:     action,
		TargetType: targetType,
		TargetID:   targetID,
		Method:     method,
		Status:     status,
		IPAddress:  ipAddress,
	}

	// Fire and forget - don't block the main operation
	// Use WithoutCancel to ensure audit log completes even if request context is cancelled
	detachedCtx := context.WithoutCancel(ctx)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				logger.Error("audit log panic", "error", r, "action", action, "targetType", targetType)
			}
		}()
		if err := s.repo.Create(detachedCtx, log); err != nil {
			logger.Error("failed to create audit log", "error", err, "action", action, "targetType", targetType)
		}
	}()
}

func (s *AuditLogService) GetAll(ctx context.Context, limit, offset int) (dtos.PaginatedResponse[models.AuditLog], error) {
	logs, total, err := s.repo.GetAll(ctx, limit, offset)
	if err != nil {
		return dtos.PaginatedResponse[models.AuditLog]{}, err
	}

	page := (offset / limit) + 1
	return dtos.NewPaginatedResponse(logs, total, page, limit), nil
}

func (s *AuditLogService) GetByActorID(ctx context.Context, actorID int64, limit, offset int) (dtos.PaginatedResponse[models.AuditLog], error) {
	logs, total, err := s.repo.GetByActorID(ctx, actorID, limit, offset)
	if err != nil {
		return dtos.PaginatedResponse[models.AuditLog]{}, err
	}

	page := (offset / limit) + 1
	return dtos.NewPaginatedResponse(logs, total, page, limit), nil
}

func (s *AuditLogService) GetByTarget(ctx context.Context, targetType string, targetID int64, limit, offset int) (dtos.PaginatedResponse[models.AuditLog], error) {
	logs, total, err := s.repo.GetByTarget(ctx, targetType, targetID, limit, offset)
	if err != nil {
		return dtos.PaginatedResponse[models.AuditLog]{}, err
	}

	page := (offset / limit) + 1
	return dtos.NewPaginatedResponse(logs, total, page, limit), nil
}
