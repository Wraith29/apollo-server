package list

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wraith29/apollo/internal/data"
)

func listArtists() error {
	listAll := viper.GetBool("all")

	db, err := data.GetDB()
	if err != nil {
		return err
	}

	artists, err := data.GetArtists(db, listAll)
	if err != nil {
		return err
	}

	for _, artist := range artists.Results {
		fmt.Printf("%s: %d\n", artist.Name, artist.Rating)
	}

	switch {
	case !listAll && artists.Count > 10:
		fmt.Printf("Displaying 10/%d artists\n", artists.Count)
	case !listAll:
		fmt.Printf("Displaying your top 10 artists\n")
	case listAll:
		fmt.Printf("Displaying %d artists\n", artists.Count)
	}

	return nil
}

var artistCmd = &cobra.Command{
	Use: "artist",
	Run: func(cmd *cobra.Command, args []string) {
		if err := listArtists(); err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			os.Exit(1)
		}
	},
}
