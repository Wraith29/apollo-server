package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wraith29/apollo/internal/ctx"
	"github.com/wraith29/apollo/internal/db"
)

func Get_Artists(w http.ResponseWriter, req *http.Request) {
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
