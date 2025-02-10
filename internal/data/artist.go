package data

import (
	"strings"

	"github.com/spf13/viper"
	"github.com/wraith29/apollo/internal/config"
	mb "github.com/wraith29/apollo/internal/data/musicbrainz"
	"github.com/wraith29/apollo/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func ArtistExists(db *gorm.DB, artistId string) bool {
	var exists int64 = 0

	db.Raw("SELECT EXISTS (SELECT 1 FROM `artist` WHERE `id` = ?)", artistId).Scan(&exists)

	return exists != 0
}

func SaveMusicBrainzArtist(db *gorm.DB, mbArtist *mb.Artist) error {
	ignoreWithSecondaryTypes := viper.GetBool(config.IgnoreWithSecondaryTypesKey)

	genres := make([]model.Genre, 0)

	for _, genre := range mbArtist.Genres {
		genres = append(genres, model.Genre{
			Id:     genre.Id,
			Name:   genre.Name,
			Rating: 0,
		})
	}

	albums := make([]model.Album, 0)

	for _, album := range mbArtist.ReleaseGroups {
		if strings.ToLower(album.PrimaryType) != "album" || (ignoreWithSecondaryTypes && len(album.SecondaryTypes) > 0) {
			continue
		}

		albumGenres := make([]model.Genre, 0)

		for _, genre := range album.Genres {
			albumGenres = append(albumGenres, model.Genre{
				Id:     genre.Id,
				Name:   genre.Name,
				Rating: 0,
			})
		}

		albums = append(albums, model.Album{
			Id:           album.Id,
			ArtistId:     mbArtist.Id,
			Name:         album.Title,
			Listened:     false,
			ListenedDate: nil,
			Genres:       albumGenres,
		})
	}

	if len(genres) > 0 {
		db.Clauses(clause.OnConflict{DoNothing: true}).Create(&genres)
	}

	if len(albums) > 0 {
		db.Create(&albums)
	}

	artist := model.Artist{
		Id:     mbArtist.Id,
		Name:   mbArtist.Name,
		Genres: genres,
	}

	db.Create(&artist)

	return db.Error
}

func GetArtists(db *gorm.DB, listAll bool) (model.ListResult[model.Artist], error) {
	artists := make([]model.Artist, 0)

	var count int
	db.Raw("SELECT COUNT(id) FROM `artist`").Scan(&count)

	query := db.Order(
		clause.OrderByColumn{Column: clause.Column{Name: "rating"}, Desc: true},
	)

	if !listAll {
		query.Limit(10)
	}

	query.Find(&artists)

	return model.ListResult[model.Artist]{
		Count:   count,
		Results: artists,
	}, db.Error
}
