package file

import (
	"errors"
	"io/fs"
	"os"
	"path"
)

const AppConfigDir = "/etc/apollo"

func Exists(filePath string) bool {
	_, err := os.Stat(filePath)

	return errors.Is(err, fs.ErrExist)
}

func CreateFileAndParents(filePath string) error {
	if err := os.MkdirAll(path.Dir(filePath), 0770); err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	file.Close()

	return nil
}
