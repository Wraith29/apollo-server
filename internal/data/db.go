package data

import (
	"github.com/spf13/viper"
	"github.com/wraith29/apollo/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _db *gorm.DB

func Init() error {
	dbUri := viper.GetString("database-uri")

	db, err := gorm.Open(sqlite.Open(dbUri), &gorm.Config{})
	if err != nil {
		return err
	}

	if err = db.AutoMigrate(&model.Genre{}, &model.Artist{}, &model.Album{}); err != nil {
		return err
	}

	_db = db

	return nil
}

func GetDB() (*gorm.DB, error) {
	if _db != nil {
		return _db, nil
	}

	if err := Init(); err != nil {
		return nil, err
	}

	return _db, nil
}
