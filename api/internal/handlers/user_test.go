package handlers_test

import (
	"context"
	"testing"

	"github.com/m-dango/demo-api/internal/generated"
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
		reqBody := generated.User{
			Name:  &c.name,
			Email: &c.email,
		}
		req := generated.PostUserRequestObject{Body: &reqBody}

		server := handlers.NewServer()
		got, err := server.PostUser(context.Background(), req)

		assert.Nil(err)
		if assert.NotNil(got) {
			id := "1"
			want := generated.PostUser201JSONResponse{
				UserJSONResponse: generated.UserJSONResponse{
					Id:   &id,
					Name: &c.name,
				},
			}

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
		reqBody generated.User
	}{
		"Missing email": {
			generated.User{Name: &name},
		},
		"Empty email": {
			generated.User{Name: &name, Email: &emptyString},
		},
		"Missing name": {
			generated.User{Email: &email},
		},
		"Empty name": {
			generated.User{Email: &email, Name: &emptyString},
		},
	}

	resType := "example"
	resTitle := "Bad Request"
	resDetail := "There was a problem with the request."
	want := generated.PostUser400ApplicationProblemPlusJSONResponse{
		ProblemApplicationProblemPlusJSONResponse: generated.ProblemApplicationProblemPlusJSONResponse{Type: &resType, Title: &resTitle, Detail: &resDetail},
	}

	for name, c := range cases {
		req := generated.PostUserRequestObject{Body: &c.reqBody}

		server := handlers.NewServer()
		got, err := server.PostUser(context.Background(), req)

		assert.Nil(err)
		if assert.NotNil(got) {
			assert.Equal(want, got, name)
		}
	}
}
