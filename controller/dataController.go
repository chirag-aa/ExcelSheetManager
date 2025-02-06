package controller

import (
	"log"
	"net/http"

	"excelsheetmanager.com/models"
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

	employeesData, parsingErr := utils.ParseExcelSheet(file)

	if parsingErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Excel sheet cannot be parsed"})
		return
	}

	isDataSaved, dataSaveErr := dc.DataService.SaveExcelDataToDatabase(employeesData)

	if dataSaveErr != nil {
		log.Println("Error in saving the data ", dataSaveErr)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to insert the data"})
		return
	}
	if !isDataSaved {
		log.Println("Error in saving the data ", dataSaveErr)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to insert the data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data imported sucessfully"})

}

func (dc *DataController) GetData(c *gin.Context) {
	employeesData, data, err := dc.DataService.GetDataFromDatabaseOrRedis()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to get the data"})
		return
	}
	if data == nil {
		c.JSON(http.StatusOK, gin.H{"Data": employeesData})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data": data})
}

func (dc *DataController) UpdateDataByEmail(c *gin.Context) {
	var requestBody models.Request

	bindErr := c.ShouldBindJSON(&requestBody)

	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to bind the data"})
		return
	}
	updatedEmployeeData, err := dc.DataService.UpdateEmployeeByEmail(requestBody)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to get the data"})
		return
	}
	c.JSON(200, gin.H{"data": updatedEmployeeData})

}
