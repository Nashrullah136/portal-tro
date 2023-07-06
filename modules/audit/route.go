package audit

import (
	"github.com/adjust/rmq/v5"
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/middleware"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/utils/filesystem"
	"nashrul-be/crm/utils/session"
)

func NewRoute(
	auditRepo repositories.AuditRepositoryInterface,
	exportCsvRepo repositories.ExportCsvRepositoryInterface,
	queue rmq.Queue,
	folder filesystem.Folder,
) Route {
	auditUseCase := NewUseCase(auditRepo, exportCsvRepo, queue, folder)
	auditController := NewController(auditUseCase)
	auditRequestHandler := NewRequestHandler(auditController)
	return Route{auditRequestHandler: auditRequestHandler}
}

type Route struct {
	auditRequestHandler RequestHandlerInterface
}

func (r Route) Handle(router *gin.Engine, manager session.Manager) {
	auditRoute := router.Group("/audits", middleware.Authenticate(manager), middleware.Refresh(manager))
	auditRoute.GET("", middleware.CheckNewUser(), middleware.AuthorizationUserOnly(), r.auditRequestHandler.GetAll)
	auditRoute.POST("", r.auditRequestHandler.CreateAudit)
	auditRoute.GET("/export", middleware.CheckNewUser(), middleware.AuthorizationUserOnly(), r.auditRequestHandler.ExportCSV)
}
