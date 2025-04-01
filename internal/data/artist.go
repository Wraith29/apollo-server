package data

import (
	"database/sql"
	"errors"

	mb "github.com/wraith29/apollo/internal/data/musicbrainz"
)

func PersistArtistForUser(userId string, artist *mb.Artist) error {
	txn, err := getTransaction()
	if err != nil {
		return err
	}

	err = errors.Join(
		saveArtist(txn, artist.Id, artist.Name, artist.Genres),
		addArtistToUser(txn, artist.Id, userId, artist.Genres),
		saveAlbums(txn, artist.Id, artist.ReleaseGroups),
		addAlbumsToUser(txn, userId, artist.ReleaseGroups),
	)

	if err != nil {
		return errors.Join(err, txn.Rollback())
	}

	return txn.Commit()
}

func addArtistToUser(txn *sql.Tx, artistId, userId string, genres []mb.Genre) error {
	statement, err := txn.Prepare(`
			INSERT INTO "user_artists" ("user_id", "artist_id")
			VALUES ($1, $2)
	`)
	if err != nil {
		return err
	}

	if _, err := statement.Exec(userId, artistId); err != nil {
		return err
	}

	if err := addGenresToUser(txn, userId, genres); err != nil {
		return err
	}

	return statement.Close()
}

func saveArtist(txn *sql.Tx, artistId, artistName string, genres []mb.Genre) error {
	statement, err := txn.Prepare(`
			INSERT INTO "artists" ("artist_id", "name")
			VALUES ($1, $2)
			ON CONFLICT ("artist_id") DO NOTHING
	`)
	if err != nil {
		return err
	}

	if _, err := statement.Exec(artistId, artistName); err != nil {
		return err
	}

	if err := saveGenres(txn, genres); err != nil {
		return err
	}

	return statement.Close()
}
