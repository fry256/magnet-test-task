package main

import (
	"log"

	"magnet-test-task/internal/config"
	"magnet-test-task/internal/generated/restapi"
	"magnet-test-task/internal/generated/restapi/operations"

	apiUser "magnet-test-task/internal/generated/restapi/operations/user"

	"github.com/go-openapi/loads"

	"magnet-test-task/internal/app"
)

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	srv := app.New()
	api := operations.NewMagnetTestTaskAPI(swaggerSpec)

	api.UserCreateUserHandler = apiUser.CreateUserHandlerFunc(srv.CreateUserHandler)
	api.UserDeleteUserHandler = apiUser.DeleteUserHandlerFunc(srv.DeleteUserHandler)
	api.UserGetUserByIDHandler = apiUser.GetUserByIDHandlerFunc(srv.GetUserByIDHandler)
	api.UserUpdateUserHandler = apiUser.UpdateUserHandlerFunc(srv.UpdateUserHandler)
	api.ServerShutdown = srv.OnShutdown
	server := restapi.NewServer(api)
	defer server.Shutdown()

	cfg, err := config.InitConfig("magnet_test_task")
	if err != nil {
		log.Fatalln(err)
	}

	server.ConfigureAPI()

	server.Port = cfg.HTTPPort
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
