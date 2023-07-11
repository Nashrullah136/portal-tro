package app

import (
	"github.com/adjust/rmq/v5"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"gorm.io/gorm"
	"nashrul-be/crm/modules/audit"
	"nashrul-be/crm/modules/authentication"
	"nashrul-be/crm/modules/briva"
	"nashrul-be/crm/modules/configuration"
	exportCsv "nashrul-be/crm/modules/export-csv"
	"nashrul-be/crm/modules/rdn"
	server_utilization "nashrul-be/crm/modules/server-utilization"
	"nashrul-be/crm/modules/span"
	"nashrul-be/crm/modules/user"
	"nashrul-be/crm/modules/worker"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/utils/crypto"
	"nashrul-be/crm/utils/filesystem"
	redisUtils "nashrul-be/crm/utils/redis"
	"nashrul-be/crm/utils/session"
	"nashrul-be/crm/utils/zabbix"
	"os"
)

func Handle(dbMain *gorm.DB, dbBriva *gorm.DB, dbRdn *gorm.DB, dbSpan *gorm.DB,
	engine *gin.Engine, sessionManager session.Manager, redisConn rmq.Connection,
	zabbixApi zabbix.API, cache zabbix.Cache, scheduler *gocron.Scheduler) error {

	actorRepo := repositories.NewActorRepository(dbMain)
	roleRepo := repositories.NewRoleRepository(dbMain)
	auditRepo := repositories.NewAuditRepository(dbMain)
	exportCsvRepo := repositories.NewExportCsvRepository(dbMain)
	brivaRepo := repositories.NewBrivaRepository(dbBriva)
	rdnRepo := repositories.NewRdnRepository(dbRdn)
	spanRepo := repositories.NewSpanRepository(dbSpan)

	reportFolder := filesystem.NewFolder(os.Getenv("EXPORT_CSV_FOLDER"))

	exportCsvWorker := worker.NewExportCSV(auditRepo, exportCsvRepo, reportFolder)

	queueCsv, err := redisUtils.MakeQueue(redisConn, "csv-export", "csv-export-worker", 10, exportCsvWorker)
	if err != nil {
		return err
	}

	bcryptHash := crypto.NewBcryptHash()

	actorRoute := user.NewRoute(actorRepo, roleRepo, bcryptHash)
	actorRoute.Handle(engine, sessionManager)

	auditRoute := audit.NewRoute(auditRepo, exportCsvRepo, queueCsv, reportFolder)
	auditRoute.Handle(engine, sessionManager)

	actorUseCase := user.NewUseCase(actorRepo, roleRepo, bcryptHash)
	auditUseCase := audit.NewUseCase(auditRepo, exportCsvRepo, queueCsv, reportFolder)
	authRoute := authentication.NewRoute(actorUseCase, auditUseCase, sessionManager, bcryptHash)
	authRoute.Handle(engine)

	exportCsvRoute := exportCsv.NewRoute(exportCsvRepo, auditRepo, reportFolder)
	exportCsvRoute.Handle(engine, sessionManager)

	auditWorker := worker.NewAudit(auditRepo)

	queueAudit, err := redisUtils.MakeQueue(redisConn, "audit-log", "audit-log-worker", 10, auditWorker)
	if err != nil {
		return err
	}

	brivaRoute := briva.NewRoute(brivaRepo, auditRepo, queueAudit)
	brivaRoute.Handle(engine, sessionManager)

	configRequestHandler := configuration.NewRequestHandler()
	configRoute := configuration.NewRoute(configRequestHandler)
	configRoute.Handle(engine, sessionManager)

	exportCsvUseCase := exportCsv.NewUseCase(exportCsvRepo, auditRepo, reportFolder)
	if _, err = scheduler.Every(1).Day().At("00:00").Do(worker.CleanerCsv(exportCsvUseCase)); err != nil {
		return err
	}

	rdnRoute := rdn.NewRoute(rdnRepo, auditRepo, queueAudit)
	rdnRoute.Handle(engine, sessionManager)

	spanRoute := span.NewRoute(spanRepo, auditRepo, queueAudit)
	spanRoute.Handle(engine, sessionManager)

	serverUtilController := server_utilization.NewController(cache, zabbixApi)
	if err = serverUtilController.RefreshHostList(); err != nil {
		return err
	}
	serverUtilRoute := server_utilization.NewRoute(cache, zabbixApi)
	serverUtilRoute.Handle(engine, sessionManager)

	updateLastData := worker.UpdateLastDataServerUtil(cache, zabbixApi)
	updateLastData()
	if _, err = scheduler.Every(1).Minute().Do(updateLastData); err != nil {
		return err
	}

	scheduler.StartAsync()
	return nil
}
