package handlers

import "context"

func (*Server) PostUser(ctx context.Context, request PostUserRequestObject) (PostUserResponseObject, error) {
	if request.Body.Email == nil || *request.Body.Email == "" || request.Body.Name == nil || *request.Body.Name == "" {
		resType := "example"
		resTitle := "Bad Request"
		resDetail := "There was a problem with the request."
		return PostUser400ApplicationProblemPlusJSONResponse{
			ProblemApplicationProblemPlusJSONResponse{
				Type:   &resType,
				Title:  &resTitle,
				Detail: &resDetail,
			},
		}, nil
	}

	userId := "1"
	return PostUser201JSONResponse{
		UserJSONResponse{
			Id:   &userId,
			Name: request.Body.Name,
		},
	}, nil
}
