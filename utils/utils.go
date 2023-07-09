package utils

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/entities"
)

func GetUser(c *gin.Context) (entities.User, error) {
	return GetUserFromContext(c.Copy())
}

func GetUserFromContext(ctx context.Context) (entities.User, error) {
	userContext := ctx.Value("user")
	if userContext == nil {
		return entities.User{}, errors.New("can't get user from context")
	}
	user, ok := userContext.(*entities.User)
	if !ok {
		return entities.User{}, errors.New("user in context is not type of entities.user ")
	}
	return *user, nil
}

func SetUser(c *gin.Context, user entities.User) {
	c.Set("user", &user)
}

func SetUserToContext(ctx context.Context, user entities.User) context.Context {
	return context.WithValue(ctx, "user", &user)
}
