package api

import (
	"net/http"
	"sync"
	"time"

	"github.com/wraith29/apollo/internal/ctx"
	"github.com/wraith29/apollo/internal/db"
)

type updateQueue struct {
	queue    chan string
	mutex    sync.Mutex
	previous time.Time
}

func newUpdateQueue() updateQueue {
	return updateQueue{
		queue: make(chan string, 50),
		mutex: sync.Mutex{},
	}
}

func (u *updateQueue) push(artistId string) {
	u.mutex.Lock()
	u.queue <- artistId

	u.mutex.Unlock()
}

func (u *updateQueue) run() {
	for {
		select {
		case artistId := <-u.queue:
			u.mutex.Lock()

			u.mutex.Unlock()
		}
	}
}

func Update(w http.ResponseWriter, req *http.Request, s *server) {
	userId := req.Context().Value(ctx.ContextKeyUserId).(string)

	artistIds, err := db.GetUserArtistIds(userId)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	for _, id := range artistIds {
		s.queue.push(id)
	}
}
