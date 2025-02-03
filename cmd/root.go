package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/wraith29/apollo/internal"
	"github.com/wraith29/apollo/internal/config"
)

var rootCmd = &cobra.Command{
	Use:   "apollo",
	Short: "Apollo is a music management and recommendation software",
}

func setupApolloDir() error {
	apolloDir, err := internal.GetStorageDir()
	if err != nil {
		return err
	}
	if err = internal.MkdirIfNotExists(apolloDir); err != nil {
		return err
	}

	config.AppRoot = apolloDir
	config.DataFile = path.Join(config.AppRoot, "apollo.json")

	return nil
}

func setupDataFile() error {
	if err := internal.CreateWithDataIfNotExists(config.DataFile, "[]"); err != nil {
		return err
	}

	return nil
}

func init() {
	if err := setupApolloDir(); err != nil {
		fmt.Printf("setup apollo dir: %+v\n", err)
		os.Exit(1)
	}
	if err := setupDataFile(); err != nil {
		fmt.Printf("setup data file: %+v\n", err)
		os.Exit(1)
	}

	rootCmd.AddCommand(addCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
