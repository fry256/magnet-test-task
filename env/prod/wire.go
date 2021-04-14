//+build wireinject

package prod

import (
	"fmt"
	"log"
	"os"

	"magnet-test-task/env"
	"magnet-test-task/internal/config"
	"magnet-test-task/pkg/storage"

	"github.com/google/wire"
)

//go:generate wire wire.go

// InitializeBox wire injection
func InitializeBox(cfg *config.Config) (*env.Box, func(), error) {
	wire.Build(
		NewBox,
		provideStorage,
	)

	return nil, nil, nil
}

func provideStorage(cfg *config.Config) *storage.Storage {
	if err := os.MkdirAll(cfg.SQLLite.Path, 0750); err != nil {
		log.Fatal("database directory is not accessible: ", err)
	}
	dsn := fmt.Sprintf("%s/%s?mode=rw", cfg.SQLLite.Path, cfg.SQLLite.Name)
	s, err := storage.New(dsn)
	if err != nil {
		log.Fatal(err)
	}

	return s
}
