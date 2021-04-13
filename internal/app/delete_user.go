package app

import (
	apiUser "magnet-test-task/internal/generated/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
)

func (srv *Service) DeleteUserHandler(params apiUser.DeleteUserParams) middleware.Responder {
	return middleware.NotImplemented("operation user DeleteUser has not yet been implemented")
}
