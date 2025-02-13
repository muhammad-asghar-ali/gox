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
	Performer struct {
		PerformerID   uuid.UUID `json:"performer_id"`
		PerformerName string    `json:"performer_name"`
		Genre         string    `json:"genre"`
		Bio           string    `json:"bio"`
	}

	GetEventByID struct {
		Event      entities.Event  `json:"event"`
		Ticket     entities.Ticket `json:"ticket"`
		Performers []Performer     `json:"performers"`
	}
)
