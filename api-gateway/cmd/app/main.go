package main

import (
	"log"

	"github.com/alibekkenny/simple-marketplace/api-gateway/internal/app"
)

func main() {
	app := app.NewApp()

	if err := app.Run(); err != nil {
		log.Fatalf("failed to run application: %v", err)
	}
}
