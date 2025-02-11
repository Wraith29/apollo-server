package storage

import (
	"os"
	"path"
)

func GetStorageDir() (string, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(userHome, "AppData", "Local", "apollo"), nil
}
