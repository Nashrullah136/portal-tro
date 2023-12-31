package configuration

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/middleware"
	"nashrul-be/crm/utils/session"
)

type Route struct {
	configRequestHandler RequestHandlerInterface
}

func NewRoute(configRequestHandler RequestHandlerInterface) Route {
	return Route{configRequestHandler: configRequestHandler}
}

func (r Route) Handle(engine *gin.Engine, manager session.Manager) {
	route := engine.Group("/config", middleware.Authenticate(manager),
		middleware.Refresh(manager), middleware.AuthorizationAdminOnly(), middleware.CheckNewUser())
	route.POST("/session-duration", r.configRequestHandler.UpdateSessionDuration)
	route.GET("/session-duration", r.configRequestHandler.GetSessionDuration)
}
