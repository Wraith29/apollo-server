package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wraith29/apollo/internal/data"
	mb "github.com/wraith29/apollo/internal/data/musicbrainz"
)

func getUserIdxSelection(artists []mb.Artist) (int, error) {
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

func addArtist(artistName string) error {
	db, err := data.GetDB()
	if err != nil {
		return err
	}

	searchData, err := mb.SearchArtist(artistName)
	if err != nil {
		return err
	}

	selectedArtistIdx := 0
	if len(searchData.Artists) > 1 {
		selectedArtistIdx, err = getUserIdxSelection(searchData.Artists)
		if err != nil {
			return err
		}
	}

	selectedArtist := searchData.Artists[selectedArtistIdx]

	if data.ArtistExists(db, selectedArtist.Id) {
		return errors.New("cannot re-add existing artist")
	}

	lookupData, err := mb.LookupArtist(selectedArtist.Id)
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
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("missing required argument: artist_name")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return addArtist(args[0])
	},
}
