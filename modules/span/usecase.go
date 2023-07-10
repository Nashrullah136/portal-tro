package span

import (
	"context"
	"encoding/json"
	"github.com/adjust/rmq/v5"
	"log"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
)

type UseCaseInterface interface {
	ValidateSpan(span entities.SPAN, validations ...validateFunc) (error, error)
	GetByDocumentNumber(ctx context.Context, documentNumber string) (entities.SPAN, error)
	UpdatePatchBankRiau(ctx context.Context, span entities.SPAN) error
}

func NewUseCase(
	spanRepo repositories.SpanRepositoryInterface,
	auditRepo repositories.AuditRepositoryInterface,
	queue rmq.Queue,
) UseCaseInterface {
	return useCase{
		spanRepo:  spanRepo,
		auditRepo: auditRepo,
		queue:     queue,
	}
}

type useCase struct {
	spanRepo  repositories.SpanRepositoryInterface
	auditRepo repositories.AuditRepositoryInterface
	queue     rmq.Queue
}

func (uc useCase) ValidateSpan(span entities.SPAN, validations ...validateFunc) (error, error) {
	for _, validation := range validations {
		validationError, err := validation(span, uc.spanRepo)
		if err != nil {
			return nil, err
		}
		if validationError != nil {
			return validationError, nil
		}
	}
	return nil, nil
}

func (uc useCase) GetByDocumentNumber(ctx context.Context, documentNumber string) (entities.SPAN, error) {
	return uc.spanRepo.GetBySpanDocumentNumber(ctx, documentNumber)
}

func (uc useCase) UpdatePatchBankRiau(ctx context.Context, span entities.SPAN) error {
	oldSpan, err := uc.GetByDocumentNumber(ctx, span.DocumentNumber)
	if err != nil {
		return err
	}
	newSpan := PatchBankRiau(oldSpan)
	auditEntities, err := uc.spanRepo.MakeAuditUpdateWithOldData(ctx, oldSpan, newSpan)
	if err != nil {
		return err
	}
	spanTx := uc.spanRepo.Begin()
	auditTx := uc.auditRepo.Begin()
	spanRepoTx := uc.spanRepo.New(spanTx)
	auditRepoTx := uc.auditRepo.New(auditTx)
	if err = spanRepoTx.Update(ctx, newSpan); err != nil {
		spanTx.Rollback()
		auditTx.Rollback()
		return err
	}
	if err = auditRepoTx.Create(auditEntities); err != nil {
		spanTx.Rollback()
		auditTx.Rollback()
		return err
	}
	if err = spanTx.Commit().Error; err != nil {
		spanTx.Rollback()
		auditTx.Rollback()
		return err
	}
	if err = auditTx.Commit().Error; err != nil {
		log.Println("Failed to commit audit table, proceed to publish data to queue.")
		auditTx.Rollback()
		auditJson, err := json.Marshal(auditEntities)
		if err != nil {
			log.Println("Failed on marshalling audit")
			return nil
		}
		if err = uc.queue.Publish(string(auditJson)); err != nil {
			log.Println("Failed to publish data to the queue")
			return nil
		}
	}
	return nil
}
