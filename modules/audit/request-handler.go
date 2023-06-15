package audit

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/dto"
	"net/http"
)

type RequestHandlerInterface interface {
	GetAll(c *gin.Context)
	CreateAudit(c *gin.Context)
}

func NewRequestHandler(auditController ControllerInterface) RequestHandlerInterface {
	return requestHandler{auditController: auditController}
}

type requestHandler struct {
	auditController ControllerInterface
}

func (r requestHandler) GetAll(c *gin.Context) {
	var request GetAllRequest
	ctx := c.Copy()
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorValidation(err))
		return
	}
	if request.PerPage <= 0 {
		request.PerPage = 10
	}
	if request.Page <= 0 {
		request.PerPage = 1
	}
	response, err := r.auditController.GetAll(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(response.Code, response)
}

func (r requestHandler) CreateAudit(c *gin.Context) {
	var request CreateAuditRequest
	ctx := c.Copy()
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorValidation(err))
		return
	}
	response, err := r.auditController.CreateAudit(ctx, request.Action)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(response.Code, response)
}
