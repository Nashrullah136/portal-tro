package rdn

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/dto"
	"net/http"
)

type RequestHandlerInterface interface {
	GetRdnNew(c *gin.Context)
	GetRdnExisting(c *gin.Context)
	UpdateRdnExisting(c *gin.Context)
	UpdateRdnNew(c *gin.Context)
}

func NewRequestHandler(rdnController ControllerInterface) RequestHandlerInterface {
	return requestHandler{rdnController: rdnController}
}

type requestHandler struct {
	rdnController ControllerInterface
}

func (h requestHandler) GetRdnNew(c *gin.Context) {
	ctx := c.Copy()
	response, err := h.rdnController.GetRdnNew(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(response.Code, response)
}

func (h requestHandler) GetRdnExisting(c *gin.Context) {
	ctx := c.Copy()
	response, err := h.rdnController.GetRdnExisting(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(response.Code, response)
}

func (h requestHandler) UpdateRdnExisting(c *gin.Context) {
	ctx := c.Copy()
	response, err := h.rdnController.UpdateRdnExisting(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(response.Code, response)
}

func (h requestHandler) UpdateRdnNew(c *gin.Context) {
	ctx := c.Copy()
	response, err := h.rdnController.UpdateRdnNew(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(response.Code, response)
}
