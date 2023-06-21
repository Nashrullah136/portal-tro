package briva

import (
	"context"
	"log"
	"nashrul-be/crm/dto"
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
		log.Println(err)
		return dto.ErrorNotFound("brivano"), nil
	}
	return dto.Success("Brivano has been found", result), nil
}

func (c controller) Update(ctx context.Context, request UpdateRequest) (dto.BaseResponse, error) {
	briva := mapUpdateRequestToBriva(request)
	if err := c.brivaUseCase.Update(ctx, briva); err != nil {
		return dto.ErrorInternalServerError(), err
	}
	return dto.Success("Success update briva", nil), nil
}
