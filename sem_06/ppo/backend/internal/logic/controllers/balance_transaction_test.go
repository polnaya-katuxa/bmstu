package controllers

import (
	"context"
	"errors"
	"testing"
	"time"

	userContext "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/context"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/interfaces"
	errors2 "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/errors"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/mocks"
	"github.com/gojuno/minimock/v3"
	"github.com/google/uuid"
)

func TestBalanceTransaction_Increase(t *testing.T) {
	mc := minimock.NewController(t)

	type fields struct {
		userRepo               interfaces.IUserRepository
		balanceTransactionRepo interfaces.IBalanceTransactionRepository
	}
	type args struct {
		ctx    context.Context
		value  int
		reason string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "successful balance increase",
			fields: fields{
				userRepo:               mocks.NewIUserRepositoryMock(mc).UpdateMock.Return(nil),
				balanceTransactionRepo: mocks.NewBalanceTransactionRepositoryMock(mc).CreateMock.Return(nil),
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:        uuid.New(),
					Login:       "uehfkjs",
					Balance:     0,
					Mail:        "mail@gmail.com",
					EnterTime:   time.Now().UTC(),
					Picture:     "jkjlkjlkj",
					Description: "fgfdydfj",
				}),
				value:  10,
				reason: "daily bonus",
			},
			wantErr: false,
		},
		{
			name: "unsuccessful update balance increase ",
			fields: fields{
				userRepo:               mocks.NewIUserRepositoryMock(mc).UpdateMock.Return(errors.New("error")),
				balanceTransactionRepo: mocks.NewBalanceTransactionRepositoryMock(mc).CreateMock.Return(nil),
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:        uuid.New(),
					Login:       "uehfkjs",
					Balance:     0,
					Mail:        "mail@gmail.com",
					EnterTime:   time.Now().UTC(),
					Picture:     "jkjlkjlkj",
					Description: "fgfdydfj",
				}),
				value:  10,
				reason: "daily bonus",
			},
			wantErr: true,
		},
		{
			name: "unsuccessful create balance increase",
			fields: fields{
				userRepo:               mocks.NewIUserRepositoryMock(mc).UpdateMock.Return(nil),
				balanceTransactionRepo: mocks.NewBalanceTransactionRepositoryMock(mc).CreateMock.Return(errors.New("error")),
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:        uuid.New(),
					Login:       "uehfkjs",
					Balance:     0,
					Mail:        "mail@gmail.com",
					EnterTime:   time.Now().UTC(),
					Picture:     "jkjlkjlkj",
					Description: "fgfdydfj",
				}),
				value:  10,
				reason: "daily bonus",
			},
			wantErr: true,
		},
		{
			name: "unsuccessful balance increase: empty context",
			fields: fields{
				userRepo:               mocks.NewIUserRepositoryMock(mc).UpdateMock.Return(nil),
				balanceTransactionRepo: mocks.NewBalanceTransactionRepositoryMock(mc).CreateMock.Return(nil),
			},
			args: args{
				ctx:    context.Background(),
				value:  10,
				reason: "daily bonus",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := &BalanceTransaction{
				userRepo:               tt.fields.userRepo,
				balanceTransactionRepo: tt.fields.balanceTransactionRepo,
			}
			if err := bt.Increase(tt.args.ctx, tt.args.value, tt.args.reason); (err != nil) != tt.wantErr {
				t.Errorf("Increase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBalanceTransaction_Decrease(t *testing.T) {
	mc := minimock.NewController(t)

	type fields struct {
		userRepo               interfaces.IUserRepository
		balanceTransactionRepo interfaces.IBalanceTransactionRepository
	}
	type args struct {
		ctx    context.Context
		value  int
		reason string
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		wantErr         bool
		wantConcreteErr error
	}{
		{
			name: "successful balance decrease",
			fields: fields{
				userRepo:               mocks.NewIUserRepositoryMock(mc).UpdateMock.Return(nil),
				balanceTransactionRepo: mocks.NewBalanceTransactionRepositoryMock(mc).CreateMock.Return(nil),
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:        uuid.New(),
					Login:       "uehfkjs",
					Balance:     10,
					Mail:        "mail@gmail.com",
					EnterTime:   time.Now().UTC(),
					Picture:     "jkjlkjlkj",
					Description: "fgfdydfj",
				}),
				value:  10,
				reason: "subscription",
			},
			wantErr: false,
		},
		{
			name: "unsuccessful update balance decrease ",
			fields: fields{
				userRepo:               mocks.NewIUserRepositoryMock(mc).UpdateMock.Return(errors.New("error")),
				balanceTransactionRepo: mocks.NewBalanceTransactionRepositoryMock(mc).CreateMock.Return(nil),
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:        uuid.New(),
					Login:       "uehfkjs",
					Balance:     10,
					Mail:        "mail@gmail.com",
					EnterTime:   time.Now().UTC(),
					Picture:     "jkjlkjlkj",
					Description: "fgfdydfj",
				}),
				value:  10,
				reason: "subscription",
			},
			wantErr: true,
		},
		{
			name: "unsuccessful create balance decrease ",
			fields: fields{
				userRepo:               mocks.NewIUserRepositoryMock(mc).UpdateMock.Return(nil),
				balanceTransactionRepo: mocks.NewBalanceTransactionRepositoryMock(mc).CreateMock.Return(errors.New("error")),
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:        uuid.New(),
					Login:       "uehfkjs",
					Balance:     10,
					Mail:        "mail@gmail.com",
					EnterTime:   time.Now().UTC(),
					Picture:     "jkjlkjlkj",
					Description: "fgfdydfj",
				}),
				value:  10,
				reason: "subscription",
			},
			wantErr: true,
		},
		{
			name: "unsuccessful balance decrease: not enough",
			fields: fields{
				userRepo:               mocks.NewIUserRepositoryMock(mc).UpdateMock.Return(nil),
				balanceTransactionRepo: mocks.NewBalanceTransactionRepositoryMock(mc).CreateMock.Return(nil),
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:        uuid.New(),
					Login:       "uehfkjs",
					Balance:     0,
					Mail:        "mail@gmail.com",
					EnterTime:   time.Now().UTC(),
					Picture:     "jkjlkjlkj",
					Description: "fgfdydfj",
				}),
				value:  10,
				reason: "subscription",
			},
			wantErr: true,
			wantConcreteErr: &errors2.InsufficientBalanceError{
				Want: 10,
				Got:  0,
			},
		},
		{
			name: "unsuccessful balance decrease: empty context",
			fields: fields{
				userRepo:               mocks.NewIUserRepositoryMock(mc).UpdateMock.Return(nil),
				balanceTransactionRepo: mocks.NewBalanceTransactionRepositoryMock(mc).CreateMock.Return(nil),
			},
			args: args{
				ctx:    context.Background(),
				value:  10,
				reason: "daily bonus",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := &BalanceTransaction{
				userRepo:               tt.fields.userRepo,
				balanceTransactionRepo: tt.fields.balanceTransactionRepo,
			}
			err := bt.Decrease(tt.args.ctx, tt.args.value, tt.args.reason)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrease() error = %v, wantErr %v", err, tt.wantErr)
			}
			if (tt.wantConcreteErr != nil) && (!errors.Is(err, tt.wantConcreteErr)) {
				t.Errorf("Decrease() error = %v, wantConcreteErr %v", err, tt.wantConcreteErr)
			}
		})
	}
}

func TestBalanceTransaction_IncreaseSub(t *testing.T) {
	mc := minimock.NewController(t)

	id := uuid.New()
	user := &models.User{
		UUID:        id,
		Login:       "uehfkjsyg",
		Balance:     0,
		Mail:        "mail@gmail.com",
		EnterTime:   time.Now().UTC(),
		Picture:     "jkjlkjlkj",
		Description: "fgfdydfj",
	}

	type fields struct {
		userRepo               interfaces.IUserRepository
		balanceTransactionRepo interfaces.IBalanceTransactionRepository
	}
	type args struct {
		ctx    context.Context
		value  int
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "successful balance increase",
			fields: fields{
				userRepo:               mocks.NewIUserRepositoryMock(mc).UpdateMock.Return(nil).GetByIDMock.Return(user, nil),
				balanceTransactionRepo: mocks.NewBalanceTransactionRepositoryMock(mc).CreateMock.Return(nil),
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:        uuid.New(),
					Login:       "uehfkjs",
					Balance:     0,
					Mail:        "mail@gmail.com",
					EnterTime:   time.Now().UTC(),
					Picture:     "jkjlkjlkj",
					Description: "fgfdydfj",
				}),
				value:  10,
				userID: id,
			},
			wantErr: false,
		},
		{
			name: "unsuccessful get by id balance increase",
			fields: fields{
				userRepo:               mocks.NewIUserRepositoryMock(mc).UpdateMock.Return(nil).GetByIDMock.Return(nil, errors.New("error")),
				balanceTransactionRepo: mocks.NewBalanceTransactionRepositoryMock(mc).CreateMock.Return(nil),
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:        uuid.New(),
					Login:       "uehfkjs",
					Balance:     0,
					Mail:        "mail@gmail.com",
					EnterTime:   time.Now().UTC(),
					Picture:     "jkjlkjlkj",
					Description: "fgfdydfj",
				}),
				value:  10,
				userID: id,
			},
			wantErr: true,
		},
		{
			name: "unsuccessful update balance increase",
			fields: fields{
				userRepo:               mocks.NewIUserRepositoryMock(mc).UpdateMock.Return(errors.New("error")).GetByIDMock.Return(user, nil),
				balanceTransactionRepo: mocks.NewBalanceTransactionRepositoryMock(mc).CreateMock.Return(nil),
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:        uuid.New(),
					Login:       "uehfkjs",
					Balance:     0,
					Mail:        "mail@gmail.com",
					EnterTime:   time.Now().UTC(),
					Picture:     "jkjlkjlkj",
					Description: "fgfdydfj",
				}),
				value:  10,
				userID: id,
			},
			wantErr: true,
		},
		{
			name: "unsuccessful create balance increase",
			fields: fields{
				userRepo:               mocks.NewIUserRepositoryMock(mc).UpdateMock.Return(nil).GetByIDMock.Return(user, nil),
				balanceTransactionRepo: mocks.NewBalanceTransactionRepositoryMock(mc).CreateMock.Return(errors.New("error")),
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:        uuid.New(),
					Login:       "uehfkjs",
					Balance:     0,
					Mail:        "mail@gmail.com",
					EnterTime:   time.Now().UTC(),
					Picture:     "jkjlkjlkj",
					Description: "fgfdydfj",
				}),
				value:  10,
				userID: id,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := &BalanceTransaction{
				userRepo:               tt.fields.userRepo,
				balanceTransactionRepo: tt.fields.balanceTransactionRepo,
			}
			if err := bt.IncreaseSub(tt.args.ctx, tt.args.value, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("IncreaseSub() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
