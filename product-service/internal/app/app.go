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

	validator := validator.New()

	categoryRepo := postgres_repo.NewCategoryPostgresRepository(db)
	categoryService := service.NewCategoryService(categoryRepo, validator)
	categoryHandler := grpcTransport.NewCategoryHandler(categoryService)

	productRepo := postgres_repo.NewProductPostgresRepository(db)
	productService := service.NewProductService(productRepo, validator)
	productHandler := grpcTransport.NewProductHandler(productService)

	productOfferRepo := postgres_repo.NewProductOfferPostgresRepository(db)
	productOfferService := service.NewProductOfferService(productOfferRepo, validator)
	productOfferHandler := grpcTransport.NewProductOfferHandler(productOfferService)

	return NewGRPCServer(a.cfg.ServiceAddr, categoryHandler, productHandler, productOfferHandler)
}
