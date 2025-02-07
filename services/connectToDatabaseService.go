package services

import (
	"database/sql"
	"fmt"

	"excelsheetmanager.com/utils"
	_ "github.com/lib/pq"
)

type mySqlConnection struct {
	db *sql.DB
}

func NewMySqlConnection(connectionString string) (*mySqlConnection, error) {

	connectionString = fmt.Sprintf(connectionString, utils.Database_User, utils.Database_Name, utils.Database_Password, utils.Databse_SSL_Mode)
	db, dbConnectionerror := sql.Open(utils.Database_Driver, connectionString)
	if dbConnectionerror != nil {
		return nil, dbConnectionerror
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}
	return &mySqlConnection{
		db: db,
	}, nil
}
