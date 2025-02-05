package data

import (
	"fmt"
	"math/rand/v2"

	"gorm.io/gorm"
)

type recommendedAlbum struct {
	Id         string
	Name       string
	ArtistName string
}

func GetRandomAlbum(db *gorm.DB, genreNames []string, includeListened bool) (*recommendedAlbum, error) {
	var albums []recommendedAlbum

	albumQuery := db.
		Table("album AL").
		Select("AL.`id`, AL.`name`, AR.`name` AS artist_name").
		Joins("INNER JOIN artist AR ON AR.`id` = AL.`artist_id`").
		Joins("INNER JOIN album_genre AG ON AG.`album_id` = AL.`id`").
		Joins("INNER JOIN genre G ON AG.`genre_id` = G.`id`").
		Where("AL.`listened`", includeListened)

	if len(genreNames) > 0 {
		albumQuery.Where("G.`name` IN ?", genreNames)
	}

	albumQuery.Scan(&albums)

	if db.Error != nil {
		return nil, db.Error
	}

	if len(albums) == 0 {
		if !includeListened {
			fmt.Print("I wasn't able to find any albums you haven't heard! Here's one you have listened to already:\n")
			return GetRandomAlbum(db, genreNames, true)
		}

		return nil, nil
	}

	albumIdx := rand.IntN(len(albums))

	return &albums[albumIdx], nil
}
