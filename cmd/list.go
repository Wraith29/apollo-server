package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:       "list [category]",
	Short:     "List the items in the given category.",
	ValidArgs: []string{"artist", "genre", "recommendations"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
}

func init() {
	listCmd.PersistentFlags().BoolP("all", "a", false, "List all items")

	if err := viper.BindPFlag("all", listCmd.PersistentFlags().Lookup("all")); err != nil {
		panic(err)
	}

	listCmd.AddCommand(artistCmd)
}
