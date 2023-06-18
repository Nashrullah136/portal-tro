package testutil

import (
	"github.com/gavv/httpexpect/v2"
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
	engine := SetUpGin(db)
	return SetHttpExpect(t, engine), nil
}
