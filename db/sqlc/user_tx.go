package db

import "context"

func (store *SQLStore) DeleteUserTx(ctx context.Context, userId int32) error {

	err := store.execTx(ctx, func(q *Queries) error {
		// 删除用户
		err := q.DeleteUser(ctx, userId)
		if err != nil {
			return err
		}
		// 删除审核表中记录
		// err = q.DeleteFriendGroupApply(ctx, userId)
		// if err != nil {
		// 	return err
		// }

		return err
	})
	return err
}
