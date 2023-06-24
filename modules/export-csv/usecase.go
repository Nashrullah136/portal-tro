package export_csv

import (
	"context"
	"errors"
	"log"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/utils/filesystem"
)

type UseCaseInterface interface {
	GetAll(ctx context.Context, query repositories.ExportCsvQuery, limit, offset int) ([]entities.ExportCsv, error)
	DownloadCsv(ctx context.Context, id uint) (filesystem.File, error)
	Delete(ctx context.Context, id uint) error
	CountAll(ctx context.Context, query repositories.ExportCsvQuery) (int, error)
}

func NewUseCase(
	exportCsvRepo repositories.ExportCsvRepositoryInterface,
	auditRepo repositories.AuditRepositoryInterface,
	folder filesystem.Folder,
) UseCaseInterface {
	return useCase{
		exportCsvRepo: exportCsvRepo,
		auditRepo:     auditRepo,
		folder:        folder,
	}
}

// TODO: abstract file access
type useCase struct {
	exportCsvRepo repositories.ExportCsvRepositoryInterface
	auditRepo     repositories.AuditRepositoryInterface
	folder        filesystem.Folder
}

func (uc useCase) GetAll(ctx context.Context, query repositories.ExportCsvQuery, limit, offset int) ([]entities.ExportCsv, error) {
	return uc.exportCsvRepo.GetAll(ctx, query, limit, offset)
}

func (uc useCase) CountAll(ctx context.Context, query repositories.ExportCsvQuery) (int, error) {
	return uc.exportCsvRepo.CountAll(ctx, query)
}

func (uc useCase) DownloadCsv(ctx context.Context, id uint) (filesystem.File, error) {
	exportCsv, err := uc.exportCsvRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	if !uc.folder.IsExist(exportCsv.Filename) {
		return nil, errors.New("csv file doesn't exist")
	}
	return filesystem.NewFile(exportCsv.Filename, uc.folder), nil
}

func (uc useCase) Delete(ctx context.Context, id uint) error {
	exportCsv, err := uc.exportCsvRepo.GetById(id)
	if err != nil {
		return err
	}
	if err = uc.folder.Remove(exportCsv.Filename); err != nil {
		log.Println(err)
		return err
	}
	return uc.exportCsvRepo.Delete(ctx, id)
}
