package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"nashrul-be/crm/dto"
	"nashrul-be/crm/entities"
	jwtUtil "nashrul-be/crm/utils/jwt"
	"net/http"
	"strings"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.GetHeader("Authorization")
		if authToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorUnauthorizedDefault())
			return
		}
		splitToken := strings.Split(authToken, " ")
		if len(splitToken) != 2 && strings.ToLower(splitToken[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorUnauthorizedDefault())
			return
		}
		jwtToken := splitToken[1]
		claims, err := jwtUtil.AuthenticateJWT(jwtToken)
		if err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorUnauthorizedDefault())
			return
		}
		c.Set("user", entities.User{
			ID:       claims.ID,
			Username: claims.Subject,
			Role:     entities.Role{RoleName: claims.Role},
		})
	}
}
