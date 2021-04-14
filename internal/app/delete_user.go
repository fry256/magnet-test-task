package app

import (
	"context"
	"magnet-test-task/internal/generated/restapi/operations/user"
	apiUser "magnet-test-task/internal/generated/restapi/operations/user"
	"magnet-test-task/pkg/storage"

	"github.com/go-openapi/runtime/middleware"
)

func (srv *Service) DeleteUserHandler(params apiUser.DeleteUserParams) middleware.Responder {
	err := srv.Box.Storage.Delete(context.Background(), params.UID)
	if err != nil {
		if err == storage.UserNotFound {
			return user.NewDeleteUserNotFound()
		}

		return user.NewDeleteUserInternalServerError()
	}

	return user.NewDeleteUserOK()
}
