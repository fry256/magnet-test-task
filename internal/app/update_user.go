package app

import (
	"context"
	"magnet-test-task/internal/generated/restapi/operations/user"
	apiUser "magnet-test-task/internal/generated/restapi/operations/user"
	"magnet-test-task/pkg/storage"

	"github.com/go-openapi/runtime/middleware"
)

func (srv *Service) UpdateUserHandler(params apiUser.UpdateUserParams) middleware.Responder {
	params.Body.UID = params.UID
	err := srv.Box.Storage.Store(context.Background(), params.Body)
	if err != nil {
		if err == storage.UserNotFound {
			return user.NewUpdateUserNotFound()
		}

		return user.NewUpdateUserInternalServerError()
	}

	return user.NewUpdateUserOK().WithPayload(params.Body)
}
