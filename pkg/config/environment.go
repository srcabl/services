package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Environment is the struct containing all environment specific configuration
type Environment struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

// New news up an environment configuration
func New() (*Environment, error) {
	cfg := Environment{}
	envconfig.Process("", &cfg)
	return &cfg, nil
}
