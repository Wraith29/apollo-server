package db

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SaveAlbum(txn *gorm.DB, album Album) error {
	return txn.Clauses(clause.OnConflict{UpdateAll: true}).Create(&album).Error
}

func AddAlbumsToUser(txn *gorm.DB, albums []Album, userId string) error {
	userAlbums := make([]UserAlbum, len(albums))

	for idx, album := range albums {
		userAlbums[idx] = UserAlbum{UserId: userId, AlbumId: album.Id}
	}

	return txn.Clauses(clause.OnConflict{UpdateAll: true}).Create(&userAlbums).Error
}

func UpdateUserAlbumRating(txn *gorm.DB, userId, albumId string, rating int) error {
	userAlbum := UserAlbum{UserId: userId, AlbumId: albumId}

	return txn.
		Model(&userAlbum).
		Update("rating", rating).
		Error
}

type RecommendedAlbum struct {
	AlbumId    string `json:"albumId"`
	AlbumName  string `json:"albumName"`
	ArtistName string `json:"artistName"`
}

func GetUserAlbums(userId string, genres []string, includeRecommended bool) ([]RecommendedAlbum, error) {
	query := conn.Table("user_albums").
		Select("albums.id AS album_id, albums.name AS album_name, artists.name AS artist_name").
		Joins("INNER JOIN albums ON albums.id = user_albums.album_id").
		Joins("INNER JOIN artists ON artists.id = albums.artist_id").
		Where("user_albums.user_id = ?", userId)

	if !includeRecommended {
		query = query.Where("user_albums.recommended = false")
	}

	if len(genres) != 0 {
		query = query.
			Joins("INNER JOIN album_genres ON album_genres.album_id = albums.id").
			Joins("INNER JOIN genres ON genres.id = album_genres.genre_id").
			Where("genres.name IN ?", genres)
	}

	var albums []RecommendedAlbum

	err := query.Find(&albums).Error

	return albums, err
}
