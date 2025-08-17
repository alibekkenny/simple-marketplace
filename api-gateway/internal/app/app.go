package app

import (
	"net/http"

	"github.com/alibekkenny/simple-marketplace/api-gateway/internal/config"
	"github.com/alibekkenny/simple-marketplace/api-gateway/internal/transport/grpc"
	"github.com/alibekkenny/simple-marketplace/api-gateway/internal/transport/handler"
)

type App struct {
	cfg *config.Config
}

func NewApp() *App {
	cfg := config.Load()
	return &App{cfg: cfg}
}

func (a *App) Run() error {

	userClient, err := grpc.NewUserClient(a.cfg.UserServiceAddr)
	if err != nil {
		return err
	}
	categoryClient, err := grpc.NewCategoryClient(a.cfg.ProductServiceAddr)
	if err != nil {
		return err
	}
	productClient, err := grpc.NewProductClient(a.cfg.ProductServiceAddr)
	if err != nil {
		return err
	}
	productOfferClient, err := grpc.NewProductOfferClient(a.cfg.ProductServiceAddr)
	if err != nil {
		return err
	}
	cartClient, err := grpc.NewCartClient(a.cfg.OrderServiceAddr)
	if err != nil {
		return err
	}
	orderClient, err := grpc.NewOrderClient(a.cfg.OrderServiceAddr)
	if err != nil {
		return err
	}

	userHandler := handler.NewUserHandler(userClient)
	categoryHandler := handler.NewCategoryHandler(categoryClient)
	productHandler := handler.NewProductHandler(productClient)
	productOfferHandler := handler.NewProductOfferHandler(productOfferClient)
	cartHandler := handler.NewCartHandler(cartClient)
	orderHandler := handler.NewOrderHandler(orderClient)

	mux := a.routes(userHandler, categoryHandler, productHandler, productOfferHandler, cartHandler, orderHandler)

	srv := &http.Server{
		Addr:    a.cfg.Addr,
		Handler: mux,
	}

	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
