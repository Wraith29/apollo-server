package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "add <artist_name>",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("missing required argument: <artist_name>")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
