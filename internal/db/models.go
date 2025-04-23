package db

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/wraith29/apollo/internal/musicbrainz"
)

type Genre struct {
	Id     string `gorm:"primaryKey"`
	Name   string
	Rating int `gorm:"default:0"`

	Artists []Artist `gorm:"many2many:artist_genres"`
	Albums  []Album  `gorm:"many2many:album_genres"`
}

func GenresFromMusicBrainzGenres(genres []musicbrainz.Genre) []Genre {
	result := make([]Genre, len(genres))

	for idx, genre := range genres {
		result[idx] = Genre{Id: genre.Id, Name: genre.Name}
	}

	return result
}

type Artist struct {
	Id        string `gorm:"primaryKey"`
	Name      string
	Rating    int `gorm:"default:0"`
	UpdatedAt time.Time

	Albums []Album
	Genres []Genre `gorm:"many2many:artist_genres"`
}

func ArtistFromMusicBrainzArtist(artist *musicbrainz.Artist) (Artist, error) {
	albums := make([]Album, 0)

	for _, releaseGroup := range artist.ReleaseGroups {
		if !releaseGroup.IsValid() {
			continue
		}

		album, err := AlbumFromMusicBrainzReleaseGroup(releaseGroup, artist.Id)
		if err != nil {
			return Artist{}, err
		}

		albums = append(albums, album)
	}

	return Artist{
		Id:        artist.Id,
		Name:      artist.Name,
		UpdatedAt: time.Now(),
		Albums:    albums,
		Genres:    GenresFromMusicBrainzGenres(artist.Genres),
	}, nil
}

type Album struct {
	Id          string `gorm:"primaryKey"`
	Name        string
	Rating      int `gorm:"default:0"`
	ReleaseDate time.Time

	ArtistId string
	Artist   Artist

	Genres []Genre `gorm:"many2many:album_genres"`
}

func AlbumFromMusicBrainzReleaseGroup(rg musicbrainz.ReleaseGroup, artistId string) (Album, error) {
	releaseDate, err := time.Parse(dateFormat, rg.FirstReleaseDate)
	if err != nil {
		return Album{}, err
	}

	return Album{
		Id:          rg.Id,
		Name:        rg.Title,
		ReleaseDate: releaseDate,
		ArtistId:    artistId,
		Genres:      GenresFromMusicBrainzGenres(rg.Genres),
	}, nil
}

type User struct {
	Id       string `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string

	Genres  []Genre  `gorm:"many2many:user_genres"`
	Artists []Artist `gorm:"many2many:user_artists"`
	Albums  []Album  `gorm:"many2many:user_albums"`
}

func NewUser(username, password string) (User, error) {
	userId, err := generateUserIdFromUsername(username)
	if err != nil {
		return User{}, err
	}

	user := User{
		Id:       userId,
		Username: username,
		Password: password,
	}

	return user, nil
}

func generateUserIdFromUsername(username string) (string, error) {
	hash := sha256.New()
	if _, err := hash.Write([]byte(username)); err != nil {
		return "", err
	}

	userId := hex.EncodeToString(hash.Sum(nil))[:16]

	return userId, nil
}

type UserGenre struct {
	UserId    string `gorm:"primaryKey"`
	GenreId   string `gorm:"primaryKey"`
	Rating    int    `gorm:"default:0"`
	CreatedAt time.Time
}

type UserArtist struct {
	UserId    string `gorm:"primaryKey"`
	ArtistId  string `gorm:"primaryKey"`
	Rating    int    `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserAlbum struct {
	UserId      string `gorm:"primaryKey"`
	AlbumId     string `gorm:"primaryKey"`
	Rating      int    `gorm:"default:0"`
	Recommended bool   `gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Recommendation struct {
	Id        uint `gorm:"primaryKey"`
	UserId    string
	User      User
	AlbumId   string
	Album     Album
	CreatedAt time.Time
}
