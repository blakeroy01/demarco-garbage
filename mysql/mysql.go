package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDatabase struct {
	Connection *sql.DB
}

// Connect is a helper function to connect to a MySQL instance
func Connect() (*MySQLDatabase, error) {
	db, err := sql.Open("mysql", "root:astronomicstudios@tcp(127.0.0.1:3306)/disposal_schema")
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to MySQL")
	return &MySQLDatabase{
		Connection: db,
	}, nil
}
