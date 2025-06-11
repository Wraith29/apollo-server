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

// Calculates the User Artist rating based on the artists albums
func UpdateUserArtistRating(txn *gorm.DB, userId, artistId string) error {
	var albums []UserAlbum

	if err := txn.
		Table("user_albums").
		Where("user_id = ?", userId).
		Joins("INNER JOIN albums ON albums.id = user_albums.album_id").
		Where("albums.artist_id = ?", artistId).
		Find(&albums).
		Error; err != nil {
		return err
	}

	rating := 0
	for _, album := range albums {
		rating += album.Rating
	}

	return txn.
		Table("user_artists").
		Where("user_id = ?", userId).
		Where("artist_id = ?", artistId).
		Update("rating", rating).Error
}

type userArtistListItem struct {
	ArtistName string
	Rating     int
	UpdatedAt  string
}

func GetAllUserArtistsForListing(userId string) ([]userArtistListItem, error) {
	var artists []userArtistListItem

	if err := conn.
		Table("user_artists").
		Select("artists.name AS artist_name", "user_artists.rating AS rating", "user_artists.updated_at AS updated_at").
		Where("user_id = ?", userId).
		Joins("INNER JOIN artists ON artists.id = user_artists.artist_id").
		Find(&artists).Error; err != nil {
		return nil, err
	}

	return artists, nil
}
