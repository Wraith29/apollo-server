package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const dateFormat = "2006-01-02"

var _conn *sql.DB

func InitDb() error {
	username := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
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

type dbWriter interface {
	write(*sql.Tx) error
}

func Exec(execs ...dbWriter) error {
	txn, err := getTransaction()
	if err != nil {
		return err
	}

	for _, query := range execs {
		if err := query.write(txn); err != nil {
			return errors.Join(err, txn.Rollback())
		}
	}

	return txn.Commit()
}

func prepAndExec(txn *sql.Tx, query string, args ...any) error {
	stmt, err := txn.Prepare(query)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(args...); err != nil {
		return err
	}

	return stmt.Close()
}
