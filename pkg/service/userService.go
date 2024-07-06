package service

import (
	"errors"
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

	now := time.Now()

	age := now.Year() - userAcc.BirthDate.Year()

	if now.Month() < userAcc.BirthDate.Month() || (userAcc.BirthDate.Month() == now.Month() && now.Day() < userAcc.BirthDate.Day()) {
		age--
	}

	if age < 12 {
		return errors.New("Too young")
	}

	err = usrSvc.userClient.CreateUser(ctx, tx, userAcc)
	if err != nil {
		return err
	}

	return err
}
