package query

const (
	InsertArtist = `
		INSERT INTO "artist" ("id", "name")
		VALUES ($1, $2)
		ON CONFLICT ("id") DO NOTHING
	`

	InsertAlbum = `
		INSERT INTO "album" ("id", "name", "release_date", "artist_id")
		VALUES ($1, $2, $3, $4)
		ON CONFLICT ("id") DO NOTHING
	`

	InsertGenre = `
		INSERT INTO "genre" ("id", "name")
		VALUES ($1, $2)
		ON CONFLICT ("id") DO NOTHING
	`

	InsertArtistGenre = `
		INSERT INTO "artist_genre" ("artist_id", "genre_id")
		VALUES ($1, $2)
	`

	InsertAlbumGenre = `
		INSERT INTO "album_genre" ("album_id", "genre_id")
		VALUES ($1, $2)
	`

	InsertUser = `
		INSERT INTO "user" ("id", "name")
		VALUES ($1, $2)
	`

	InsertUserArtist = `
		INSERT INTO "user_artist" ("user_id", "artist_id")
		VALUES ($1, $2)
		ON CONFLICT ("user_id", "artist_id") DO NOTHING
	`

	InsertUserAlbum = `
		INSERT INTO "user_album" ("user_id", "album_id")
		VALUES ($1, $2)
		ON CONFLICT ("user_id", "album_id") DO NOTHING
	`

	InsertUserGenre = `
		INSERT INTO "user_genre" ("user_id", "genre_id")
		VALUES ($1, $2)
		ON CONFLICT ("user_id", "genre_id") DO NOTHING
	`

	InsertRecommendation = `
		INSERT INTO "recommendation" ("user_id", "album_id", "listened_date")
		VALUES ($1, $2, $3)
	`
)
