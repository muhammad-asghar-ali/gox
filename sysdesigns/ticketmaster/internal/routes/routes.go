package routes

import (
	"github.com/gofiber/fiber/v3"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/handlers"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/middlewares"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/services"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	health := handlers.NewHealthHandler()
	user := handlers.NewAuthHandler(services.UserService{})

	api.Get("/health", health.HealthCheck)
	api.Post("/register", user.Register)
	api.Post("/login", user.Login)

	venue := handlers.NewVenueHandler(services.VenueService{})
	app.Use(middlewares.Auth)
	app.Get("/venues", venue.CreateVenue)
}
