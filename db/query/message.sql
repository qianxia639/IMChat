-- name: AddMessage :one
INSERT INTO messages (
    sender_id, receiver_id, content, sender_time
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;