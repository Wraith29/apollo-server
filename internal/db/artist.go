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
