package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/wraith29/apollo/internal/ctx"
	"github.com/wraith29/apollo/internal/db"
)

type rateRequest struct {
	AlbumId string `json:"albumId"`
	Rating  int    `json:"rating"`
}

func Put_Rating(w http.ResponseWriter, req *http.Request) {
	userId := req.Context().Value(ctx.ContextKeyUserId).(string)

	var body rateRequest

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	if body.Rating <= 0 || body.Rating > 5 {
		writeError(w, http.StatusBadRequest, errors.New("rating must be between 1 and 5 inclusive"))
		return
	}

	if err := db.RateAlbumForUser(userId, body.AlbumId, body.Rating-3); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
