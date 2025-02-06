package services

type DataService struct {
	mySqlConnection *mySqlConnection
}

func NewDataService(ms *mySqlConnection) *DataService {
	return &DataService{
		mySqlConnection: ms,
	}
}

// func (ds *DataService) SaveExcelDataToDatabase([]models.Employee) (bool, error) {

// }
