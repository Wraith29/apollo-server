package db

import (
	"time"

	"github.com/wraith29/apollo/internal/musicbrainz"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SaveArtistToUser(artist *musicbrainz.Artist, userId uint) error {
	mbGenres := artist.GetUniqueGenres()
	allGenres := make([]Genre, len(mbGenres))

	for idx, genre := range mbGenres {
		allGenres[idx] = Genre{Id: genre.Id, Name: genre.Name}
	}

	conn.Transaction(func(txn *gorm.DB) error {
		txn = txn.Clauses(clause.OnConflict{DoNothing: true})

		albums := make([]Album, 0)

		for _, album := range artist.ReleaseGroups {
			if !album.IsValid() {
				continue
			}

			releaseDate, err := time.Parse(dateFormat, album.FirstReleaseDate)
			if err != nil {
				return err
			}

			albumGenres := make([]Genre, 0)
			for _, genre := range album.Genres {
				albumGenres = append(albumGenres, Genre{
					Id:   genre.Id,
					Name: genre.Name,
				})
			}

			albums = append(albums, Album{
				Id:          album.Id,
				Name:        album.Title,
				Rating:      0,
				ReleaseDate: releaseDate,
				ArtistId:    artist.Id,
				Genres:      albumGenres,
			})
		}

		user := User{
			Id:     userId,
			Genres: allGenres,
			Artists: []Artist{Artist{
				Id:     artist.Id,
				Name:   artist.Name,
				Genres: allGenres,
			}},
			Albums: albums,
		}

		return txn.Create(&user).Error
	})

	return nil
}
