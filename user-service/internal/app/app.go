package app

import (
	"github.com/alibekkenny/simple-marketplace/user-service/internal/config"
	"github.com/rs/zerolog"
)

type Application struct {
	Logger *zerolog.Logger
	Config *config.Config
}

func NewApplication(logger *zerolog.Logger, config *config.Config) *Application {
	return &Application{Logger: logger, Config: config}
}
