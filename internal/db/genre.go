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
