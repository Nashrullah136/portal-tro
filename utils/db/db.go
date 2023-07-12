package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"nashrul-be/crm/utils/localtime"
	"os"
	"time"
)

type openFunc func(string) gorm.Dialector

func GormConfig() *gorm.Config {
	return &gorm.Config{
		NowFunc: func() time.Time {
			result := localtime.Now
			return *result()
		},
	}
}

func Connect(prefix string) (connDb *gorm.DB, err error) {
	log.Printf("Connecting to DB %s...\n", prefix)
	prefixResult := prefixRule(prefix)
	switch os.Getenv(prefixResult + "DB_DRIVER") {
	case "mysql":
		dsn := DsnMySQL(prefix)
		connDb, err = ConnectMySql(dsn)
	case "mssql":
		dsn := DsnSqlServer(prefix)
		connDb, err = ConnectSqlServer(dsn)
	}
	return
}

func DsnMySQL(prefix string) string {
	prefix = prefixRule(prefix)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv(prefix+"DB_USERNAME"), os.Getenv(prefix+"DB_PASSWORD"),
		os.Getenv(prefix+"DB_HOST"), os.Getenv(prefix+"DB_PORT"), os.Getenv(prefix+"DB_NAME"))
	return dsn
}

func DsnSqlServer(prefix string) string {
	prefix = prefixRule(prefix)
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s/%s",
		os.Getenv(prefix+"DB_USERNAME"), os.Getenv(prefix+"DB_PASSWORD"),
		os.Getenv(prefix+"DB_HOST"), os.Getenv(prefix+"DB_PORT"), os.Getenv(prefix+"DB_NAME"))
	return dsn
}

func ConnectMySql(dsn string) (*gorm.DB, error) {
	return connectDB(mysql.Open, dsn)
}

func ConnectSqlServer(dsn string) (*gorm.DB, error) {
	return connectDB(sqlserver.Open, dsn)
}

func prefixRule(prefix string) string {
	if prefix != "" {
		prefix = prefix + "_"
	}
	return prefix
}

func connectDB(open openFunc, dsn string) (dbConn *gorm.DB, err error) {
	for i := 0; i < 10; i++ {
		dbConn, err = gorm.Open(open(dsn), GormConfig())
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
