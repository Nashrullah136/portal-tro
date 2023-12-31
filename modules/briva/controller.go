package briva

import (
	"context"
	"nashrul-be/crm/dto"
	"nashrul-be/crm/utils/logutils"
)

type ControllerInterface interface {
	GetByBrivaNo(ctx context.Context, brivano string) (dto.BaseResponse, error)
	Update(ctx context.Context, request UpdateRequest) (dto.BaseResponse, error)
}

func NewController(brivaUseCase UseCaseInterface) ControllerInterface {
	return controller{brivaUseCase: brivaUseCase}
}

type controller struct {
	brivaUseCase UseCaseInterface
}

func (c controller) GetByBrivaNo(ctx context.Context, brivano string) (dto.BaseResponse, error) {
	result, err := c.brivaUseCase.GetByBrivaNo(ctx, brivano)
	if err != nil {
		logutils.Get().Println(err)
		return dto.ErrorNotFound("brivano"), nil
	}
	return dto.Success("Brivano has been found", result), nil
}

func (c controller) Update(ctx context.Context, request UpdateRequest) (dto.BaseResponse, error) {
	briva := mapUpdateRequestToBriva(request)
	validateErr, err := c.brivaUseCase.ValidateBriva(briva, validateExist)
	if err != nil {
		return dto.ErrorInternalServerError(), nil
	}
	if validateErr != nil {
		return dto.ErrorBadRequest(validateErr.Error()), nil
	}
	if err := c.brivaUseCase.Update(ctx, briva); err != nil {
		return dto.ErrorInternalServerError(), err
	}
	return dto.Success("Success update briva", nil), nil
}
