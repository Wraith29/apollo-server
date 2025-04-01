package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wraith29/apollo/internal/ctx"
	"github.com/wraith29/apollo/internal/data"
	mb "github.com/wraith29/apollo/internal/data/musicbrainz"
)

type addRequest struct {
	ArtistName string `json:"artist_name"`
}

func saveArtistWithId(userId, artistId string) error {
	artistData, err := mb.LookupArtistById(artistId)
	if err != nil {
		return err
	}

	return data.PersistArtistForUser(userId, artistData)
}

func AddArtist(w http.ResponseWriter, req *http.Request) {
	userId := req.Context().Value(ctx.ContextKeyUserId).(string)

	var body addRequest
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	mbData, err := mb.SearchArtistByName(body.ArtistName)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	if mbData.Count <= 0 {
		writeError(w, http.StatusNoContent, fmt.Errorf("no artists found for %s", body.ArtistName))
		return
	}

	artist := mbData.FindArtistWithShortestDistance(body.ArtistName)

	if err := saveArtistWithId(userId, artist.Id); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	println("Committed transaction")

	w.WriteHeader(http.StatusOK)
}

func AddArtistInteractive(w http.ResponseWriter, req *http.Request) {

}
