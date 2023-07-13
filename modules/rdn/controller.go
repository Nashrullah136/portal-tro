package rdn

import (
	"context"
	"nashrul-be/crm/dto"
	"nashrul-be/crm/utils/logutils"
)

type ControllerInterface interface {
	GetRdnNew(ctx context.Context) (dto.BaseResponse, error)
	GetRdnExisting(ctx context.Context) (dto.BaseResponse, error)
	UpdateRdnExisting(ctx context.Context) (dto.BaseResponse, error)
	UpdateRdnNew(ctx context.Context) (dto.BaseResponse, error)
}

func NewController(rdnUseCase UseCaseInterface) ControllerInterface {
	return controller{rdnUseCase: rdnUseCase}
}

type controller struct {
	rdnUseCase UseCaseInterface
}

func (c controller) GetRdnNew(ctx context.Context) (dto.BaseResponse, error) {
	result, err := c.rdnUseCase.GetRdnNew(ctx)
	if err != nil {
		logutils.Get().Println(err)
		return dto.ErrorNotFound("RDN not found"), nil
	}
	return dto.Success("Success retrieve RDN", result), nil
}

func (c controller) GetRdnExisting(ctx context.Context) (dto.BaseResponse, error) {
	result, err := c.rdnUseCase.GetRdnExisting(ctx)
	if err != nil {
		logutils.Get().Println(err)
		return dto.ErrorNotFound("RDN not found"), nil
	}
	return dto.Success("Success retrieve RDN", result), nil
}

func (c controller) UpdateRdnExisting(ctx context.Context) (dto.BaseResponse, error) {
	err := c.rdnUseCase.UpdateRdnExisting(ctx)
	if err != nil {
		logutils.Get().Println(err)
		return dto.ErrorNotFound("Error when update RDN"), nil
	}
	return dto.Success("Success update RDN", nil), nil
}

func (c controller) UpdateRdnNew(ctx context.Context) (dto.BaseResponse, error) {
	err := c.rdnUseCase.UpdateRdnNew(ctx)
	if err != nil {
		logutils.Get().Println(err)
		return dto.ErrorNotFound("Error when update RDN"), nil
	}
	return dto.Success("Success update RDN", nil), nil
}
