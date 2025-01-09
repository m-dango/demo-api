package main

import (
	"log"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/m-dango/demo-api/internal/handlers"
	middleware "github.com/oapi-codegen/nethttp-middleware"
)

func CreateHandler() http.Handler {
	baseUrl := "/v0"
	server := handlers.NewServer()
	swagger, err := handlers.GetSwagger()
	if err != nil {
		log.Fatalln("spec error:", err)
	}
	swagger.Servers = openapi3.Servers{&openapi3.Server{URL: baseUrl}}

	// middleware for validating requests with spec
	options := middleware.Options{SilenceServersWarning: true}
	mw := middleware.OapiRequestValidatorWithOptions(swagger, &options)
	// Server Interface for the handler
	si := handlers.NewStrictHandler(server, nil)
	// Handler for http.Server
	h := handlers.HandlerFromMuxWithBaseURL(si, nil, baseUrl)
	// Wrap handler in validation middleware
	return mw(h)
}

func main() {
	s := &http.Server{
		Handler: CreateHandler(),
		Addr:    ":8080",
	}

	log.Fatal(s.ListenAndServe())
}
