package config

import (
	"os"
)

type Config struct {
	OrderDSN           string
	CartDSN            string
	ServiceAddr        string
	ProductServiceAddr string
}

func Load() (*Config, error) {
	cfg := &Config{
		OrderDSN:           os.Getenv("DATABASE_URL"),
		CartDSN:            os.Getenv("REDIS_URL"),
		ServiceAddr:        os.Getenv("SERVICE_ADDR"),
		ProductServiceAddr: os.Getenv("PRODUCT_SERVICE_ADDR"),
	}

	return cfg, nil
}
