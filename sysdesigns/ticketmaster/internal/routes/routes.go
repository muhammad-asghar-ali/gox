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
	api.Post("/venues", venue.CreateVenue)
	api.Get("/venues", venue.ListVenue)

	// ---------------- EVENTS
	event := handlers.NewEventHandler(services.EventService{})
	api.Post("/events", event.CreateEvent)
	api.Get("/events", event.ListEvent)
	api.Get("/events/:id", event.GetEventByID)

	// ---------------- PERFORMERS
	performer := handlers.NewPerformerHandler(services.PerformerService{})
	api.Post("/performers", performer.AddPerformer)
	api.Get("/performers", performer.ListPerformer)

	// ---------------- BOOKINGS
	booking := handlers.NewBookingHandler(services.BookingService{})
	api.Post("/bookings", booking.CreateBooking)
	api.Get("/users/bookings", booking.GetUserBookings)
	api.Post("/bookings/ticket", booking.BookTicket)
	api.Get("/bookings/confirm/:id", booking.ConfirmBooking)
	api.Get("/bookings/cancel/:id", booking.CancelBooking)
	api.Post("/bookings/:id", booking.GetBookingByID)

	// ---------------- TICKETS
	ticket := handlers.NewTicketHandler(services.TicketService{})
	api.Post("/tickets", ticket.CreateTicket)
	api.Get("/tickets/available/:id", ticket.GetAvailableTickets)

	// ---------------- EVENT_PERFORMERS
	event_performer := handlers.NewEventPerformerHandler(services.EventPerformerService{})
	api.Post("/events/add/performers", event_performer.AddPerformerToEvent)

	// ---------------- PAYMENTS
	payment := handlers.NewPaymentHandler(services.PaymentService{})
	api.Post("/payments/confirm/:id", payment.ConfirmPayment)
	api.Post("/payments/fail/:id", payment.FailPayment)

}
