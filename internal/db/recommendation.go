package db

func SaveRecommendation(userId, albumId string) error {
	recommendation := Recommendation{
		UserId:  userId,
		AlbumId: albumId,
	}

	return conn.Create(&recommendation).Error
}
