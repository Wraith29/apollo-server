package db

import (
	"database/sql"

	"github.com/wraith29/apollo/internal/db/query"
	mb "github.com/wraith29/apollo/internal/musicbrainz"
)

type artistDbWriter struct {
	artist *mb.Artist
}

func (dw *artistDbWriter) write(txn *sql.Tx) error {
	return prepAndExec(txn, query.InsertArtist, dw.artist.Id, dw.artist.Name)
}

func SaveArtist(artist *mb.Artist) dbWriter {
	return &artistDbWriter{artist}
}

type artistGenresDbWriter struct {
	artist *mb.Artist
}

func (dw *artistGenresDbWriter) write(txn *sql.Tx) error {
	stmt, err := txn.Prepare(query.InsertArtistGenre)
	if err != nil {
		return err
	}

	for _, genre := range dw.artist.Genres {
		if _, err := stmt.Exec(dw.artist.Id, genre.Id); err != nil {
			return err
		}
	}

	return stmt.Close()
}

func SaveArtistGenres(artist *mb.Artist) dbWriter {
	return &artistGenresDbWriter{artist}
}

type addArtistQuery struct {
	userId string
	artist *mb.Artist
}

func (a *addArtistQuery) write(txn *sql.Tx) error {
	allGenres := a.artist.GetUniqueGenres()

	for _, writer := range []dbWriter{
		SaveGenres(allGenres),
		SaveGenresToUser(a.userId, allGenres),

		SaveArtist(a.artist),
		SaveArtistToUser(a.userId, a.artist),
		SaveArtistGenres(a.artist),

		SaveAlbums(a.artist.Id, a.artist.ReleaseGroups),
		SaveAlbumsToUser(a.userId, a.artist.ReleaseGroups),
		SaveAlbumsGenres(a.artist.ReleaseGroups),
	} {
		if err := writer.write(txn); err != nil {
			return err
		}
	}

	return nil
}

func SaveAllArtistData(userId string, artist *mb.Artist) dbWriter {
	return &addArtistQuery{userId, artist}
}
