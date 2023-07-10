package span

import (
	"context"
	"log"
	"nashrul-be/crm/dto"
)

type ControllerInterface interface {
	GetByDocumentNumber(ctx context.Context, documentNumber string) (dto.BaseResponse, error)
	UpdateBankRiau(ctx context.Context, request UpdateRequest) (dto.BaseResponse, error)
}

func NewController(spanUseCase UseCaseInterface) ControllerInterface {
	return controller{spanUseCase: spanUseCase}
}

type controller struct {
	spanUseCase UseCaseInterface
}

func (c controller) GetByDocumentNumber(ctx context.Context, documentNumber string) (dto.BaseResponse, error) {
	result, err := c.spanUseCase.GetByDocumentNumber(ctx, documentNumber)
	if err != nil {
		log.Println(err)
		return dto.ErrorNotFound("Document Number"), nil
	}
	return dto.Success("Success retrieve span", result), nil
}

func (c controller) UpdateBankRiau(ctx context.Context, request UpdateRequest) (dto.BaseResponse, error) {
	span := mapUpdateRequestToSpan(request)
	result, err := c.spanUseCase.ValidateSpan(span, validateExist)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	if result != nil {
		return dto.ErrorBadRequest(result.Error()), nil
	}
	if err := c.spanUseCase.UpdatePatchBankRiau(ctx, span); err != nil {
		return dto.ErrorInternalServerError(), err
	}
	return dto.Success("Success update SPAN", nil), nil
}
