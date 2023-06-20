package user

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/middleware"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/utils/session"
)

type Route struct {
	actorRequestHandler RequestHandlerInterface
}

func NewRoute(actorRepository repositories.ActorRepositoryInterface,
	roleRepository repositories.RoleRepositoryInterface,
) Route {
	useCase := NewUseCase(actorRepository, roleRepository)
	actorController := NewController(useCase)
	requestHandler := NewRequestHandler(actorController)
	return Route{actorRequestHandler: requestHandler}
}

func (r Route) Handle(router *gin.Engine, sessionManager session.Manager) {
	router.PATCH("/me", middleware.Authenticate(sessionManager), r.actorRequestHandler.UpdatePasswordUser)
	actor := router.Group("/users", middleware.Authenticate(sessionManager))
	actor.GET("/:username", r.actorRequestHandler.GetByUsername)
	actor.GET("", r.actorRequestHandler.GetAll)
	actor.POST("", middleware.AuthorizationAdminOnly(), r.actorRequestHandler.CreateUser)
	actor.PATCH("/:username", middleware.AuthorizationAdminOnly(), r.actorRequestHandler.UpdateUser)
	actor.DELETE("/:username", middleware.AuthorizationAdminOnly(), r.actorRequestHandler.DeleteUser)
}
