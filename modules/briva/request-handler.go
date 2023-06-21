package briva

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/dto"
	"net/http"
)

type RequestHandlerInterface interface {
	GetByBrivaNo(c *gin.Context)
	Update(c *gin.Context)
}

func NewRequestHandler(brivaController ControllerInterface) RequestHandlerInterface {
	return requestHandler{brivaController: brivaController}
}

type requestHandler struct {
	brivaController ControllerInterface
}

func (h requestHandler) GetByBrivaNo(c *gin.Context) {
	ctx := c.Copy()
	brivano := c.Param("brivano")
	response, err := h.brivaController.GetByBrivaNo(ctx, brivano)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(response.Code, response)
}

func (h requestHandler) Update(c *gin.Context) {
	var request UpdateRequest
	ctx := c.Copy()
	request.Brivano = c.Param("brivano")
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorValidation(err))
		return
	}
	response, err := h.brivaController.Update(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
	}
	c.JSON(response.Code, response)
}
