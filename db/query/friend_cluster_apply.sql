-- name: CreateFriendClsuterApply :one
INSERT INTO friend_cluster_apply(
    apply_id, receiver_id, apply_desc, flag
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: ExistsFriendClusterApply :one
SELECT COUNT(*) FROM friend_cluster_apply
WHERE
    apply_id = $1 AND receiver_id = $2 AND status = 0 AND flag = $3;

-- name: ListFriendApply :many
-- SELECt 
--     fp.apply_id, u.nickname, fp.apply_desc, fp.created_at
-- FROM friend_apply AS fp
-- JOIN users AS u ON fp.apply_id = u.id AND fp.reply_id = $1;

-- name: DeleteFriendApply :exec
-- DELETE FROM friend_apply 
-- WHERE apply_id = $1 AND reply_id = $2;

-- name: GetFriendApply :one
-- SELECT * FROM friend_apply
-- WHERE apply_id = $1 AND reply_id = $2;