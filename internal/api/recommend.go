package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/wraith29/apollo/internal/ctx"
	"github.com/wraith29/apollo/internal/db"
)

func getGenres(req *http.Request) []string {
	genreList := req.FormValue("genres")

	if genreList == "" {
		return []string{}
	}

	return strings.Split(genreList, ",")
}

func getIncludeListened(req *http.Request) bool {
	include := req.FormValue("include-listened")

	if include == "" {
		return false
	}

	return include == "true"
}

func Recommend(w http.ResponseWriter, req *http.Request) {
	userId := req.Context().Value(ctx.ContextKeyUserId).(string)

	genres := getGenres(req)
	includeListened := getIncludeListened(req)

	albums, err := db.GetUserAlbums(userId, includeListened, genres)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	if len(albums) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
}
