package db

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SaveGenres(txn *gorm.DB, genres []Genre) error {
	return txn.Clauses(clause.OnConflict{DoNothing: true}).Create(&genres).Error
}

func AddGenresToUser(txn *gorm.DB, genres []Genre, userId string) error {
	userGenres := make([]UserGenre, len(genres))

	for idx, genre := range genres {
		userGenres[idx] = UserGenre{UserId: userId, GenreId: genre.Id}
	}

	return txn.Clauses(clause.OnConflict{DoNothing: true}).Create(&userGenres).Error
}

func UpdateUserGenreRatings(txn *gorm.DB, userId string) error {
	type result struct {
		AlbumId string
		Rating  int
		GenreId string
	}

	var results []result

	if err := txn.
		Table("albums").
		Joins("INNER JOIN user_albums ON user_albums.album_id = albums.id").
		Joins("INNER JOIN album_genres ON album_genres.album_id = albums.id").
		Where("user_albums.user_id = ?", userId).
		Scan(&results).
		Error; err != nil {
		return err
	}

	genreRatings := make(map[string]int)

	for _, result := range results {
		current := genreRatings[result.GenreId]
		genreRatings[result.GenreId] = current + result.Rating
	}

	for genreId, rating := range genreRatings {
		if err := txn.
			Table("user_genres").
			Where("user_id = ?", userId).
			Where("genre_id = ?", genreId).
			Update("rating", rating).
			Error; err != nil {
			return err
		}
	}

	return nil
}
