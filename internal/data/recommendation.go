package data

import (
	"time"

	"github.com/wraith29/apollo/internal/model"
	"gorm.io/gorm"
)

func SaveRecommendation(db *gorm.DB, recommendation *recommendedAlbum) error {
	rec := model.Recommendation{
		AlbumId: recommendation.Id,
		Date:    time.Now(),
	}

	db.Create(&rec)

	return db.Error
}
