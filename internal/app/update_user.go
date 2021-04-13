package app

import (
	apiUser "magnet-test-task/internal/generated/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
)

func (srv *Service) UpdateUserHandler(params apiUser.UpdateUserParams) middleware.Responder {
	return middleware.NotImplemented("operation user UpdateUser has not yet been implemented")
}
