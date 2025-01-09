package main_test

import (
	"testing"

	main "github.com/m-dango/demo-api"
	"github.com/stretchr/testify/assert"
)

func TestCreateHandler(t *testing.T) {
	h := main.CreateHandler()
	assert.NotNil(t, h)
}
