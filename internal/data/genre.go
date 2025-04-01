package data

import (
	"database/sql"
	mb "github.com/wraith29/apollo/internal/data/musicbrainz"
)

func addGenresToUser(txn *sql.Tx, userId string, genres []mb.Genre) error {
	statement, err := txn.Prepare(`
			INSERT INTO "user_genres" ("user_id", "genre_id")
			VALUES ($1, $2)
		`)

	if err != nil {
		return err
	}

	for _, genre := range genres {
		if _, err := statement.Exec(userId, genre.Id); err != nil {
			return err
		}
	}

	return statement.Close()
}

func saveGenres(txn *sql.Tx, genres []mb.Genre) error {
	statement, err := txn.Prepare(`
			INSERT INTO "genres" ("genre_id", "name")
			VALUES ($1, $2)
			ON CONFLICT ("genre_id") DO NOTHING
		`)
	if err != nil {
		return err
	}

	for _, genre := range genres {
		if _, err := statement.Exec(genre.Id, genre.Name); err != nil {
			return err
		}
	}

	return statement.Close()
}
