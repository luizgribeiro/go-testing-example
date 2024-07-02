package service_test

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	cliMocks "luizgribeiro/testing/mocks/luizgribeiro/testing/pkg/client"
	svcMocks "luizgribeiro/testing/mocks/luizgribeiro/testing/pkg/service"
	"luizgribeiro/testing/pkg/client"
	"luizgribeiro/testing/pkg/model"
	"luizgribeiro/testing/pkg/service"
	query "luizgribeiro/testing/pkg/sql"
)

var usrId = uuid.New()

var userAccount = &model.UserAccount{
	User: model.User{
		ID:    usrId,
		Name:  "user name",
		Email: "user@mail.com",
	},
	Account: model.Account{
		ID:        uuid.New(),
		Handler:   "some handler",
		CreatedAt: time.Now(),
		LastLogin: time.Now(),
		UserID:    usrId,
	},
	AccountID: uuid.New(),
}
var userQueryParams = pgx.NamedArgs{
	"id":    userAccount.User.ID,
	"name":  userAccount.Name,
	"email": userAccount.Email,
}

var accountQueryParams = pgx.NamedArgs{
	"id":         userAccount.AccountID,
	"handler":    userAccount.Handler,
	"created_at": userAccount.CreatedAt,
	"last_login": userAccount.LastLogin,
	"user_id":    userAccount.UserID,
}
var errMock = errors.New("mocked error")

var _ = Describe("AccountService", func() {
	var mockTx *cliMocks.MockDBTransactioner
	var mockBeginner *svcMocks.MocktransactionBeginner
	var accountClient *client.AccountClient
	var userClient *client.UserClient
	var userService *service.UserService
	var ctx context.Context

	BeforeEach(func() {
		mockBeginner = svcMocks.NewMocktransactionBeginner(GinkgoT())
		mockTx = cliMocks.NewMockDBTransactioner(GinkgoT())
		accountClient = client.NewAccountClient(mockTx)
		userClient = client.NewUserClient(mockTx, accountClient)
		userService = service.NewUserService(mockBeginner, userClient)
		ctx = context.Background()
	})

	Describe("CreateUserWithAccount", func() {
		var pgCommandTag pgconn.CommandTag

		Context("Creating User with account", func() {
			When("Database transaction creation fails", func() {
				It("Shoud fail", func() {
					mockBeginner.EXPECT().Begin(ctx).Return(nil, errMock)

					err := userService.CreateUserWithAccount(ctx, userAccount)

					Expect(err).To(Equal(errMock))
				})
			})

			When("Saving user on database fails", func() {
				It("Shoud fail", func() {
					pgCommandTag = pgconn.CommandTag{}
					mockTx.EXPECT().Exec(ctx, query.InsertUser, userQueryParams).Return(pgCommandTag, errMock)
					mockTx.EXPECT().Rollback(ctx).Return(nil)
					mockBeginner.EXPECT().Begin(ctx).Return(mockTx, nil)

					err := userService.CreateUserWithAccount(ctx, userAccount)

					Expect(err).To(Equal(errMock))
				})
			})

			When("Saving account on database fails", func() {
				It("Shoud fail", func() {
					pgCommandTag = pgconn.CommandTag{}
					mockTx.EXPECT().Exec(ctx, query.InsertUser, userQueryParams).Return(pgCommandTag, nil)
					mockTx.EXPECT().Exec(ctx, query.InsertAccount, accountQueryParams).Return(pgCommandTag, errMock)
					mockTx.EXPECT().Rollback(ctx).Return(nil)
					mockBeginner.EXPECT().Begin(ctx).Return(mockTx, nil)

					err := userService.CreateUserWithAccount(ctx, userAccount)

					Expect(err).To(Equal(errMock))
				})
			})

			When("Saving user and account on database succeeds", func() {
				It("Shoud succeed", func() {
					pgCommandTag = pgconn.CommandTag{}
					mockTx.EXPECT().Exec(ctx, query.InsertUser, userQueryParams).Return(pgCommandTag, nil)
					mockTx.EXPECT().Exec(ctx, query.InsertAccount, accountQueryParams).Return(pgCommandTag, nil)
					mockTx.EXPECT().Commit(ctx).Return(nil)
					mockBeginner.EXPECT().Begin(ctx).Return(mockTx, nil)

					err := userService.CreateUserWithAccount(ctx, userAccount)

					Expect(err).To(BeNil())
				})
			})
		})

	})
})
