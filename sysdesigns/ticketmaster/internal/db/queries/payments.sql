-- name: CreatePayment :one
INSERT INTO payments (booking_id, user_id, amount, payment_method, status)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: ConfirmPayment :exec
WITH updated_payment AS (
    UPDATE payments
    SET status = 'completed'
    WHERE payments.id = $1 AND status = 'pending'
    RETURNING booking_id
)
UPDATE bookings
SET status = 'confirmed'
WHERE bookings.id = (SELECT booking_id FROM updated_payment);

-- name: FailPayment :exec
WITH updated_payment AS (
    UPDATE payments
    SET status = 'failed'
    WHERE payments.id = $1 AND status = 'pending'
    RETURNING booking_id
)
UPDATE bookings
SET status = 'canceled'
WHERE bookings.id = (SELECT booking_id FROM updated_payment);
