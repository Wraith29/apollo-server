package db

import (
	"database/sql"
	"time"
)

type saveRecommendationQuery struct {
	userId, albumId string
	date            time.Time
}

func (s *saveRecommendationQuery) execute(txn *sql.Tx) error {
	return prepareAndExecute(
		txn,
		insertRecommendation,
		s.userId,
		s.albumId,
		s.date.Format(dateFormat),
	)
}

func SaveRecommendation(userId, albumId string, date time.Time) query {
	return &saveRecommendationQuery{userId, albumId, date}
}
