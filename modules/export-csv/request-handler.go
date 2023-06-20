package export_csv

import (
	"github.com/gin-gonic/gin"
	"log"
	"nashrul-be/crm/dto"
	csvutils "nashrul-be/crm/utils/csv"
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
		log.Println(err)
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
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest("Invalid id"))
		return
	}
	filename, err := h.exportCsvController.DownloadCsv(ctx, uint(id))
	c.FileAttachment(csvutils.Path(filename), filename)
}
