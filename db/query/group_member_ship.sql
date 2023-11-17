-- name: AddGroupMember :one
INSERT INTO group_member_ships(
    user_id, group_id, nickname, role
)
VALUES(
    $1, $2, $3, $4
)
RETURNING *;

-- name: GetGroup :one
SELECT * FROM group_member_ships
WHERE
    user_id = $1 AND group_id = $2;