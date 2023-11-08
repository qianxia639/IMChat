// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: group.sql

package db

import (
	"context"
)

const createGroup = `-- name: CreateGroup :one
INSERT INTO groups (
    creator_id, group_name, icon, description
) VALUES(
    $1, $2, $3, $4
)
RETURNING id, creator_id, group_name, icon, description, notice, created_at, updated_at
`

type CreateGroupParams struct {
	CreatorID   int32  `json:"creator_id"`
	GroupName   string `json:"group_name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
}

func (q *Queries) CreateGroup(ctx context.Context, arg *CreateGroupParams) (Group, error) {
	row := q.db.QueryRow(ctx, createGroup,
		arg.CreatorID,
		arg.GroupName,
		arg.Icon,
		arg.Description,
	)
	var i Group
	err := row.Scan(
		&i.ID,
		&i.CreatorID,
		&i.GroupName,
		&i.Icon,
		&i.Description,
		&i.Notice,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listGroup = `-- name: ListGroup :many
SELECT g.id AS group_id, g.group_name, g.icon, u.id AS user_id, u.username, u.nickname
FROM groups g
JOIN users u ON g.creator_id = u.id
WHERE
    u.id = $1
`

type ListGroupRow struct {
	GroupID   int32  `json:"group_id"`
	GroupName string `json:"group_name"`
	Icon      string `json:"icon"`
	UserID    int32  `json:"user_id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
}

func (q *Queries) ListGroup(ctx context.Context, id int32) ([]ListGroupRow, error) {
	rows, err := q.db.Query(ctx, listGroup, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListGroupRow
	for rows.Next() {
		var i ListGroupRow
		if err := rows.Scan(
			&i.GroupID,
			&i.GroupName,
			&i.Icon,
			&i.UserID,
			&i.Username,
			&i.Nickname,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
