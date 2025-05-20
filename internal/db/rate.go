package db

import (
	"gorm.io/gorm"
)

func RateAlbumForUser(userId, albumId string, rating int) error {
	var album Album

	if err := conn.
		Table("albums").
		Preload("Genres").
		Joins("LEFT JOIN album_genres ON album_genres.album_id = albums.id").
		Where("albums.id = ?", albumId).
		First(&album).
		Error; err != nil {
		return err
	}

	return conn.Transaction(func(txn *gorm.DB) error {
		if err := UpdateUserAlbumRating(txn, userId, album.Id, rating); err != nil {
			return err
		}

		if err := UpdateUserArtistRating(txn, userId, album.ArtistId); err != nil {
			return err
		}

		if err := UpdateUserGenreRatings(txn, userId); err != nil {
			return err
		}

		return nil
	})
}
