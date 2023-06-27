package span

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/dto"
	"net/http"
)

type RequestHandlerInterface interface {
	GetByDocumentNumber(c *gin.Context)
	UpdateBankRiau(c *gin.Context)
}

func NewRequestHandler(spanController ControllerInterface) RequestHandlerInterface {
	return requestHandler{spanController: spanController}
}

type requestHandler struct {
	spanController ControllerInterface
}

func (h requestHandler) GetByDocumentNumber(c *gin.Context) {
	ctx := c.Copy()
	documentNumber := c.Query("documentNumber")
	if documentNumber != "" {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest("Document number is required"))
	}
	response, err := h.spanController.GetByDocumentNumber(ctx, documentNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(response.Code, response)
}

func (h requestHandler) UpdateBankRiau(c *gin.Context) {
	var request UpdateRequest
	ctx := c.Copy()
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorValidation(err))
		return
	}
	response, err := h.spanController.UpdateBankRiau(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
	}
	c.JSON(response.Code, response)
}
