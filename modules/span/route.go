package span

import (
	"github.com/adjust/rmq/v5"
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/middleware"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/utils/session"
)

type Route struct {
	brivaRequestHandler RequestHandlerInterface
}

func NewRoute(
	spanRepo repositories.SpanRepositoryInterface,
	auditRepo repositories.AuditRepositoryInterface,
	queue rmq.Queue,
) Route {
	brivaUseCase := NewUseCase(spanRepo, auditRepo, queue)
	brivaController := NewController(brivaUseCase)
	brivaRequestHandler := NewRequestHandler(brivaController)
	return Route{
		brivaRequestHandler: brivaRequestHandler,
	}
}

func (r Route) Handle(engine *gin.Engine, manager session.Manager) {
	router := engine.Group("/span", middleware.Authenticate(manager),
		middleware.Refresh(manager), middleware.CheckNewUser(), middleware.AuthorizationUserOnly())
	router.GET("/:documentNumber", r.brivaRequestHandler.GetByDocumentNumber)
	router.POST("/:documentNumber", r.brivaRequestHandler.UpdateBankRiau)
}
