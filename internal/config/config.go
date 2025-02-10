package config

import (
	"path"

	"github.com/spf13/viper"
	"github.com/wraith29/apollo/internal/storage"
)

const (
	DatabaseUriKey              = "database-uri"
	IgnoreWithSecondaryTypesKey = "ignore-with-secondary-types"
)

func setupFiles(storageDir string) error {
	if err := storage.MkdirIfNotExists(storageDir); err != nil {
		return err
	}

	if err := storage.CreateIfNotExists(path.Join(storageDir, "apollo.db")); err != nil {
		return err
	}

	return nil
}

func init() {
	storageDir, err := storage.GetStorageDir()
	if err != nil {
		panic(err)
	}

	if err := setupFiles(storageDir); err != nil {
		panic(err)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(storageDir)

	viper.SetDefault(DatabaseUriKey, path.Join(storageDir, "apollo.db"))
	viper.SetDefault(IgnoreWithSecondaryTypesKey, true)

	err = viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		if err = viper.SafeWriteConfig(); err != nil {
			panic(err)
		}
	}
}
