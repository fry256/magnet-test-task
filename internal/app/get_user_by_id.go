package app

import (
	apiUser "magnet-test-task/internal/generated/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
)

func (srv *Service) GetUserByIDHandler(params apiUser.GetUserByIDParams) middleware.Responder {
	return middleware.NotImplemented("operation user GetUserByID has not yet been implemented")
}
