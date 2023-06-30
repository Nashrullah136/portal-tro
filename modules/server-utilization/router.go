package server_utilization

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/utils/zabbix"
)

type Route struct {
	requestHandler RequestHandler
}

func NewRoute(cache zabbix.Cache, api zabbix.API) Route {
	serverUtilController := NewController(cache, api)
	serverUtilRequestHandler := NewRequestHandler(serverUtilController)
	return Route{requestHandler: serverUtilRequestHandler}
}

func (r Route) Handle(engine *gin.Engine) {
	router := engine.Group("/server-utilization")
	router.GET("/latest-data", r.requestHandler.GetLatestData)
	router.GET("/update-host", r.requestHandler.UpdateHostList)
}
