package utils

const (
	Ecxel_Sheet_Columns       = 10
	Port                      = ":8080"
	Upload_Data_Path          = "/api/upload"
	Get_Data_Path             = "/api/getData"
	Update_Data_Path          = "/api/updateByEmail"
	Empty_String              = ""
	Message                   = "Message"
	Excel_Sheet_Parsing_Error = "Excel sheet cannot be parsed please check the extension or the file format"
	Validation_Employee_Data  = "validations of employee data failed"
	Data_Save_Error           = "Unable to save the data to database"
	Data_Imported_Sucessfully = "Data imported sucessfully"
	Data_Retrive_Failed       = "Unable to get the data from database"
	Bind_Error                = "Unable to bind the data"
	Validation_Failed         = "Data validation failed please check request body"
	Redis_Connection_Error    = "There is a error connecting to redis"
	Database_Connection_Error = "There is a error connecting to db"
)
