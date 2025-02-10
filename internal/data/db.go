package data

import (
	"errors"

	"github.com/spf13/viper"
	"github.com/wraith29/apollo/internal/config"
	"github.com/wraith29/apollo/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _db *gorm.DB

func init() {
	dbUri := viper.GetString(config.DatabaseUriKey)

	db, err := gorm.Open(sqlite.Open(dbUri), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err = db.AutoMigrate(&model.Genre{}, &model.Artist{}, &model.Album{}, &model.Recommendation{}); err != nil {
		panic(err)
	}

	_db = db
}

func GetDB() (*gorm.DB, error) {
	if _db == nil {
		return nil, errors.New("db is uninitialised")
	}

	return _db, nil
}
