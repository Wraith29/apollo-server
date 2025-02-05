package model

import "time"

type Album struct {
	Id           string `gorm:"primaryKey"`
	ArtistId     string
	Artist       Artist `gorm:"foreignKey:ArtistId"`
	Name         string
	Listened     bool
	ListenedDate *time.Time
	Rating       int
	Genres       []Genre `gorm:"many2many:album_genre"`
}

func (a Album) TableName() string {
	return "album"
}
