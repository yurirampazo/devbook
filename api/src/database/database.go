package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //MYSQL Driver
)

// Opens database connection
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DatabaseUrl)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil

}