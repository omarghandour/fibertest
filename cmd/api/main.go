package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/omarghandour/fibertest/internal/routes"
)

func main() {
	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Fiber Starter Project",
	})

	// Setup Routes
	routes.SetupRoutes(app)

	// Start Server
	log.Fatal(app.Listen(":3000"))
}
