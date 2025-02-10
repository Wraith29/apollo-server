package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wraith29/apollo/internal/data"
)

func removeArtist(artistName string) error {
	db, err := data.GetDB()
	if err != nil {
		return err
	}

	return data.RemoveArtist(db, artistName)
}

var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := removeArtist(args[0]); err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			os.Exit(1)
		}
	},
}
