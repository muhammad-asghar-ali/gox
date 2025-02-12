// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: bookings.sql

package entities

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createBooking = `-- name: CreateBooking :one
INSERT INTO bookings (user_id, ticket_id, quantity, total_price, status)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, user_id, ticket_id, quantity, total_price, status, created_at
`

type CreateBookingParams struct {
	UserID     uuid.UUID      `json:"user_id"`
	TicketID   uuid.UUID      `json:"ticket_id"`
	Quantity   int32          `json:"quantity"`
	TotalPrice pgtype.Numeric `json:"total_price"`
	Status     string         `json:"status"`
}

func (q *Queries) CreateBooking(ctx context.Context, arg CreateBookingParams) (Booking, error) {
	row := q.db.QueryRow(ctx, createBooking,
		arg.UserID,
		arg.TicketID,
		arg.Quantity,
		arg.TotalPrice,
		arg.Status,
	)
	var i Booking
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TicketID,
		&i.Quantity,
		&i.TotalPrice,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getUserBookings = `-- name: GetUserBookings :many
SELECT id, user_id, ticket_id, quantity, total_price, status, created_at FROM bookings WHERE user_id = $1
`

func (q *Queries) GetUserBookings(ctx context.Context, userID uuid.UUID) ([]Booking, error) {
	rows, err := q.db.Query(ctx, getUserBookings, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Booking
	for rows.Next() {
		var i Booking
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.TicketID,
			&i.Quantity,
			&i.TotalPrice,
			&i.Status,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
