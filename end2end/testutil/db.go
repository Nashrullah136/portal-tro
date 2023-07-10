package testutil

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"nashrul-be/crm/utils/db"
)

func GetConn() (*gorm.DB, error) {
	dsn := db.DsnMySQL()
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
