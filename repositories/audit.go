package repositories

import (
	"context"
	"gorm.io/gorm"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/utils"
	"nashrul-be/crm/utils/db"
	"nashrul-be/crm/utils/localtime"
	"time"
)

type AuditRepositoryInterface interface {
	CreateAudit(ctx context.Context, action string) (err error)
	CountGetAll(ctx context.Context, query AuditQuery) (int, error)
	GetAll(ctx context.Context, query AuditQuery, limit, offset int) (result []entities.Audit, err error)
	Create(audit entities.Audit) error
	Begin() db.Transactor
	New(transact db.Transactor) AuditRepositoryInterface
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
	if !query.FromDate.IsZero() {
		sql.Where("CONVERT(DATETIME, ?) <= date_time", query.FromDate)
	}
	if !query.ToDate.IsZero() {
		sql.Where("date_time < CONVERT(DATETIME, ?)", query.ToDate.Add(24*time.Hour))
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

func (r auditRepository) Create(audit entities.Audit) error {
	return r.db.Create(&audit).Error
}

func (r auditRepository) CreateAudit(ctx context.Context, action string) (err error) {
	user, err := utils.GetUserFromContext(ctx)
	if err != nil {
		return
	}
	result := entities.Audit{
		DateTime: localtime.Now(),
		Username: user.Identity(),
		Action:   action,
	}
	err = r.db.WithContext(ctx).Create(&result).Error
	return
}

func (r auditRepository) Begin() db.Transactor {
	return db.NewTransactor(r.db.Begin())
}

func (r auditRepository) New(transact db.Transactor) AuditRepositoryInterface {
	return auditRepository{db: transact.GetDB()}
}
