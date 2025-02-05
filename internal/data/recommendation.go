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

func GetLatestRecommendation(db *gorm.DB) (*model.Recommendation, error) {
	var rec model.Recommendation

	exists := 0

	db.Raw("SELECT EXISTS (SELECT 1 FROM `recommendation`)").Scan(&exists)

	if exists == 0 {
		return nil, db.Error
	}

	db.Last(&rec)

	return &rec, db.Error
}

func IsLatestRecommendationRated(db *gorm.DB) (bool, error) {
	latestRec, err := GetLatestRecommendation(db)
	if err != nil || latestRec == nil {
		return true, err
	}

	return latestRec.Rated, nil
}
