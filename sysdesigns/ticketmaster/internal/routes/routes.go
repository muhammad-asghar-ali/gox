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
	api.Get("/health", health.HealthCheck)

	// ---------------- AUTH
	user := handlers.NewAuthHandler(services.UserService{})
	api.Post("/register", user.Register)
	api.Post("/login", user.Login)

	// ---------------- VENUES
	venue := handlers.NewVenueHandler(services.VenueService{})
	app.Use(middlewares.Auth)
	app.Get("/venues", venue.CreateVenue)

	// ---------------- EVENTS
	event := handlers.NewEventHandler(services.EventService{})
	app.Get("/events", event.CreateEvent)
}
