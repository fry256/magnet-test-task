package app

import (
	"context"
	"magnet-test-task/internal/generated/restapi/operations/user"
	apiUser "magnet-test-task/internal/generated/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
)

func (srv *Service) CreateUserHandler(params apiUser.CreateUserParams) middleware.Responder {
	if err := srv.Box.Storage.Store(context.Background(), params.Body); err != nil {
		return user.NewCreateUserInternalServerError()
	}

	return user.NewCreateUserOK().WithPayload(params.Body)
}
