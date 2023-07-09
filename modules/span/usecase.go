package span

import (
	"context"
	"encoding/json"
	"github.com/adjust/rmq/v5"
	"log"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/utils"
	"nashrul-be/crm/utils/audit"
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
	actor, err := utils.GetUserFromContext(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	oldSpan, err := uc.GetByDocumentNumber(ctx, span.DocumentNumber)
	if err != nil {
		return err
	}
	newSpan := PatchBankRiau(oldSpan)
	auditResult, err := audit.UpdateWithOldData(&actor, &newSpan, &oldSpan)
	if err != nil {
		return err
	}
	auditEntities := entities.MapAuditResultToAuditEntities(auditResult)
	spanTx := uc.spanRepo.Begin()
	auditTx := uc.auditRepo.Begin()
	brivaRepoTx := uc.spanRepo.New(spanTx)
	auditRepoTx := uc.auditRepo.New(auditTx)
	if err = brivaRepoTx.Update(ctx, newSpan); err != nil {
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
