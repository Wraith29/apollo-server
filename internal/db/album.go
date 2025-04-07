package db

import (
	"database/sql"

	"github.com/lib/pq"
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

type userAlbum struct {
	ArtistName string `json:"artist_name"`
	AlbumName  string `json:"album_name"`
	AlbumId    string `json:"album_id"`
}

func getUserAlbums(rows *sql.Rows) ([]userAlbum, error) {
	albums := make([]userAlbum, 0)

	var arName, alName, alId string

	for rows.Next() {
		if err := rows.Scan(&arName, &alName, &alId); err != nil {
			return nil, err
		}

		albums = append(albums, userAlbum{arName, alName, alId})
	}

	return albums, rows.Err()
}

func getUserAlbumsNoFilter(userId string, includeListened bool) ([]userAlbum, error) {
	rows, err := _conn.Query(selectAlbumsForUser, userId, includeListened)
	if err != nil {
		return nil, err
	}

	return getUserAlbums(rows)
}

func getUserAlbumsWithFilter(userId string, includeListened bool, genres []string) ([]userAlbum, error) {
	rows, err := _conn.Query(
		selectAlbumsForUserWithGenres,
		userId,
		includeListened,
		pq.Array(genres),
	)
	if err != nil {
		return nil, err
	}

	return getUserAlbums(rows)
}

func GetUserAlbums(userId string, includeListened bool, genres []string) ([]userAlbum, error) {
	if len(genres) == 0 {
		return getUserAlbumsNoFilter(userId, includeListened)
	}

	return getUserAlbumsWithFilter(userId, includeListened, genres)
}
