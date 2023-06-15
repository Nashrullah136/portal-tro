package audit

import (
	"context"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
)

type UseCaseInterface interface {
	GetAll(ctx context.Context, query repositories.AuditQuery, limit, offset int) ([]entities.Audit, error)
	CreateAudit(ctx context.Context, action string) error
}

func NewUseCase(auditRepo repositories.AuditRepositoryInterface) UseCaseInterface {
	return useCase{auditRepo: auditRepo}
}

type useCase struct {
	auditRepo repositories.AuditRepositoryInterface
}

func (uc useCase) GetAll(ctx context.Context, query repositories.AuditQuery, limit, offset int) ([]entities.Audit, error) {
	result, err := uc.auditRepo.GetAll(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc useCase) CreateAudit(ctx context.Context, action string) error {
	return uc.auditRepo.CreateAudit(ctx, action)
}
