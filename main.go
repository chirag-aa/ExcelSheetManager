package main

import (
	"fmt"

	"excelsheetmanager.com/controller"
	"excelsheetmanager.com/services"
	"excelsheetmanager.com/utils"
	"github.com/gin-gonic/gin"
)

func main() {

	dbService, dbServiceErr := services.NewMySqlConnection(utils.Database_Connection_String)

	if dbServiceErr != nil {
		fmt.Println("There is a error connecting to db", dbServiceErr)
		return
	}

	dataService := services.NewDataService(dbService)

	dataController := controller.NewController(dataService)

	router := gin.Default()

	router.POST("/api/upload", dataController.UploadExcel)
	router.Run(":8080")
}
