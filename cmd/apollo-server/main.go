package main

import (
	"fmt"
	"os"

	"github.com/wraith29/apollo/internal/api"
	"github.com/wraith29/apollo/internal/db"
	"github.com/wraith29/apollo/internal/env"
)

func main() {
	if os.Getenv("APOLLO_ENV") != "prod" {
		if err := env.Load(); err != nil {
			panic(err)
		}
	}

	if err := db.InitDb(); err != nil {
		panic(err)
	}

	server := api.NewServer()

	server.AddRoute("POST /auth/register", api.Post_Register)
	server.AddRoute("POST /auth/login", api.Post_Login)

	server.AddAuthenticatedRoute("GET /auth/refresh", api.Get_Refresh)
	server.AddAuthenticatedRoute("GET /album/recommendation", api.Get_Recommendation)
	server.AddAuthenticatedRoute("PUT /album/rating", api.Put_Rating)

	server.AddAuthenticatedRoute("GET /artists", api.Get_ListArtists)
	server.AddAuthenticatedRoute("GET /albums", api.Get_ListAlbums)
	server.AddAuthenticatedRoute("GET /recommendations", api.Get_ListRecommendations)

	server.AddAuthenticatedStatefulRoute("POST /artist", api.Post_Artist)
	server.AddAuthenticatedStatefulRoute("POST /artist/update", api.Post_Update)

	fmt.Printf("Starting server on port 5000\n")

	if err := server.Run(); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}
