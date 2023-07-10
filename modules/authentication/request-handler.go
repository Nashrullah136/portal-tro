package authentication

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/dto"
	"nashrul-be/crm/middleware"
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
	newSession, err := h.sessionManager.Create(*account)
	if err != nil {
		c.JSON(http.StatusLocked, dto.UsernameAlreadyLogin())
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
	c.JSON(http.StatusOK, dto.Authenticated(*account))
}

func (h requestHandler) Logout(c *gin.Context) {
	currentSession, err := h.sessionManager.Get(c)
	if err == nil {
		_, err := h.sessionManager.Delete(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
			return
		}
		middleware.Authenticate(h.sessionManager)(c)
		if err := h.authController.Logout(c.Copy()); err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
			return
		}
	}
	if errors.Is(err, session.ErrNotExist) {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.SetCookie(session.Name, currentSession.Key, -1, "/", os.Getenv("DOMAIN"), false, true)
	}
	c.JSON(http.StatusOK, dto.Success("Log out success", nil))
}
