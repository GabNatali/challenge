package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/subosito/gotenv"
)

type Config struct {
	HttpHost              string `envconfig:"Http_HOST"`
	HttpPort              int    `envconfig:"Http_PORT"`
	DatabaseURL           string `envconfig:"DATABASE_URL"`
	AccessTokenExpiresTTL int    `envconfig:"ACCESS_TOKEN_EXPIRES_TTL"`
	AccessTokenSecret     string `envconfig:"ACCESS_TOKEN_SECRET"`
}

func ParseEnv(envPath string) (*Config, error) {
	if envPath != "" {
		if err := gotenv.OverLoad(envPath); err != nil {
			return nil, err
		}
	}

	var config Config

	if err := envconfig.Process("", &config); err != nil {
		return nil, err
	}

	return &config, nil
}
