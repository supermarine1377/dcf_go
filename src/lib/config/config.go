package config

import (
	"fmt"

	"github.com/caarlos0/env/v8"
)

type Config struct {
	// Financial Modeling Prep API key
	FMPAPIKey string `env:"FMP_API_KEY,notEmpty"`
}

func (c *Config) GetFMPAPIKey() string {
	return c.FMPAPIKey
}

func NewConfig() (*Config, error) {
	var conf Config
	if err := env.Parse(&conf); err != nil {
		return nil, fmt.Errorf("failed to env.Parse(): %w", err)
	}
	return &conf, nil
}
