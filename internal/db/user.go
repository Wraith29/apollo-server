package db

import (
	"errors"

	"gorm.io/gorm"
)

func UsernameTaken(username string) bool {
	var user User

	err := conn.Select("1").Where("id = ?", username).First(&user).Error

	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func SaveUser(username, password string) (string, error) {
	user, err := NewUser(username, password)
	if err != nil {
		return "", err
	}

	err = conn.Create(&user).Error

	return user.Id, err
}
