package db

import (
	"database/sql"

	"github.com/wraith29/apollo/internal/db/query"
	mb "github.com/wraith29/apollo/internal/musicbrainz"
)

type genresDbWriter struct {
	genres []mb.Genre
}

func (dw *genresDbWriter) write(txn *sql.Tx) error {
	stmt, err := txn.Prepare(query.InsertGenre)
	if err != nil {
		return err
	}

	for _, genre := range dw.genres {
		if _, err := stmt.Exec(genre.Id, genre.Name); err != nil {
			return err
		}
	}

	return stmt.Close()
}

func SaveGenres(genres []mb.Genre) dbWriter {
	return &genresDbWriter{genres}
}
