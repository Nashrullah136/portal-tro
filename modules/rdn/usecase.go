package rdn

import (
	"context"
	"encoding/json"
	"github.com/adjust/rmq/v5"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/utils"
	"nashrul-be/crm/utils/auditUtils"
	"nashrul-be/crm/utils/logutils"
)

type UseCaseInterface interface {
	MakeAuditBatch(ctx context.Context, rdnData []entities.RDN, patchFunc patch) (result []entities.Audit, err error)
	GetRdnExisting(ctx context.Context) ([]entities.RDN, error)
	GetRdnNew(ctx context.Context) ([]entities.RDN, error)
	UpdateRdnExisting(ctx context.Context) error
	UpdateRdnNew(ctx context.Context) error
	Update(ctx context.Context, rdnPatch entities.RDN, whereCond map[string]any, audits []entities.Audit) (err error)
}

func NewUseCase(
	rdnRepo repositories.RdnRepositoryInterface,
	auditRepo repositories.AuditRepositoryInterface,
	queue rmq.Queue,
) UseCaseInterface {
	return useCase{
		rdnRepo:   rdnRepo,
		auditRepo: auditRepo,
		queue:     queue,
	}
}

type useCase struct {
	rdnRepo   repositories.RdnRepositoryInterface
	auditRepo repositories.AuditRepositoryInterface
	queue     rmq.Queue
}

func (uc useCase) MakeAuditBatch(ctx context.Context, rdnData []entities.RDN, patchFunc patch) (result []entities.Audit, err error) {
	actor, err := utils.GetUserFromContext(ctx)
	if err != nil {
		logutils.Get().Println(err)
		return nil, err
	}
	for _, rdn := range rdnData {
		rdnPatched := patchFunc(rdn)
		auditResult, err := auditUtils.UpdateWithOldData(&actor, &rdnPatched, &rdn)
		if err != nil {
			return nil, err
		}
		result = append(result, entities.MapAuditResultToAuditEntities(auditResult))
	}
	return result, nil
}

func (uc useCase) GetRdnExisting(ctx context.Context) ([]entities.RDN, error) {
	whereCond := GetRdnExistCondition()
	return uc.rdnRepo.GetWithCond(ctx, whereCond)
}

func (uc useCase) GetRdnNew(ctx context.Context) ([]entities.RDN, error) {
	whereCond := GetRdnNewCondition()
	return uc.rdnRepo.GetWithCond(ctx, whereCond)
}

func (uc useCase) UpdateRdnExisting(ctx context.Context) error {
	whereCond := GetRdnExistCondition()
	rdnPatch := PatchRdnExisting(entities.RDN{})
	rdnData, err := uc.GetRdnExisting(ctx)
	if err != nil {
		return err
	}
	audits, err := uc.MakeAuditBatch(ctx, rdnData, PatchRdnExisting)
	if err != nil {
		return err
	}
	return uc.Update(ctx, rdnPatch, whereCond, audits)
}

func (uc useCase) UpdateRdnNew(ctx context.Context) error {
	whereCond := GetRdnExistCondition()
	rdnPatch := PatchRdnNew(entities.RDN{})
	rdnData, err := uc.GetRdnExisting(ctx)
	if err != nil {
		return err
	}
	audits, err := uc.MakeAuditBatch(ctx, rdnData, PatchRdnNew)
	if err != nil {
		return err
	}
	return uc.Update(ctx, rdnPatch, whereCond, audits)
}

func (uc useCase) Update(ctx context.Context, rdnPatch entities.RDN, whereCond map[string]any, audits []entities.Audit) (err error) {
	rdnTx := uc.rdnRepo.Begin()
	auditTx := uc.auditRepo.Begin()
	rdnRepoTx := uc.rdnRepo.New(rdnTx)
	auditRepoTx := uc.auditRepo.New(auditTx)
	if err = rdnRepoTx.UpdateWithWhereCond(ctx, rdnPatch, whereCond); err != nil {
		rdnTx.Rollback()
		auditTx.Rollback()
		return err
	}
	for _, auditData := range audits {
		if err = auditRepoTx.Create(auditData); err != nil {
			rdnTx.Rollback()
			auditTx.Rollback()
			return err
		}
	}
	if err = rdnTx.Commit().Error; err != nil {
		rdnTx.Rollback()
		auditTx.Rollback()
		return err
	}
	if err = auditTx.Commit().Error; err != nil {
		logutils.Get().Println("Failed to commit audit table, proceed to publish data to queue.")
		auditTx.Rollback()
		for _, auditData := range audits {
			auditJson, err := json.Marshal(auditData)
			if err != nil {
				logutils.Get().Println("Failed on marshalling audit")
			}
			if err = uc.queue.Publish(string(auditJson)); err != nil {
				logutils.Get().Println("Failed to publish data to the queue")
			}
		}
	}
	return nil
}
