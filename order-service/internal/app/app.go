package app

import "github.com/alibekkenny/simple-marketplace/order-service/internal/config"

type App struct {
	cfg *config.Config
}

func NewApp() (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	return &App{cfg: cfg}, nil
}

func (a *App) Run() {

}
