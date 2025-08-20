package config

import "os"

type Config struct {
	DSN    string
	JWTKey string
	Addr   string
}

func Load() *Config {
	dsn := os.Getenv("DATABASE_URL")
	jwtKey := os.Getenv("JWT_SECRET")
	addr := os.Getenv("SERVICE_ADDR")

	return &Config{
		DSN:    dsn,
		JWTKey: jwtKey,
		Addr:   addr,
	}
}
