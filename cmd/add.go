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

func selectArtistIdx(artists []mb.Artist) (int, error) {
	fmt.Printf("Select Artist:\n")

	for idx, artist := range artists {
		extraDetails := ""

		if artist.Disambiguation != "" {
			// Pad the string with a space on the left to not leave random whitespace
			extraDetails = fmt.Sprintf(" (%s)", artist.Disambiguation)
		}

		fmt.Printf("Artist %d: %s%s\n", idx+1, artist.Name, extraDetails)
	}

	fmt.Printf("> ")

	reader := bufio.NewReader(os.Stdin)

	userInput, err := reader.ReadString('\n')
	if err != nil {
		return -1, err
	}

	parsed, err := strconv.Atoi(strings.Trim(userInput, "\n"))
	if err != nil {
		return -1, err
	}

	if parsed > len(artists) {
		return -1, fmt.Errorf("invalid index. max index = %d", len(artists))
	}

	return parsed - 1, nil
}

func addArtist(artistName string) error {
	currentArtists, err := data.LoadAllArtists()
	if err != nil {
		return err
	}

	searchResults, err := mb.SearchArtist(artistName)
	if err != nil {
		return err
	}

	artistChoice := 0
	if len(searchResults.Artists) > 1 {
		artistChoice, err = selectArtistIdx(searchResults.Artists)
		if err != nil {
			return err
		}
	}

	selectedArtist := searchResults.Artists[artistChoice]

	for _, savedArtist := range currentArtists {
		if savedArtist.Id == selectedArtist.Id {
			return fmt.Errorf("cannot add %s. artist already saved", selectedArtist.Name)
		}
	}

	artistData, err := mb.LookupArtist(selectedArtist.Id)
	if err != nil {
		return err
	}

	currentArtists = append(currentArtists, artistData.ToCustomArtist())

	return data.SaveArtists(currentArtists)
}

var addCmd = &cobra.Command{
	Use:   "add artist_name",
	Short: "Add an artist to your library",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("missing required argument: <artist_name>")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return addArtist(args[0])
	},
}
