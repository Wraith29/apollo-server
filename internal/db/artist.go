package db

import (
	"github.com/wraith29/apollo/internal/musicbrainz"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SaveArtist(artist *musicbrainz.Artist) error {
	allGenres := artist.GetUniqueGenres()

	return conn.Transaction(func(txn *gorm.DB) error {
		if err := SaveGenres(txn, GenresFromMusicBrainzGenres(allGenres)); err != nil {
			return err
		}

		artistData, err := ArtistFromMusicBrainzArtist(artist)
		if err != nil {
			return err
		}

		if err := txn.Clauses(clause.OnConflict{UpdateAll: true}).Create(&artistData).Error; err != nil {
			return err
		}

		return nil
	})
}

func AddArtistToUser(artist *musicbrainz.Artist, userId string) error {
	artistData, err := ArtistFromMusicBrainzArtist(artist)
	if err != nil {
		return err
	}

	allGenres := artist.GetUniqueGenres()

	return conn.Transaction(func(txn *gorm.DB) error {
		if err := AddGenresToUser(txn, GenresFromMusicBrainzGenres(allGenres), userId); err != nil {
			return err
		}

		if err := AddAlbumsToUser(txn, artistData.Albums, userId); err != nil {
			return err
		}

		userArtist := UserArtist{UserId: userId, ArtistId: artist.Id}

		return txn.Clauses(clause.OnConflict{DoNothing: true}).Create(&userArtist).Error
	})
}

func UpdateUserArtistRating(txn *gorm.DB, userId, ArtistId string, rating int) error {
	userArtist := UserArtist{UserId: userId, ArtistId: ArtistId}

	return txn.
		Model(&userArtist).
		Update("rating", rating).Error
}

func UpdateGlobalArtistRating(txn *gorm.DB, artistId string, rating int) error {
	var artist Artist

	if err := txn.First(&artist).Where("id = ?", artistId).Error; err != nil {
		return err
	}

	artist.Rating += rating

	return txn.Save(&artist).Error
}
