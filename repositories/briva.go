package repositories

import (
	"context"
	"gorm.io/gorm"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/utils/db"
)

type BrivaRepositoryInterface interface {
	GetByBrivaNo(ctx context.Context, brivano string) (briva entities.Briva, err error)
	Update(ctx context.Context, briva entities.Briva) error
	MakeAuditUpdate(ctx context.Context, briva entities.Briva) (entities.Audit, error)
	Begin() db.Transactor
	New(transact db.Transactor) BrivaRepositoryInterface
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

func (r brivaRepository) MakeAuditUpdate(ctx context.Context, briva entities.Briva) (entities.Audit, error) {
	return entities.AuditUpdate(r.db.WithContext(ctx), &briva)
}

func (r brivaRepository) Begin() db.Transactor {
	return db.NewTransactor(r.db.Begin())
}

func (r brivaRepository) New(transact db.Transactor) BrivaRepositoryInterface {
	return brivaRepository{db: transact.GetDB()}
}
