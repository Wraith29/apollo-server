package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wraith29/apollo/internal/data"
)

func getAlbumToRecommend(genres []string) error {
	includeListened := viper.GetBool("listened")

	db, err := data.GetDB()
	if err != nil {
		return err
	}

	album, err := data.GetRandomAlbum(db, genres, includeListened)
	if err != nil {
		return err
	}

	if album == nil {
		switch len(genres) {
		case 0:
			fmt.Print("You don't seem to have any albums. Try `apollo add <artist_name>` to add some\n")
		case 1:
			fmt.Print("You don't seem to have any albums matching that genre. Try again with a different genre, or add a new artist with `apollo add <artist_name>`\n")
		default:
			fmt.Print("You don't seem to have any albums matching those genres. Try again with some different genres, or add a new artist with `apollo add <artist_name>`\n")
		}
		return nil
	}

	if err := data.SaveRecommendation(db, album); err != nil {
		return err
	}

	fmt.Printf("You should listen to: '%s' by %s\n", album.Name, album.ArtistName)

	return nil
}

func init() {
	recCmd.PersistentFlags().BoolP("listened", "l", false, "Include albums that have already been listened to")

	if err := viper.BindPFlag("listened", recCmd.PersistentFlags().Lookup("listened")); err != nil {
		panic(err)
	}
}

var recCmd = &cobra.Command{
	Use:     "recommend [genres... (Max 3)]",
	Aliases: []string{"rec"},
	Short:   "Recommend an album, based on your tastes. Include up to 3 genres to filter down what you want to listen to.",
	Args:    cobra.MaximumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		if err := getAlbumToRecommend(args); err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			os.Exit(1)
		}
	},
}
