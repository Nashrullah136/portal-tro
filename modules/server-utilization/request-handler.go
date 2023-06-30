package server_utilization

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/dto"
	"net/http"
)

type RequestHandler interface {
	UpdateHostList(c *gin.Context)
	GetLatestData(c *gin.Context)
}

func NewRequestHandler(serverUtilController Controller) RequestHandler {
	return requestHandler{serverUtilController: serverUtilController}
}

type requestHandler struct {
	serverUtilController Controller
}

func (h requestHandler) UpdateHostList(c *gin.Context) {
	if err := h.serverUtilController.RefreshHostList(); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(http.StatusOK, dto.Success("Success update host list", nil))
}

func (h requestHandler) GetLatestData(c *gin.Context) {
	response, err := h.serverUtilController.GetLastData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(response.Code, response)
}
