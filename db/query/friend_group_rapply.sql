-- name: CreateFriendGroupApply :one
INSERT INTO friend_group_applys(
    sender_id, receiver_id, apply_desc, apply_type
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: ExistsFriendGroupApply :one
SELECT COUNT(*) FROM friend_group_applys
WHERE
    sender_id = $1 AND receiver_id = $2 AND status = 0 AND apply_type = $3;

-- name: ListFriendGroupApply :many
SELECT fca.sender_id, u.nickname, fca.apply_desc,  fca.apply_type, fca.apply_time
FROM friend_group_applys fca
JOIN users u ON fca.sender_id = u.id
WHERE fca.receiver_id = $1 AND fca.status = 0;

-- name: UpdateFriendGroupApply :exec
UPDATE friend_group_applys
SET
    status = @status,
    reply_time = now()
WHERE 
    (sender_id = @sender_id AND receiver_id = @receiver_id AND status = 0 AND apply_type = @apply_type)
    OR
    (sender_id = @receiver_id AND receiver_id = @sender_id AND status = 0 AND apply_type = @apply_type);

-- name: DeleteFriendGroupApply :exec
DELETE FROM friend_group_applys
WHERE sender_id = $1 OR receiver_id = $1;