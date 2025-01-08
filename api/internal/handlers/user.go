package handlers

import (
	"context"

	"github.com/m-dango/demo-api/internal/generated"
)

func (*Server) PostUser(ctx context.Context, request generated.PostUserRequestObject) (generated.PostUserResponseObject, error) {
	if request.Body.Email == nil || *request.Body.Email == "" || request.Body.Name == nil || *request.Body.Name == "" {
		resType := "example"
		resTitle := "Bad Request"
		resDetail := "There was a problem with the request."
		return generated.PostUser400ApplicationProblemPlusJSONResponse{
			ProblemApplicationProblemPlusJSONResponse: generated.ProblemApplicationProblemPlusJSONResponse{
				Type:   &resType,
				Title:  &resTitle,
				Detail: &resDetail,
			},
		}, nil
	}

	userId := "1"
	return generated.PostUser201JSONResponse{
		UserJSONResponse: generated.UserJSONResponse{
			Id:   &userId,
			Name: request.Body.Name,
		},
	}, nil
}
