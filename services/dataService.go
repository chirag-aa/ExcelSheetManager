package services

import (
	"encoding/json"

	"excelsheetmanager.com/models"
	"excelsheetmanager.com/utils"
)

type DataService struct {
	mySqlConnection *mySqlConnection
	redisConnection *RedisService
}

func NewDataService(ms *mySqlConnection, rs *RedisService) *DataService {
	return &DataService{
		mySqlConnection: ms,
		redisConnection: rs,
	}
}

func (ds *DataService) SaveExcelDataToDatabase(employeesData []models.Employee) (bool, error) {

	for _, employee := range employeesData {
		_, insertionErr := ds.mySqlConnection.db.Exec(utils.Insert_Data_Into_Employees, employee.First_name, employee.Last_name, employee.Company_name, employee.Address, employee.City, employee.Country, employee.Postal, employee.Phone, employee.Email, employee.Web)
		if insertionErr != nil {
			return false, insertionErr
		}
	}

	_, err := ds.redisConnection.SaveDataToRedis(employeesData)

	if err != nil {
		return false, err
	}

	return true, nil
}
func (ds *DataService) GetDataFromDatabaseOrRedis(isInternalCall bool) ([]models.Employee, []map[string]interface{}, error) {

	data, redisErr := ds.redisConnection.GetDataFromRedis()

	if data == "" || isInternalCall {
		var employeesData []models.Employee
		rows, err := ds.mySqlConnection.db.Query(utils.Select_All_From_Employee)

		if err != nil {
			return nil, nil, err
		}

		for rows.Next() {
			var employee models.Employee
			rows.Scan(&employee.First_name, &employee.Last_name, &employee.Company_name, &employee.Address, &employee.City, &employee.Country, &employee.Postal, &employee.Phone, &employee.Email, &employee.Web)
			employeesData = append(employeesData, employee)
		}
		return employeesData, nil, nil
	}

	if redisErr != nil {
		return nil, nil, redisErr
	}
	var unMarshaledData []map[string]interface{}
	err := json.Unmarshal([]byte(data), &unMarshaledData)
	if err != nil {
		return nil, nil, err
	}
	return nil, unMarshaledData, err

}

func (ds *DataService) UpdateEmployeeByEmail(requestBody models.Request) (models.Request, error) {

	_, err := ds.mySqlConnection.db.Exec(utils.Update_Employee_By_Email, requestBody.Companyname, requestBody.Firstname, requestBody.Email)
	if err != nil {
		return models.Request{}, err
	}
	employeesData, _, err := ds.GetDataFromDatabaseOrRedis(true)

	if err != nil {
		return models.Request{}, err
	}

	_, err = ds.redisConnection.SaveDataToRedis(employeesData)

	if err != nil {
		return models.Request{}, err
	}

	return models.Request{
		Firstname:   requestBody.Firstname,
		Companyname: requestBody.Companyname,
		Email:       requestBody.Email,
	}, nil

}
