package musicbrainz

import (
	"strings"

	"github.com/wraith29/apollo/internal/data"
)

type ReleaseGroup struct {
	Id          string       `json:"id"`
	Title       string       `json:"title"`
	PrimaryType string       `json:"primary-type"`
	Genres      []data.Genre `json:"genres"`
}

type Artist struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	Disambiguation string         `json:"disambiguation"`
	Genres         []data.Genre   `json:"genres"`
	ReleaseGroups  []ReleaseGroup `json:"release-groups"`
}

func (a *Artist) ToCustomArtist() data.Artist {
	albums := make([]data.Album, 0)

	for _, release := range a.ReleaseGroups {
		if strings.ToLower(release.PrimaryType) != "album" {
			continue
		}

		album := data.Album{
			Id:               release.Id,
			Name:             release.Title,
			Genres:           release.Genres,
			Listened:         false,
			LatestListenDate: "",
		}

		albums = append(albums, album)
	}

	return data.Artist{
		Id:     a.Id,
		Name:   a.Name,
		Genres: a.Genres,
		Albums: albums,
	}
}

type Search struct {
	Count   int      `json:"count"`
	Artists []Artist `json:"artists"`
}
