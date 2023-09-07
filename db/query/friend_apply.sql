-- name: CreateFriendApply :one
INSERT INTO friend_apply(
    apply_id, reply_id, apply_desc, note
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: ListFriendApply :many
SELECt 
    fp.apply_id, u.nickname, fp.apply_desc, fp.created_at
FROM friend_apply AS fp
JOIN users AS u ON fp.apply_id = u.id AND fp.reply_id = $1;

-- name: DeleteFriendApply :exec
DELETE FROM friend_apply 
WHERE apply_id = $1 AND reply_id = $2;

-- name: GetFriendApply :one
SELECT * FROM friend_apply
WHERE apply_id = $1 AND reply_id = $2;