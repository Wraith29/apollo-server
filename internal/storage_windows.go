package internal

import "errors"

func GetStorageDir() (string, error) {
	return "", errors.New("not implemented for windows")
}
