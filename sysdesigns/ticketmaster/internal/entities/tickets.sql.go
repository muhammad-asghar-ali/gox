// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: tickets.sql

package entities

import (
	"context"

	"github.com/google/uuid"
)

const createTicket = `-- name: CreateTicket :one
INSERT INTO tickets (event_id, ticket_type, price, total_tickets, available_tickets)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, event_id, ticket_type, price, total_tickets, available_tickets, created_at
`

type CreateTicketParams struct {
	EventID          uuid.UUID `json:"event_id"`
	TicketType       string    `json:"ticket_type"`
	Price            float64   `json:"price"`
	TotalTickets     int32     `json:"total_tickets"`
	AvailableTickets int32     `json:"available_tickets"`
}

func (q *Queries) CreateTicket(ctx context.Context, arg CreateTicketParams) (Ticket, error) {
	row := q.db.QueryRow(ctx, createTicket,
		arg.EventID,
		arg.TicketType,
		arg.Price,
		arg.TotalTickets,
		arg.AvailableTickets,
	)
	var i Ticket
	err := row.Scan(
		&i.ID,
		&i.EventID,
		&i.TicketType,
		&i.Price,
		&i.TotalTickets,
		&i.AvailableTickets,
		&i.CreatedAt,
	)
	return i, err
}
