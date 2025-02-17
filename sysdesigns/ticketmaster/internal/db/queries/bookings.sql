-- name: CreateBooking :one
-- @optional user_id
INSERT INTO bookings (user_id, ticket_id, quantity, total_price, status)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserBookings :many
SELECT * FROM bookings WHERE user_id = $1;

-- name: BookTicket :exec
WITH updated_ticket AS (
    UPDATE tickets
    SET available_tickets = available_tickets - $3
    WHERE id = $1 AND available_tickets >= $3
    RETURNING id
)
INSERT INTO bookings (user_id, ticket_id, quantity, total_price, status)
SELECT $2, $1, $3, $4, 'pending'
FROM updated_ticket;

-- name: CancelBooking :exec
WITH updated_booking AS (
    UPDATE bookings
    SET status = 'canceled'
    WHERE bookings.id = $1 AND status != 'canceled'
    RETURNING ticket_id, quantity
)
UPDATE tickets
SET available_tickets = tickets.available_tickets + updated_booking.quantity
FROM updated_booking
WHERE tickets.id = updated_booking.ticket_id;

-- name: ConfirmBooking :exec
UPDATE bookings
SET status = 'confirmed'
WHERE id = $1 AND status = 'pending';

-- name: GetBookingStatus :one
SELECT status
FROM bookings
WHERE id = $1;
