package list

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wraith29/apollo/internal/data"
)

func listRecommendations() error {
	listAll := viper.GetBool("all")

	db, err := data.GetDB()
	if err != nil {
		return err
	}

	recommendations, err := data.GetRecommendations(db, listAll)
	if err != nil {
		return err
	}

	for _, recommendation := range recommendations {
		fmt.Printf("%d: %s\n", recommendation.Id, recommendation.AlbumName)
	}

	return nil
}

var recommendationCmd = &cobra.Command{
	Use:     "recommendation",
	Aliases: []string{"rec"},
	Run: func(cmd *cobra.Command, args []string) {
		if err := listRecommendations(); err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			os.Exit(1)
		}
	},
}
