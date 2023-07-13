package main

import (
	"fmt"
	"github.com/adjust/rmq/v5"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	"log"
	"nashrul-be/crm/app"
	"nashrul-be/crm/middleware"
	"nashrul-be/crm/utils/db"
	"nashrul-be/crm/utils/logutils"
	redisUtils "nashrul-be/crm/utils/redis"
	"nashrul-be/crm/utils/session"
	"nashrul-be/crm/utils/translate"
	"nashrul-be/crm/utils/zabbix"
	"os"
	"time"
)

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

func main() {
	errChan := make(chan error, 10)
	go logErrors(errChan)

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if err := logutils.Init(); err != nil {
		log.Panicf("Can't init log. error: %s", err)
	}

	if err := translate.RegisterTranslator(); err != nil {
		panic(err)
	}
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engine.Use(middleware.CORS())

	dbMain, err := db.Connect("TRO")
	if err != nil {
		panic(err)
	}
	logutils.Get().Println("Success connect to DB TRO")

	dbBriva, err := db.Connect("BRIVA")
	if err != nil {
		panic(err)
	}
	logutils.Get().Println("Success connect to DB BRIVA")

	//dbRdn, err := db.Connect("TRO")
	//if err != nil {
	//	panic(err)
	//}

	dbSpan, err := db.Connect("SPAN")
	if err != nil {
		panic(err)
	}
	logutils.Get().Println("Success connect to DB SPAN")

	redisConn, err := redisUtils.Connect()
	if err != nil {
		panic(err)
	}

	sessionManager := session.NewManager(redisConn)

	messageQueue, err := rmq.OpenConnectionWithRedisClient("default-client", redisConn, errChan)
	if err != nil {
		panic(err)
	}

	zabbixServer := zabbix.NewServer(os.Getenv("ZABBIX_URL"), os.Getenv("ZABBIX_USERNAME"), os.Getenv("ZABBIX_PASSWORD"))
	zabbixApi := zabbix.NewAPI(zabbixServer)
	logutils.Get().Println("Success login to zabbix server")

	zabbixCache := zabbix.NewCache()

	wib, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := gocron.NewScheduler(wib)

	if err = app.Handle(dbMain, dbBriva, dbMain, dbSpan, engine, sessionManager, messageQueue, zabbixApi, zabbixCache, scheduler); err != nil {
		panic(err)
	}

	urlServe := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	logutils.Get().Printf("Serve on %s\n", urlServe)
	if err = engine.Run(urlServe); err != nil {
		panic(err)
	}
}
