package audit

import (
	"context"
	"encoding/json"
	"github.com/adjust/rmq/v5"
	"log"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
	csvutils "nashrul-be/crm/utils/csv"
	"nashrul-be/crm/utils/filesystem"
)

type UseCaseInterface interface {
	GetAll(ctx context.Context, query repositories.AuditQuery, limit, offset int) ([]entities.Audit, error)
	CreateAudit(ctx context.Context, action string) error
	CountAll(ctx context.Context, query repositories.AuditQuery) (int, error)
	ExportCsvAsync(ctx context.Context, query repositories.AuditQuery) error
	ExportCsv(ctx context.Context, query repositories.AuditQuery) (*csvutils.FileCsv, error)
}

func NewUseCase(
	auditRepo repositories.AuditRepositoryInterface,
	exportCsvRepo repositories.ExportCsvRepositoryInterface,
	queue rmq.Queue,
	folder filesystem.Folder,
) UseCaseInterface {
	return useCase{
		auditRepo:     auditRepo,
		exportCsvRepo: exportCsvRepo,
		queue:         queue,
		folder:        folder,
	}
}

type useCase struct {
	auditRepo     repositories.AuditRepositoryInterface
	exportCsvRepo repositories.ExportCsvRepositoryInterface
	queue         rmq.Queue
	folder        filesystem.Folder
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

func (uc useCase) ExportCsv(ctx context.Context, query repositories.AuditQuery) (*csvutils.FileCsv, error) {
	audits, err := uc.auditRepo.GetAll(ctx, query, 0, 0)
	if err != nil {
		log.Printf("Failed to get audit trail data. error: %s\n", err)
		return nil, err
	}
	csvFile, err := csvutils.NewCSV(uc.folder)
	defer csvFile.Finish()
	if err != nil {
		log.Printf("Failed to create csv file. error: %s\n", err)
		return nil, err
	}
	if err = csvFile.Write(entities.Audit{}.HeaderCSV()); err != nil {
		log.Printf("Failed to write to csv file. error: %s\n", err)
		return nil, err
	}
	for _, auditData := range audits {
		if err = csvFile.Write(auditData.CsvRepresentation()); err != nil {
			log.Printf("Failed to write to csv file. error: %s\n", err)
			return nil, err
		}
	}
	return csvFile, nil
}

func (uc useCase) ExportCsvAsync(ctx context.Context, query repositories.AuditQuery) error {
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
