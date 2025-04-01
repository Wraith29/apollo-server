package data

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var _conn *sql.DB

func InitDb() error {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf(
		"postgres://%s:%s@localhost:%s/apollo?sslmode=disable",
		username,
		password,
		dbPort,
	)

	db, err := sql.Open("postgres", connStr)

	_conn = db

	return err
}

func getTransaction() (*sql.Tx, error) {
	return _conn.Begin()
}
