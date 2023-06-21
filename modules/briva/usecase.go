package briva

import (
	"context"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
)

type UseCaseInterface interface {
	GetByBrivaNo(ctx context.Context, brivano string) (entities.Briva, error)
	Update(ctx context.Context, briva entities.Briva) error
}

func NewUseCase(brivaRepo repositories.BrivaRepositoryInterface) UseCaseInterface {
	return useCase{
		brivaRepo: brivaRepo,
	}
}

type useCase struct {
	brivaRepo repositories.BrivaRepositoryInterface
}

func (uc useCase) GetByBrivaNo(ctx context.Context, brivano string) (entities.Briva, error) {
	return uc.brivaRepo.GetByBrivaNo(ctx, brivano)
}

func (uc useCase) Update(ctx context.Context, briva entities.Briva) error {
	return uc.brivaRepo.Update(ctx, briva)
}
