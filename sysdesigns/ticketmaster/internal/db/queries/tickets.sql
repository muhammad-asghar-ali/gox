-- name: CreateTicket :one
INSERT INTO tickets (event_id, ticket_type, price, total_tickets, available_tickets)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetAvailableTickets :one
SELECT available_tickets
FROM tickets
WHERE id = $1;

-- name: GetRemainingTickets :one
SELECT 
    t.available_tickets AS ticket_available,
    e.available_tickets AS event_available
FROM tickets t
JOIN events e ON e.id = t.event_id
WHERE t.id = $1;
