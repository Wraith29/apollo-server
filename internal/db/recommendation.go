package db

import "time"

func SaveRecommendation(userId, albumId string) error {
	recommendation := Recommendation{
		UserId:  userId,
		AlbumId: albumId,
	}

	return conn.Create(&recommendation).Error
}

type userRecommendationListItem struct {
	AlbumName string
	CreatedAt *time.Time
}

func GetAllRecommendationsForUser(userId string) ([]userRecommendationListItem, error) {
	var recommendations []userRecommendationListItem

	if err := conn.
		Table("recommendations").
		Select("albums.name AS album_name", "recommendations.created_at AS created_at").
		Joins("INNER JOIN albums ON albums.id = recommendations.album_id").
		Where("user_id = ?", userId).
		Find(&recommendations).
		Error; err != nil {
		return nil, err
	}

	return recommendations, nil
}
