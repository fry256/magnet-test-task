package app

import (
	apiUser "magnet-test-task/internal/generated/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
)

func (srv *Service) CreateUserHandler(params apiUser.CreateUserParams) middleware.Responder {
	return middleware.NotImplemented("operation user CreateUser has not yet been implemented")
}
