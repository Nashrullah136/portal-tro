package audit

import (
	"context"
	"nashrul-be/crm/dto"
)

type ControllerInterface interface {
	GetAll(ctx context.Context, request GetAllRequest) (dto.BaseResponse, error)
	CreateAudit(ctx context.Context, action string) (dto.BaseResponse, error)
	ExportCSV(ctx context.Context, request ExportRequest) error
}

func NewController(auditUseCase UseCaseInterface) ControllerInterface {
	return controller{auditUseCase: auditUseCase}
}

type controller struct {
	auditUseCase UseCaseInterface
}

func (c controller) GetAll(ctx context.Context, request GetAllRequest) (dto.BaseResponse, error) {
	auditQuery := mapGetAllRequestToAuditQuery(request)
	offset := (request.Page - 1) * request.PerPage
	result, err := c.auditUseCase.GetAll(ctx, auditQuery, request.PerPage, offset)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	totalRow, err := c.auditUseCase.CountAll(ctx, auditQuery)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	return dto.SuccessPagination("Success retrieve audit", request.Page, totalRow/request.PerPage+1, totalRow, result), nil
}

func (c controller) CreateAudit(ctx context.Context, action string) (dto.BaseResponse, error) {
	err := c.auditUseCase.CreateAudit(ctx, action)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	return dto.Success("Success create audit", nil), nil
}

func (c controller) ExportCSV(ctx context.Context, request ExportRequest) error {
	auditQuery := mapExportRequestToAuditQuery(request)
	return c.auditUseCase.ExportCSV(ctx, auditQuery)
}
