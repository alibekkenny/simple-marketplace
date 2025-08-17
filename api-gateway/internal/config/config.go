package config

import "os"

type Config struct {
	Addr               string
	UserServiceAddr    string
	ProductServiceAddr string
	OrderServiceAddr   string
	JWTSecret          string
}

func Load() *Config {
	addr := os.Getenv("SERVICE_ADDR")
	userServiceAddr := os.Getenv("USER_SERVICE_ADDR")
	productServiceAddr := os.Getenv("PRODUCT_SERVICE_ADDR")
	orderServiceAddr := os.Getenv("ORDER_SERVICE_ADDR")
	jwtSecret := os.Getenv("JWT_SECRET")

	return &Config{
		Addr:               addr,
		UserServiceAddr:    userServiceAddr,
		ProductServiceAddr: productServiceAddr,
		OrderServiceAddr:   orderServiceAddr,
		JWTSecret:          jwtSecret,
	}
}
