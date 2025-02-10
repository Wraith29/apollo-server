package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wraith29/apollo/internal/data"
	mb "github.com/wraith29/apollo/internal/data/musicbrainz"
)

func artistSelection(artists []mb.Artist) (int, error) {
	fmt.Printf("Select Artist:\n")

	for idx, artist := range artists {
		extraDetails := ""

		if artist.Disambiguation != "" {
			extraDetails = fmt.Sprintf(" (%s)", artist.Disambiguation)
		}

		fmt.Printf("Artist %d: %s%s\n", idx+1, artist.Name, extraDetails)
	}

	fmt.Printf("> ")

	reader := bufio.NewReader(os.Stdin)

	selected := -1

	for selected <= 0 || selected > len(artists) {
		userInput, err := reader.ReadString('\n')
		if err != nil {
			return -1, err
		}

		selected, err = strconv.Atoi(strings.Trim(userInput, "\n"))
		if err != nil {
			return -1, err
		}

		if selected > len(artists) {
			fmt.Printf("Invalid Choice. Please try again\n")
		}
	}

	return selected - 1, nil
}

func artistNameWithShortestDist(artists []mb.Artist, artistName string) int {
	artistIndex := 0
	minDistance := 100

	for index, artist := range artists {
		distance := data.LevenshteinDistance(artist.Name, artistName)

		if distance < minDistance {
			artistIndex = index
			minDistance = distance
		}
	}

	return artistIndex
}

func getArtistSelection(artists []mb.Artist, artistName string) (int, error) {
	if viper.GetBool("interactive") {
		return artistSelection(artists)
	}

	return artistNameWithShortestDist(artists, artistName), nil
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

	artistIndex, err := getArtistSelection(searchData.Artists, artistName)
	if err != nil {
		return err
	}

	artist := searchData.Artists[artistIndex]

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
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := addArtist(args[0]); err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	addCmd.PersistentFlags().BoolP("interactive", "i", false, "Check the artist is who you expect before saving them.")
	if err := viper.BindPFlag("interactive", addCmd.PersistentFlags().Lookup("interactive")); err != nil {
		panic(err)
	}
}
