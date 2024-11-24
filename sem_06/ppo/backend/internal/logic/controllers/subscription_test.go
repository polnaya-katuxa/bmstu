package controllers

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	userContext "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/context"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/interfaces"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/mocks"
	"github.com/gojuno/minimock/v3"
	"github.com/google/uuid"
)

func TestSubscription_Subscribe(t *testing.T) {
	mc := minimock.NewController(t)
	sub := &models.Subscription{
		UUID:     uuid.New(),
		ReaderID: uuid.New(),
		WriterID: uuid.New(),
	}
	type fields struct {
		userRepo                interfaces.IUserRepository
		subscriptionRepo        interfaces.ISubscriptionRepository
		balanceTransactionLogic interfaces.IBalanceTransactionLogic
		cost                    int
	}
	type args struct {
		ctx     context.Context
		writeID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "successful subscribe",
			fields: fields{
				userRepo:                nil,
				subscriptionRepo:        mocks.NewSubscriptionRepositoryMock(mc).GetMock.Return(nil, nil).CreateMock.Return(nil),
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).DecreaseMock.Return(nil).IncreaseSubMock.Return(nil),
				cost:                    0,
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
					IsAdmin:     true,
				}),
				writeID: uuid.New(),
			},
			wantErr: false,
		},
		{
			name: "successful subscribe: unsubscribe",
			fields: fields{
				userRepo:                nil,
				subscriptionRepo:        mocks.NewSubscriptionRepositoryMock(mc).GetMock.Return(sub, nil).CreateMock.Return(nil).DeleteMock.Return(nil),
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).DecreaseMock.Return(nil),
				cost:                    0,
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
					IsAdmin:     true,
				}),
				writeID: uuid.New(),
			},
			wantErr: false,
		},
		{
			name: "unsuccessful subscribe: empty context",
			fields: fields{
				userRepo:                nil,
				subscriptionRepo:        mocks.NewSubscriptionRepositoryMock(mc).GetMock.Return(nil, nil).CreateMock.Return(nil),
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).DecreaseMock.Return(nil),
				cost:                    0,
			},
			args: args{
				ctx:     context.Background(),
				writeID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "unsuccessful subscribe: get",
			fields: fields{
				userRepo:                nil,
				subscriptionRepo:        mocks.NewSubscriptionRepositoryMock(mc).GetMock.Return(nil, errors.New("error")).CreateMock.Return(nil),
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).DecreaseMock.Return(nil),
				cost:                    0,
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
					IsAdmin:     true,
				}),
				writeID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "unsuccessful subscribe: unsubscribe delete",
			fields: fields{
				userRepo:                nil,
				subscriptionRepo:        mocks.NewSubscriptionRepositoryMock(mc).GetMock.Return(sub, nil).CreateMock.Return(nil).DeleteMock.Return(errors.New("error")),
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).DecreaseMock.Return(nil),
				cost:                    0,
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
					IsAdmin:     true,
				}),
				writeID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "unsuccessful subscribe: create",
			fields: fields{
				userRepo:                nil,
				subscriptionRepo:        mocks.NewSubscriptionRepositoryMock(mc).GetMock.Return(nil, nil).CreateMock.Return(errors.New("error")),
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).DecreaseMock.Return(nil).IncreaseSubMock.Return(nil),
				cost:                    0,
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
					IsAdmin:     true,
				}),
				writeID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "unsuccessful subscribe: balance transaction decrease",
			fields: fields{
				userRepo:                nil,
				subscriptionRepo:        mocks.NewSubscriptionRepositoryMock(mc).GetMock.Return(nil, nil).CreateMock.Return(nil),
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).DecreaseMock.Return(errors.New("error")),
				cost:                    0,
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
					IsAdmin:     true,
				}),
				writeID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "unsuccessful subscribe: balance transaction increase",
			fields: fields{
				userRepo:                nil,
				subscriptionRepo:        mocks.NewSubscriptionRepositoryMock(mc).GetMock.Return(nil, nil).CreateMock.Return(nil),
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).DecreaseMock.Return(nil).IncreaseSubMock.Return(errors.New("error")),
				cost:                    0,
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
					IsAdmin:     true,
				}),
				writeID: uuid.New(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Subscription{
				userRepo:                tt.fields.userRepo,
				subscriptionRepo:        tt.fields.subscriptionRepo,
				balanceTransactionLogic: tt.fields.balanceTransactionLogic,
				cost:                    tt.fields.cost,
			}
			if err := s.Subscribe(tt.args.ctx, tt.args.writeID); (err != nil) != tt.wantErr {
				t.Errorf("Subscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSubscription_GetSubscriptions(t *testing.T) {
	mc := minimock.NewController(t)
	subs := []*models.Subscription{{
		UUID:     uuid.New(),
		ReaderID: uuid.New(),
		WriterID: uuid.New(),
	}}
	user := &models.User{
		UUID:        uuid.New(),
		Login:       "uehfkjsyg",
		Balance:     0,
		Mail:        "mail@gmail.com",
		EnterTime:   time.Now().UTC(),
		Picture:     "jkjlkjlkj",
		Description: "fgfdydfj",
	}
	users := []*models.User{{
		UUID:        user.UUID,
		Login:       "uehfkjsyg",
		Balance:     0,
		Mail:        "mail@gmail.com",
		EnterTime:   user.EnterTime,
		Picture:     "jkjlkjlkj",
		Description: "fgfdydfj",
	}}
	type fields struct {
		userRepo                interfaces.IUserRepository
		subscriptionRepo        interfaces.ISubscriptionRepository
		balanceTransactionLogic interfaces.IBalanceTransactionLogic
		cost                    int
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.User
		wantErr bool
	}{
		{
			name: "successful get all",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetByIDMock.Return(user, nil),
				subscriptionRepo:        mocks.NewSubscriptionRepositoryMock(mc).GetAllMock.Return(subs, nil),
				balanceTransactionLogic: nil,
				cost:                    0,
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
					IsAdmin:     true,
				}),
			},
			want:    users,
			wantErr: false,
		},
		{
			name: "unsuccessful get all: empty context",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetByIDMock.Return(user, nil),
				subscriptionRepo:        mocks.NewSubscriptionRepositoryMock(mc).GetAllMock.Return(subs, nil),
				balanceTransactionLogic: nil,
				cost:                    0,
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful get all: get all",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetByIDMock.Return(user, nil),
				subscriptionRepo:        mocks.NewSubscriptionRepositoryMock(mc).GetAllMock.Return(subs, errors.New("error")),
				balanceTransactionLogic: nil,
				cost:                    0,
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
					IsAdmin:     true,
				}),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful get all: get by id",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetByIDMock.Return(user, errors.New("error")),
				subscriptionRepo:        mocks.NewSubscriptionRepositoryMock(mc).GetAllMock.Return(subs, nil),
				balanceTransactionLogic: nil,
				cost:                    0,
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
					IsAdmin:     true,
				}),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Subscription{
				userRepo:                tt.fields.userRepo,
				subscriptionRepo:        tt.fields.subscriptionRepo,
				balanceTransactionLogic: tt.fields.balanceTransactionLogic,
				cost:                    tt.fields.cost,
			}
			got, err := s.GetSubscriptions(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSubscriptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSubscriptions() got = %v, want %v", got, tt.want)
			}
		})
	}
}
