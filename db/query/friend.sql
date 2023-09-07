-- name: AddFriend :one
INSERT INTO friends (user_id, friend_id, note)
VALUEs ($1, $2, $3) 
RETURNING *;

-- name: GetFriend :one
SELECT * FROM friends
WHERE 
    user_id = $1 AND friend_id = $2;

-- name: UpdateFriendNote :one
UPDATE friends
SET
    note = @note
WHERE
    user_id = @user_id AND friend_id = @friend_id
RETURNING *;

-- name: DeleteFriend :exec
DELETE FROM friends
WHERE
    user_id = $1 AND friend_id = $2;

-- name: ListFriends :many
SELECT f.friend_id, f.note, u.avatar FROM friends AS f
JOIN users AS u ON f.user_id = u.id
WHERE u.id = $1;