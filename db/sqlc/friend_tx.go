package db

import (
	"context"
)

type AddFriendTxParams struct {
	UserID   int32  `json:"user_id"`
	FriendID int32  `json:"friend_id"`
	Comment  string `json:"comment"`
}

func (store *SQLStore) AddFriendTx(ctx context.Context, arg *AddFriendTxParams) (friend Friendship, err error) {
	err = store.execTx(ctx, func(q *Queries) error {
		var err error
		friend, err = q.AddFriend(ctx, &AddFriendParams{
			UserID:   arg.UserID,
			FriendID: arg.FriendID,
			Comment:  arg.Comment,
		})
		if err != nil {
			return err
		}

		friend, err = q.AddFriend(ctx, &AddFriendParams{
			UserID:   arg.FriendID,
			FriendID: arg.UserID,
			Comment:  arg.Comment,
		})
		if err != nil {
			return err
		}

		return err
	})

	return friend, err
}

type DeleteFriendTxParams struct {
	UserID   int32 `json:"user_id"`
	FriendID int32 `json:"friend_id"`
}

func (store *SQLStore) DeleteFriendTx(ctx context.Context, arg *DeleteFriendTxParams) error {
	err := store.execTx(ctx, func(q *Queries) error {
		err := q.DeleteFriend(ctx, &DeleteFriendParams{
			UserID:   arg.UserID,
			FriendID: arg.FriendID,
		})
		if err != nil {
			return err
		}

		err = q.DeleteFriend(ctx, &DeleteFriendParams{
			UserID:   arg.FriendID,
			FriendID: arg.UserID,
		})
		if err != nil {
			return err
		}

		return err
	})

	return err
}
