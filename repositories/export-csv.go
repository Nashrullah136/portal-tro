package repositories

import (
	"gorm.io/gorm"
	"nashrul-be/crm/entities"
)

type ExportCsvRepositoryInterface interface {
	Create(exportCsv entities.ExportCsv) (entities.ExportCsv, error)
	Update(exportCsv entities.ExportCsv) error
	GetById(id uint) (entities.ExportCsv, error)
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

func (r exportCsvRepository) Create(exportCsv entities.ExportCsv) (entities.ExportCsv, error) {
	export := exportCsv
	err := r.db.Create(&export).Error
	return export, err
}

func (r exportCsvRepository) Update(exportCsv entities.ExportCsv) error {
	return r.db.Updates(exportCsv).Error
}
