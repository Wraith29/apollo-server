package list

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wraith29/apollo/internal/data"
)

func listGenres() error {
	listAll := viper.GetBool("all")

	db, err := data.GetDB()
	if err != nil {
		return err
	}

	genres, err := data.GetGenres(db, listAll)
	if err != nil {
		return err
	}

	for _, genre := range genres.Results {
		fmt.Printf("%s: %d\n", genre.Name, genre.Rating)
	}

	switch {
	case !listAll && genres.Count > 10:
		fmt.Printf("Displaying 10/%d genres\n", genres.Count)
	case !listAll:
		fmt.Printf("Displaying your top 10 genres\n")
	case listAll:
		fmt.Printf("Displaying %d genres\n", genres.Count)
	}

	return nil
}

var genreCmd = &cobra.Command{
	Use: "genre",
	Run: func(cmd *cobra.Command, args []string) {
		if err := listGenres(); err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			os.Exit(1)
		}
	},
}
