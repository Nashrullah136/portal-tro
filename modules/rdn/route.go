package rdn

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
	rdnRepo repositories.RdnRepositoryInterface,
	auditRepo repositories.AuditRepositoryInterface,
	queue rmq.Queue,
) Route {
	brivaUseCase := NewUseCase(rdnRepo, auditRepo, queue)
	brivaController := NewController(brivaUseCase)
	brivaRequestHandler := NewRequestHandler(brivaController)
	return Route{
		brivaRequestHandler: brivaRequestHandler,
	}
}

func (r Route) Handle(engine *gin.Engine, manager session.Manager) {
	router := engine.Group("/rdn", middleware.Authenticate(manager))
	router.GET("/existing", r.brivaRequestHandler.GetRdnExisting)
	router.GET("/new", r.brivaRequestHandler.GetRdnNew)
	router.POST("/existing", r.brivaRequestHandler.UpdateRdnExisting)
	router.POST("/new", r.brivaRequestHandler.UpdateRdnNew)
}
