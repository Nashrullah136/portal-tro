package repositories

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/utils"
	"nashrul-be/crm/utils/auditUtils"
	"nashrul-be/crm/utils/db"
	"nashrul-be/crm/utils/localtime"
	"nashrul-be/crm/utils/logutils"
)

type SpanRepositoryInterface interface {
	IsSpanExist(span entities.SPAN) (exist bool, err error)
	GetBySpanDocumentNumber(ctx context.Context, documentNumber string) (span entities.SPAN, err error)
	Update(ctx context.Context, span entities.SPAN) error
	MakeAuditUpdate(ctx context.Context, span entities.SPAN) (entities.Audit, error)
	MakeAuditUpdateWithOldData(ctx context.Context, oldSpan entities.SPAN, newSpan entities.SPAN) (entities.Audit, error)
	Begin() db.Transactor
	New(transact db.Transactor) SpanRepositoryInterface
}

func NewSpanRepository(db *gorm.DB) SpanRepositoryInterface {
	return spanRepository{db: db}
}

type spanRepository struct {
	db *gorm.DB
}

func (r spanRepository) IsSpanExist(span entities.SPAN) (exist bool, err error) {
	var count int64
	err = r.db.Model(&entities.SPAN{}).Where("DOCUMENTNUMBER = ?", span.DocumentNumber).Count(&count).Error
	if err != nil {
		return
	}
	exist = count > 0
	return
}

func (r spanRepository) GetBySpanDocumentNumber(ctx context.Context, documentNumber string) (span entities.SPAN, err error) {
	span.DocumentNumber = documentNumber
	err = r.db.WithContext(ctx).Where("documentdate = ?", localtime.Now().Format("2006-01-02")).
		Where("statuscode not in ('0001','void')").First(&span).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entities.SPAN{}, utils.ErrNotFound
	}
	return span, err
}

func (r spanRepository) Update(ctx context.Context, span entities.SPAN) error {
	return r.db.WithContext(ctx).Updates(&span).Error
}

func (r spanRepository) MakeAuditUpdate(ctx context.Context, span entities.SPAN) (entities.Audit, error) {
	actor, err := utils.GetUserFromContext(ctx)
	if err != nil {
		logutils.Get().Println(err)
		return entities.Audit{}, err
	}
	result, err := auditUtils.Update(r.db, &actor, &span)
	if err != nil {
		return entities.Audit{}, err
	}
	return entities.MapAuditResultToAuditEntities(result), nil
}

func (r spanRepository) MakeAuditUpdateWithOldData(ctx context.Context, oldSpan entities.SPAN, newSpan entities.SPAN) (entities.Audit, error) {
	actor, err := utils.GetUserFromContext(ctx)
	if err != nil {
		logutils.Get().Println(err)
		return entities.Audit{}, err
	}
	result, err := auditUtils.UpdateWithOldData(&actor, &newSpan, &oldSpan)
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
