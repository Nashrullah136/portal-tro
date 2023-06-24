package user

import (
	"context"
	"errors"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/utils/crypto"
)

type UseCaseInterface interface {
	validateActor(actor entities.User, validations ...validateFunc) (error, error)
	CountAll(ctx context.Context, username, role string) (int, error)
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
	hash crypto.Hash,
) UseCaseInterface {
	return useCase{
		actorRepository: repositoryInterface,
		roleRepository:  roleRepositoryInterface,
		hash:            hash,
	}
}

type useCase struct {
	actorRepository repositories.ActorRepositoryInterface
	roleRepository  repositories.RoleRepositoryInterface
	hash            crypto.Hash
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

func (uc useCase) CountAll(ctx context.Context, username, role string) (int, error) {
	return uc.actorRepository.CountAll(ctx, username, role)
}

func (uc useCase) GetAll(ctx context.Context, username, role string, limit, offset uint) ([]entities.User, error) {
	return uc.actorRepository.GetAll(ctx, username, role, limit, offset)
}

func (uc useCase) GetByUsername(ctx context.Context, username string) (actor entities.User, err error) {
	return uc.actorRepository.GetByUsername(ctx, username)
}

func (uc useCase) CreateUser(ctx context.Context, actor entities.User) (result entities.User, err error) {
	actor.Password, err = uc.hash.Hash(actor.Password)
	if err != nil {
		return
	}
	result, err = uc.actorRepository.Create(ctx, actor)
	if err != nil {
		return
	}
	role, err := uc.roleRepository.GetByID(result.RoleID)
	if err == nil {
		result.Role = role
	}
	return
}

func (uc useCase) UpdateUser(ctx context.Context, actor entities.User) (result entities.User, err error) {
	if actor.Password != "" {
		actor.Password, err = uc.hash.Hash(actor.Password)
		if err != nil {
			return
		}
	}
	err = uc.actorRepository.Update(ctx, actor)
	if err != nil {
		return
	}
	result, err = uc.actorRepository.GetByUsername(ctx, actor.Username)
	if err != nil {
		return actor, nil
	}
	return
}

func (uc useCase) ChangePassword(ctx context.Context, oldPassword string, user entities.User) (result error, err error) {
	dbUser, err := uc.actorRepository.GetByUsername(ctx, user.Username)
	if err != nil {
		return nil, err
	}
	if result = uc.hash.Compare(oldPassword, dbUser.Password); result != nil {
		return errors.New("wrong password"), nil
	}
	_, err = uc.UpdateUser(ctx, user)
	return nil, err
}

func (uc useCase) DeleteUser(ctx context.Context, username string) (err error) {
	return uc.actorRepository.Delete(ctx, username)
}
