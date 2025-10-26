package main

import (
	"fmt"
	"log"

	"github.com/sorfian/go-contact-management-api/app"
)

func main() {
	// Load configuration
	config := app.LoadConfig()

	// Initialize the app with all dependencies using Wire
	fiberApp := InitializeApp()

	// Start server
	log.Printf("Starting server on port %s in %s mode...", config.AppPort, config.AppEnv)
	log.Fatal(fiberApp.Listen(fmt.Sprintf(":%s", config.AppPort)))
}
