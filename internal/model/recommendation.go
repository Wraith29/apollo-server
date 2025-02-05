package model

import "time"

type Recommendation struct {
	Id      uint `gorm:"primaryKey"`
	AlbumId string
	Album   Album `gorm:"foreignKey:AlbumId"`
	Date    time.Time
}

func (r Recommendation) TableName() string {
	return "recommendation"
}
