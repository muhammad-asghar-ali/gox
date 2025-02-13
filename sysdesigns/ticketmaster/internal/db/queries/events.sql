-- name: CreateEvent :one
INSERT INTO events (name, description, added_by, venue_id, event_date)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: ListEvent :many
SELECT * FROM events;

-- name: GetEventByID :one
SELECT 
    sqlc.embed(e), 
    sqlc.embed(t),
    JSON_AGG(
        JSON_BUILD_OBJECT(
            'performer_id', p.id,
            'performer_name', p.name,
            'genre', p.genre,
            'bio', p.bio
        )
    ) AS performers
FROM events e
INNER JOIN tickets t ON t.event_id = e.id
INNER JOIN event_performers ep ON ep.event_id = e.id
INNER JOIN performers p ON p.id = ep.performer_id
WHERE e.id = $1
GROUP BY e.id, t.id;



