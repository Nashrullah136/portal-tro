package repositories

import (
	"context"
	"gorm.io/gorm"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/utils/db"
)

type RdnRepositoryInterface interface {
	GetWithCond(ctx context.Context, whereCond map[string]any) ([]entities.RDN, error)
	Update(ctx context.Context, briva entities.RDN) error
	UpdateWithWhereCond(ctx context.Context, rdn entities.RDN, whereCond map[string]any) error
	MakeAuditUpdate(ctx context.Context, briva entities.RDN) (entities.Audit, error)
	Begin() db.Transactor
	New(transact db.Transactor) RdnRepositoryInterface
}

func NewRdnRepository(db *gorm.DB) RdnRepositoryInterface {
	return rdnRepository{db: db}
}

type rdnRepository struct {
	db *gorm.DB
}

func (r rdnRepository) GetWithCond(ctx context.Context, whereCond map[string]any) (result []entities.RDN, err error) {
	err = r.db.WithContext(ctx).Model(&entities.RDN{}).Where(whereCond).Where("RDN is not null").Find(&result).Error
	return
}

func (r rdnRepository) Update(ctx context.Context, briva entities.RDN) error {
	return r.db.WithContext(ctx).Updates(&briva).Error
}

func (r rdnRepository) UpdateWithWhereCond(ctx context.Context, rdn entities.RDN, whereCond map[string]any) error {
	return r.db.WithContext(ctx).Model(&entities.RDN{}).Where(whereCond).Where("RDN is not null").Updates(rdn).Error
}

func (r rdnRepository) MakeAuditUpdate(ctx context.Context, rdn entities.RDN) (entities.Audit, error) {
	return entities.AuditUpdate(r.db.WithContext(ctx), &rdn)
}

func (r rdnRepository) Begin() db.Transactor {
	return db.NewTransactor(r.db.Begin())
}

func (r rdnRepository) New(transact db.Transactor) RdnRepositoryInterface {
	return rdnRepository{db: transact.GetDB()}
}
