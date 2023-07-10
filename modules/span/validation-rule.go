package span

import (
	"errors"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
)

var ErrNotFound = errors.New("data span not found")

type validateFunc func(span entities.SPAN, repositoryInterface repositories.SpanRepositoryInterface) (error, error)

func validateExist(span entities.SPAN, spanRepo repositories.SpanRepositoryInterface) (error, error) {
	exist, err := spanRepo.IsSpanExist(span)
	if err != nil {
		return nil, err
	}
	if !exist {
		return ErrNotFound, nil
	}
	return nil, nil
}
