package types

import (
	"github.com/google/uuid"
	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/internal/entities"
)

type (
	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
)

type (
	GetEventByID struct {
		Event      entities.Event `json:"event"`
		Venue      entities.Venue `json:"venue"`
		Ticket     []Ticket       `json:"tickets"`
		Performers []Performer    `json:"performers"`
	}

	Performer struct {
		ID    uuid.UUID `json:"id"`
		Name  string    `json:"name"`
		Genre string    `json:"genre"`
		Bio   string    `json:"bio"`
	}

	Ticket struct {
		ID               uuid.UUID `json:"id"`
		TicketType       string    `json:"ticket_type"`
		Price            float64   `json:"price"`
		TotalTickets     int32     `json:"total_tickets"`
		AvailableTickets int32     `json:"available_tickets"`
	}
)
