-- name: ListFriendGroupApplyPending :many
-- SELECT u.id, u.nickname, fga.description, fga.apply_type, fga.created_at, g.group_name
-- FROM friend_group_applys fga
-- JOIN users u ON fga.user_id = u.id
-- JOIN groups g ON fga.user_id = g.id
-- WHERE fga.target_id = $1 AND fga.status = 1;

-- name: CreateFriendGroupApply :one
-- INSERT INTO friend_group_applys(
--     sender_id, receiver_id, apply_desc, apply_type
-- ) VALUES (
--     $1, $2, $3, $4
-- ) RETURNING *;

-- name: ExistsFriendGroupApply :one
-- SELECT COUNT(*) FROM friend_group_applys
-- WHERE
--     sender_id = $1 AND receiver_id = $2 AND status = 0 AND apply_type = $3;

-- name: UpdateFriendGroupApply :exec
-- UPDATE friend_group_applys
-- SET
--     status = @status,
--     reply_time = now()
-- WHERE 
--     (sender_id = @sender_id AND receiver_id = @receiver_id AND status = 0 AND apply_type = @apply_type)
--     OR
--     (sender_id = @receiver_id AND receiver_id = @sender_id AND status = 0 AND apply_type = @apply_type);

-- name: DeleteFriendGroupApply :exec
-- DELETE FROM friend_group_applys
-- WHERE sender_id = $1 OR receiver_id = $1;