package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DSN         string
	ServiceAddr string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		DSN:         os.Getenv("DSN"),
		ServiceAddr: os.Getenv("SERVICE_ADDR"),
	}

	return cfg, nil
}
