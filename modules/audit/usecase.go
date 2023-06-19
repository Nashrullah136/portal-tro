package audit

import (
	"context"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
	csvutils "nashrul-be/crm/utils/csv"
)

type UseCaseInterface interface {
	GetAll(ctx context.Context, query repositories.AuditQuery, limit, offset int) ([]entities.Audit, error)
	CreateAudit(ctx context.Context, action string) error
	CountAll(ctx context.Context, query repositories.AuditQuery) (int, error)
	ExportCSV(ctx context.Context, query repositories.AuditQuery) (*csvutils.FileCsv, error)
}

func NewUseCase(auditRepo repositories.AuditRepositoryInterface) UseCaseInterface {
	return useCase{auditRepo: auditRepo}
}

type useCase struct {
	auditRepo repositories.AuditRepositoryInterface
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

func (uc useCase) ExportCSV(ctx context.Context, query repositories.AuditQuery) (*csvutils.FileCsv, error) {
	audits, err := uc.auditRepo.GetAll(ctx, query, 0, 0)
	if err != nil {
		return nil, err
	}
	csvFile, err := csvutils.NewCSV()
	defer csvFile.Finish()
	if err != nil {
		return nil, err
	}
	if err := csvFile.Write(entities.Audit{}.HeaderCSV()); err != nil {
		return nil, err
	}
	for _, audit := range audits {
		if err := csvFile.Write(audit.CsvRepresentation()); err != nil {
			return nil, err
		}
	}
	return csvFile, nil
}
