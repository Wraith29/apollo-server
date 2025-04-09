package db

import (
	"database/sql"

	"github.com/wraith29/apollo/internal/db/query"
	mb "github.com/wraith29/apollo/internal/musicbrainz"
)

type saveUserQuery struct {
	userId, username string
}

func (s *saveUserQuery) write(txn *sql.Tx) error {
	stmt, err := txn.Prepare(`INSERT INTO "user" ("id", "name") VALUES ($1, $2)`)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(s.userId, s.username); err != nil {
		return err
	}

	return stmt.Close()
}

func SaveUser(userId, username string) dbWriter {
	return &saveUserQuery{userId, username}
}

func saveAlbumsToUser(txn *sql.Tx, userId string, albums []mb.ReleaseGroup) error {
	stmt, err := txn.Prepare(query.InsertUserAlbum)
	if err != nil {
		return err
	}

	for _, album := range albums {
		if !album.IsValid() {
			continue
		}

		if _, err := stmt.Exec(userId, album.Id); err != nil {
			return err
		}
	}

	return stmt.Close()
}

func saveArtistToUser(txn *sql.Tx, artist *mb.Artist, userId string) error {
	return prepAndExec(txn, query.InsertUserArtist, userId, artist.Id)
}

func saveGenresToUser(txn *sql.Tx, userId string, genres []mb.Genre) error {
	stmt, err := txn.Prepare(query.InsertUserGenre)
	if err != nil {
		return err
	}

	for _, genre := range genres {
		if _, err := stmt.Exec(userId, genre.Id); err != nil {
			return err
		}
	}

	return stmt.Close()
}
