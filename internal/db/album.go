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
