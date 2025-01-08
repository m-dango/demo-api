package handlers_test

import (
	"context"
	"testing"

	"github.com/m-dango/demo-api/internal/handlers"
	"github.com/stretchr/testify/assert"
)

func TestPostUserSuccess(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		name, email string
	}{
		{name: "Alice", email: "alice@example.com"},
		{name: "Bob", email: "bob@example.com"},
	}

	for _, c := range cases {
		reqBody := handlers.User{
			Name:  &c.name,
			Email: &c.email,
		}
		req := handlers.PostUserRequestObject{Body: &reqBody}

		server := handlers.NewServer()
		got, err := server.PostUser(context.Background(), req)

		assert.Nil(err)
		if assert.NotNil(got) {
			id := "1"
			want := handlers.PostUser201JSONResponse{handlers.UserJSONResponse{
				Id:   &id,
				Name: &c.name,
			}}

			assert.Equal(want, got, c.name)
		}
	}
}

func TestPostUserBadFields(t *testing.T) {
	assert := assert.New(t)

	name := "Charlie"
	email := "charlie@example.com"
	emptyString := ""
	cases := map[string]struct {
		reqBody handlers.User
	}{
		"Missing email": {
			handlers.User{Name: &name},
		},
		"Empty email": {
			handlers.User{Name: &name, Email: &emptyString},
		},
		"Missing name": {
			handlers.User{Email: &email},
		},
		"Empty name": {
			handlers.User{Email: &email, Name: &emptyString},
		},
	}

	resType := "example"
	resTitle := "Bad Request"
	resDetail := "There was a problem with the request."
	want := handlers.PostUser400ApplicationProblemPlusJSONResponse{
		handlers.ProblemApplicationProblemPlusJSONResponse{Type: &resType, Title: &resTitle, Detail: &resDetail},
	}

	for name, c := range cases {
		req := handlers.PostUserRequestObject{Body: &c.reqBody}

		server := handlers.NewServer()
		got, err := server.PostUser(context.Background(), req)

		assert.Nil(err)
		if assert.NotNil(got) {
			assert.Equal(want, got, name)
		}
	}
}
