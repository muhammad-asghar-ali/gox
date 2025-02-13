// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: payments.sql

package entities

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createPayment = `-- name: CreatePayment :one
INSERT INTO payments (booking_id, user_id, amount, payment_method, status)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, booking_id, user_id, amount, payment_method, status, created_at
`

type CreatePaymentParams struct {
	BookingID     uuid.UUID      `json:"booking_id"`
	UserID *uuid.UUID      `json:"user_id"`
	Amount        pgtype.Numeric `json:"amount"`
	PaymentMethod string         `json:"payment_method"`
	Status        string         `json:"status"`
}

func (q *Queries) CreatePayment(ctx context.Context, arg CreatePaymentParams) (Payment, error) {
	row := q.db.QueryRow(ctx, createPayment,
		arg.BookingID,
		arg.UserID,
		arg.Amount,
		arg.PaymentMethod,
		arg.Status,
	)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.BookingID,
		&i.UserID,
		&i.Amount,
		&i.PaymentMethod,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}
