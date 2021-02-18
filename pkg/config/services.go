package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Service is the struct containing all service specific configuration
type Service struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

// DatabaseConfiguration has all configs for the database
type DatabaseConfiguration struct {
	User     string
	Password string
	Address  string
	Port     string
	Name     string
}

// NewService news up a service configuration
func NewService(configLocation string) (*Service, error) {
	cfg := Service{}
	viper.SetConfigFile(configLocation)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrapf(err, "failed to read configuration from %s", configLocation)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal config to struct")
	}

	return &cfg, nil
}
