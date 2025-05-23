package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wraith29/apollo/internal/ctx"
	"github.com/wraith29/apollo/internal/db"
)

func Get_ListArtists(w http.ResponseWriter, req *http.Request) {
	userId := req.Context().Value(ctx.ContextKeyUserId).(string)

	artists, err := db.GetAllUserArtistsForListing(userId)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	bytes, err := json.Marshal(artists)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	if _, err := w.Write(bytes); err != nil {
		fmt.Printf("Error writing response: %+v\n", err)
	}
}

func Get_ListAlbums(w http.ResponseWriter, req *http.Request) {
	userId := req.Context().Value(ctx.ContextKeyUserId).(string)

	artists, err := db.GetAllUserAlbumsForListing(userId)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	bytes, err := json.Marshal(artists)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	if _, err := w.Write(bytes); err != nil {
		fmt.Printf("Error writing response: %+v\n", err)
	}
}

func Get_ListRecommendations(w http.ResponseWriter, req *http.Request) {
	userId := req.Context().Value(ctx.ContextKeyUserId).(string)

	recommendations, err := db.GetAllRecommendationsForUser(userId)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	bytes, err := json.Marshal(recommendations)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	if _, err := w.Write(bytes); err != nil {
		fmt.Printf("Error writing response: %+v\n", err)
	}
}
