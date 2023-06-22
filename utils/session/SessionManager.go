package session

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
)

const Name = "SESSION_ID"

type Manager struct {
	redisConn *redis.Client
}

func NewManager(redisConn *redis.Client) Manager {
	return Manager{redisConn: redisConn}
}

func (m Manager) Create() (*Session, error) {
	randByte := make([]byte, 32)
	_, err := rand.Read(randByte)
	if err != nil {
		return nil, err
	}
	key := base64.StdEncoding.EncodeToString(randByte)
	return &Session{
		redisConn: m.redisConn,
		Key:       key,
	}, nil
}

func (m Manager) Get(c *gin.Context) (*Session, error) {
	cookie, err := c.Cookie(Name)
	if err != nil {
		return nil, err
	}
	redisLen := m.redisConn.HLen(c.Copy(), cookie)
	if redisLen.Err() != nil {
		return nil, err
	}
	if redisLen.Val() == 0 {
		return nil, errors.New("cookies doesn't exist")
	}
	return &Session{
		redisConn: m.redisConn,
		Key:       cookie,
	}, nil
}

func (m Manager) Delete(c *gin.Context) (string, error) {
	cookie, err := c.Cookie(Name)
	if err != nil {
		log.Println("cookie doesn't exist")
		return "", nil
	}
	result := m.redisConn.Del(context.Background(), cookie)
	if result.Err() != nil {
		return "", result.Err()
	}
	return cookie, nil
}
