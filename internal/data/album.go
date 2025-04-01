package data

import (
	"database/sql"

	mb "github.com/wraith29/apollo/internal/data/musicbrainz"
)

func addAlbumsToUser(txn *sql.Tx, userId string, albums []mb.ReleaseGroup) error {
	statement, err := txn.Prepare(`
			INSERT INTO "user_albums" ("user_id", "album_id")
			VALUES ($1, $2)
		`)
	if err != nil {
		return err
	}

	for _, album := range albums {
		if !album.IsValid() {
			continue
		}

		if _, err := statement.Exec(userId, album.Id); err != nil {
			return err
		}

		if err := addGenresToUser(txn, userId, album.Genres); err != nil {
			return err
		}
	}

	return statement.Close()
}

func saveAlbums(txn *sql.Tx, artistId string, albums []mb.ReleaseGroup) error {
	statement, err := txn.Prepare(`
			INSERT INTO "albums" ("album_id", "artist_id", "name")
			VALUES ($1, $2, $3)
			ON CONFLICT ("album_id") DO NOTHING
	`)
	if err != nil {
		return err
	}

	for _, album := range albums {
		if !album.IsValid() {
			continue
		}

		if _, err := statement.Exec(album.Id, artistId, album.Title); err != nil {
			return err
		}

		if err := saveGenres(txn, album.Genres); err != nil {
			return err
		}
	}

	return statement.Close()
}
