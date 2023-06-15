package actor

import (
	"context"
	"errors"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/utils/hash"
)

type UseCaseInterface interface {
	validateActor(actor entities.User, validations ...validateFunc) (error, error)
	GetAll(ctx context.Context, username, role string, limit, offset uint) ([]entities.User, error)
	GetByUsername(ctx context.Context, username string) (actor entities.User, err error)
	CreateUser(ctx context.Context, actor entities.User) (result entities.User, err error)
	UpdateUser(ctx context.Context, actor entities.User) (result entities.User, err error)
	ChangePassword(ctx context.Context, oldPassword string, user entities.User) (result error, err error)
	DeleteUser(ctx context.Context, username string) (err error)
}

func NewUseCase(
	repositoryInterface repositories.ActorRepositoryInterface,
	roleRepositoryInterface repositories.RoleRepositoryInterface,
) UseCaseInterface {
	return useCase{
		actorRepository: repositoryInterface,
		roleRepository:  roleRepositoryInterface,
	}
}

type useCase struct {
	actorRepository repositories.ActorRepositoryInterface
	roleRepository  repositories.RoleRepositoryInterface
}

func (uc useCase) validateActor(actor entities.User, validations ...validateFunc) (error, error) {
	for _, validation := range validations {
		validationError, err := validation(actor, uc.actorRepository)
		if err != nil {
			return nil, err
		}
		if validationError != nil {
			return validationError, nil
		}
	}
	return nil, nil
}

func (uc useCase) GetByUsername(ctx context.Context, username string) (actor entities.User, err error) {
	actor, err = uc.actorRepository.GetByUsername(ctx, username)
	if err != nil {
		return
	}
	role, err := uc.roleRepository.GetByID(actor.RoleID)
	actor.Role = role
	return
}

func (uc useCase) GetAll(ctx context.Context, username, role string, limit, offset uint) ([]entities.User, error) {
	actors, err := uc.actorRepository.GetAll(ctx, username, role, limit, offset)
	return actors, err
}

func (uc useCase) CreateUser(ctx context.Context, actor entities.User) (result entities.User, err error) {
	actor.Password, err = hash.Hash(actor.Password)
	if err != nil {
		return
	}
	result, err = uc.actorRepository.Create(ctx, actor)
	if err != nil {
		return
	}
	return
}

func (uc useCase) UpdateUser(ctx context.Context, actor entities.User) (result entities.User, err error) {
	if actor.Password != "" {
		actor.Password, err = hash.Hash(actor.Password)
		if err != nil {
			return
		}
	}
	result, err = uc.actorRepository.Update(ctx, actor)
	return
}

func (uc useCase) ChangePassword(ctx context.Context, oldPassword string, user entities.User) (result error, err error) {
	dbUser, err := uc.actorRepository.GetByUsername(ctx, user.Username)
	if err != nil {
		return nil, err
	}
	if result = hash.Compare(oldPassword, dbUser.Password); result != nil {
		return errors.New("wrong password"), nil
	}
	_, err = uc.UpdateUser(ctx, user)
	return nil, err
}

func (uc useCase) DeleteUser(ctx context.Context, username string) (err error) {
	return uc.actorRepository.Delete(ctx, username)
}
