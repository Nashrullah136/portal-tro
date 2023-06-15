package migrations

import (
	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/gorm"
	"log"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/modules/actor"
	"nashrul-be/crm/repositories"
	"os"
	"path"
	"strconv"
	"time"
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

func Seed(db *gorm.DB, actorRepo repositories.ActorRepositoryInterface,
	roleRepo repositories.RoleRepositoryInterface,
	approvalRepo repositories.RegisterApprovalRepositoryInterface,
) error {
	isMigrate, err := strconv.ParseBool(os.Getenv("MIGRATE"))
	if err != nil {
		log.Printf("Invalid MIGRATE valua of enviroment variable")
		return err
	}
	isSeed, err := strconv.ParseBool(os.Getenv("MIGRATE"))
	if err != nil {
		log.Printf("Invalid MIGRATE valua of enviroment variable")
		return err
	}
	if isSeed || isMigrate {
		superRole := entities.Role{
			ID:       1,
			RoleName: "super_admin",
		}
		adminRole := entities.Role{
			ID:       2,
			RoleName: "admin",
		}
		db.Save(adminRole)
		db.Save(superRole)
		var oldSuperAdmin entities.User
		db.Model(&entities.User{}).InnerJoins("Role").
			Where("role_name = ?", "super_admin").First(&oldSuperAdmin)
		actorUsecase := actor.NewUseCase(actorRepo, roleRepo, approvalRepo)
		if err := actorUsecase.DeleteUser(oldSuperAdmin.ID); err != nil {
			panic("can't delete old super admin")
		}
		superAdminUsername := os.Getenv("SUPER_ADMIN_USERNAME")
		if superAdminUsername == "" {
			superAdminUsername = "super_admin"
		}
		superAdminPassword := os.Getenv("SUPER_ADMIN_PASSWORD")
		if superAdminPassword == "" {
			superAdminPassword = superAdminUsername
		}
		superAdmin := &entities.User{
			Username:  superAdminUsername,
			Password:  superAdminPassword,
			RoleID:    1,
			Verified:  true,
			Active:    true,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}
		if err := actorUsecase.CreateUser(nil, superAdmin); err != nil {
			panic("Can't create super admin")
		}
		approval, err := approvalRepo.GetByAdminID(superAdmin.ID)
		if err != nil {
			panic("Can't update register approval super admin")
		}
		approval.Status = "approved"
		approval.SuperAdminID = superAdmin.ID
		approvalRepo.Update(approval)
	}
	return nil
}
