package data

import (
	"encoding/json"
	"os"

	"github.com/wraith29/apollo/internal/config"
)

type Genre struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Album struct {
	Id               string  `json:"id"`
	Name             string  `json:"name"`
	Genres           []Genre `json:"genres"`
	Listened         bool    `json:"listened"`
	LatestListenDate string  `json:"latest-listen-date"`
}

type Artist struct {
	Id     string  `json:"id"`
	Name   string  `json:"string"`
	Genres []Genre `json:"genres"`
	Albums []Album `json:"albums"`
}

func LoadAllArtists() ([]Artist, error) {
	fileContents, err := os.ReadFile(config.DataFile)
	if err != nil {
		return nil, err
	}

	artists := make([]Artist, 0)

	err = json.Unmarshal(fileContents, &artists)
	return artists, err
}

func SaveArtists(artists []Artist) error {
	jsonData, err := json.Marshal(&artists)
	if err != nil {
		return err
	}

	return os.WriteFile(config.DataFile, jsonData, os.ModeExclusive)
}
