package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"nashrul-be/crm/modules/audit"
	"nashrul-be/crm/modules/authentication"
	"nashrul-be/crm/modules/user"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/utils/db"
	"nashrul-be/crm/utils/translate"
	"reflect"
	"strings"
	"time"
)

func registerTranslator() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		if err := en.RegisterDefaultTranslations(v, translate.DefaultTranslator()); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	if err := registerTranslator(); err != nil {
		panic(err.Error())
	}

	engine := gin.Default()

	var dbConn *gorm.DB
	var err error
	for i := 0; i < 10; i++ {
		dbConn, err = db.DefaultConnection()
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		panic(err.Error())
	}

	actorRepo := repositories.NewActorRepository(dbConn)
	roleRepo := repositories.NewRoleRepository(dbConn)
	auditRepo := repositories.NewAuditRepository(dbConn)

	actorRoute := user.NewRoute(actorRepo, roleRepo)
	actorRoute.Handle(engine)

	auditRoute := audit.NewRoute(auditRepo)
	auditRoute.Handle(engine)

	actorUseCase := user.NewUseCase(actorRepo, roleRepo)
	auditUseCase := audit.NewUseCase(auditRepo)
	authRoute := authentication.NewRoute(actorUseCase, auditUseCase)
	authRoute.Handle(engine)

	if err := engine.Run(); err != nil {
		panic(err.Error())
	}
}
