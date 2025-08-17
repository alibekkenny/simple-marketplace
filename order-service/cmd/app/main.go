package main

import (
	"log"

	"github.com/alibekkenny/simple-marketplace/order-service/internal/app"
)

func main() {
	a, err := app.NewApp()
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}

	if err := a.Run(); err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
