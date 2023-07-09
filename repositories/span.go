package repositories

import (
	"context"
	"gorm.io/gorm"
	"log"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/utils"
	"nashrul-be/crm/utils/audit"
	"nashrul-be/crm/utils/db"
)

type SpanRepositoryInterface interface {
	GetBySpanDocumentNumber(ctx context.Context, documentNumber string) (span entities.SPAN, err error)
	Update(ctx context.Context, span entities.SPAN) error
	MakeAuditUpdate(ctx context.Context, span entities.SPAN) (entities.Audit, error)
	Begin() db.Transactor
	New(transact db.Transactor) SpanRepositoryInterface
}

func NewSpanRepository(db *gorm.DB) SpanRepositoryInterface {
	return spanRepository{db: db}
}

type spanRepository struct {
	db *gorm.DB
}

func (r spanRepository) GetBySpanDocumentNumber(ctx context.Context, documentNumber string) (span entities.SPAN, err error) {
	span.DocumentNumber = documentNumber
	err = r.db.WithContext(ctx).Where("documentdate = substring(CONVERT(varchar,getdate(),126),1,10) " +
		"and statuscode not in ('0001','void')").First(&span).Error
	return span, err
}

func (r spanRepository) Update(ctx context.Context, span entities.SPAN) error {
	return r.db.WithContext(ctx).Updates(&span).Error
}

func (r spanRepository) MakeAuditUpdate(ctx context.Context, span entities.SPAN) (entities.Audit, error) {
	actor, err := utils.GetUserFromContext(ctx)
	if err != nil {
		log.Println(err)
		return entities.Audit{}, err
	}
	result, err := audit.Update(r.db, &actor, &span)
	if err != nil {
		return entities.Audit{}, err
	}
	return entities.MapAuditResultToAuditEntities(result), nil
}

func (r spanRepository) Begin() db.Transactor {
	return db.NewTransactor(r.db.Begin())
}

func (r spanRepository) New(transact db.Transactor) SpanRepositoryInterface {
	return spanRepository{db: transact.GetDB()}
}
