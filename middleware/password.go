package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"nashrul-be/crm/dto"
	"nashrul-be/crm/entities"
	"net/http"
)

func CheckNewUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		actor, ok := c.MustGet("user").(entities.User)
		if !ok {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		if actor.IsNewUser() {
			log.Println(actor, actor.IsNewUser())
			c.AbortWithStatusJSON(http.StatusForbidden, dto.NeedChangePassword())
		}
	}
}
