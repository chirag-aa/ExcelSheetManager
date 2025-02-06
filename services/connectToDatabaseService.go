package services

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type mySqlConnection struct {
	db *sql.DB
}

func NewMySqlConnection(connectionString string) (*mySqlConnection, error) {

	db, dbConnectionerror := sql.Open("postgres", connectionString)
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
