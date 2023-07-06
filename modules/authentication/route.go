package authentication

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/middleware"
	"nashrul-be/crm/modules/audit"
	"nashrul-be/crm/modules/user"
	"nashrul-be/crm/utils/crypto"
	"nashrul-be/crm/utils/session"
)

type Route struct {
	authRequestHandler RequestHandlerInterface
}

func NewRoute(actorUseCase user.UseCaseInterface,
	auditUseCase audit.UseCaseInterface,
	sessionManager session.Manager,
	hash crypto.Hash,
) Route {
	controller := NewAuthController(actorUseCase, auditUseCase, hash)
	requestHandler := NewRequestHandler(controller, sessionManager)
	return Route{
		authRequestHandler: requestHandler,
	}
}

func (r Route) Handle(router *gin.Engine, manager session.Manager) {
	router.POST("/login", r.authRequestHandler.Login)
	router.GET("/logout", middleware.Authenticate(manager), r.authRequestHandler.Logout)
}
