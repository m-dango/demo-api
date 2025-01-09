//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=server.config.yaml ../../openapi.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=types.config.yaml ../../openapi.yaml
package handlers

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	middleware "github.com/oapi-codegen/nethttp-middleware"
)

var _ StrictServerInterface = (*Server)(nil)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func CreateHandler() http.Handler {
	baseUrl := "/v0"
	server := NewServer()
	swagger, _ := GetSwagger()
	swagger.Servers = openapi3.Servers{&openapi3.Server{URL: baseUrl}}

	// middleware for validating requests with spec
	options := middleware.Options{SilenceServersWarning: true}
	mw := middleware.OapiRequestValidatorWithOptions(swagger, &options)
	// Server Interface for the handler
	si := NewStrictHandler(server, nil)
	// Handler for http.Server
	h := HandlerFromMuxWithBaseURL(si, nil, baseUrl)
	// Wrap handler in validation middleware
	return mw(h)
}
