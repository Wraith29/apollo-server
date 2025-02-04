package model

type Artist struct {
	Id     string `gorm:"primaryKey"`
	Name   string
	Genres []Genre `gorm:"many2many:artist_genre"`
}
