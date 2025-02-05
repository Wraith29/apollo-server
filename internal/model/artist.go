package model

type Artist struct {
	Id            string `gorm:"primaryKey"`
	Name          string
	OverallRating int
	Genres        []Genre `gorm:"many2many:artist_genre"`
}

func (a Artist) TableName() string {
	return "artist"
}
