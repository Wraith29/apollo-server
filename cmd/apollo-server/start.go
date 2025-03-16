package main

import (
	"fmt"
	"net/http"

	"github.com/wraith29/apollo/pkg/config"
	"github.com/wraith29/apollo/pkg/server"
)

func start(args []string, cfg *config.Config) error {
	app := server.NewServer(cfg.Server.Port)

	app.Add("GET /ping", ping)

	return app.Run()
}

func ping(w http.ResponseWriter, req *http.Request) {
	if _, err := w.Write([]byte("Pong")); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", req)
}
