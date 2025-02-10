package list

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ListCmd = &cobra.Command{
	Use:       "list [category]",
	Short:     "List the items in the given category.",
	Aliases:   []string{"ls"},
	ValidArgs: []string{"artist", "genre", "recommendations"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
}

func init() {
	ListCmd.PersistentFlags().BoolP("all", "a", false, "List all items")
	if err := viper.BindPFlag("all", ListCmd.PersistentFlags().Lookup("all")); err != nil {
		panic(err)
	}

	ListCmd.AddCommand(artistCmd)
	ListCmd.AddCommand(genreCmd)
	ListCmd.AddCommand(recommendationCmd)
}
