package main

import (
	"fmt"

	"github.com/wraith29/apollo/internal/api"
	"github.com/wraith29/apollo/internal/db"
	"github.com/wraith29/apollo/internal/env"
)

func main() {
	if err := env.Load(); err != nil {
		panic(err)
	}

	if err := db.InitDb(); err != nil {
		panic(err)
	}

	server := api.NewServer()

	server.AddRoute("POST /init", api.Init)
	server.AddAuthenticatedRoute("POST /artist", api.AddArtist)
	server.AddAuthenticatedRoute("GET /recommendation", api.Recommend)
	server.AddAuthenticatedRoute("PUT /rate", api.Rate)
	server.AddAuthenticatedStatefulRoute("POST /update", api.Update)

	fmt.Printf("Starting server on port 5000\n")

	// if err := server.Run(); err != nil {
	// 	fmt.Printf("%+v\n", err)
	// 	os.Exit(1)
	// }
}
