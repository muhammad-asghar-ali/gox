-- name: CreateTicket :one
INSERT INTO tickets (event_id, ticket_type, price, total_tickets, available_tickets)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
