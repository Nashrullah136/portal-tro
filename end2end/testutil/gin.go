package testutil

import (
	"github.com/adjust/rmq/v5"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"nashrul-be/crm/app"
	"nashrul-be/crm/utils/session"
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

func SetUpGin(db *gorm.DB, redisConn *redis.Client) (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	errChan := make(chan error, 10)
	go logErrors(errChan)
	engine := gin.New()
	sessionManager := session.NewManager(redisConn)

	messageQueue, err := rmq.OpenConnectionWithRedisClient("default-client", redisConn, errChan)
	if err != nil {
		return nil, err
	}

	queue, err := messageQueue.OpenQueue("export-csv")
	if err != nil {
		return nil, err
	}
	if err = queue.StartConsuming(10, 5*time.Second); err != nil {
		return nil, err
	}
	zabbixServer := zabbix.NewServer(os.Getenv("ZABBIX_URL"), os.Getenv("ZABBIX_USERNAME"), os.Getenv("ZABBIX_PASSWORD"))
	//if err = zabbixServer.Login(); err != nil {
	//	panic("can't login to zabbix server")
	//}
	zabbixApi := zabbix.NewAPI(zabbixServer)

	zabbixCache := zabbix.NewCache()

	wib, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := gocron.NewScheduler(wib)
	if err = app.Handle(db, db, db, db, engine, sessionManager, messageQueue, zabbixApi, zabbixCache, scheduler); err != nil {
		return nil, err
	}
	return engine, nil
}
