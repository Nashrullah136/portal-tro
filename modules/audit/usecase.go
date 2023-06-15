package audit

import (
	"context"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
)

type UseCaseInterface interface {
	GetAll(ctx context.Context, query repositories.AuditQuery, limit, offset int) ([]entities.Audit, error)
	CreateAudit(ctx context.Context, action string) error
	CountAll(ctx context.Context, query repositories.AuditQuery) (int, error)
}

func NewUseCase(auditRepo repositories.AuditRepositoryInterface) UseCaseInterface {
	return useCase{auditRepo: auditRepo}
}

type useCase struct {
	auditRepo repositories.AuditRepositoryInterface
}

func (uc useCase) GetAll(ctx context.Context, query repositories.AuditQuery, limit, offset int) ([]entities.Audit, error) {
	return uc.auditRepo.GetAll(ctx, query, limit, offset)
}

func (uc useCase) CountAll(ctx context.Context, query repositories.AuditQuery) (int, error) {
	return uc.auditRepo.CountGetAll(ctx, query)
}

func (uc useCase) CreateAudit(ctx context.Context, action string) error {
	return uc.auditRepo.CreateAudit(ctx, action)
}
