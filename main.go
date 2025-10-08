package main

import (
	"os"

	"github.com/dcmcand/camp-quansight-2025-go-bof/pkg/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.ListenaAndServe(port)
}
