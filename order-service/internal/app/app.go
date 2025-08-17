package app

import (
	"log"

	"github.com/alibekkenny/simple-marketplace/order-service/internal/config"
	"github.com/alibekkenny/simple-marketplace/order-service/internal/db"
	"github.com/alibekkenny/simple-marketplace/order-service/internal/repository/postgres_repo"
	"github.com/alibekkenny/simple-marketplace/order-service/internal/repository/redis_repo"
	"github.com/alibekkenny/simple-marketplace/order-service/internal/service"
	transport_grpc "github.com/alibekkenny/simple-marketplace/order-service/internal/transport/grpc"
	"github.com/alibekkenny/simple-marketplace/order-service/internal/transport/grpc/client"
	"github.com/go-playground/validator/v10"
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
	orderDB, err := db.NewPostgresDB(a.cfg.OrderDSN)
	if err != nil {
		log.Fatalf("couldn't connect to order db: %v", err.Error())
	}
	defer orderDB.Close()

	redis, err := db.NewRedisDB(a.cfg.CartDSN)
	if err != nil {
		log.Fatalf("couldn't connect to cart db: %v", err.Error())
	}
	defer redis.Close()

	validator := validator.New()

	cartRepo := redis_repo.NewCartRedisRepository(redis)
	cartService := service.NewCartService(validator, cartRepo)
	cartHandler := transport_grpc.NewCartHandler(cartService)

	productOfferClient, err := client.NewProductOfferClient(a.cfg.ProductServiceAddr)
	if err != nil {
		return err
	}

	orderRepo := postgres_repo.NewOrderPostgresRepository(orderDB)
	orderService := service.NewOrderService(validator, orderRepo, cartService, productOfferClient)
	orderHandler := transport_grpc.NewOrderHandler(orderService)

	return NewGRPCServer(a.cfg.ServiceAddr, cartHandler, orderHandler)
}
