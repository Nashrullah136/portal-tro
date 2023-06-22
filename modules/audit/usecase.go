package audit

import (
	"context"
	"encoding/json"
	"github.com/adjust/rmq/v5"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
)

type UseCaseInterface interface {
	GetAll(ctx context.Context, query repositories.AuditQuery, limit, offset int) ([]entities.Audit, error)
	CreateAudit(ctx context.Context, action string) error
	CountAll(ctx context.Context, query repositories.AuditQuery) (int, error)
	ExportCSV(ctx context.Context, query repositories.AuditQuery) error
}

func NewUseCase(
	auditRepo repositories.AuditRepositoryInterface,
	exportCsvRepo repositories.ExportCsvRepositoryInterface,
	queue rmq.Queue,
) UseCaseInterface {
	return useCase{
		auditRepo:     auditRepo,
		exportCsvRepo: exportCsvRepo,
		queue:         queue,
	}
}

type useCase struct {
	auditRepo     repositories.AuditRepositoryInterface
	exportCsvRepo repositories.ExportCsvRepositoryInterface
	queue         rmq.Queue
}

func (uc useCase) GetAll(ctx context.Context, query repositories.AuditQuery, limit, offset int) ([]entities.Audit, error) {
	return uc.auditRepo.GetAll(ctx, query, limit, offset)
}

func (uc useCase) CountAll(ctx context.Context, query repositories.AuditQuery) (int, error) {
	return uc.auditRepo.CountGetAll(ctx, query)
}

func (uc useCase) CreateAudit(ctx context.Context, action string) error {
	return uc.auditRepo.CreateAudit(ctx, action)
}

func (uc useCase) ExportCSV(ctx context.Context, query repositories.AuditQuery) error {
	user, err := entities.ExtractActorFromContext(ctx)
	if err != nil {
		return err
	}
	csvReq := entities.InitExportCsv(user.Username)
	csvReq, err = uc.exportCsvRepo.Create(csvReq)
	if err != nil {
		return err
	}
	payload := PayloadQueue{
		RequestID: csvReq.ID,
		Query:     query,
	}
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	if err = uc.queue.Publish(string(payloadJson)); err != nil {
		return err
	}
	return nil
}
