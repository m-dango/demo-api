//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=server.config.yaml ../../openapi.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=types.config.yaml ../../openapi.yaml
package handlers

import (
	"github.com/m-dango/demo-api/internal/generated"
)

var _ generated.StrictServerInterface = (*Server)(nil)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}
