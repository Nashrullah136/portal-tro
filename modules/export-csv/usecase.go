package export_csv

import (
	"context"
	"errors"
	"log"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
	csvutils "nashrul-be/crm/utils/csv"
	"os"
)

type UseCaseInterface interface {
	GetAll(ctx context.Context, query repositories.ExportCsvQuery, limit, offset int) ([]entities.ExportCsv, error)
	DownloadCsv(ctx context.Context, id uint) (string, error)
	Delete(ctx context.Context, id uint) error
	CountAll(ctx context.Context, query repositories.ExportCsvQuery) (int, error)
}

func NewUseCase(
	exportCsvRepo repositories.ExportCsvRepositoryInterface,
	auditRepo repositories.AuditRepositoryInterface,
) UseCaseInterface {
	return useCase{
		exportCsvRepo: exportCsvRepo,
		auditRepo:     auditRepo,
	}
}

type useCase struct {
	exportCsvRepo repositories.ExportCsvRepositoryInterface
	auditRepo     repositories.AuditRepositoryInterface
}

func (uc useCase) GetAll(ctx context.Context, query repositories.ExportCsvQuery, limit, offset int) ([]entities.ExportCsv, error) {
	return uc.exportCsvRepo.GetAll(ctx, query, limit, offset)
}

func (uc useCase) CountAll(ctx context.Context, query repositories.ExportCsvQuery) (int, error) {
	return uc.exportCsvRepo.CountAll(ctx, query)
}

func (uc useCase) DownloadCsv(ctx context.Context, id uint) (string, error) {
	exportCsv, err := uc.exportCsvRepo.GetById(id)
	if err != nil {
		return "", err
	}
	if _, err := os.Stat(csvutils.Path(exportCsv.Filename)); err != nil {
		return "", errors.New("csv file doesn't exist")
	}
	return exportCsv.Filename, nil
}

func (uc useCase) Delete(ctx context.Context, id uint) error {
	exportCsv, err := uc.exportCsvRepo.GetById(id)
	if err != nil {
		return err
	}
	csvPath := csvutils.Path(exportCsv.Filename)
	if _, err = os.Stat(csvPath); err == nil {
		if err = os.Remove(csvPath); err != nil {
			log.Println("failed to delete csv File")
		}
	}
	return uc.exportCsvRepo.Delete(ctx, id)
}
