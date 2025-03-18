package main

import (
	"fmt"
	"net/http"
	"os"
)

func ping(w http.ResponseWriter, req *http.Request) {
	_ = req

	if _, err := w.Write([]byte("Pong!")); err != nil {
		panic(err)
	}
}

func main() {
	server := http.NewServeMux()

	server.HandleFunc("GET /ping", ping)

	fmt.Printf("Starting server on port 5000\n")

	if err := http.ListenAndServe(":5000", server); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}
