package authentication

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/modules/audit"
	"nashrul-be/crm/modules/user"
)

type Route struct {
	authRequestHandler RequestHandlerInterface
}

func NewRoute(actorUseCase user.UseCaseInterface, auditUseCase audit.UseCaseInterface) Route {
	controller := NewAuthController(actorUseCase, auditUseCase)
	requestHandler := NewRequestHandler(controller)
	return Route{
		authRequestHandler: requestHandler,
	}
}

func (r Route) Handle(router *gin.Engine) {
	router.POST("/login", r.authRequestHandler.Login)
}
