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

func (uc useCase) GetByDocumentNumber(ctx context.Context, documentNumber string) (entities.SPAN, error) {
	return uc.spanRepo.GetBySpanDocumentNumber(ctx, documentNumber)
}

func (uc useCase) UpdatePatchBankRiau(ctx context.Context, span entities.SPAN) error {
	oldSpan, err := uc.GetByDocumentNumber(ctx, span.DocumentNumber)
	if err != nil {
		return err
	}
	newSpan := PatchBankRiau(oldSpan)
	audit, err := entities.AuditUpdateWithOldData(ctx, &newSpan, &oldSpan)
	if err != nil {
		return err
	}
	brivaTx := uc.spanRepo.Begin()
	auditTx := uc.auditRepo.Begin()
	brivaRepoTx := uc.spanRepo.New(brivaTx)
	auditRepoTx := uc.auditRepo.New(auditTx)
	if err = brivaRepoTx.Update(ctx, span); err != nil {
		brivaTx.Rollback()
		auditTx.Rollback()
		return err
	}
	if err = auditRepoTx.Create(audit); err != nil {
		brivaTx.Rollback()
		auditTx.Rollback()
		return err
	}
	if err = brivaTx.Commit().Error; err != nil {
		brivaTx.Rollback()
		auditTx.Rollback()
		return err
	}
	if err = auditTx.Commit().Error; err != nil {
		log.Println("Failed to commit audit table, proceed to publish data to queue.")
		auditTx.Rollback()
		auditJson, err := json.Marshal(audit)
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
