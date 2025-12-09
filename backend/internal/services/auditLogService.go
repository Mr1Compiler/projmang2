package services

import (
	"context"

	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/models"
	"github.com/mustafaameen91/project-managment/backend/internal/repository"
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
	go func() {
		_ = s.repo.Create(context.Background(), log)
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