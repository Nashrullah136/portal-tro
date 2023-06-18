package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func DsnWithConfig(config Config) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.User, config.Password,
		config.Host, config.Port, config.DBName)
	return dsn
}

func DsnWithEnv() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	return dsn
}

func Connect(dsn string) (*gorm.DB, error) {
	var dbConn *gorm.DB
	var err error
	for i := 0; i < 10; i++ {
		dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}
