package db

import (
	"database/sql"
	"errors"

	"github.com/wraith29/apollo/internal/db/query"
	mb "github.com/wraith29/apollo/internal/musicbrainz"
)

func saveArtist(txn *sql.Tx, artist *mb.Artist) error {
	return prepAndExec(txn, query.InsertArtist, artist.Id, artist.Name)
}

func saveGenresToArtist(txn *sql.Tx, artist *mb.Artist) error {
	stmt, err := txn.Prepare(query.InsertArtistGenre)
	if err != nil {
		return err
	}

	for _, genre := range artist.Genres {
		if _, err := stmt.Exec(artist.Id, genre.Id); err != nil {
			return err
		}
	}

	return stmt.Close()
}

type addArtistQuery struct {
	userId string
	artist *mb.Artist
}

func (a *addArtistQuery) write(txn *sql.Tx) error {
	allGenres := a.artist.GetUniqueGenres()

	return errors.Join(
		saveGenres(txn, allGenres),
		saveGenresToUser(txn, a.userId, allGenres),

		saveArtist(txn, a.artist),
		saveArtistToUser(txn, a.artist, a.userId),
		saveGenresToArtist(txn, a.artist),

		saveAlbums(txn, a.artist.Id, a.artist.ReleaseGroups),
		saveAlbumsToUser(txn, a.userId, a.artist.ReleaseGroups),
		saveGenresToAlbums(txn, a.artist.ReleaseGroups),
	)
}

func SaveArtist(userId string, artist *mb.Artist) dbWriter {
	return &addArtistQuery{userId, artist}
}
