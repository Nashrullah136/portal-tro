package user

import (
	"errors"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
)

type validateFunc func(actor entities.User, actorRepo repositories.ActorRepositoryInterface) (error, error)

func validateUsername(actor entities.User, actorRepo repositories.ActorRepositoryInterface) (error, error) {
	exist, err := actorRepo.IsUsernameExist(actor)
	if err != nil {
		return nil, err
	}
	if exist {
		return errors.New("username already taken"), nil
	}
	return nil, nil
}

func validateExist(actor entities.User, actorRepo repositories.ActorRepositoryInterface) (error, error) {
	exist, err := actorRepo.IsUsernameExist(actor)
	if err != nil {
		return nil, err
	}
	if !exist {
		return errors.New("user doesn't exist"), nil
	}
	return nil, nil
}
