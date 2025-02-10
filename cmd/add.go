package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wraith29/apollo/internal/data"
	mb "github.com/wraith29/apollo/internal/data/musicbrainz"
)

func getArtistWithShortestDistance(artists []mb.Artist, searchTerm string) int {
	artistIndex := 0
	minDistance := 100

	for index, artist := range artists {
		distance := data.LevenshteinDistance(artist.Name, searchTerm)

		if distance < minDistance {
			artistIndex = index
			minDistance = distance
		}
	}

	return artistIndex
}

func addArtist(artistName string) error {
	db, err := data.GetDB()
	if err != nil {
		return err
	}

	searchData, err := mb.SearchArtist(artistName)
	if err != nil {
		return err
	}

	if searchData.Count <= 0 {
		fmt.Print("No artists found. Please try again\n")
		return nil
	}

	artist := searchData.Artists[getArtistWithShortestDistance(searchData.Artists, artistName)]

	if data.ArtistExists(db, artist.Id) {
		return errors.New("cannot re-add existing artist")
	}

	lookupData, err := mb.LookupArtist(artist.Id)
	if err != nil {
		return err
	}

	if err = data.SaveMusicBrainzArtist(db, lookupData); err != nil {
		return err
	}

	return nil
}

var addCmd = &cobra.Command{
	Use:   "add artist_name",
	Short: "Add an artist to your library",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := addArtist(args[0]); err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			os.Exit(1)
		}
	},
}
