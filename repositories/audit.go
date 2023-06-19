package repositories

import (
	"context"
	"gorm.io/gorm"
	"nashrul-be/crm/entities"
	"time"
)

type AuditRepositoryInterface interface {
	CreateAudit(ctx context.Context, action string) (err error)
	CountGetAll(ctx context.Context, query AuditQuery) (int, error)
	GetAll(ctx context.Context, query AuditQuery, limit, offset int) (result []entities.Audit, err error)
}

func NewAuditRepository(db *gorm.DB) AuditRepositoryInterface {
	return auditRepository{db: db}
}

type auditRepository struct {
	db *gorm.DB
}

func (r auditRepository) buildGetAllQuery(ctx context.Context, query AuditQuery) *gorm.DB {
	sql := r.db.WithContext(ctx).Model(&entities.Audit{})
	if query.Username != "" {
		sql.Where("username = ?", query.Username)
	}
	if query.Object != "" {
		sql.Where("entity = ?", query.Object)
	}
	if query.ObjectId != "" {
		sql.Where("entity_id = ?", query.ObjectId)
	}
	if !query.AfterDate.IsZero() {
		sql.Where("date_time >= ?", query.AfterDate)
	}
	if !query.BeforeDate.IsZero() {
		sql.Where("date_time <= ?", query.BeforeDate)
	}
	return sql
}

func (r auditRepository) CountGetAll(ctx context.Context, query AuditQuery) (int, error) {
	var result int64
	err := r.buildGetAllQuery(ctx, query).Count(&result).Error
	return int(result), err
}

func (r auditRepository) GetAll(ctx context.Context, query AuditQuery, limit, offset int) (result []entities.Audit, err error) {
	sql := r.buildGetAllQuery(ctx, query)
	if limit > 0 {
		sql = sql.Limit(limit)
	}
	if offset > 0 {
		sql = sql.Offset(offset)
	}
	err = sql.Order("date_time desc").Find(&result).Error
	return
}

func (r auditRepository) CreateAudit(ctx context.Context, action string) (err error) {
	user, err := entities.ExtractActorFromContext(ctx)
	if err != nil {
		return
	}
	audit := entities.Audit{
		DateTime: time.Now(),
		Username: user.Username,
		Action:   action,
	}
	err = r.db.WithContext(ctx).Create(&audit).Error
	return
}
