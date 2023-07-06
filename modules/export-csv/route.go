package export_csv

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/middleware"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/utils/filesystem"
	"nashrul-be/crm/utils/session"
)

type Route struct {
	csvRequestHandler RequestHandlerInterface
}

func NewRoute(
	exportCsvRepo repositories.ExportCsvRepositoryInterface,
	auditRepo repositories.AuditRepositoryInterface,
	folder filesystem.Folder,
) Route {
	csvUseCase := NewUseCase(exportCsvRepo, auditRepo, folder)
	csvController := NewController(csvUseCase)
	csvRequestHandler := NewRequestHandler(csvController)
	return Route{csvRequestHandler: csvRequestHandler}
}

func (r Route) Handle(engine *gin.Engine, manager session.Manager) {
	route := engine.Group("/exports/csv", middleware.Authenticate(manager),
		middleware.Refresh(manager), middleware.CheckNewUser())
	route.GET("/:id", r.csvRequestHandler.DownloadCsv)
	route.GET("", r.csvRequestHandler.GetAll)
}
