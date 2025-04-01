package musicbrainz

import "github.com/wraith29/apollo/internal/strutil"

type PrimaryType string

const (
	Album     PrimaryType = "Album"
	Single    PrimaryType = "Single"
	EP        PrimaryType = "EP"
	Broadcast PrimaryType = "Broadcast"
	Other     PrimaryType = "Other"
)

type SecondaryType string

const (
	Compilation    SecondaryType = "Compilation"
	Soundtrack     SecondaryType = "Soundtrack"
	Spokenword     SecondaryType = "Spokenword"
	Interview      SecondaryType = "Interview"
	Audiobook      SecondaryType = "Audiobook"
	AudioDrama     SecondaryType = "Audio Drama"
	Live           SecondaryType = "Live"
	Remix          SecondaryType = "Remix"
	DJMix          SecondaryType = "DJ Mix"
	MixtapeStreet  SecondaryType = "Mixtape/Street"
	Demo           SecondaryType = "Demo"
	FieldRecording SecondaryType = "Field Recording"
)

type Genre struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ReleaseGroup struct {
	Id             string          `json:"id"`
	Title          string          `json:"title"`
	PrimaryType    PrimaryType     `json:"primary-type"`
	SecondaryTypes []SecondaryType `json:"secondary-types"`
	Genres         []Genre         `json:"genres"`
}

func (r *ReleaseGroup) IsValid() bool {
	return r.PrimaryType == Album && len(r.SecondaryTypes) != 0
}

type Artist struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	Disambiguation string         `json:"disambiguation"`
	Genres         []Genre        `json:"genres"`
	ReleaseGroups  []ReleaseGroup `json:"release-groups"`
}

type SearchResult struct {
	Count   int      `json:"count"`
	Artists []Artist `json:"artists"`
}

func (s *SearchResult) FindArtistWithShortestDistance(artistName string) *Artist {
	artistIndex := 0
	minimumDistance := 100

	for index, artist := range s.Artists {
		distance := strutil.Distance(artistName, artist.Name)

		if distance < minimumDistance {
			artistIndex = index
			minimumDistance = distance
		}
	}

	return &s.Artists[artistIndex]
}
