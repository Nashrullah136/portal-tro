package migrations

import (
	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/gorm"
	"log"
	"os"
	"path"
	"strconv"
)

func Migrate(db *gorm.DB) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Can't find current working directory")
		return
	}
	isMigrate, err := strconv.ParseBool(os.Getenv("MIGRATE"))
	if err != nil {
		log.Printf("Invalid MIGRATE valua of enviroment variable")
		return
	}
	if isMigrate {
		migrationSource := &migrate.FileMigrationSource{
			Dir: path.Join(dir, "migrations"),
		}
		sqlDb, _ := db.DB()
		n, err := migrate.Exec(sqlDb, "mysql", migrationSource, migrate.Up)
		if err != nil {
			log.Fatal("Can't do migration!\n", err.Error())
		}
		log.Printf("Success applied %d migrations", n)
	}
}
