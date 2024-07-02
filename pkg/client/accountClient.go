package client

import (
	"context"

	"github.com/jackc/pgx/v5"
	"luizgribeiro/testing/pkg/model"
	query "luizgribeiro/testing/pkg/sql"
)

type AccountClient struct {
	dbClient Execer
}

func NewAccountClient(dbClient Execer) *AccountClient {
	return &AccountClient{
		dbClient: dbClient,
	}
}

func (ac *AccountClient) CreateAccount(ctx context.Context, tx Execer, userAcc *model.UserAccount) error {

	args := pgx.NamedArgs{
	"id":         userAcc.AccountID,
	"handler":    userAcc.Handler,
	"created_at": userAcc.CreatedAt,
	"last_login": userAcc.LastLogin,
	"user_id":    userAcc.UserID,
}

	_, err := tx.Exec(ctx, query.InsertAccount, args)
	if err != nil {
		return err
	}

	return nil
}
