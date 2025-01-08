//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=server.config.yaml ../../openapi.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=types.config.yaml ../../openapi.yaml
package handlers

var _ StrictServerInterface = (*Server)(nil)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}
