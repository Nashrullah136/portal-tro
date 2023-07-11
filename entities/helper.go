package entities

import (
	"context"
	"errors"
	"nashrul-be/crm/utils/auditUtils"
)

func MapAuditResultToAuditEntities(result auditUtils.Result) Audit {
	return Audit{
		DateTime:   result.DateTime,
		Username:   result.Username,
		Action:     result.Action,
		Entity:     result.Entity,
		EntityID:   result.EntityID,
		DataBefore: result.DataBefore,
		DataAfter:  result.DataAfter,
	}
}

func getUserFromContext(ctx context.Context) (User, error) {
	userContext := ctx.Value("user")
	if userContext == nil {
		return User{}, errors.New("can't get user from context")
	}
	user, ok := userContext.(*User)
	if !ok {
		return User{}, errors.New("user in context is not type of entities.user ")
	}
	return *user, nil
}
