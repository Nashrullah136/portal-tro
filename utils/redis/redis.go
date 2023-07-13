package redisUtils

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"nashrul-be/crm/utils/logutils"
	"os"
)

func Connect() (*redis.Client, error) {
	var (
		redisConn *redis.Client
		err       error
	)
	logutils.Get().Println("Connecting to redis...")
	for i := 0; i < 10; i++ {
		redisConn = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		})
		if err = redisConn.Ping(context.Background()).Err(); err == nil {
			logutils.Get().Println("Success connect to redis")
			break
		}
	}
	return redisConn, err
}
