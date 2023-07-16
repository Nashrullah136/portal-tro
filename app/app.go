package app

import (
	"fmt"
	"github.com/adjust/rmq/v5"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"nashrul-be/crm/middleware"
	"nashrul-be/crm/modules/audit"
	"nashrul-be/crm/modules/authentication"
	"nashrul-be/crm/modules/briva"
	"nashrul-be/crm/modules/configuration"
	exportCsv "nashrul-be/crm/modules/export-csv"
	serverUtilization "nashrul-be/crm/modules/server-utilization"
	"nashrul-be/crm/modules/span"
	"nashrul-be/crm/modules/user"
	"nashrul-be/crm/modules/worker"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/utils/crypto"
	"nashrul-be/crm/utils/db"
	"nashrul-be/crm/utils/filesystem"
	"nashrul-be/crm/utils/logutils"
	redisUtils "nashrul-be/crm/utils/redis"
	"nashrul-be/crm/utils/session"
	"nashrul-be/crm/utils/translate"
	"nashrul-be/crm/utils/zabbix"
	"net/http"
	"os"
	"time"
)

func Init(envPath string) *http.Server {
	errChan := make(chan error, 10)
	go logErrors(errChan)

	if err := godotenv.Load(envPath); err != nil {
		log.Panicf("can't load envPath.\nenv path : %s.\nerror: %s", envPath, err)
	}

	if !isDir(os.Getenv("LOG_FOLDER")) {
		log.Panicln("log path is not directory/folder.")
	}
	if err := logutils.Init(os.Getenv("LOG_FOLDER")); err != nil {
		log.Panicf("Can't init log. error: %s", err)
	}

	if !isDir(os.Getenv("EXPORT_CSV_FOLDER")) {
		log.Panicln("report path is not directory/folder.")
	}
	reportFolder := filesystem.NewFolder(os.Getenv("EXPORT_CSV_FOLDER"))

	if err := translate.RegisterTranslator(); err != nil {
		logutils.Get().Panicf("can't register translator. error: %s\n", err)
	}
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engine.Use(middleware.CORS())

	dbMain, err := db.Connect("TRO")
	if err != nil {
		logutils.Get().Panicf("can't connect to DB TRO. erro:r %s\n", err)
	}
	logutils.Get().Println("Success connect to DB TRO")

	dbBriva, err := db.Connect("BRIVA")
	if err != nil {
		logutils.Get().Panicf("Can't connect to DB BRIVA. error: %s\n", err)
	}
	logutils.Get().Println("Success connect to DB BRIVA")

	//dbRdn, err := db.Connect("TRO")
	//if err != nil {
	//	panic(err)
	//}

	dbSpan, err := db.Connect("SPAN")
	if err != nil {
		logutils.Get().Panicf("Can't connect to DB SPAN. error: %s\n", err)
	}
	logutils.Get().Println("Success connect to DB SPAN")

	redisConn, err := redisUtils.Connect()
	if err != nil {
		logutils.Get().Panicf("Can't connect to redis. error: %s\n", err)
	}

	sessionManager := session.NewManager(redisConn)

	messageQueue, err := rmq.OpenConnectionWithRedisClient("default-client", redisConn, errChan)
	if err != nil {
		logutils.Get().Panicf("Can't make connection rmq. error: %s\n", err)
	}

	zabbixServer := zabbix.NewServer(os.Getenv("ZABBIX_URL"), os.Getenv("ZABBIX_USERNAME"), os.Getenv("ZABBIX_PASSWORD"))
	zabbixApi := zabbix.NewAPI(zabbixServer)

	zabbixCache := zabbix.NewCache()

	wib, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := gocron.NewScheduler(wib)

	if err = Handle(dbMain, dbBriva, dbMain, dbSpan, engine,
		sessionManager, messageQueue, zabbixApi, zabbixCache, scheduler, reportFolder); err != nil {
		logutils.Get().Panicf("Error when handle call app.Handle. error: %s\n", err)
	}

	urlServe := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	logutils.Get().Printf("Serve on %s\n", urlServe)
	srv := &http.Server{
		Addr:    urlServe,
		Handler: engine,
	}
	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logutils.Get().Panicf("Can't serve. error: %s\n", err)
		}
	}()

	return srv
}

func Handle(dbMain *gorm.DB, dbBriva *gorm.DB, dbRdn *gorm.DB, dbSpan *gorm.DB,
	engine *gin.Engine, sessionManager session.Manager, redisConn rmq.Connection,
	zabbixApi zabbix.API, cache zabbix.Cache, scheduler *gocron.Scheduler, reportFolder filesystem.Folder) error {

	actorRepo := repositories.NewActorRepository(dbMain)
	roleRepo := repositories.NewRoleRepository(dbMain)
	auditRepo := repositories.NewAuditRepository(dbMain)
	exportCsvRepo := repositories.NewExportCsvRepository(dbMain)
	brivaRepo := repositories.NewBrivaRepository(dbBriva)
	//rdnRepo := repositories.NewRdnRepository(dbRdn)
	spanRepo := repositories.NewSpanRepository(dbSpan)

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

	if _, err = scheduler.Every(1).Day().At("00:00").Do(worker.CleanerCsv(reportFolder)); err != nil {
		return err
	}
	if _, err = scheduler.Every(1).Day().At("00:00").Do(logutils.Init, os.Getenv("LOG_FOLDER")); err != nil {
		return err
	}

	//rdnRoute := rdn.NewRoute(rdnRepo, auditRepo, queueAudit)
	//rdnRoute.Handle(engine, sessionManager)

	spanRoute := span.NewRoute(spanRepo, auditRepo, queueAudit)
	spanRoute.Handle(engine, sessionManager)

	if os.Getenv("SERVER_UTIL") != "off" {
		serverUtilController := serverUtilization.NewController(cache, zabbixApi)
		if err = serverUtilController.RefreshHostList(); err != nil {
			return err
		}
		serverUtilRoute := serverUtilization.NewRoute(cache, zabbixApi)
		serverUtilRoute.Handle(engine, sessionManager)

		updateLastData := worker.UpdateLastDataServerUtil(cache, zabbixApi)
		updateLastData()
		if _, err = scheduler.Every(1).Minute().Do(updateLastData); err != nil {
			return err
		}
	}

	if err := redisUtils.CreateCleaner(redisConn, scheduler); err != nil {
		return err
	}
	scheduler.StartAsync()
	return nil
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		logutils.Get().Printf("error when access path. error : %v\n", err)
	}
	return info.IsDir()
}

func logErrors(errChan <-chan error) {
	for err := range errChan {
		switch err := err.(type) {
		case *rmq.HeartbeatError:
			if err.Count == rmq.HeartbeatErrorLimit {
				logutils.Get().Print("heartbeat error (limit): ", err)
			} else {
				logutils.Get().Print("heartbeat error: ", err)
			}
		case *rmq.ConsumeError:
			logutils.Get().Print("consume error: ", err)
		case *rmq.DeliveryError:
			logutils.Get().Print("delivery error: ", err.Delivery, err)
		default:
			logutils.Get().Print("other error: ", err)
		}
	}
}
