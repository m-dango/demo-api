package main

import (
	"log"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/m-dango/demo-api/internal/generated"
	"github.com/m-dango/demo-api/internal/handlers"
	middleware "github.com/oapi-codegen/nethttp-middleware"
)

func main() {
	server := handlers.NewServer()
	swagger, err := generated.GetSwagger()
	if err != nil {
		log.Fatalln("spec error:", err)
	}
	swagger.Servers = openapi3.Servers{&openapi3.Server{URL: "/v0"}}

	// middleware for validating requests with spec
	options := middleware.Options{SilenceServersWarning: true}
	mw := middleware.OapiRequestValidatorWithOptions(swagger, &options)
	// Server Interface for the handler
	si := generated.NewStrictHandler(server, nil)
	// Handler for http.Server
	h := generated.HandlerFromMuxWithBaseURL(si, nil, "/v0")
	// Wrap handler in validation middleware
	h = mw(h)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(s.ListenAndServe())
}
