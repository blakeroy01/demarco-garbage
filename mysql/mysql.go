package mysql

import (
	"database/sql"
	"log"
)

type MySQLDatabase struct {
	connection *sql.DB
}

func Connect() (*MySQLDatabase, error) {
	db, err := sql.Open("mysql", "root:astronomicstudios@tcp(127.0.0.1:3306)/disposal_schema")
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to MySQL")
	return &MySQLDatabase{
		connection: db,
	}, nil
}
