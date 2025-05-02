package api

import (
	"net/http"
)

type server struct {
	mux   *http.ServeMux
	queue musicBrainzQueue
}

type statefulHandler func(http.ResponseWriter, *http.Request, *server)

func NewServer() server {
	return server{
		mux:   http.NewServeMux(),
		queue: newMusicBrainzQueue(),
	}
}

func (s *server) Run() error {
	go s.queue.poll()

	return http.ListenAndServe(":5000", s.mux)
}

func (s *server) AddRoute(path string, handler http.HandlerFunc) {
	s.mux.HandleFunc(path, LoggingMiddleware(handler))
}

func (s *server) AddAuthenticatedRoute(path string, handler http.HandlerFunc) {
	s.AddRoute(path, AuthenticationMiddleware(handler))
}

func (s *server) AddAuthenticatedStatefulRoute(path string, handler statefulHandler) {
	s.AddAuthenticatedRoute(path, func(w http.ResponseWriter, req *http.Request) {
		handler(w, req, s)
	})
}
