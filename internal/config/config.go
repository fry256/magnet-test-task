package config

import (
	"github.com/pkg/errors"
	"github.com/vrischmann/envconfig"
)

// Config struct
type Config struct {
	HTTPPort int `envconfig:"default=8001"`
	SQLLite  struct {
		Name          string `envconfig:"default=database.sqlite"`
		Path          string `envconfig:"default=/etc/sqlite"`
		MigrationPath string `envconfig:"default=/etc/sqlite/migrations"`
	}
}

// InitConfig func
func InitConfig(prefix string) (*Config, error) {
	config := &Config{}
	if err := envconfig.InitWithPrefix(config, prefix); err != nil {
		return nil, errors.Wrap(err, "init config failed")
	}

	return config, nil
}
