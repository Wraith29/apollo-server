package query

const (
	UpdateUserAlbumRating = `
		UPDATE "user_album" SET "rating" = $1 WHERE "user_id" = $2 AND "album_id" = $3
	`
)
