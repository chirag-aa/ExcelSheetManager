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

	redisClient, redisConnectionErr := services.NewRedisService()

	if redisConnectionErr != nil {
		fmt.Println(utils.Redis_Connection_Error, redisConnectionErr)
		return
	}

	if dbServiceErr != nil {
		fmt.Println(utils.Database_Connection_Error, dbServiceErr)
		return
	}

	dataService := services.NewDataService(dbService, redisClient)

	dataController := controller.NewController(dataService)

	router := gin.Default()

	router.POST(utils.Upload_Data_Path, dataController.UploadExcel)
	router.GET(utils.Get_Data_Path, dataController.GetData)
	router.PUT(utils.Update_Data_Path, dataController.UpdateDataByEmail)
	router.Run(utils.Port)
}
