package export_csv

import (
	"context"
	"nashrul-be/crm/dto"
)

type ControllerInterface interface {
	GetAll(ctx context.Context, request GetAllRequest) (dto.BaseResponse, error)
	DownloadCsv(ctx context.Context, id uint) (string, error)
}

func NewController(exportUseCase UseCaseInterface) ControllerInterface {
	return controller{exportCsvUseCase: exportUseCase}
}

type controller struct {
	exportCsvUseCase UseCaseInterface
}

func (c controller) GetAll(ctx context.Context, request GetAllRequest) (dto.BaseResponse, error) {
	query := mapGetAllRequestToExportCsvQuery(request)
	offset := request.PerPage * (request.Page - 1)
	results, err := c.exportCsvUseCase.GetAll(ctx, query, request.PerPage, offset)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	totalRow, err := c.exportCsvUseCase.CountAll(ctx, query)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	return dto.SuccessPagination("Success retrieve export request", request.Page,
		totalRow/request.PerPage, totalRow, results), nil
}

func (c controller) DownloadCsv(ctx context.Context, id uint) (string, error) {
	return c.exportCsvUseCase.DownloadCsv(ctx, id)
}
