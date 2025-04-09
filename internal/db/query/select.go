package query

const (
	selectAlbumsForUserBase = `
		SELECT
			AR."name" AS "artist_name",
			AL."name" AS "album_name",
			AL."id" AS "album_id"
		FROM "user_album" UA
		INNER JOIN "album" AL ON AL."id" = UA."album_id"
		INNER JOIN "artist" AR ON AR."id" = AL."artist_id"
	`

	SelectAlbumsForUserNoFilters              = selectAlbumsForUserBase + `WHERE UA."user_id" = $1`
	SelectAlbumsForUserAnyGenreNotRecommended = SelectAlbumsForUserNoFilters + ` AND UA."recommended" = false`
	SelectAlbumsForUserWithGenres             = selectAlbumsForUserBase + `
		INNER JOIN "album_genre" AG ON AG."album_id" = UA."album_id"
		INNER JOIN "genre" G ON G."id" = AG."genre_id"
		WHERE UA."user_id" = $1
		AND G."name" = ANY($2)
	`
	SelectAlbumsForUserWithGenresNotRecommended = SelectAlbumsForUserWithGenres + ` AND UA."recommended" = false`

	SelectAllArtists = `SELECT id FROM "artist"`
)
