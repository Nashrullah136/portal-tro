package briva

import (
	"context"
	"encoding/json"
	"github.com/adjust/rmq/v5"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/utils/logutils"
)

//go:generate mockery --name UseCaseInterface --inpackage
type UseCaseInterface interface {
	ValidateBriva(briva entities.Briva, validations ...validateFunc) (error, error)
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

func (uc useCase) ValidateBriva(briva entities.Briva, validations ...validateFunc) (error, error) {
	for _, validation := range validations {
		validationError, err := validation(briva, uc.brivaRepo)
		if err != nil {
			return nil, err
		}
		if validationError != nil {
			return validationError, nil
		}
	}
	return nil, nil
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
		logutils.Get().Println("Failed to commit audit table, proceed to publish data to queue.")
		auditTx.Rollback()
		auditJson, err := json.Marshal(audit)
		if err != nil {
			logutils.Get().Println("Failed on marshalling audit")
			return nil
		}
		if err = uc.queue.Publish(string(auditJson)); err != nil {
			logutils.Get().Println("Failed to publish data to the queue")
			return nil
		}
	}
	return nil
}
