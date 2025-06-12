package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/wraith29/apollo/internal/ctx"
	"github.com/wraith29/apollo/internal/db"
	"github.com/wraith29/apollo/internal/musicbrainz"
)

type command interface {
	execute(*musicbrainz.Client) error
}

type updateCommand struct {
	artistId string
}

func (u *updateCommand) execute(client *musicbrainz.Client) error {
	fmt.Printf("Updating %s\n", u.artistId)

	artistData, err := client.LookupArtistById(u.artistId)
	if err != nil {
		return err
	}

	if err := db.SaveArtist(artistData); err != nil {
		return err
	}

	return nil
}

type addCommand struct {
	artistName, userId string
}

func (a *addCommand) execute(client *musicbrainz.Client) error {
	fmt.Printf("Adding %s\n", a.artistName)

	artistSearch, err := client.SearchArtistByName(a.artistName)
	if err != nil {
		return err
	}

	artistData, err := client.LookupArtistById(artistSearch.FindArtistWithShortestDistance(a.artistName).Id)
	if err != nil {
		return err
	}

	if err := db.SaveArtist(artistData); err != nil {
		return err
	}

	if err := db.AddArtistToUser(artistData, a.userId); err != nil {
		return err
	}

	return nil
}

// TODO: Try come up with a better name for this
type musicBrainzQueue struct {
	queue  chan command
	mutex  sync.Mutex
	client musicbrainz.Client
}

func newMusicBrainzQueue() musicBrainzQueue {
	return musicBrainzQueue{
		queue:  make(chan command, 50),
		mutex:  sync.Mutex{},
		client: musicbrainz.NewClient(),
	}
}

func (m *musicBrainzQueue) Add(userId, artistId string) {
	m.mutex.Lock()

	m.queue <- &addCommand{artistId, userId}

	m.mutex.Unlock()
}

func (m *musicBrainzQueue) Update(artistId string) {
	m.mutex.Lock()

	m.queue <- &updateCommand{artistId}

	m.mutex.Unlock()
}

func (m *musicBrainzQueue) poll() {
	for {
		select {
		case cmd := <-m.queue:
			if err := cmd.execute(&m.client); err != nil {
				fmt.Printf("%+v\n", err)
			}
		}
	}
}

type addRequest struct {
	ArtistName string `json:"artistName"`
}

func Post_Artist(w http.ResponseWriter, req *http.Request, s *server) {
	userId := req.Context().Value(ctx.ContextKeyUserId).(string)

	var body addRequest

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	s.queue.Add(userId, body.ArtistName)

	w.WriteHeader(http.StatusAccepted)
}

type updateRequest struct {
	ArtistId string `json:"artistId"`
}

func Post_Update(w http.ResponseWriter, req *http.Request, s *server) {
	var body updateRequest

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	s.queue.Update(body.ArtistId)

	w.WriteHeader(http.StatusAccepted)
}
