-- name: AddPerformerToEvent :one
INSERT INTO event_performers (event_id, performer_id)
VALUES ($1, $2)
RETURNING *;