package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"nashrul-be/crm/utils/localtime"
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

func GormConfig() *gorm.Config {
	return &gorm.Config{
		NowFunc: func() time.Time {
			result := localtime.Now
			return *result()
		},
	}
}

func DsnWithConfig(config Config) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.User, config.Password,
		config.Host, config.Port, config.DBName)
	return dsn
}

func DsnMySQL() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	return dsn
}

func DsnSqlServer(prefix string) string {
	if prefix != "" {
		prefix = prefix + "_"
	}
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s/%s",
		os.Getenv(prefix+"DB_USERNAME"), os.Getenv(prefix+"DB_PASSWORD"),
		os.Getenv(prefix+"DB_HOST"), os.Getenv(prefix+"DB_PORT"), os.Getenv(prefix+"DB_NAME"))
	return dsn
}

func ConnectMySql(dsn string) (*gorm.DB, error) {
	var dbConn *gorm.DB
	var err error
	for i := 0; i < 10; i++ {
		dbConn, err = gorm.Open(mysql.Open(dsn), GormConfig())
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

func ConnectSqlServer(dsn string) (*gorm.DB, error) {
	var dbConn *gorm.DB
	var err error
	for i := 0; i < 10; i++ {
		dbConn, err = gorm.Open(sqlserver.Open(dsn), GormConfig())
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
