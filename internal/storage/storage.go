package storage

import (
	"errors"
	"io/fs"
	"os"
)

func PathExists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func MkdirIfNotExists(path string) error {
	exists, err := PathExists(path)
	if err != nil {
		return err
	}

	if !exists {
		err := os.MkdirAll(path, os.ModeDir|os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateWithDataIfNotExists(path, contents string) error {
	exists, err := PathExists(path)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	if _, err := f.WriteString(contents); err != nil {
		return err
	}

	f.Close()

	return nil
}

func CreateIfNotExists(path string) error {
	exists, err := PathExists(path)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	f.Close()

	return nil
}
