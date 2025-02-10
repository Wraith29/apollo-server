package data

import (
	"github.com/wraith29/apollo/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetGenres(db *gorm.DB, listAll bool) (model.ListResult[model.Genre], error) {
	genres := make([]model.Genre, 0)

	var count int
	db.Raw("SELECT COUNT(id) FROM genre").Scan(&count)

	query := db.Order(
		clause.OrderByColumn{Column: clause.Column{Name: "rating"}, Desc: true},
	)

	if !listAll {
		query.Limit(10)
	}

	query.Find(&genres)

	return model.ListResult[model.Genre]{
		Count:   count,
		Results: genres,
	}, db.Error
}
