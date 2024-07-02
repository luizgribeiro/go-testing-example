package service

import (
	"luizgribeiro/testing/pkg/client"
	"luizgribeiro/testing/pkg/model"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/net/context"
)

type transactionBeginner interface {
	Begin(ctx context.Context) (pgx.Tx, error)
}

type UserService struct {
	dbClient   transactionBeginner
	userClient *client.UserClient
}

func NewUserService(dbClient transactionBeginner, userClient *client.UserClient) *UserService {
	return &UserService{
		dbClient:   dbClient,
		userClient: userClient,
	}
}

type Account struct {
	ID        uuid.UUID
	createdAt time.Time
}

func (usrSvc *UserService) CreateUserWithAccount(ctx context.Context, userAcc *model.UserAccount) (err error) {
	tx, err := usrSvc.dbClient.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if tx != nil {
			if err != nil {
				tx.Rollback(ctx)
			} else {
				tx.Commit(ctx)
			}
		}

	}()

	err = usrSvc.userClient.CreateUser(ctx, tx, userAcc)
	if err != nil {
		return err
	}

	return err
}
