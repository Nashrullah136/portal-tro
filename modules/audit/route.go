package audit

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/middleware"
	"nashrul-be/crm/repositories"
)

func NewRoute(auditRepo repositories.AuditRepositoryInterface) Route {
	auditUseCase := NewUseCase(auditRepo)
	auditController := NewController(auditUseCase)
	auditRequestHandler := NewRequestHandler(auditController)
	return Route{auditRequestHandler: auditRequestHandler}
}

type Route struct {
	auditRequestHandler RequestHandlerInterface
}

func (r Route) Handle(router *gin.Engine) {
	auditRoute := router.Group("/audits", middleware.Authenticate())
	auditRoute.GET("", r.auditRequestHandler.GetAll) //TODO: Add page, total page,
	auditRoute.POST("", r.auditRequestHandler.CreateAudit)
}
