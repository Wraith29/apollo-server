package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dateFormat = "2006-01-02"

var conn *gorm.DB

func InitDb() error {
	username := os.Getenv("APOLLO_POSTGRES_USERNAME")
	password := os.Getenv("APOLLO_POSTGRES_PASSWORD")
	dbPort := os.Getenv("APOLLO_DB_PORT")

	dsn := fmt.Sprintf(
		"host=localhost user=%s password=%s dbname=apollo sslmode=disable port=%s",
		username,
		password,
		dbPort,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return err
	}

	conn = db

	return migrateModels()
}

func migrateModels() error {
	if err := conn.AutoMigrate(&Genre{}); err != nil {
		return err
	}

	if err := conn.AutoMigrate(&Artist{}); err != nil {
		return err
	}

	if err := conn.AutoMigrate(&Album{}); err != nil {
		return err
	}

	if err := conn.AutoMigrate(&User{}); err != nil {
		return err
	}

	if err := conn.AutoMigrate(&UserGenre{}); err != nil {
		return err
	}

	if err := conn.SetupJoinTable(&User{}, "Genres", &UserGenre{}); err != nil {
		return err
	}

	if err := conn.AutoMigrate(&UserArtist{}); err != nil {
		return err
	}

	if err := conn.SetupJoinTable(&User{}, "Artists", &UserArtist{}); err != nil {
		return err
	}

	if err := conn.AutoMigrate(&UserAlbum{}); err != nil {
		return err
	}

	if err := conn.SetupJoinTable(&User{}, "Albums", &UserAlbum{}); err != nil {
		return err
	}

	if err := conn.AutoMigrate(&Recommendation{}); err != nil {
		return err
	}

	return nil
}

func Collect[TIn, TOut any](collection []TIn, selector func(TIn) TOut) []TOut {
	result := make([]TOut, len(collection))

	for idx, value := range collection {
		result[idx] = selector(value)
	}

	return result
}
