package db

import (
	"gorm.io/gorm"
)

// Rating has already been changed to be the correct scale
func RateAlbumForUser(userId, albumId string, rating int) error {
	_ = userId

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

	genreIds := Collect(album.Genres, func(g Genre) string { return g.Id })

	return conn.Transaction(func(txn *gorm.DB) error {
		if err := UpdateUserArtistRating(txn, userId, album.ArtistId, rating); err != nil {
			return err
		}

		if err := UpdateGlobalArtistRating(txn, album.ArtistId, rating); err != nil {
			return err
		}

		if err := UpdateUserAlbumRating(txn, userId, album.Id, rating); err != nil {
			return err
		}

		if err := UpdateGlobalAlbumRating(txn, album.Id, rating); err != nil {
			return err
		}

		if err := UpdateUserGenreRatings(txn, userId, genreIds, rating); err != nil {
			return err
		}

		if err := UpdateGlobalGenreRatings(txn, genreIds, rating); err != nil {
			return err
		}

		return nil
	})
}
