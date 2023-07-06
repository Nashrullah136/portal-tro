package user

import (
	"nashrul-be/crm/entities"
)

func mapCreateRequestToActor(request CreateRequest) entities.User {
	var role uint
	switch request.Role {
	case "admin":
		role = 1
	case "user":
		role = 2
	default:
		role = 2
	}
	return entities.User{
		Name:     request.Name,
		Username: request.Username,
		Password: request.Password,
		RoleID:   role,
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
		NewUser:   actor.IsNewUser(),
	}
}

func mapChangePasswordToUser(request ChangePasswordRequest) entities.User {
	return entities.User{
		Username: request.Username,
		Password: request.Password,
	}
}

func mapUpdateProfileToUser(profile UpdateProfile) entities.User {
	return entities.User{
		Name:     profile.Name,
		Username: profile.Username,
	}
}
