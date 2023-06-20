package session

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Session struct {
	redisConn *redis.Client
	Key       string
}

func (s Session) UpdateExpire(second int) error {
	return s.redisConn.Expire(context.Background(), s.Key, time.Duration(second)*time.Second).Err()
}

func (s Session) Set(key string, val string) error {
	return s.redisConn.HSet(context.Background(), s.Key, key, val).Err()
}

func (s Session) Get(key string) (string, error) {
	redisResult := s.redisConn.HGet(context.Background(), s.Key, key)
	if redisResult.Err() != nil {
		return "", redisResult.Err()
	}
	return redisResult.Val(), nil
}
