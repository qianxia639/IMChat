package db

import (
	"context"
)

type ReplyFriendClusterApplyTxTxParams struct {
	UserID   int32  `json:"user_id"`
	FriendID int32  `json:"friend_id"`
	Status   int32  `json:"status"`
	Flag     int32  `json:"flag"`
	Note     string `json:"note"`
}

func (store *SQLStore) ReplyFriendClusterApplyTx(ctx context.Context, arg *ReplyFriendClusterApplyTxTxParams) (friend Friend, err error) {
	err = store.execTx(ctx, func(q *Queries) error {
		if arg.Status == 2 { // 拒绝
			// 更新申请表中的状态
			return q.UpdateFriendClusterApply(ctx, &UpdateFriendClusterApplyParams{
				Status:     int16(arg.Status),
				ApplyID:    arg.UserID,
				ReceiverID: arg.FriendID,
			})
		}
		var err error
		friend, err = q.AddFriend(ctx, &AddFriendParams{
			UserID:   arg.UserID,
			FriendID: arg.FriendID,
			Note:     arg.Note,
		})
		if err != nil {
			return err
		}

		friend, err = q.AddFriend(ctx, &AddFriendParams{
			UserID:   arg.FriendID,
			FriendID: arg.UserID,
			Note:     arg.Note,
		})
		if err != nil {
			return err
		}

		// 更新申请表中的状态
		err = q.UpdateFriendClusterApply(ctx, &UpdateFriendClusterApplyParams{
			Status:     int16(arg.Status),
			ApplyID:    arg.UserID,
			ReceiverID: arg.FriendID,
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
