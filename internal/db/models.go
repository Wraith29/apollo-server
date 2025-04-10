package db

import "time"

type Genre struct {
	Id     string `gorm:"primaryKey"`
	Name   string
	Rating int `gorm:"default:0"`

	Artists []Artist `gorm:"many2many:artist_genres"`
	Albums  []Album  `gorm:"many2many:album_genres"`
}

type Artist struct {
	Id        string `gorm:"primaryKey"`
	Name      string
	Rating    int `gorm:"default:0"`
	UpdatedAt time.Time

	Genres []Genre `gorm:"many2many:artist_genres"`
}

type Album struct {
	Id          string `gorm:"primaryKey"`
	Name        string
	Rating      int `gorm:"default:0"`
	ReleaseDate time.Time

	ArtistId string
	Artist   Artist

	Genres []Genre `gorm:"many2many:album_genres"`
}

type User struct {
	Id       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string

	Genres  []Genre  `gorm:"many2many:user_genres"`
	Artists []Artist `gorm:"many2many:user_artists"`
	Albums  []Album  `gorm:"many2many:user_albums"`
}

type UserGenre struct {
	UserId    uint   `gorm:"primaryKey"`
	GenreId   string `gorm:"primaryKey"`
	Rating    int    `gorm:"default:0"`
	CreatedAt time.Time
}

type UserArtist struct {
	UserId    uint   `gorm:"primaryKey"`
	ArtistId  string `gorm:"primaryKey"`
	Rating    int    `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserAlbum struct {
	UserId      uint   `gorm:"primaryKey"`
	AlbumId     string `gorm:"primaryKey"`
	Rating      int    `gorm:"default:0"`
	Recommended bool   `gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Recommendation struct {
	Id        uint `gorm:"primaryKey"`
	UserId    uint
	User      User
	GenreId   string
	Genre     Genre
	CreatedAt time.Time
}
