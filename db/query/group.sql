-- name: CreateGroup :one
INSERT INTO groups (
    creator_id, group_name, icon, description
) VALUES(
    $1, $2, $3, $4
)
RETURNING *;

-- name: ListGroup :many
SELECT g.id AS group_id, g.group_name, g.icon, u.id AS user_id, u.username, u.nickname
FROM groups g
JOIN users u ON g.creator_id = u.id
WHERE
    u.id = $1;
