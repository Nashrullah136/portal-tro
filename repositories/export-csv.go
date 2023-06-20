package repositories

import (
	"context"
	"gorm.io/gorm"
	"nashrul-be/crm/entities"
)

type ExportCsvRepositoryInterface interface {
	Create(exportCsv entities.ExportCsv) (entities.ExportCsv, error)
	Update(exportCsv entities.ExportCsv) error
	GetAll(ctx context.Context, query ExportCsvQuery, limit, offset int) (exportCsv []entities.ExportCsv, err error)
	CountAll(ctx context.Context, query ExportCsvQuery) (int, error)
	GetById(id uint) (entities.ExportCsv, error)
	Delete(ctx context.Context, id uint) error
}

func NewExportCsvRepository(db *gorm.DB) ExportCsvRepositoryInterface {
	return exportCsvRepository{db: db}
}

type exportCsvRepository struct {
	db *gorm.DB
}

func (r exportCsvRepository) GetById(id uint) (entities.ExportCsv, error) {
	exportCsv := entities.ExportCsv{
		ID: id,
	}
	err := r.db.Find(&exportCsv).Error
	return exportCsv, err
}

func (r exportCsvRepository) buildQuery(ctx context.Context, query ExportCsvQuery) *gorm.DB {
	sql := r.db.WithContext(ctx).Model(&entities.ExportCsv{})
	if query.Username != "" {
		sql = sql.Where("username LIKE ?", query.Username+"%")
	}
	if !query.AfterDate.IsZero() {
		sql = sql.Where("created_at >= ?", query.AfterDate)
	}
	if !query.BeforeDate.IsZero() {
		sql = sql.Where("created_at <= ?", query.BeforeDate)
	}
	return sql
}

func (r exportCsvRepository) GetAll(ctx context.Context, query ExportCsvQuery, limit, offset int) (exportCsv []entities.ExportCsv, err error) {
	err = r.buildQuery(ctx, query).Limit(limit).Offset(offset).Find(&exportCsv).Error
	return
}

func (r exportCsvRepository) CountAll(ctx context.Context, query ExportCsvQuery) (int, error) {
	var total int64
	err := r.buildQuery(ctx, query).Count(&total).Error
	return int(total), err
}

func (r exportCsvRepository) Create(exportCsv entities.ExportCsv) (entities.ExportCsv, error) {
	export := exportCsv
	err := r.db.Create(&export).Error
	return export, err
}

func (r exportCsvRepository) Update(exportCsv entities.ExportCsv) error {
	return r.db.Updates(exportCsv).Error
}

func (r exportCsvRepository) Delete(ctx context.Context, id uint) error {
	exportCsv := &entities.ExportCsv{ID: id}
	return r.db.WithContext(ctx).Delete(exportCsv).Error
}
