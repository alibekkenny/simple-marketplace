package app

import (
	"github.com/alibekkenny/simple-marketplace/product-service/internal/config"
	"github.com/alibekkenny/simple-marketplace/product-service/internal/db"
	"github.com/alibekkenny/simple-marketplace/product-service/internal/repository/postgres_repo"
	"github.com/alibekkenny/simple-marketplace/product-service/internal/service"
	grpcTransport "github.com/alibekkenny/simple-marketplace/product-service/internal/transport/grpc"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

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

func (a *App) Run() error {
	db, err := db.NewPostgresDB(a.cfg.DSN)
	if err != nil {
		return err
	}
	defer db.Close()

	categoryRepo := postgres_repo.NewCategoryPostgresRepository(db)
	validator := validator.New()
	categoryService := service.NewCategoryService(categoryRepo, validator)
	categoryHandler := grpcTransport.NewCategoryHandler(categoryService)

	return NewGRPCServer(a.cfg.ServiceAddr, categoryHandler)
}
