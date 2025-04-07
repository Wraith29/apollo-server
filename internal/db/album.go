package db

import (
	"database/sql"

	mb "github.com/wraith29/apollo/internal/musicbrainz"
)

func saveAlbums(txn *sql.Tx, artistId string, albums []mb.ReleaseGroup) error {
	stmt, err := txn.Prepare(insertAlbum)
	if err != nil {
		return err
	}

	for _, album := range albums {
		if !album.IsValid() {
			continue
		}

		if _, err := stmt.Exec(album.Id, album.Title, album.FirstReleaseDate, artistId); err != nil {
			return err
		}
	}

	return stmt.Close()
}

func saveGenresToAlbums(txn *sql.Tx, albums []mb.ReleaseGroup) error {
	stmt, err := txn.Prepare(insertAlbumGenre)
	if err != nil {
		return err
	}

	for _, album := range albums {
		if !album.IsValid() {
			continue
		}

		for _, genre := range album.Genres {
			if _, err := stmt.Exec(album.Id, genre.Id); err != nil {
				return err
			}
		}
	}

	return stmt.Close()
}

type getAlbumsForUserQuery struct {
	userId          string
	includeListened bool
	genres          []string
}
