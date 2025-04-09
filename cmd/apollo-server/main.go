package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/wraith29/apollo/internal/api"
	"github.com/wraith29/apollo/internal/db"
	"github.com/wraith29/apollo/internal/env"
)

type statefulHandler func(*api.State, http.ResponseWriter, *http.Request)

type server struct {
	mux   *http.ServeMux
	state *api.State
}

func newServer() server {
	return server{
		mux:   http.NewServeMux(),
		state: api.NewState(),
	}
}

func (s *server) run() error {
	return http.ListenAndServe(":5000", s.mux)
}

func (s *server) addRoute(path string, handler http.HandlerFunc) {
	s.mux.HandleFunc(path, handler)
}

func (s *server) addAuthenticatedRoute(path string, handler http.HandlerFunc) {
	s.addRoute(path, api.UserIdMiddleware(handler))
}

func (s *server) addAuthenticatedStatefulRoute(path string, handler statefulHandler) {
	s.addAuthenticatedRoute(path, func(w http.ResponseWriter, req *http.Request) {
		handler(s.state, w, req)
	})
}

func main() {
	if err := env.Load(); err != nil {
		panic(err)
	}

	if err := db.InitDb(); err != nil {
		panic(err)
	}

	server := newServer()

	server.addRoute("POST /init", api.Init)
	server.addAuthenticatedRoute("POST /artist", api.AddArtist)
	server.addAuthenticatedRoute("GET /recommendation", api.Recommend)
	server.addAuthenticatedRoute("PUT /rate", api.Rate)
	server.addAuthenticatedStatefulRoute("POST /update", api.UpdateArtists)

	fmt.Printf("Starting server on port 5000\n")

	if err := server.run(); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}
