package repositories

import (
	"context"
	"gorm.io/gorm"
	"log"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/utils"
	"nashrul-be/crm/utils/auditUtils"
	"nashrul-be/crm/utils/db"
)

type BrivaRepositoryInterface interface {
	IsBrivaExist(briva entities.Briva) (exist bool, err error)
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

func (r brivaRepository) IsBrivaExist(briva entities.Briva) (exist bool, err error) {
	var count int64
	err = r.db.Model(&entities.Briva{}).Where("brivano = ?", briva.Brivano).Count(&count).Error
	if err != nil {
		return
	}
	exist = count > 0
	return
}

func (r brivaRepository) GetByBrivaNo(ctx context.Context, brivano string) (briva entities.Briva, err error) {
	err = r.db.WithContext(ctx).First(&briva, brivano).Error
	return briva, err
}

func (r brivaRepository) Update(ctx context.Context, briva entities.Briva) error {
	return r.db.WithContext(ctx).Updates(&briva).Error
}

func (r brivaRepository) MakeAuditUpdate(ctx context.Context, briva entities.Briva) (entities.Audit, error) {
	actor, err := utils.GetUserFromContext(ctx)
	if err != nil {
		log.Println(err)
		return entities.Audit{}, err
	}
	result, err := auditUtils.Update(r.db, &actor, &briva)
	if err != nil {
		return entities.Audit{}, err
	}
	return entities.MapAuditResultToAuditEntities(result), nil
}

func (r brivaRepository) Begin() db.Transactor {
	return db.NewTransactor(r.db.Begin())
}

func (r brivaRepository) New(transact db.Transactor) BrivaRepositoryInterface {
	return brivaRepository{db: transact.GetDB()}
}
