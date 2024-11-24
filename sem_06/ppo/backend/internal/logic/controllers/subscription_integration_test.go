package controllers

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/containers"
	userContext "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/context"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/interfaces"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/repository/postgres"
	"github.com/google/uuid"
)

func TestIntegrationSubscription_Subscribe(t *testing.T) {
	dbContainer, db, err := containers.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = dbContainer.Terminate(context.Background())
	}()

	subRepo := postgres.NewSR(db)
	userRepo := postgres.NewUR(db)
	btRepo := postgres.NewBTR(db)
	btLog := NewBTL(userRepo, btRepo)

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
		check   func(*Subscription) error
	}{
		{
			name: "successful subscribe",
			fields: fields{
				userRepo:                userRepo,
				subscriptionRepo:        subRepo,
				balanceTransactionLogic: btLog,
				cost:                    10,
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:     uuid.MustParse("a52b8aea-d751-4933-91bb-691132e3b760"),
					Login:    "1",
					Password: "password",
					Balance:  100,
					Mail:     "1@mail.ru",
				}),
				writeID: uuid.MustParse("77dcd288-79b2-4655-9584-cc9b5329665d"),
			},
			wantErr: false,
			check: func(s *Subscription) error {
				rc, err := s.IsSubscribed(
					context.Background(),
					uuid.MustParse("a52b8aea-d751-4933-91bb-691132e3b760"),
					uuid.MustParse("77dcd288-79b2-4655-9584-cc9b5329665d"),
				)
				if err != nil {
					return fmt.Errorf("is subscribed: %w", err)
				}
				if !rc {
					return errors.New("not subscribed")
				}
				u, err := s.userRepo.GetByID(context.Background(), uuid.MustParse("a52b8aea-d751-4933-91bb-691132e3b760"))
				if err != nil {
					return fmt.Errorf("get by id: %w", err)
				}
				if u.Balance != 90 {
					return fmt.Errorf("get by id: %w", errors.New("wrong balance"))
				}
				return nil
			},
		},
		{
			name: "successful unsubscribe",
			fields: fields{
				userRepo:                userRepo,
				subscriptionRepo:        subRepo,
				balanceTransactionLogic: btLog,
				cost:                    10,
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID: uuid.MustParse("a52b8aea-d751-4933-91bb-691132e3b760"),
				}),
				writeID: uuid.MustParse("77dcd288-79b2-4655-9584-cc9b5329665d"),
			},
			wantErr: false,
			check: func(s *Subscription) error {
				rc, err := s.IsSubscribed(
					context.Background(),
					uuid.MustParse("a52b8aea-d751-4933-91bb-691132e3b761"),
					uuid.MustParse("77dcd288-79b2-4655-9584-cc9b5329665d"),
				)
				if err != nil {
					return err
				}
				if !rc {
					return nil
				}
				return errors.New("subscribed")
			},
		},
		{
			name: "unsuccessful subscribe: insufficient balance",
			fields: fields{
				userRepo:                userRepo,
				subscriptionRepo:        subRepo,
				balanceTransactionLogic: btLog,
				cost:                    10,
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:     uuid.MustParse("a52b8aea-d751-4933-91bb-691132e3b760"),
					Login:    "1",
					Password: "password",
					Balance:  1,
					Mail:     "1@mail.ru",
				}),
				writeID: uuid.MustParse("77dcd288-79b2-4655-9584-cc9b5329665d"),
			},
			wantErr: true,
			check: func(s *Subscription) error {
				rc, err := s.IsSubscribed(
					context.Background(),
					uuid.MustParse("a52b8aea-d751-4933-91bb-691132e3b760"),
					uuid.MustParse("77dcd288-79b2-4655-9584-cc9b5329665d"),
				)
				if err != nil {
					return err
				}
				if rc {
					return errors.New("subscribed")
				}
				return nil
			},
		},
		{
			name: "unsuccessful subscribe: non existing writer",
			fields: fields{
				userRepo:                userRepo,
				subscriptionRepo:        subRepo,
				balanceTransactionLogic: btLog,
				cost:                    10,
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:     uuid.MustParse("a52b8aea-d751-4933-91bb-691132e3b760"),
					Login:    "1",
					Password: "password",
					Balance:  1,
					Mail:     "1@mail.ru",
				}),
				writeID: uuid.MustParse("77dcd288-79b2-4655-9584-cc9b5329665a"),
			},
			wantErr: true,
			check: func(s *Subscription) error {
				rc, err := s.IsSubscribed(
					context.Background(),
					uuid.MustParse("a52b8aea-d751-4933-91bb-691132e3b760"),
					uuid.MustParse("77dcd288-79b2-4655-9584-cc9b5329665a"),
				)
				if err != nil {
					return err
				}
				if rc {
					return errors.New("subscribed")
				}
				return nil
			},
		},
		{
			name: "unsuccessful subscribe: non existing reader",
			fields: fields{
				userRepo:                userRepo,
				subscriptionRepo:        subRepo,
				balanceTransactionLogic: btLog,
				cost:                    10,
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:     uuid.MustParse("a52b8aea-d751-4933-91bb-691132e3b766"),
					Login:    "1",
					Password: "password",
					Balance:  1,
					Mail:     "1@mail.ru",
				}),
				writeID: uuid.MustParse("77dcd288-79b2-4655-9584-cc9b5329665d"),
			},
			wantErr: true,
			check: func(s *Subscription) error {
				rc, err := s.IsSubscribed(
					context.Background(),
					uuid.MustParse("a52b8aea-d751-4933-91bb-691132e3b766"),
					uuid.MustParse("77dcd288-79b2-4655-9584-cc9b5329665d"),
				)
				if err != nil {
					return err
				}
				if rc {
					return errors.New("subscribed")
				}
				return nil
			},
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
			if err := tt.check(s); err != nil {
				t.Errorf("Subscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
