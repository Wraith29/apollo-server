package main

import (
	"fmt"
	"os"

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

	server.AddRoute("POST /auth/register", api.Post_Register)
	server.AddRoute("POST /auth/login", api.Post_Login)
	server.AddAuthenticatedRoute("GET /auth/refresh", api.Get_Refresh)
	server.AddAuthenticatedStatefulRoute("POST /artist", api.Post_Artist)
	server.AddAuthenticatedStatefulRoute("POST /artist/update", api.Post_Update)

	fmt.Printf("Starting server on port 5000\n")

	if err := server.Run(); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}
