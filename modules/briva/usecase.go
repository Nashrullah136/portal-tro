package briva

import (
	"context"
	"encoding/json"
	"github.com/adjust/rmq/v5"
	"log"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
)

type UseCaseInterface interface {
	GetByBrivaNo(ctx context.Context, brivano string) (entities.Briva, error)
	Update(ctx context.Context, briva entities.Briva) error
}

func NewUseCase(
	brivaRepo repositories.BrivaRepositoryInterface,
	auditRepo repositories.AuditRepositoryInterface,
	queue rmq.Queue,
) UseCaseInterface {
	return useCase{
		brivaRepo: brivaRepo,
		auditRepo: auditRepo,
		queue:     queue,
	}
}

type useCase struct {
	brivaRepo repositories.BrivaRepositoryInterface
	auditRepo repositories.AuditRepositoryInterface
	queue     rmq.Queue
}

func (uc useCase) GetByBrivaNo(ctx context.Context, brivano string) (entities.Briva, error) {
	return uc.brivaRepo.GetByBrivaNo(ctx, brivano)
}

func (uc useCase) Update(ctx context.Context, briva entities.Briva) error {
	audit, err := uc.brivaRepo.MakeAuditUpdate(ctx, briva)
	if err != nil {
		return err
	}
	brivaTx := uc.brivaRepo.Begin()
	auditTx := uc.auditRepo.Begin()
	brivaRepoTx := uc.brivaRepo.New(brivaTx)
	auditRepoTx := uc.auditRepo.New(auditTx)
	if err = brivaRepoTx.Update(ctx, briva); err != nil {
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
		}
		if err = uc.queue.Publish(string(auditJson)); err != nil {
			log.Println("Failed to publish data to the queue")
		}
	}
	return nil
}
