package model

type Genre struct {
	Id      string `gorm:"primaryKey"`
	Name    string
	Rating  int
	Albums  []Album  `gorm:"many2many:album_genre"`
	Artists []Artist `gorm:"many2many:artist_genre"`
}

func (g Genre) TableName() string {
	return "genre"
}
