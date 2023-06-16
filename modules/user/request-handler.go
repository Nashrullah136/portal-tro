package user

import (
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/dto"
	"nashrul-be/crm/entities"
	"net/http"
)

type RequestHandlerInterface interface {
	GetByUsername(c *gin.Context)
	GetAll(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	UpdatePasswordUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

func NewRequestHandler(controllerInterface ControllerInterface) RequestHandlerInterface {
	return requestHandler{actorController: controllerInterface}
}

type requestHandler struct {
	actorController ControllerInterface
}

func (h requestHandler) GetByUsername(c *gin.Context) {
	ctx := c.Copy()
	username := c.Param("username")
	response, err := h.actorController.GetByUsername(ctx, username)
	if err != nil {
		c.JSON(http.StatusNotFound, actorNotFound())
		return
	}
	c.JSON(response.Code, response)
}

func (h requestHandler) GetAll(c *gin.Context) {
	ctx := c.Copy()
	var request PaginationRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorValidation(err))
		return
	}
	if request.PerPage < 1 {
		request.PerPage = 10
	}
	if request.Page < 1 {
		request.Page = 1
	}
	response, err := h.actorController.GetAll(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(response.Code, response)
}

func (h requestHandler) CreateUser(c *gin.Context) {
	var request CreateRequest
	ctx := c.Copy()
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorValidation(err))
		return
	}
	response, err := h.actorController.CreateActor(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(response.Code, response)
}

func (h requestHandler) UpdateUser(c *gin.Context) {
	var request UpdateRequest
	ctx := c.Copy()
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorValidation(err))
		return
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorValidation(err))
		return
	}
	response, err := h.actorController.UpdateActor(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(response.Code, response)
}

func (h requestHandler) UpdatePasswordUser(c *gin.Context) {
	var request ChangePasswordRequest
	ctx := c.Copy()
	actor, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest(err.Error()))
		return
	}
	request.Username = actor.(entities.User).Username
	response, err := h.actorController.ChangePassword(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(response.Code, response)
}

func (h requestHandler) DeleteUser(c *gin.Context) {
	ctx := c.Copy()
	username := c.Param("username")
	if err := h.actorController.DeleteActor(ctx, username); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
