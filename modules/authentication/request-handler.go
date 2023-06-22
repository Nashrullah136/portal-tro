package authentication

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/dto"
	"nashrul-be/crm/utils/session"
	"net/http"
	"os"
	"strconv"
)

type RequestHandlerInterface interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

func NewRequestHandler(authController ControllerInterface, manager session.Manager) RequestHandlerInterface {
	return requestHandler{
		authController: authController,
		sessionManager: manager,
	}
}

type requestHandler struct {
	authController ControllerInterface
	sessionManager session.Manager
}

func (h requestHandler) Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest("Invalid Username/Password"))
		return
	}
	account, err := h.authController.Login(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest("Invalid Username/Password"))
		return
	}
	newSession, err := h.sessionManager.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	accountJson, err := json.Marshal(account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	if err = newSession.Set("user", string(accountJson)); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	duration, _ := strconv.Atoi(os.Getenv("SESSION_DURATION"))
	if err = newSession.UpdateExpire(duration); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.SetCookie(session.Name, newSession.Key, 0, "/", os.Getenv("DOMAIN"), false, true)
	c.JSON(http.StatusOK, dto.Authenticated(account.Username, account.Role.RoleName))
}

func (h requestHandler) Logout(c *gin.Context) {
	key, err := h.sessionManager.Delete(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.Header("Access-Control-Allow-Credentials", "true")
	c.SetCookie(session.Name, key, -1, "/", os.Getenv("DOMAIN"), false, true)
	c.JSON(http.StatusOK, dto.Success("Log out success", nil))
}
