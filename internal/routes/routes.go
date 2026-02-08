package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omarghandour/fibertest/internal/handlers"
)

// SetupRoutes registers all application routes
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api") // Grouping routes under /api

	// Health Check
	api.Get("/health", handlers.HealthCheck)
}
