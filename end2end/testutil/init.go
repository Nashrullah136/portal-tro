package testutil

import (
	"github.com/gavv/httpexpect/v2"
	redisUtils "nashrul-be/crm/utils/redis"
	"testing"
)

func InitTest(t *testing.T) (*httpexpect.Expect, error) {
	db, err := GetConn()
	if err != nil {
		return nil, err
	}
	if err = LoadEnv(); err != nil {
		return nil, err
	}
	redisConn, err := redisUtils.Connect()
	if err != nil {
		return nil, err
	}
	engine, err := SetUpGin(db, redisConn)
	if err != nil {
		return nil, err
	}
	return SetHttpExpect(t, engine), nil
}
