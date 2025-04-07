package db

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
	`

	insertUserAlbum = `
		INSERT INTO "user_album" ("user_id", "album_id")
		VALUES ($1, $2)
	`

	insertUserGenre = `
		INSERT INTO "user_genre" ("user_id", "genre_id")
		VALUES ($1, $2)
	`

	insertRecommendation = `
		INSERT INTO "recommendation" ("user_id", "album_id", "listened_date")
		VALUES ($1, $2, $3)
	`
)
