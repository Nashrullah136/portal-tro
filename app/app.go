package app

import (
	"github.com/adjust/rmq/v5"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"nashrul-be/crm/modules/audit"
	"nashrul-be/crm/modules/authentication"
	"nashrul-be/crm/modules/briva"
	"nashrul-be/crm/modules/configuration"
	exportCsv "nashrul-be/crm/modules/export-csv"
	"nashrul-be/crm/modules/user"
	"nashrul-be/crm/modules/worker"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/utils/filesystem"
	redisUtils "nashrul-be/crm/utils/redis"
	"nashrul-be/crm/utils/session"
	"os"
)

func Handle(dbMain *gorm.DB, dbBriva *gorm.DB, engine *gin.Engine, sessionManager session.Manager, redisConn rmq.Connection) error {
	actorRepo := repositories.NewActorRepository(dbMain)
	roleRepo := repositories.NewRoleRepository(dbMain)
	auditRepo := repositories.NewAuditRepository(dbMain)
	exportCsvRepo := repositories.NewExportCsvRepository(dbMain)
	brivaRepo := repositories.NewBrivaRepository(dbBriva)

	reportFolder := filesystem.NewFolder(os.Getenv("EXPORT_CSV_FOLDER"))

	exportCsvWorker := worker.NewExportCSV(auditRepo, exportCsvRepo, reportFolder)

	queueCsv, err := redisUtils.MakeQueue(redisConn, "csv-export", "csv-export-worker", 10, exportCsvWorker)
	if err != nil {
		return err
	}

	actorRoute := user.NewRoute(actorRepo, roleRepo)
	actorRoute.Handle(engine, sessionManager)

	auditRoute := audit.NewRoute(auditRepo, exportCsvRepo, queueCsv)
	auditRoute.Handle(engine, sessionManager)

	actorUseCase := user.NewUseCase(actorRepo, roleRepo)
	auditUseCase := audit.NewUseCase(auditRepo, exportCsvRepo, queueCsv)
	authRoute := authentication.NewRoute(actorUseCase, auditUseCase, sessionManager)
	authRoute.Handle(engine)

	exportCsvRoute := exportCsv.NewRoute(exportCsvRepo, auditRepo, reportFolder)
	exportCsvRoute.Handle(engine, sessionManager)

	auditWorker := worker.NewAudit(auditRepo)

	queueAudit, err := redisUtils.MakeQueue(redisConn, "audit-log", "audit-log-worker", 10, auditWorker)

	brivaRoute := briva.NewRoute(brivaRepo, auditRepo, queueAudit)
	brivaRoute.Handle(engine, sessionManager)

	configRequestHandler := configuration.NewRequestHandler()
	configRoute := configuration.NewRoute(configRequestHandler)
	configRoute.Handle(engine, sessionManager)
	return nil
}
