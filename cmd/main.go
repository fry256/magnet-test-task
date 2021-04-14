package main

import (
	"fmt"
	"log"

	"magnet-test-task/env/prod"
	"magnet-test-task/internal/config"

	"magnet-test-task/internal/generated/restapi"
	"magnet-test-task/internal/generated/restapi/operations"

	apiUser "magnet-test-task/internal/generated/restapi/operations/user"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"

	"github.com/go-openapi/loads"

	"magnet-test-task/internal/app"
)

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	cfg, err := config.InitConfig("magnet_test_task")
	if err != nil {
		log.Fatalln(err)
	}

	box, cleanup, err := prod.InitializeBox(cfg)
	if cleanup != nil {
		defer cleanup()
	}

	if err != nil {
		log.Fatalf("InitializeBox error: %v", err)
	}

	if cfg.SQLLite.MigrationPath != "" {
		driver, err := sqlite3.WithInstance(box.Storage.GetInstance(), &sqlite3.Config{})
		if err != nil {
			log.Fatal("can't get database driver: ", err)
		}
		migrations, err := migrate.NewWithDatabaseInstance(
			fmt.Sprintf("file://%s/", cfg.SQLLite.MigrationPath),
			cfg.SQLLite.Name,
			driver,
		)
		if err != nil {
			log.Fatal("can't get migrations: ", err)
		}
		if err = migrations.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("can't roll up migrations: ", err)
		}
	}

	srv := app.New(box)
	api := operations.NewMagnetTestTaskAPI(swaggerSpec)

	api.UserCreateUserHandler = apiUser.CreateUserHandlerFunc(srv.CreateUserHandler)
	api.UserDeleteUserHandler = apiUser.DeleteUserHandlerFunc(srv.DeleteUserHandler)
	api.UserGetUserByIDHandler = apiUser.GetUserByIDHandlerFunc(srv.GetUserByIDHandler)
	api.UserUpdateUserHandler = apiUser.UpdateUserHandlerFunc(srv.UpdateUserHandler)
	api.ServerShutdown = srv.OnShutdown
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.ConfigureAPI()

	server.Port = cfg.HTTPPort
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
