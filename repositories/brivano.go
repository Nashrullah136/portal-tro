package repositories

import (
	"context"
	"gorm.io/gorm"
	"nashrul-be/crm/entities"
)

type BrivaRepositoryInterface interface {
	GetByBrivaNo(ctx context.Context, brivano string) (briva entities.Briva, err error)
	Update(ctx context.Context, briva entities.Briva) error
}

func NewBrivaRepository(db *gorm.DB) BrivaRepositoryInterface {
	return brivaRepository{db: db}
}

type brivaRepository struct {
	db *gorm.DB
}

func (r brivaRepository) GetByBrivaNo(ctx context.Context, brivano string) (briva entities.Briva, err error) {
	err = r.db.WithContext(ctx).First(&briva, brivano).Error
	return briva, err
}

func (r brivaRepository) Update(ctx context.Context, briva entities.Briva) error {
	return r.db.WithContext(ctx).Updates(&briva).Error
}
