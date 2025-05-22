package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/http"
	"strings"

	"github.com/wraith29/apollo/internal/ctx"
	"github.com/wraith29/apollo/internal/db"
)

func getGenreFilters(req *http.Request) []string {
	genres := req.FormValue("genres")

	if genres == "" {
		return []string{}
	}

	return strings.Split(genres, ",")
}

func getIncludeRecommended(req *http.Request) bool {
	includeRecommended := req.FormValue("include-recommended")

	return includeRecommended == "true"
}

func getRandomAlbum(albums []db.RecommendedAlbum) db.RecommendedAlbum {
	if len(albums) == 1 {
		return albums[0]
	}

	return albums[rand.IntN(len(albums)-1)]
}

func Get_Recommendation(w http.ResponseWriter, req *http.Request) {
	userId := req.Context().Value(ctx.ContextKeyUserId).(string)
	genres := getGenreFilters(req)
	includeRec := getIncludeRecommended(req)

	albums, err := db.GetUserAlbums(userId, genres, includeRec)
	if err != nil {
		writeError(w, http.StatusInternalServerError, errors.New("error retrieving user's albums"))
		return
	}

	if len(albums) == 0 {
		writeError(w, http.StatusBadRequest, errors.New("no albums found"))
		return
	}

	album := getRandomAlbum(albums)

	if err := db.SaveRecommendation(userId, album.AlbumId); err != nil {
		writeError(w, http.StatusInternalServerError, errors.New("failed to save recommendation"))
		return
	}

	bytes, err := json.Marshal(album)
	if err != nil {
		writeError(w, http.StatusInternalServerError, errors.New("failed to load response"))
		return
	}

	if _, err := w.Write(bytes); err != nil {
		fmt.Printf("Error writing response: %+v\n", err)
	}
}
