package user

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/middleware"
	"nashrul-be/crm/repositories"
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

func (r Route) Handle(router *gin.Engine) {
	router.PATCH("/me", middleware.Authenticate(), r.actorRequestHandler.UpdatePasswordUser)
	actor := router.Group("/users", middleware.Authenticate())
	actor.GET("/:username", r.actorRequestHandler.GetByUsername)
	actor.GET("", r.actorRequestHandler.GetAll) //TODO: set page and perpage default when only page and perpage is wrong, total page and current page
	actor.POST("", middleware.AuthorizationAdminOnly(), r.actorRequestHandler.CreateUser)
	actor.PATCH("/:username", middleware.AuthorizationAdminOnly(), r.actorRequestHandler.UpdateUser)
	actor.DELETE("/:username", middleware.AuthorizationAdminOnly(), r.actorRequestHandler.DeleteUser)
}
