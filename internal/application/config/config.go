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
	OAuth struct {
		GitHub struct {
			ClientID     string `env:"GITHUB_CLIENT_ID"`
			ClientSecret string `env:"GITHUB_CLIENT_SECRET"`
			RedirectURL  string `env:"GITHUB_REDIRECT_URL"`
		}
	}
}

func NewConfig() (*Config, error) {
	godotenv.Load()

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("config parse error: %w", err)
	}

	return &cfg, nil
}
