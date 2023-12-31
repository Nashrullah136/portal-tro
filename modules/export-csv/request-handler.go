package export_csv

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nashrul-be/crm/dto"
	"nashrul-be/crm/utils/logutils"
	"net/http"
	"strconv"
)

type RequestHandlerInterface interface {
	GetAll(c *gin.Context)
	DownloadCsv(c *gin.Context)
}

func NewRequestHandler(exportCsvController ControllerInterface) RequestHandlerInterface {
	return requestHandler{exportCsvController: exportCsvController}
}

type requestHandler struct {
	exportCsvController ControllerInterface
}

func (h requestHandler) GetAll(c *gin.Context) {
	var request GetAllRequest
	ctx := c.Copy()
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorValidation(err))
		return
	}
	response, err := h.exportCsvController.GetAll(ctx, request)
	if err != nil {
		logutils.Get().Println(err)
		c.JSON(http.StatusInternalServerError, dto.ErrorInternalServerError())
		return
	}
	c.JSON(response.Code, response)
}

func (h requestHandler) DownloadCsv(c *gin.Context) {
	ctx := c.Copy()
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		c.JSON(http.StatusNotFound, dto.ErrorNotFound(fmt.Sprintf("CSV with id %s not found", idParam)))
		return
	}
	file, err := h.exportCsvController.DownloadCsv(ctx, uint(id))
	c.FileAttachment(file.Path(), file.Filename())
}
