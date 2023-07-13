package user

import (
	"context"
	"nashrul-be/crm/dto"
)

type ControllerInterface interface {
	GetByUsername(ctx context.Context, username string) (dto.BaseResponse, error)
	GetAll(ctx context.Context, req PaginationRequest) (dto.BaseResponse, error)
	CreateActor(ctx context.Context, req CreateRequest) (dto.BaseResponse, error)
	UpdateProfile(ctx context.Context, req UpdateProfile) (dto.BaseResponse, error)
	UpdateActor(ctx context.Context, req UpdateRequest) (dto.BaseResponse, error)
	ChangePassword(ctx context.Context, req ChangePasswordRequest) (dto.BaseResponse, error)
	DeleteActor(ctx context.Context, username string) error
}

func NewController(useCaseInterface UseCaseInterface) ControllerInterface {
	return controller{
		actorUseCase: useCaseInterface,
	}
}

type controller struct {
	actorUseCase UseCaseInterface
}

func (c controller) UpdateProfile(ctx context.Context, req UpdateProfile) (dto.BaseResponse, error) {
	user := mapUpdateProfileToUser(req)
	updatedUser, err := c.actorUseCase.UpdateUser(ctx, user)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	response := mapActorToResponse(updatedUser)
	return dto.Success("Success update user", response), nil
}

func (c controller) GetByUsername(ctx context.Context, username string) (dto.BaseResponse, error) {
	actor, err := c.actorUseCase.GetByUsername(ctx, username)
	if err != nil {
		return actorNotFound(), err
	}
	actorRepresentation := mapActorToResponse(actor)
	return dto.Success("Success retrieve user", actorRepresentation), nil
}

func (c controller) GetAll(ctx context.Context, req PaginationRequest) (dto.BaseResponse, error) {
	offset := (req.Page - 1) * req.PerPage
	actors, err := c.actorUseCase.GetAll(ctx, req.Username+"%", req.Role, uint(req.PerPage), uint(offset))
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	actorResponse := make([]Representation, 0)
	for _, actor := range actors {
		actorResponse = append(actorResponse, mapActorToResponse(actor))
	}
	totalRow, err := c.actorUseCase.CountAll(ctx, req.Username+"%", req.Role)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	return dto.SuccessPagination("Success retrieve user", req.Page, req.PerPage, totalRow, actorResponse), err
}

func (c controller) CreateActor(ctx context.Context, req CreateRequest) (dto.BaseResponse, error) {
	actor := mapCreateRequestToActor(req)
	validationErr, err := c.actorUseCase.validateActor(actor, validateUsername)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	if validationErr != nil {
		return dto.ErrorBadRequest(validationErr.Error()), nil
	}
	createdActor, err := c.actorUseCase.CreateUser(ctx, actor)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	response := mapActorToResponse(createdActor)
	return dto.Created("Success create user", response), nil
}

func (c controller) UpdateActor(ctx context.Context, req UpdateRequest) (dto.BaseResponse, error) {
	actor := mapUpdateRequestToActor(req)
	validationErr, err := c.actorUseCase.validateActor(actor, validateExist)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	if validationErr != nil {
		return dto.ErrorBadRequest(validationErr.Error()), nil
	}
	updatedActor, err := c.actorUseCase.UpdateUser(ctx, actor)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	response := mapActorToResponse(updatedActor)
	return dto.Success("Success update user", response), nil
}

func (c controller) ChangePassword(ctx context.Context, req ChangePasswordRequest) (dto.BaseResponse, error) {
	user := mapChangePasswordToUser(req)
	validationErr, err := c.actorUseCase.validateActor(user, validateExist)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	if validationErr != nil {
		return dto.ErrorBadRequest(validationErr.Error()), nil
	}
	result, err := c.actorUseCase.ChangePassword(ctx, req.OldPassword, user)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	if result != nil {
		return dto.ErrorBadRequest(result.Error()), nil
	}
	return dto.Success("Success update password", nil), nil
}

func (c controller) DeleteActor(ctx context.Context, username string) error {
	if err := c.actorUseCase.DeleteUser(ctx, username); err != nil {
		return err
	}
	return nil
}
