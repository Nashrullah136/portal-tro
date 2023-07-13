package middleware

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/dto"
	"nashrul-be/crm/utils"
	"nashrul-be/crm/utils/logutils"
	"net/http"
)

func CheckNewUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		actor, err := utils.GetUser(c)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		if actor.IsNewUser() {
			logutils.Get().Println(actor, actor.IsNewUser())
			c.AbortWithStatusJSON(http.StatusForbidden, dto.NeedChangePassword())
		}
	}
}
