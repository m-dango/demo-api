package handlers_test

import (
	"testing"

	"github.com/m-dango/demo-api/internal/handlers"
	"github.com/stretchr/testify/assert"
)

func TestCreateHandler(t *testing.T) {
	h := handlers.CreateHandler()
	assert.NotNil(t, h)
}
