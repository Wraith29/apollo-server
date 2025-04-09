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

// TODO:
// Figure out how to make HTTP Handlers be methods on the server type
// ^ Ensure this is optional, only on endpoints that _need_ server information
// ^ Namely the Update endpoint, as it will have access to the update queue from the server
// Remove old server implementation
