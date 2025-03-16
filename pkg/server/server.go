package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	port int
	mux  http.ServeMux
}

func NewServer(port int) Server {
	return Server{
		port,
		*http.NewServeMux(),
	}
}

func (s *Server) Add(path string, handler http.HandlerFunc) {
	s.mux.HandleFunc(path, handler)
}

func (s *Server) Run() error {
	return http.ListenAndServe(
		fmt.Sprintf(":%d", s.port),
		&s.mux,
	)
}
