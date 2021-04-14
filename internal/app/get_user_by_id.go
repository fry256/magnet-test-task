package app

import (
	"context"
	"magnet-test-task/internal/generated/restapi/operations/user"
	apiUser "magnet-test-task/internal/generated/restapi/operations/user"
	"magnet-test-task/pkg/storage"

	"github.com/go-openapi/runtime/middleware"
)

func (srv *Service) GetUserByIDHandler(params apiUser.GetUserByIDParams) middleware.Responder {
	u, err := srv.Box.Storage.OneByUID(context.Background(), params.UID)
	if err != nil {
		if err == storage.UserNotFound {
			return user.NewGetUserByIDNotFound()
		}

		return user.NewGetUserByIDInternalServerError()
	}

	return user.NewGetUserByIDOK().WithPayload(u)
}
