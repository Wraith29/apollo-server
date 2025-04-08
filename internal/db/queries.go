package db

// TODO: Investigate a better way to do this.
const (
	insertArtist = `
		INSERT INTO "artist" ("id", "name")
		VALUES ($1, $2)
		ON CONFLICT ("id") DO NOTHING
	`

	insertAlbum = `
		INSERT INTO "album" ("id", "name", "release_date", "artist_id")
		VALUES ($1, $2, $3, $4)
		ON CONFLICT ("id") DO NOTHING
	`

	insertGenre = `
		INSERT INTO "genre" ("id", "name")
		VALUES ($1, $2)
		ON CONFLICT ("id") DO NOTHING
	`

	insertArtistGenre = `
		INSERT INTO "artist_genre" ("artist_id", "genre_id")
		VALUES ($1, $2)
	`

	insertAlbumGenre = `
		INSERT INTO "album_genre" ("album_id", "genre_id")
		VALUES ($1, $2)
	`

	insertUser = `
		INSERT INTO "user" ("id", "name")
		VALUES ($1, $2)
	`

	insertUserArtist = `
		INSERT INTO "user_artist" ("user_id", "artist_id")
		VALUES ($1, $2)
		ON CONFLICT ("user_id", "artist_id") DO NOTHING
	`

	insertUserAlbum = `
		INSERT INTO "user_album" ("user_id", "album_id")
		VALUES ($1, $2)
		ON CONFLICT ("user_id", "album_id") DO NOTHING
	`

	insertUserGenre = `
		INSERT INTO "user_genre" ("user_id", "genre_id")
		VALUES ($1, $2)
		ON CONFLICT ("user_id", "genre_id") DO NOTHING
	`

	insertRecommendation = `
		INSERT INTO "recommendation" ("user_id", "album_id", "listened_date")
		VALUES ($1, $2, $3)
	`

	select_AlbumsForUserBase = `
		SELECT
			AR."name" AS "artist_name",
			AL."name" AS "album_name",
			AL."id" AS "album_id"
		FROM "user_album" UA
		INNER JOIN "album" AL ON AL."id" = UA."album_id"
		INNER JOIN "artist" AR ON AR."id" = AL."artist_id"
	`

	select_AlbumsForUserNoFilters              = select_AlbumsForUserBase + `WHERE UA."user_id" = $1`
	select_AlbumsForUserAnyGenreNotRecommended = select_AlbumsForUserNoFilters + ` AND UA."recommended" = false`
	select_AlbumsForUserWithGenres             = select_AlbumsForUserBase + `
		INNER JOIN "album_genre" AG ON AG."album_id" = UA."album_id"
		INNER JOIN "genre" G ON G."id" = AG."genre_id"
		WHERE UA."user_id" = $1
		AND G."name" = ANY($2)
	`
	select_AlbumsForUserWithGenresNotRecommended = select_AlbumsForUserWithGenres + ` AND UA."recommended" = false`
)
