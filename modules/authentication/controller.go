package authentication

import (
	"context"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/modules/audit"
	"nashrul-be/crm/modules/user"
	"nashrul-be/crm/utils"
	"nashrul-be/crm/utils/crypto"
)

type ControllerInterface interface {
	Login(request LoginRequest) (*entities.User, error)
	Logout(ctx context.Context) error
}

func NewAuthController(
	actorUseCase user.UseCaseInterface,
	auditUseCase audit.UseCaseInterface,
	hash crypto.Hash,
) ControllerInterface {
	return controller{
		actorUseCase: actorUseCase,
		auditUseCase: auditUseCase,
		hash:         hash,
	}
}

type controller struct {
	actorUseCase user.UseCaseInterface
	auditUseCase audit.UseCaseInterface
	hash         crypto.Hash
}

func (c controller) Login(request LoginRequest) (*entities.User, error) {
	account, err := c.actorUseCase.GetByUsername(nil, request.Username)
	if err != nil {
		return nil, err
	}
	if err = c.hash.Compare(request.Password, account.Password); err != nil {
		return nil, err
	}
	ctx := utils.SetUserToContext(context.Background(), account)
	if err = c.auditUseCase.CreateAudit(ctx, "Login"); err != nil {
		return nil, err
	}
	return &account, nil
}

func (c controller) Logout(ctx context.Context) error {
	return c.auditUseCase.CreateAudit(ctx, "Logout")
}
