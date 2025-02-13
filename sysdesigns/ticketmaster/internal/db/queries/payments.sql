-- name: CreatePayment :one
INSERT INTO payments (booking_id, user_id, amount, payment_method, status)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
