package middleware

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/entities"
	"net/http"
)

func AuthorizationWithRole(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		actor, ok := c.MustGet("user").(entities.User)
		if !ok {
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
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}

func AuthorizationAdminOnly() gin.HandlerFunc {
	return AuthorizationWithRole([]string{entities.ADMIN})
}
