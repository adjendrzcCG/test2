// Package main is the entry point for the task API server.
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/adjendrzcCG/test2/internal/handler"
	"github.com/adjendrzcCG/test2/internal/store"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	s := store.New()
	mux := handler.Tasks(s)

	log.Printf("server listening on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
