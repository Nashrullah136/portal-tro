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
	me := router.Group("/me", middleware.Authenticate(sessionManager),
		middleware.Refresh(sessionManager))
	me.PATCH("/password", r.actorRequestHandler.UpdatePasswordUser)
	me.PATCH("", middleware.CheckNewUser(), r.actorRequestHandler.UpdateProfile)
	me.GET("", r.actorRequestHandler.ProfileUser)
	actor := router.Group("/users", middleware.Authenticate(sessionManager),
		middleware.Refresh(sessionManager), middleware.AuthorizationAdminOnly())
	actor.GET("/:username", r.actorRequestHandler.GetByUsername)
	actor.GET("", r.actorRequestHandler.GetAll)
	actor.POST("", r.actorRequestHandler.CreateUser)
	actor.PATCH("/:username", r.actorRequestHandler.UpdateUser)
	actor.DELETE("/:username", r.actorRequestHandler.DeleteUser)
}
