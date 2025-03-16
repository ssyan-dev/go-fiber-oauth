package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Address string `env:"SERVER_ADDRESS" envDefault:":8080"`
	}
}

func NewConfig() (*Config, error) {
	godotenv.Load()

	var cfg Config
	if err := env.Parse(&cfg.Server); err != nil {
		return nil, fmt.Errorf("config parse error: %w", err)
	}

	return &cfg, nil
}
