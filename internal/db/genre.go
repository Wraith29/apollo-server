package db

import (
	"database/sql"

	"github.com/wraith29/apollo/internal/db/query"
	mb "github.com/wraith29/apollo/internal/musicbrainz"
)

func saveGenres(txn *sql.Tx, genres []mb.Genre) error {
	stmt, err := txn.Prepare(query.InsertGenre)
	if err != nil {
		return err
	}

	for _, genre := range genres {
		if _, err := stmt.Exec(genre.Id, genre.Name); err != nil {
			return err
		}
	}

	return stmt.Close()
}
