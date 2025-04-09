package db

import (
	"database/sql"
	"time"

	"github.com/wraith29/apollo/internal/db/query"
)

type saveRecommendationQuery struct {
	userId, albumId string
	date            time.Time
}

func (s *saveRecommendationQuery) write(txn *sql.Tx) error {
	return prepAndExec(
		txn,
		query.InsertRecommendation,
		s.userId,
		s.albumId,
		s.date.Format(dateFormat),
	)
}

func SaveRecommendation(userId, albumId string, date time.Time) dbWriter {
	return &saveRecommendationQuery{userId, albumId, date}
}
