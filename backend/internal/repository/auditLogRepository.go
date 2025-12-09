package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafaameen91/project-managment/backend/internal/models"
)

type AuditLogRepositoryInterface interface {
	Create(ctx context.Context, log *models.AuditLog) error
	GetAll(ctx context.Context, limit, offset int) ([]models.AuditLog, int64, error)
	GetByActorID(ctx context.Context, actorID int64, limit, offset int) ([]models.AuditLog, int64, error)
	GetByTarget(ctx context.Context, targetType string, targetID int64, limit, offset int) ([]models.AuditLog, int64, error)
}

type AuditLogRepository struct {
	db *pgxpool.Pool
}

func NewAuditLogRepository(db *pgxpool.Pool) *AuditLogRepository {
	return &AuditLogRepository{db: db}
}

func (r *AuditLogRepository) Create(ctx context.Context, log *models.AuditLog) error {
	query := `
		INSERT INTO auditLogs (actorId, action, targetType, targetId, method, status, ipAddress)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, createdAt
	`

	return r.db.QueryRow(ctx, query,
		log.ActorID, log.Action, log.TargetType, log.TargetID,
		log.Method, log.Status, log.IPAddress,
	).Scan(&log.ID, &log.CreatedAt)
}

func (r *AuditLogRepository) GetAll(ctx context.Context, limit, offset int) ([]models.AuditLog, int64, error) {
	var total int64
	countQuery := `SELECT COUNT(*) FROM auditLogs`
	if err := r.db.QueryRow(ctx, countQuery).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT id, actorId, action, targetType, targetId, method, status, ipAddress, createdAt
		FROM auditLogs
		ORDER BY createdAt DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var logs []models.AuditLog
	for rows.Next() {
		var l models.AuditLog
		err := rows.Scan(
			&l.ID, &l.ActorID, &l.Action, &l.TargetType, &l.TargetID,
			&l.Method, &l.Status, &l.IPAddress, &l.CreatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		logs = append(logs, l)
	}

	return logs, total, rows.Err()
}

func (r *AuditLogRepository) GetByActorID(ctx context.Context, actorID int64, limit, offset int) ([]models.AuditLog, int64, error) {
	var total int64
	countQuery := `SELECT COUNT(*) FROM auditLogs WHERE actorId = $1`
	if err := r.db.QueryRow(ctx, countQuery, actorID).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT id, actorId, action, targetType, targetId, method, status, ipAddress, createdAt
		FROM auditLogs
		WHERE actorId = $1
		ORDER BY createdAt DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.Query(ctx, query, actorID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var logs []models.AuditLog
	for rows.Next() {
		var l models.AuditLog
		err := rows.Scan(
			&l.ID, &l.ActorID, &l.Action, &l.TargetType, &l.TargetID,
			&l.Method, &l.Status, &l.IPAddress, &l.CreatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		logs = append(logs, l)
	}

	return logs, total, rows.Err()
}

func (r *AuditLogRepository) GetByTarget(ctx context.Context, targetType string, targetID int64, limit, offset int) ([]models.AuditLog, int64, error) {
	var total int64
	countQuery := `SELECT COUNT(*) FROM auditLogs WHERE targetType = $1 AND targetId = $2`
	if err := r.db.QueryRow(ctx, countQuery, targetType, targetID).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT id, actorId, action, targetType, targetId, method, status, ipAddress, createdAt
		FROM auditLogs
		WHERE targetType = $1 AND targetId = $2
		ORDER BY createdAt DESC
		LIMIT $3 OFFSET $4
	`

	rows, err := r.db.Query(ctx, query, targetType, targetID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var logs []models.AuditLog
	for rows.Next() {
		var l models.AuditLog
		err := rows.Scan(
			&l.ID, &l.ActorID, &l.Action, &l.TargetType, &l.TargetID,
			&l.Method, &l.Status, &l.IPAddress, &l.CreatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		logs = append(logs, l)
	}

	return logs, total, rows.Err()
}