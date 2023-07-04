package user

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/middleware"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/utils/crypto"
	"nashrul-be/crm/utils/session"
)

type Route struct {
	actorRequestHandler RequestHandlerInterface
}

func NewRoute(actorRepository repositories.ActorRepositoryInterface,
	roleRepository repositories.RoleRepositoryInterface,
	hash crypto.Hash,
) Route {
	useCase := NewUseCase(actorRepository, roleRepository, hash)
	actorController := NewController(useCase)
	requestHandler := NewRequestHandler(actorController)
	return Route{actorRequestHandler: requestHandler}
}

func (r Route) Handle(router *gin.Engine, sessionManager session.Manager) {
	router.PATCH("/me/password", middleware.Authenticate(sessionManager), r.actorRequestHandler.UpdatePasswordUser)
	router.PATCH("/me", middleware.Authenticate(sessionManager), r.actorRequestHandler.UpdateProfile)
	router.GET("/me", middleware.Authenticate(sessionManager), r.actorRequestHandler.ProfileUser)
	actor := router.Group("/users", middleware.Authenticate(sessionManager))
	actor.GET("/:username", middleware.AuthorizationAdminOnly(), r.actorRequestHandler.GetByUsername)
	actor.GET("", middleware.AuthorizationAdminOnly(), r.actorRequestHandler.GetAll)
	actor.POST("", middleware.AuthorizationAdminOnly(), r.actorRequestHandler.CreateUser)
	actor.PATCH("/:username", middleware.AuthorizationAdminOnly(), r.actorRequestHandler.UpdateUser)
	actor.DELETE("/:username", middleware.AuthorizationAdminOnly(), r.actorRequestHandler.DeleteUser)
}
