package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/wraith29/apollo/internal/data"
)

func rateLatestRecommendation(rating int) error {
	db, err := data.GetDB()
	if err != nil {
		return err
	}

	rec, err := data.GetLatestRecommendation(db)
	if err != nil {
		return err
	}

	if err = data.RateAlbum(db, rec.AlbumId, rec.Id, rating); err != nil {
		return err
	}

	if err = data.MarkAlbumAsListened(db, rec.AlbumId); err != nil {
		return err
	}

	return nil
}

var rateCmd = &cobra.Command{
	Use:   "rate [rating (1-3)]",
	Short: "Rate your latest recommendation on a scale from 1-3",
	Long: `Rate your latest album recommendation from 1-3
	1: I didn't like this album
	2: This album was OK
	3: I liked this album`,
	ValidArgs: []string{"1", "2", "3"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		rating, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		if err := rateLatestRecommendation(rating - 2); err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			os.Exit(1)
		}
	},
}
