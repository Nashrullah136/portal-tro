package configuration

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/dto"
	"net/http"
	"os"
)

type RequestHandlerInterface interface {
	UpdateSessionDuration(c *gin.Context)
	GetSessionDuration(c *gin.Context)
}

func NewRequestHandler() RequestHandlerInterface {
	return requestHandler{}
}

type requestHandler struct {
}

func (h requestHandler) UpdateSessionDuration(c *gin.Context) {
	var request SessionDurationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorValidation(err))
		return
	}
	if err := os.Setenv("SESSION_DURATION", request.Duration); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(http.StatusOK, dto.Success("Success update configuration", nil))
}

func (h requestHandler) GetSessionDuration(c *gin.Context) {
	c.JSON(http.StatusOK, dto.Success("Success get configuration", map[string]string{"session_duration": os.Getenv("SESSION_DURATION")}))
}
