package middleware

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/dto"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/utils"
	"net/http"
)

func AuthorizationWithRole(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		actor, err := utils.GetUser(c)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		found := false
		for _, role := range roles {
			if role == actor.Role.RoleName {
				found = true
				break
			}
		}
		if !found {
			c.AbortWithStatusJSON(http.StatusForbidden, dto.ErrorForbidden())
		}
	}
}

func AuthorizationAdminOnly() gin.HandlerFunc {
	return AuthorizationWithRole([]string{entities.ADMIN})
}

func AuthorizationUserOnly() gin.HandlerFunc {
	return AuthorizationWithRole([]string{entities.USER})
}
