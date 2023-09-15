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

-- name: ListFriendClusterApply :many
SELECT fca.apply_id, u.nickname, fca.apply_desc,  fca.flag, fca.apply_time
FROM friend_cluster_apply fca
JOIN users u ON fca.apply_id = u.id
WHERE fca.receiver_id = $1 AND fca.status = 0;

-- name: UpdateFriendClusterApply :exec
-- DELETE FROM friend_cluster_apply 
-- WHERE apply_id = $1 AND reply_id = $2;
UPDATE friend_cluster_apply
SET
    status = @status
WHERE 
    (apply_id = @apply_id AND receiver_id = @receiver_id AND status = 0)
    OR
    (apply_id = @receiver_id AND receiver_id = @apply_id AND status = 0);

-- name: GetFriendApply :one
-- SELECT * FROM friend_apply
-- WHERE apply_id = $1 AND reply_id = $2;