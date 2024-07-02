package client

import (
	"context"

	"luizgribeiro/testing/pkg/model"
	query "luizgribeiro/testing/pkg/sql"

	"github.com/jackc/pgx/v5"
)


type UserClient struct {
	dbClient Execer
	accClient *AccountClient
}

func NewUserClient(dbClient Execer, accClient *AccountClient) *UserClient {
	return &UserClient{
		dbClient: dbClient,
		accClient: accClient,
	}
}

func (uc *UserClient) CreateUser(ctx context.Context, tx DBTransactioner, userAcc *model.UserAccount) error {
	args := pgx.NamedArgs{
		"id": userAcc.User.ID,
		"name": userAcc.Name,
		"email": userAcc.Email,
	}

	_, err := tx.Exec(ctx, query.InsertUser, args)
	if err != nil {
		return err
	}


	return uc.accClient.CreateAccount(ctx, tx, userAcc)
}
