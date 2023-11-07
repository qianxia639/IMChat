-- name: CreateGroup :one
INSERT INTO groups (
    creator_id, group_name, icon, description
) VALUES(
    $1, $2, $3, $4
)
RETURNING *;