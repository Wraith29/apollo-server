package config

import (
	"encoding/json"
	"os"
	"path"

	"github.com/wraith29/apollo/pkg/file"
)

type server struct {
	Port int `json:"port"`
}

type Config struct {
	Server server `json:"server"`
}

func newConfig() Config {
	return Config{
		Server: server{
			Port: 1128,
		},
	}
}

func createAndLoadConfig(filePath string) (*Config, error) {
	config := newConfig()

	if err := file.CreateFileAndParents(filePath); err != nil {
		return nil, err
	}

	return &config, nil
}

func loadConfig(filePath string) (*Config, error) {
	fileContents, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(fileContents, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func Load() (*Config, error) {
	configFilePath := path.Join(file.AppConfigDir, "config.json")
	if !file.Exists(configFilePath) {
		return createAndLoadConfig(configFilePath)
	}

	return loadConfig(configFilePath)
}
