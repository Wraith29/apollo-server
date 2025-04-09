package db

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"github.com/wraith29/apollo/internal/db/query"
	mb "github.com/wraith29/apollo/internal/musicbrainz"

	sq "github.com/Masterminds/squirrel"
)

type albumDbWriter struct {
	artistId string
	albums   []mb.ReleaseGroup
}

func (dw *albumDbWriter) write(txn *sql.Tx) error {
	stmt, err := txn.Prepare(query.InsertAlbum)
	if err != nil {
		return err
	}

	for _, album := range dw.albums {
		if !album.IsValid() {
			continue
		}

		if _, err := stmt.Exec(album.Id, album.Title, album.FirstReleaseDate, dw.artistId); err != nil {
			return err
		}
	}

	return stmt.Close()
}

func SaveAlbums(artistId string, albums []mb.ReleaseGroup) dbWriter {
	return &albumDbWriter{artistId, albums}
}

type albumGenresDbWriter struct {
	albums []mb.ReleaseGroup
}

func (dw *albumGenresDbWriter) write(txn *sql.Tx) error {
	stmt, err := txn.Prepare(query.InsertAlbumGenre)
	if err != nil {
		return err
	}

	for _, album := range dw.albums {
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

func SaveAlbumsGenres(albums []mb.ReleaseGroup) dbWriter {
	return &albumGenresDbWriter{albums}
}

func scanUserAlbums(rows *sql.Rows) ([]userAlbum, error) {
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

func getUserAlbumsNoFilters(userId string) ([]userAlbum, error) {
	rows, err := _conn.Query(query.SelectAlbumsForUserNoFilters, userId)
	if err != nil {
		return nil, err
	}

	return scanUserAlbums(rows)
}

func getUserAlbumsAnyGenreNotRecommended(userId string) ([]userAlbum, error) {
	rows, err := _conn.Query(query.SelectAlbumsForUserAnyGenreNotRecommended, userId)
	if err != nil {
		return nil, err
	}

	return scanUserAlbums(rows)
}

func getUserAlbumsWithGenres(userId string, genres []string) ([]userAlbum, error) {
	rows, err := _conn.Query(query.SelectAlbumsForUserWithGenres, userId, pq.Array(genres))
	if err != nil {
		return nil, err
	}

	return scanUserAlbums(rows)
}

func getUserAlbumsWithGenresNotRecommended(userId string, genres []string) ([]userAlbum, error) {
	rows, err := _conn.Query(query.SelectAlbumsForUserWithGenresNotRecommended, userId, pq.Array(genres))
	if err != nil {
		return nil, err
	}

	return scanUserAlbums(rows)
}

func GetUserAlbums(userId string, includeListened bool, genres []string) ([]userAlbum, error) {
	qb := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("artist.name", "album.name", "album.id").
		From("user_album").
		InnerJoin("album ON album.id = user_album.album_id").
		InnerJoin("artist ON artist.id = album.artist_id").
		Where(sq.Eq{"user_album.user_id": userId})

	if len(genres) > 0 {
		qb = qb.InnerJoin("album_genre ON album_genre.album_id = user_album.album_id").
			InnerJoin("genre ON genre.id = album_genre.genre_id").
			// Where("genre.name = ANY($2)", pq.Array(genres))
	}

	if !includeListened {
		qb = qb.Where(sq.Eq{"user_album.recommended": false})
	}

	query, args, err := qb.ToSql()

	if err != nil {
		return nil, err
	}

	fmt.Println(query)

	for _, arg := range args {
		fmt.Println(arg)
	}

	rows, err := _conn.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return scanUserAlbums(rows)

	// switch {
	// case !includeListened && len(genres) == 0:
	// 	return getUserAlbumsAnyGenreNotRecommended(userId)
	// case includeListened && len(genres) != 0:
	// 	return getUserAlbumsWithGenres(userId, genres)
	// case !includeListened && len(genres) != 0:
	// 	return getUserAlbumsWithGenresNotRecommended(userId, genres)
	// default:
	// 	return getUserAlbumsNoFilters(userId)
	// }
}
