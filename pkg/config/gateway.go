package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Gateway is the struct containing all services specific configuration
type Gateway struct {
	Services ServicesConfiguration
	Server   ServerConfiguration
}

// ServicesConfiguration has all configs for the services
type ServicesConfiguration struct {
	UsersPort   int
	PostsPort   int
	SourcesPort int
}

// NewGateway news up a gateway configuration
func NewGateway(configLocation string) (*Gateway, error) {
	cfg := Gateway{}
	viper.SetConfigFile(configLocation)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrapf(err, "failed to read configuration from %s", configLocation)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal config to struct")
	}

	return &cfg, nil
}
