package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omarghandour/fibertest/internal/models"
)

// HealthCheck mimics a simple Elysia handler
func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(models.GenericResponse{
		Success: true,
		Message: "Server is up and running!",
		Data: fiber.Map{
			"status": "ok",
		},
	})
}
