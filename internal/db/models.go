package db

type userAlbum struct {
	ArtistName string `json:"artist_name"`
	AlbumName  string `json:"album_name"`
	AlbumId    string `json:"album_id"`
}
