package controller

import (
	"net/http"

	"excelsheetmanager.com/services"
	"excelsheetmanager.com/utils"
	"github.com/gin-gonic/gin"
)

type DataController struct {
	DataService *services.DataService
}

func NewController(ds *services.DataService) *DataController {
	return &DataController{
		DataService: ds,
	}
}
func (dc *DataController) UploadExcel(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read the file"})
		return
	}

	_, parsingErr := utils.ParseExcelSheet(file)

	if parsingErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Excel sheet cannot be parsed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data imported sucessfully"})

}
