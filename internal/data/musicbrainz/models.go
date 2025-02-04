package musicbrainz

type Genre struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ReleaseGroup struct {
	Id             string   `json:"id"`
	Title          string   `json:"title"`
	PrimaryType    string   `json:"primary-type"`
	SecondaryTypes []string `json:"secondary-types"`
	Genres         []Genre  `json:"genres"`
}

type Artist struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	Disambiguation string         `json:"disambiguation"`
	Genres         []Genre        `json:"genres"`
	ReleaseGroups  []ReleaseGroup `json:"release-groups"`
}

type Search struct {
	Count   int      `json:"count"`
	Artists []Artist `json:"artists"`
}
