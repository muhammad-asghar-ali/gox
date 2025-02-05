package routes

import (
	"github.com/gofiber/fiber/v3"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	health := handlers.NewHealthHandler()
	api.Get("/health", health.HealthCheck)
}
