package db

import "context"

type AddMessageTxParams struct {
	AddMessageParams
	AfterFunc func(message Message) error
}

type AddMessageTxResult struct {
	Message Message
}

func (store *SQLStore) AddMessageTx(ctx context.Context, arg *AddMessageTxParams) (AddMessageTxResult, error) {
	var result AddMessageTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Message, err = q.AddMessage(ctx, &arg.AddMessageParams)
		if err != nil {
			return err
		}

		return arg.AfterFunc(result.Message)
	})

	return result, err
}
