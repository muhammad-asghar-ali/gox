package routes

import (
	"github.com/gofiber/fiber/v3"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/handlers"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/services"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	health := handlers.NewHealthHandler()
	user := handlers.NewAuthHandler(services.UserService{})

	api.Get("/health", health.HealthCheck)
	api.Get("/register", user.Register)
	api.Get("/login", user.Login)
}
