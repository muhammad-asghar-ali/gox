-- name: CreateBooking :one
INSERT INTO bookings (user_id, ticket_id, quantity, total_price, status)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
