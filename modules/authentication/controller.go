package authentication

import (
	"context"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/modules/audit"
	"nashrul-be/crm/modules/user"
	"nashrul-be/crm/utils/hash"
)

type ControllerInterface interface {
	Login(request LoginRequest) (*entities.User, error)
}

func NewAuthController(actorUseCase user.UseCaseInterface, auditUseCase audit.UseCaseInterface) ControllerInterface {
	return controller{
		actorUseCase: actorUseCase,
		auditUseCase: auditUseCase,
	}
}

type controller struct {
	actorUseCase user.UseCaseInterface
	auditUseCase audit.UseCaseInterface
}

func (c controller) Login(request LoginRequest) (*entities.User, error) {
	account, err := c.actorUseCase.GetByUsername(nil, request.Username)
	if err != nil {
		return nil, err
	}
	if err = hash.Compare(request.Password, account.Password); err != nil {
		return nil, err
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, "user", account)
	if err = c.auditUseCase.CreateAudit(ctx, "Login"); err != nil {
		return nil, err
	}
	return &account, nil
}
