package config

import (
	"os"
)

type Config struct {
	DSN         string
	ServiceAddr string
}

func Load() (*Config, error) {
	cfg := &Config{
		DSN:         os.Getenv("DATABASE_URL"),
		ServiceAddr: os.Getenv("SERVICE_ADDR"),
	}

	return cfg, nil
}
