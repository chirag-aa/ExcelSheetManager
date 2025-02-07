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
		c.JSON(http.StatusBadRequest, gin.H{utils.Message: utils.Excel_Sheet_Parsing_Error})
		return
	}
	if !validateEmployeeData(employeesData) {
		c.JSON(http.StatusBadRequest, gin.H{utils.Message: utils.Validation_Failed})
		return
	}
	isDataSaved, dataSaveErr := dc.DataService.SaveExcelDataToDatabase(employeesData)

	if dataSaveErr != nil {
		log.Println("Error in saving the data ", dataSaveErr)
		c.JSON(http.StatusBadRequest, gin.H{utils.Message: utils.Data_Save_Error})
		return
	}
	if !isDataSaved {
		log.Println("Error in saving the data ", dataSaveErr)
		c.JSON(http.StatusBadRequest, gin.H{utils.Message: utils.Data_Save_Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{utils.Message: utils.Data_Imported_Sucessfully})

}

func (dc *DataController) GetData(c *gin.Context) {
	employeesData, data, err := dc.DataService.GetDataFromDatabaseOrRedis(false)
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
		c.JSON(http.StatusBadRequest, gin.H{utils.Message: utils.Bind_Error})
		return
	}
	if !validateRequestBody(requestBody) {
		c.JSON(http.StatusBadRequest, gin.H{utils.Message: utils.Validation_Failed})
		return
	}
	updatedEmployeeData, err := dc.DataService.UpdateEmployeeByEmail(requestBody)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to get the data"})
		return
	}
	c.JSON(200, gin.H{"data": updatedEmployeeData})

}

func validateRequestBody(requestBody models.Request) bool {
	if requestBody.Companyname == "" || requestBody.Email == "" || requestBody.Firstname == "" {
		return false
	}
	return true
}

func validateEmployeeData(employeesData []models.Employee) bool {
	if len(employeesData) == 0 || employeesData == nil {
		return false
	}
	return true
}
