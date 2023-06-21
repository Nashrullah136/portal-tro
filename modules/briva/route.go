package briva

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/middleware"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/utils/session"
)

type Route struct {
	brivaRequestHandler RequestHandlerInterface
}

func NewRoute(brivaRepo repositories.BrivaRepositoryInterface,
) Route {
	brivaUseCase := NewUseCase(brivaRepo)
	brivaController := NewController(brivaUseCase)
	brivaRequestHandler := NewRequestHandler(brivaController)
	return Route{
		brivaRequestHandler: brivaRequestHandler,
	}
}

func (r Route) Handle(engine *gin.Engine, manager session.Manager) {
	router := engine.Group("/briva", middleware.Authenticate(manager))
	router.GET("/:brivano", r.brivaRequestHandler.GetByBrivaNo)
	router.POST("/:brivano", r.brivaRequestHandler.Update)
}
