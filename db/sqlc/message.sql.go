// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: message.sql

package db

import (
	"context"
)

const addMessage = `-- name: AddMessage :one
INSERT INTO messages (
    sender_id, receiver_id, content
) VALUES (
    $1, $2, $3
)
RETURNING id, sender_id, receiver_id, content, type, send_time, receive_time
`

type AddMessageParams struct {
	SenderID   int32  `json:"sender_id"`
	ReceiverID int32  `json:"receiver_id"`
	Content    string `json:"content"`
}

func (q *Queries) AddMessage(ctx context.Context, arg *AddMessageParams) (Message, error) {
	row := q.db.QueryRow(ctx, addMessage, arg.SenderID, arg.ReceiverID, arg.Content)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.SenderID,
		&i.ReceiverID,
		&i.Content,
		&i.Type,
		&i.SendTime,
		&i.ReceiveTime,
	)
	return i, err
}
