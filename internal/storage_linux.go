package internal

import (
	"os"
	"path"
)

func GetStorageDir() (string, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(userHome, ".local", "share", "apollo"), nil
}
