package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"nashrul-be/crm/dto"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/utils"
	"nashrul-be/crm/utils/session"
	"net/http"
	"os"
	"strconv"
)

func Authenticate(manager session.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentSession, err := manager.Get(c)
		if err != nil {
			log.Println(err.Error())
			if errors.Is(err, session.ErrNotExist) {
				c.Header("Access-Control-Allow-Credentials", "true")
				c.SetCookie(session.Name, currentSession.Key, -1, "/", os.Getenv("DOMAIN"), false, true)
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorUnauthorizedDefault())
			return
		}
		accountJson, err := currentSession.Get("user")
		if err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorUnauthorizedDefault())
			return
		}
		var user entities.User
		if err := json.Unmarshal([]byte(accountJson), &user); err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
			return
		}
		utils.SetUser(c, user)
	}
}

func Refresh(manager session.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentSession, _ := manager.Get(c)
		duration, _ := strconv.Atoi(os.Getenv("SESSION_DURATION"))
		if err := currentSession.UpdateExpire(duration); err != nil {
			log.Println(fmt.Sprintf("Can't update redis expire. error: %s", err))
		}
	}
}
