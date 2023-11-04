// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: friendship.sql

package db

import (
	"context"
)

const addFriend = `-- name: AddFriend :one
INSERT INTO friendships (user_id, friend_id, comment)
VALUEs ($1, $2, $3) 
RETURNING user_id, friend_id, comment, created_at
`

type AddFriendParams struct {
	UserID   int32  `json:"user_id"`
	FriendID int32  `json:"friend_id"`
	Comment  string `json:"comment"`
}

func (q *Queries) AddFriend(ctx context.Context, arg *AddFriendParams) (Friendship, error) {
	row := q.db.QueryRow(ctx, addFriend, arg.UserID, arg.FriendID, arg.Comment)
	var i Friendship
	err := row.Scan(
		&i.UserID,
		&i.FriendID,
		&i.Comment,
		&i.CreatedAt,
	)
	return i, err
}

const deleteFriend = `-- name: DeleteFriend :exec
DELETE FROM friendships
WHERE
    user_id = $1 AND friend_id = $2
`

type DeleteFriendParams struct {
	UserID   int32 `json:"user_id"`
	FriendID int32 `json:"friend_id"`
}

func (q *Queries) DeleteFriend(ctx context.Context, arg *DeleteFriendParams) error {
	_, err := q.db.Exec(ctx, deleteFriend, arg.UserID, arg.FriendID)
	return err
}

const getFriend = `-- name: GetFriend :one
SELECT user_id, friend_id, comment, created_at FROM friendships
WHERE 
    user_id = $1 AND friend_id = $2
`

type GetFriendParams struct {
	UserID   int32 `json:"user_id"`
	FriendID int32 `json:"friend_id"`
}

func (q *Queries) GetFriend(ctx context.Context, arg *GetFriendParams) (Friendship, error) {
	row := q.db.QueryRow(ctx, getFriend, arg.UserID, arg.FriendID)
	var i Friendship
	err := row.Scan(
		&i.UserID,
		&i.FriendID,
		&i.Comment,
		&i.CreatedAt,
	)
	return i, err
}

const listFriends = `-- name: ListFriends :many
SELECT f.friend_id, f.comment, u.profile_picture_url FROM friendships AS f
JOIN users AS u ON f.user_id = u.id
WHERE u.id = $1
`

type ListFriendsRow struct {
	FriendID          int32  `json:"friend_id"`
	Comment           string `json:"comment"`
	ProfilePictureUrl string `json:"profile_picture_url"`
}

func (q *Queries) ListFriends(ctx context.Context, id int32) ([]ListFriendsRow, error) {
	rows, err := q.db.Query(ctx, listFriends, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListFriendsRow
	for rows.Next() {
		var i ListFriendsRow
		if err := rows.Scan(&i.FriendID, &i.Comment, &i.ProfilePictureUrl); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateFriendComment = `-- name: UpdateFriendComment :one
UPDATE friendships
SET
    comment = $1
WHERE
    user_id = $2 AND friend_id = $3
RETURNING user_id, friend_id, comment, created_at
`

type UpdateFriendCommentParams struct {
	Comment  string `json:"comment"`
	UserID   int32  `json:"user_id"`
	FriendID int32  `json:"friend_id"`
}

func (q *Queries) UpdateFriendComment(ctx context.Context, arg *UpdateFriendCommentParams) (Friendship, error) {
	row := q.db.QueryRow(ctx, updateFriendComment, arg.Comment, arg.UserID, arg.FriendID)
	var i Friendship
	err := row.Scan(
		&i.UserID,
		&i.FriendID,
		&i.Comment,
		&i.CreatedAt,
	)
	return i, err
}
