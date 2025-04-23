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

func UpdateUserGenreRatings(txn *gorm.DB, userId string, genreIds []string, rating int) error {
	return txn.
		Model(&UserGenre{}).
		Where("user_id = ?", userId).
		Where("genre_id IN ?", genreIds).
		Update("rating", rating).Error
}

func UpdateGlobalGenreRatings(txn *gorm.DB, genreIds []string, rating int) error {
	var genres []Genre

	if err := txn.Find(&genres).Where("id = ?", genreIds).Error; err != nil {
		return err
	}

	for _, genre := range genres {
		genre.Rating += rating
	}

	return txn.Save(&genres).Error
}
