package db

import "context"

type CreateGroupTxParams struct {
	Nickname string `json:"nickname"`
	CreateGroupParams
}

func (store *SQLStore) CreateGroupTx(ctx context.Context, arg *CreateGroupTxParams) error {
	err := store.execTx(ctx, func(q *Queries) error {
		group, err := q.CreateGroup(ctx, &CreateGroupParams{
			CreatorID:   arg.CreatorID,
			GroupName:   arg.GroupName,
			Icon:        arg.Icon,
			Description: arg.Description,
		})
		if err != nil {
			return err
		}

		_, err = q.AddGroupMember(ctx, &AddGroupMemberParams{
			UserID:   group.CreatorID,
			GroupID:  group.ID,
			Nickname: arg.Nickname,
			Role:     1,
		})
		if err != nil {
			return err
		}

		return nil
	})
	return err
}
