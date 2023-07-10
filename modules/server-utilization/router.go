package server_utilization

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/middleware"
	"nashrul-be/crm/utils/session"
	"nashrul-be/crm/utils/zabbix"
)

type Route struct {
	requestHandler RequestHandlerInterface
}

func NewRoute(cache zabbix.Cache, api zabbix.API) Route {
	serverUtilController := NewController(cache, api)
	serverUtilRequestHandler := NewRequestHandler(serverUtilController)
	return Route{requestHandler: serverUtilRequestHandler}
}

func (r Route) Handle(engine *gin.Engine, manager session.Manager) {
	router := engine.Group("/server-utilization", middleware.Authenticate(manager), middleware.CheckNewUser(), middleware.AuthorizationUserOnly())
	router.GET("/latest-data", r.requestHandler.GetLatestData)
	router.GET("/update-host", r.requestHandler.UpdateHostList)
}
