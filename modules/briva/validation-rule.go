package briva

import (
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/utils"
)

type validateFunc func(briva entities.Briva, brivaRepo repositories.BrivaRepositoryInterface) (error, error)

func validateExist(briva entities.Briva, brivaRepo repositories.BrivaRepositoryInterface) (error, error) {
	exist, err := brivaRepo.IsBrivaExist(briva)
	if err != nil {
		return nil, err
	}
	if !exist {
		return utils.ErrNotFound, nil
	}
	return nil, nil
}
