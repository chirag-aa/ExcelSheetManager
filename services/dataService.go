package services

import (
	"encoding/json"

	"excelsheetmanager.com/models"
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

	query := "insert into employees (first_name,last_name,company_name,address,city,country,postal,phone,email,web) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)"

	for _, employee := range employeesData {
		_, insertionErr := ds.mySqlConnection.db.Exec(query, employee.First_name, employee.Last_name, employee.Company_name, employee.Address, employee.City, employee.Country, employee.Postal, employee.Phone, employee.Email, employee.Web)
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
func (ds *DataService) GetDataFromDatabaseOrRedis() ([]models.Employee, []map[string]interface{}, error) {

	data, err := ds.redisConnection.GetDataFromRedis()

	if data == "" {
		query := "select * from employees"
		var employeesData []models.Employee
		rows, err := ds.mySqlConnection.db.Query(query)

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
	var unMarshaledData []map[string]interface{}
	err = json.Unmarshal([]byte(data), &unMarshaledData)
	if err != nil {
		return nil, nil, err
	}
	return nil, unMarshaledData, err

}

func (ds *DataService) UpdateEmployeeByEmail(requestBody models.Request) (models.Request, error) {
	query := "update employees set company_name = $1 , first_name = $2 where email = $3"
	_, err := ds.mySqlConnection.db.Exec(query, requestBody.Companyname, requestBody.Firstname, requestBody.Email)
	if err != nil {
		return models.Request{}, err
	}
	return models.Request{
		Firstname:   requestBody.Firstname,
		Companyname: requestBody.Companyname,
		Email:       requestBody.Email,
	}, nil

}
