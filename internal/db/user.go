package db

import (
	"database/sql"

	"github.com/wraith29/apollo/internal/db/query"
	mb "github.com/wraith29/apollo/internal/musicbrainz"
)

type userDbWriter struct {
	userId, username string
}

func (dw *userDbWriter) write(txn *sql.Tx) error {
	return prepAndExec(txn, query.InsertUser, dw.userId, dw.username)
}

func SaveUser(userId, username string) dbWriter {
	return &userDbWriter{userId, username}
}

type userAlbumDbWriter struct {
	userId string
	albums []mb.ReleaseGroup
}

func (dw *userAlbumDbWriter) write(txn *sql.Tx) error {
	stmt, err := txn.Prepare(query.InsertUserAlbum)
	if err != nil {
		return err
	}

	for _, album := range dw.albums {
		if !album.IsValid() {
			continue
		}

		if _, err := stmt.Exec(dw.userId, album.Id); err != nil {
			return err
		}
	}

	return stmt.Close()
}

func SaveAlbumsToUser(userId string, albums []mb.ReleaseGroup) dbWriter {
	return &userAlbumDbWriter{userId, albums}
}

type userArtistDbWriter struct {
	userId string
	artist *mb.Artist
}

func (dw *userArtistDbWriter) write(txn *sql.Tx) error {
	return prepAndExec(txn, query.InsertUserArtist, dw.userId, dw.artist.Id)
}

func SaveArtistToUser(userId string, artist *mb.Artist) dbWriter {
	return &userArtistDbWriter{userId, artist}
}

type userGenresDbWriter struct {
	userId string
	genres []mb.Genre
}

func (dw *userGenresDbWriter) write(txn *sql.Tx) error {
	stmt, err := txn.Prepare(query.InsertUserGenre)
	if err != nil {
		return err
	}

	for _, genre := range dw.genres {
		if _, err := stmt.Exec(dw.userId, genre.Id); err != nil {
			return err
		}
	}

	return stmt.Close()
}

func SaveGenresToUser(userId string, genres []mb.Genre) dbWriter {
	return &userGenresDbWriter{userId, genres}
}

type userAlbumRatingWriter struct {
	userId, albumId string
	rating          int
}

func (r *userAlbumRatingWriter) write(txn *sql.Tx) error {
	return prepAndExec(txn, query.UpdateUserAlbumRating, r.rating, r.userId, r.albumId)
}

func RateAlbum(userId, albumId string, rating int) dbWriter {
	return &userAlbumRatingWriter{userId, albumId, rating}
}

func GetUserArtistIds(userId string) ([]string, error) {
	rows, err := _conn.Query(query.SelectAllArtistsForUser, userId)
	if err != nil {
		return nil, err
	}

	var artistId string
	artistIds := make([]string, 0)

	for rows.Next() {
		if err := rows.Scan(&artistId); err != nil {
			return nil, err
		}

		artistIds = append(artistIds, artistId)
	}

	return artistIds, rows.Err()
}
