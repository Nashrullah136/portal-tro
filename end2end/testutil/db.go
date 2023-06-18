package testutil

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func GetConn() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("tro.db"), &gorm.Config{})
}
