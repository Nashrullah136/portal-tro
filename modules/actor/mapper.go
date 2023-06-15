package actor

import (
	"nashrul-be/crm/entities"
)

func mapCreateRequestToActor(request CreateRequest) entities.User {
	return entities.User{
		Name:     request.Name,
		Username: request.Username,
		Password: request.Password,
		RoleID:   2,
	}
}

func mapUpdateRequestToActor(request UpdateRequest) entities.User {
	return entities.User{
		Username: request.Username,
		Name:     request.Name,
		Password: request.Password,
	}
}

func mapActorToResponse(actor entities.User) Representation {
	return Representation{
		Name:      actor.Name,
		Username:  actor.Username,
		Role:      actor.Role.RoleName,
		CreatedAt: actor.CreatedAt,
		CreatedBy: actor.CreatedBy,
		UpdatedAt: actor.UpdatedAt,
		UpdatedBy: actor.UpdatedBy,
	}
}
