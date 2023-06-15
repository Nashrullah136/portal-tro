package authentication

import (
	"context"
	"fmt"
	"log"
	"nashrul-be/crm/dto"
	"nashrul-be/crm/modules/audit"
	"nashrul-be/crm/modules/user"
	"nashrul-be/crm/utils/hash"
	jwtUtil "nashrul-be/crm/utils/jwt"
)

type ControllerInterface interface {
	Login(request LoginRequest) (dto.BaseResponse, error)
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

func (c controller) Login(request LoginRequest) (dto.BaseResponse, error) {
	account, err := c.actorUseCase.GetByUsername(nil, request.Username)
	defaultResponse := dto.ErrorUnauthorized("Wrong Username/Password")
	if err != nil {
		return defaultResponse, err
	}
	if err := hash.Compare(request.Password, account.Password); err != nil {
		return defaultResponse, err
	}
	token, err := jwtUtil.GenerateJWT(account)
	if err != nil {
		return dto.ErrorInternalServerError(), err
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, "user", account)
	if err := c.auditUseCase.CreateAudit(ctx, "Login"); err != nil {
		log.Println(err)
		return dto.ErrorInternalServerError(), err
	}
	result := LoginResponse{Token: fmt.Sprintf("Bearer %s", token)}
	return dto.Success("Authenticated success", result), nil
}
