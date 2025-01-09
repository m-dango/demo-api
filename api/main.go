package main

import (
	"log"
	"net/http"

	"github.com/m-dango/demo-api/internal/handlers"
)

func main() {
	s := &http.Server{
		Handler: handlers.CreateHandler(),
		Addr:    ":8080",
	}

	log.Fatal(s.ListenAndServe())
}
