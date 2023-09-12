-- name: AddMessage :one
INSERT INTO messages (
    send_id, receive_id, content
) VALUES (
    $1, $2, $3
)
RETURNING *;