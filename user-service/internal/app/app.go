package app

import (
	"github.com/rs/zerolog"
)

type Application struct {
	Logger *zerolog.Logger
}

func NewApplication(logger *zerolog.Logger) *Application {
	return &Application{Logger: logger}
}
