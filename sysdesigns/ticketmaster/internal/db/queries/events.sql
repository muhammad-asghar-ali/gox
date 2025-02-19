-- name: CreateEvent :one
INSERT INTO events (name, description, added_by, venue_id, event_date, total_tickets, available_tickets)
VALUES ($1, $2, $3, $4, $5, $6, $6)
RETURNING *;

-- name: ListEvent :many
SELECT * FROM events;

-- name: GetEventByID :one
SELECT 
    sqlc.embed(e),
    sqlc.embed(v),
    JSON_AGG(
        JSON_BUILD_OBJECT(
            'id', p.id,
            'name', p.name,
            'genre', p.genre,
            'bio', p.bio
        )
    ) AS performers,
    JSON_AGG(
        JSON_BUILD_OBJECT(
            'id', t.id,
            'ticket_type', t.ticket_type,
            'price', t.price,
            'total_tickets', t.total_tickets,
            'available_tickets', t.available_tickets
        )
    ) AS tickets
FROM events e
INNER JOIN venues v ON v.id = e.venue_id
LEFT JOIN tickets t ON t.event_id = e.id
LEFT JOIN event_performers ep ON ep.event_id = e.id
LEFT JOIN performers p ON p.id = ep.performer_id
WHERE e.id = $1
GROUP BY e.id, t.id, v.id;

-- name: SearchEvents :many
SELECT *
FROM events
WHERE 
  COALESCE($1, '') = '' OR name ILIKE '%' || COALESCE($1, '') || '%' OR description ILIKE '%' || COALESCE($1, '') || '%'
  AND event_date BETWEEN $2 AND $3
ORDER BY event_date ASC
LIMIT $4
OFFSET $5;
