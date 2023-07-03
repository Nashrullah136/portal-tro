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
				log.Print("heartbeat error (limit): ", err)
			} else {
				log.Print("heartbeat error: ", err)
			}
		case *rmq.ConsumeError:
			log.Print("consume error: ", err)
		case *rmq.DeliveryError:
			log.Print("delivery error: ", err.Delivery, err)
		default:
			log.Print("other error: ", err)
		}
	}
}

func main() {
	errChan := make(chan error, 10)
	go logErrors(errChan)

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if err := translate.RegisterTranslator(); err != nil {
		panic(err)
	}
	engine := gin.Default()
	engine.Use(middleware.CORS())

	dsnMain := db.DsnMySQL()
	dbMain, err := db.ConnectMySql(dsnMain)
	if err != nil {
		panic(err)
	}

	dsnBriva := db.DsnSqlServer("BRIVA")
	dbBriva, err := db.ConnectSqlServer(dsnBriva)
	if err != nil {
		panic(err)
	}

	dsnRdn := db.DsnSqlServer("RDN")
	dbRdn, err := db.ConnectSqlServer(dsnRdn)
	if err != nil {
		panic(err)
	}

	dsnSpan := db.DsnSqlServer("SPAN")
	dbSpan, err := db.ConnectSqlServer(dsnSpan)
	if err != nil {
		panic(err)
	}

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
	if err = zabbixServer.Login(); err != nil {
		panic("can't login to zabbix server")
	}
	zabbixApi := zabbix.NewAPI(zabbixServer)

	zabbixCache := zabbix.NewCache()

	wib, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := gocron.NewScheduler(wib)

	if err = app.Handle(dbMain, dbBriva, dbRdn, dbSpan, engine, sessionManager, messageQueue, zabbixApi, zabbixCache, scheduler); err != nil {
		panic(err)
	}

	urlServe := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	if err = engine.Run(urlServe); err != nil {
		panic(err)
	}
}
