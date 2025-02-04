-- name: CreateEvent :one
INSERT INTO events (name, description, added_by, venue_id, event_date)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
