package model

type Album struct {
	Id           string `gorm:"primaryKey"`
	ArtistId     string
	Artist       Artist `gorm:"foreignKey:ArtistId"`
	Name         string
	Listened     bool
	ListenedDate string
	Rating       int
	Genres       []Genre `gorm:"many2many:album_genre"`
}
