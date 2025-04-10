package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wraith29/apollo/internal/ctx"
)

type rateRequest struct {
	AlbumId string `json:"album_id"`
	Rating  int    `json:"rating"`
}

func Rate(w http.ResponseWriter, req *http.Request) {
	userId := req.Context().Value(ctx.ContextKeyUserId).(string)

	var body rateRequest
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	fmt.Printf("%s -> %+v\n", userId, body)

	// if err := db.Exec(db.RateAlbum(userId, body.AlbumId, body.Rating)); err != nil {
	// 	writeError(w, http.StatusInternalServerError, err)
	// 	return
	// }
}
