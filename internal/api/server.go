package api

import (
	"net/http"
)

type server struct {
	mux   *http.ServeMux
	queue updateQueue
}

type statefulHandler func(http.ResponseWriter, *http.Request, *server)

func NewServer() server {
	return server{
		mux:   http.NewServeMux(),
		queue: newUpdateQueue(),
	}
}

func (s *server) Run() error {
	go s.queue.run()

	return http.ListenAndServe(":5000", s.mux)
}

func (s *server) AddRoute(path string, handler http.HandlerFunc) {
	s.mux.HandleFunc(path, handler)
}

func (s *server) AddAuthenticatedRoute(path string, handler http.HandlerFunc) {
	s.AddRoute(path, UserIdMiddleware(handler))
}

func (s *server) AddAuthenticatedStatefulRoute(path string, handler statefulHandler) {
	s.AddAuthenticatedRoute(path, func(w http.ResponseWriter, req *http.Request) {
		handler(w, req, s)
	})
}
