package session

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"nashrul-be/crm/entities"
	"os"
	"strings"
	"sync"
)

const Name = "SESSION_ID"

var ErrNotExist = errors.New("session not found")
var createLock = &sync.Mutex{}

type Manager struct {
	redisConn *redis.Client
}

func NewManager(redisConn *redis.Client) Manager {
	return Manager{redisConn: redisConn}
}

func (m Manager) generatePrefix(username string) string {
	rest := 3 - (len(username) % 3)
	prefix := username + strings.Repeat(":", rest)
	return base64.StdEncoding.EncodeToString([]byte(prefix))
}

func (m Manager) Create(user entities.User) (*Session, error) {
	prefix := m.generatePrefix(user.Username)
	createLock.Lock()
	defer createLock.Unlock()
	if os.Getenv("ONE_USER_ONE_SESSION") != "false" {
		if keys := m.redisConn.Keys(context.Background(), prefix+"*"); len(keys.Val()) > 0 {
			return nil, errors.New("user already login")
		}
	}
	randByte := make([]byte, 32)
	_, err := rand.Read(randByte)
	if err != nil {
		return nil, err
	}
	key := base64.StdEncoding.EncodeToString(randByte)
	return &Session{
		redisConn: m.redisConn,
		Key:       prefix + key,
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
		return &Session{
			Key: cookie,
		}, ErrNotExist
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
