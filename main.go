package main

import (
	"fmt"
	"github.com/adjust/rmq/v5"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"nashrul-be/crm/app"
	"nashrul-be/crm/middleware"
	"nashrul-be/crm/utils/db"
	redisUtils "nashrul-be/crm/utils/redis"
	"nashrul-be/crm/utils/session"
	"nashrul-be/crm/utils/translate"
	"os"
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

	redisConn, err := redisUtils.Connect()
	if err != nil {
		panic(err)
	}

	sessionManager := session.NewManager(redisConn)

	messageQueue, err := rmq.OpenConnectionWithRedisClient("default-client", redisConn, errChan)
	if err != nil {
		panic(err)
	}

	if err = app.Handle(dbMain, dbBriva, engine, sessionManager, messageQueue); err != nil {
		panic(err)
	}

	urlServe := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	if err = engine.Run(urlServe); err != nil {
		panic(err)
	}
}
