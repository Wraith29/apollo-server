package api

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type command struct {
	isUpdate bool
	artistId string
	userId   string
}

// TODO: Try come up with a better name for this
type musicBrainzQueue struct {
	queue    chan command
	mutex    sync.Mutex
	previous time.Time
}

func newMusicBrainzQueue() musicBrainzQueue {
	return musicBrainzQueue{
		queue:    make(chan command, 50),
		mutex:    sync.Mutex{},
		previous: time.Now(),
	}
}

func (m *musicBrainzQueue) Push(userId, artistId string) {
	m.mutex.Lock()
	m.queue <- command{
		isUpdate: false,
		artistId: artistId,
		userId:   userId,
	}

	m.mutex.Unlock()
}

func (m *musicBrainzQueue) PushUpdate(artistId string) {
	m.mutex.Lock()
	m.queue <- command{
		isUpdate: true,
		artistId: artistId,
	}

	m.mutex.Unlock()
}

func (m *musicBrainzQueue) poll() {
	for {
		select {
		case cmd := <-m.queue:
			fmt.Printf("%+v\n", cmd)
		}
	}
}

func Post_Artist(w http.ResponseWriter, req *http.Request, s *server) {
	writeError(w, http.StatusInternalServerError, errors.New("not implemented"))
}

func Post_Update(w http.ResponseWriter, req *http.Request, s *server) {
	writeError(w, http.StatusInternalServerError, errors.New("not implemented"))
}
