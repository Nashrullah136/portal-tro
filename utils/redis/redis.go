package redisUtils

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

func Connect() (*redis.Client, error) {
	var (
		redisConn *redis.Client
		err       error
	)
	for i := 0; i < 10; i++ {
		redisConn = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		})
		if err = redisConn.Ping(context.Background()).Err(); err == nil {
			break
		}
	}
	return redisConn, err
}
