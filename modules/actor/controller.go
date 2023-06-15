package actor

import (
	"context"
	"nashrul-be/crm/dto"
	"nashrul-be/crm/entities"
)

type ControllerInterface interface {
	GetByUsername(ctx context.Context, username string) (dto.BaseResponse, error)
	GetAll(ctx context.Context, req PaginationRequest) (dto.BaseResponse, error)
	CreateActor(ctx context.Context, req CreateRequest) (dto.BaseResponse, error)
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

func (c controller) GetByUsername(ctx context.Context, username string) (dto.BaseResponse, error) {
	actor, err := c.actorUseCase.GetByUsername(ctx, username)
	if err != nil {
		return actorNotFound(), err
	}
	actorRepresentation := mapActorToResponse(actor)
	return dto.Success("Success retrieve data", actorRepresentation), nil
}

func (c controller) GetAll(ctx context.Context, req PaginationRequest) (dto.BaseResponse, error) {
	offset := (req.Page - 1) * req.PerPage
	if req.Role != "admin" && req.Role != "user" {
		return dto.ErrorBadRequest("Invalid role"), nil
	}
	actors, err := c.actorUseCase.GetAll(ctx, req.Username+"%", req.Role, req.PerPage, offset)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	actorResponse := make([]Representation, 0)
	for _, actor := range actors {
		actorResponse = append(actorResponse, mapActorToResponse(actor))
	}
	return dto.Success("Success retrieve actor", actorResponse), err
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
	return dto.Created("Success create actor", response), nil
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
	return dto.Success("Success update actor", response), nil
}

func (c controller) ChangePassword(ctx context.Context, req ChangePasswordRequest) (dto.BaseResponse, error) {
	var user entities.User
	user.Password = req.Password
	user.Username = req.Username
	validationErr, err := c.actorUseCase.validateActor(user, validateExist)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	if validationErr != nil {
		return dto.ErrorBadRequest(validationErr.Error()), nil
	}
	ctx = context.WithValue(ctx, "except", "change password")
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
