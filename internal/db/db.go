package db

import (
	// "fmt"
	// "os"

	// "gorm.io/driver/postgres"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dateFormat = "2006-01-02"

var _db *gorm.DB

func InitDb() error {
	// username := os.Getenv("POSTGRES_USERNAME")
	// password := os.Getenv("POSTGRES_PASSWORD")
	// dbPort := os.Getenv("DB_PORT")

	// dsn := fmt.Sprintf(
	// 	"host=localhost user=%s password=%s dbname=apollo sslmode=disable port=%s",
	// 	username,
	// 	password,
	// 	dbPort,
	// )

	f, err := os.Create("/home/iacna/dev/apollo/apollo.db")
	if err != nil {
		return err
	}
	f.Close()

	db, err := gorm.Open(sqlite.Open("apollo.db"))
	if err != nil {
		return err
	}

	_db = db

	return migrateModels()
}

func migrateModels() error {
	if err := _db.AutoMigrate(&Genre{}); err != nil {
		return err
	}

	if err := _db.AutoMigrate(&Artist{}); err != nil {
		return err
	}

	if err := _db.AutoMigrate(&Album{}); err != nil {
		return err
	}

	if err := _db.AutoMigrate(&User{}); err != nil {
		return err
	}

	if err := _db.AutoMigrate(&UserGenre{}); err != nil {
		return err
	}

	if err := _db.SetupJoinTable(&User{}, "Genres", &UserGenre{}); err != nil {
		return err
	}

	if err := _db.AutoMigrate(&UserArtist{}); err != nil {
		return err
	}

	if err := _db.SetupJoinTable(&User{}, "Artists", &UserArtist{}); err != nil {
		return err
	}

	if err := _db.AutoMigrate(&UserAlbum{}); err != nil {
		return err
	}

	if err := _db.SetupJoinTable(&User{}, "Albums", &UserAlbum{}); err != nil {
		return err
	}

	if err := _db.AutoMigrate(&Recommendation{}); err != nil {
		return err
	}

	return nil
}
