-- name: AddMessage :one
INSERT INTO messages (
    sender_id, receiver_id, message_type, content, send_type, sending_time
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;