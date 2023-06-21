package app

import (
	"github.com/adjust/rmq/v5"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"nashrul-be/crm/modules/audit"
	"nashrul-be/crm/modules/authentication"
	"nashrul-be/crm/modules/briva"
	exportCsv "nashrul-be/crm/modules/export-csv"
	"nashrul-be/crm/modules/user"
	"nashrul-be/crm/modules/worker"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/utils/session"
)

func Handle(dbConn *gorm.DB, engine *gin.Engine, sessionManager session.Manager, queue rmq.Queue) error {
	actorRepo := repositories.NewActorRepository(dbConn)
	roleRepo := repositories.NewRoleRepository(dbConn)
	auditRepo := repositories.NewAuditRepository(dbConn)
	exportCsvRepo := repositories.NewExportCsvRepository(dbConn)
	brivaRepo := repositories.NewBrivaRepository(dbConn)

	exportCsvWorker := worker.NewExportCSV(auditRepo, exportCsvRepo)

	if _, err := queue.AddConsumer("default-csv-export-consumer", exportCsvWorker); err != nil {
		return err
	}

	actorRoute := user.NewRoute(actorRepo, roleRepo)
	actorRoute.Handle(engine, sessionManager)

	auditRoute := audit.NewRoute(auditRepo, exportCsvRepo, queue)
	auditRoute.Handle(engine, sessionManager)

	actorUseCase := user.NewUseCase(actorRepo, roleRepo)
	auditUseCase := audit.NewUseCase(auditRepo, exportCsvRepo, queue)
	authRoute := authentication.NewRoute(actorUseCase, auditUseCase, sessionManager)
	authRoute.Handle(engine)

	exportCsvRoute := exportCsv.NewRoute(exportCsvRepo, auditRepo)
	exportCsvRoute.Handle(engine, sessionManager)

	brivaRoute := briva.NewRoute(brivaRepo)
	brivaRoute.Handle(engine, sessionManager)
	return nil
}
