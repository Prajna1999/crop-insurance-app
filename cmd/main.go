package main

import (
	"log"

	"github.com/Prajna1999/dataspace-be/internal/app"
)

func main() {
	application, err := app.NewApp()

	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	if err := application.Run(); err != nil {
		log.Fatalf("Failed to run application: %v", err)
	}
}
