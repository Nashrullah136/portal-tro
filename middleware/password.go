package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"nashrul-be/crm/dto"
	"nashrul-be/crm/utils"
	"net/http"
)

func CheckNewUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		actor, err := utils.GetUser(c)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		if actor.IsNewUser() {
			log.Println(actor, actor.IsNewUser())
			c.AbortWithStatusJSON(http.StatusForbidden, dto.NeedChangePassword())
		}
	}
}
