package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/wraith29/apollo/internal/ctx"
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

}
