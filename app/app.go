package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"nashrul-be/crm/modules/audit"
	"nashrul-be/crm/modules/authentication"
	"nashrul-be/crm/modules/user"
	"nashrul-be/crm/repositories"
)

func Handle(dbConn *gorm.DB, engine *gin.Engine) {
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
}
