package controllers

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	userContext "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/context"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/interfaces"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/mocks"
	"github.com/gojuno/minimock/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func TestProfile_AuthByToken(t *testing.T) {
	mc := minimock.NewController(t)
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ID:        uuid.New().String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("secretkey"))
	if err != nil {
		fmt.Println(err)
	}

	claims1 := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ID:        "kejhrfbe",
	}

	token1 := jwt.NewWithClaims(jwt.SigningMethodHS256, claims1)
	ss1, err := token1.SignedString([]byte("secretkey"))
	if err != nil {
		fmt.Println(err)
	}

	user := &models.User{
		UUID:        uuid.New(),
		Login:       "uehfkjs",
		Balance:     0,
		Mail:        "mail@gmail.com",
		EnterTime:   time.Now().UTC(),
		Picture:     "jkjlkjlkj",
		Description: "fgfdydfj",
	}
	type fields struct {
		userRepo                interfaces.IUserRepository
		subscriptionRepo        interfaces.ISubscriptionRepository
		balanceTransactionLogic interfaces.IBalanceTransactionLogic
		postLogic               interfaces.IPostLogic
		dailyBonus              int
		tokenExpiration         time.Duration
		secretKey               string
	}
	type args struct {
		ctx   context.Context
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "successful auth",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetByIDMock.Return(user, nil),
				subscriptionRepo:        nil,
				balanceTransactionLogic: nil,
				postLogic:               nil,
				dailyBonus:              10,
				tokenExpiration:         time.Hour,
				secretKey:               "secretkey",
			},
			args: args{
				ctx:   context.Background(),
				token: ss,
			},
			want:    user,
			wantErr: false,
		},
		{
			name: "unsuccessful auth: parse",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetByIDMock.Return(user, nil),
				subscriptionRepo:        nil,
				balanceTransactionLogic: nil,
				postLogic:               nil,
				dailyBonus:              10,
				tokenExpiration:         time.Hour,
				secretKey:               "secretkey",
			},
			args: args{
				ctx:   context.Background(),
				token: "ejrfnlekrfl",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful auth: id",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetByIDMock.Return(user, nil),
				subscriptionRepo:        nil,
				balanceTransactionLogic: nil,
				postLogic:               nil,
				dailyBonus:              10,
				tokenExpiration:         time.Hour,
				secretKey:               "secretkey",
			},
			args: args{
				ctx:   context.Background(),
				token: ss1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful auth: id",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetByIDMock.Return(user, errors.New("error")),
				subscriptionRepo:        nil,
				balanceTransactionLogic: nil,
				postLogic:               nil,
				dailyBonus:              10,
				tokenExpiration:         time.Hour,
				secretKey:               "secretkey",
			},
			args: args{
				ctx:   context.Background(),
				token: ss,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Profile{
				userRepo:                tt.fields.userRepo,
				subscriptionRepo:        tt.fields.subscriptionRepo,
				balanceTransactionLogic: tt.fields.balanceTransactionLogic,
				postLogic:               tt.fields.postLogic,
				dailyBonus:              tt.fields.dailyBonus,
				tokenExpiration:         tt.fields.tokenExpiration,
				secretKey:               tt.fields.secretKey,
			}
			got, err := p.AuthByToken(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthByToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthByToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProfile_Register(t *testing.T) {
	mc := minimock.NewController(t)
	user := &models.User{
		UUID:        uuid.New(),
		Login:       "uehfkjsyg",
		Balance:     0,
		Mail:        "mail@gmail.com",
		EnterTime:   time.Now().UTC(),
		Picture:     "jkjlkjlkj",
		Description: "fgfdydfj",
	}
	user2 := &models.User{
		UUID:        uuid.New(),
		Login:       "g",
		Balance:     0,
		Mail:        "mail@gmail.com",
		EnterTime:   time.Now().UTC(),
		Picture:     "jkjlkjlkj",
		Description: "fgfdydfj",
	}
	user3 := &models.User{
		UUID:        uuid.New(),
		Login:       "ggggggggg",
		Balance:     0,
		Mail:        "mailacom",
		EnterTime:   time.Now().UTC(),
		Picture:     "jkjlkjlkj",
		Description: "fgfdydfj",
	}
	type fields struct {
		userRepo                interfaces.IUserRepository
		subscriptionRepo        interfaces.ISubscriptionRepository
		balanceTransactionLogic interfaces.IBalanceTransactionLogic
		postLogic               interfaces.IPostLogic
		dailyBonus              int
		tokenExpiration         time.Duration
		secretKey               string
	}
	type args struct {
		ctx      context.Context
		user     *models.User
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "successful register",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetByLoginMock.Return(nil, nil).CreateMock.Return(user, nil),
				subscriptionRepo:        nil,
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).IncreaseMock.Return(nil),
				postLogic:               nil,
				dailyBonus:              0,
				tokenExpiration:         time.Hour,
				secretKey:               "secretkey",
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
				user:     user,
				password: "ftioypue",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "unsuccessful register: get by login",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetByLoginMock.Return(nil, errors.New("error")).CreateMock.Return(user, nil),
				subscriptionRepo:        nil,
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).IncreaseMock.Return(nil),
				postLogic:               nil,
				dailyBonus:              0,
				tokenExpiration:         time.Hour,
				secretKey:               "secretkey",
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
				user:     user,
				password: "ftioypue",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "unsuccessful register: already exists",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetByLoginMock.Return(user, nil).CreateMock.Return(user, nil),
				subscriptionRepo:        nil,
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).IncreaseMock.Return(nil),
				postLogic:               nil,
				dailyBonus:              0,
				tokenExpiration:         time.Hour,
				secretKey:               "secretkey",
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
				user:     user,
				password: "ftioypue",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "unsuccessful register: login",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetByLoginMock.Return(nil, nil).CreateMock.Return(user, nil),
				subscriptionRepo:        nil,
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).IncreaseMock.Return(nil),
				postLogic:               nil,
				dailyBonus:              0,
				tokenExpiration:         time.Hour,
				secretKey:               "secretkey",
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
				user:     user2,
				password: "ftioypue",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "unsuccessful register: password",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetByLoginMock.Return(nil, nil).CreateMock.Return(user, nil),
				subscriptionRepo:        nil,
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).IncreaseMock.Return(nil),
				postLogic:               nil,
				dailyBonus:              0,
				tokenExpiration:         time.Hour,
				secretKey:               "secretkey",
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:        uuid.New(),
					Login:       "uehfkjjs",
					Balance:     0,
					Mail:        "mail@gmail.com",
					EnterTime:   time.Now().UTC(),
					Picture:     "jkjlkjlkj",
					Description: "fgfdydfj",
				}),
				user:     user,
				password: "ft",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "unsuccessful register: email",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetByLoginMock.Return(nil, nil).CreateMock.Return(user, nil),
				subscriptionRepo:        nil,
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).IncreaseMock.Return(nil),
				postLogic:               nil,
				dailyBonus:              0,
				tokenExpiration:         time.Hour,
				secretKey:               "secretkey",
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
				user:     user3,
				password: "fttuytuty",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "unsuccessful register: balance transaction",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetByLoginMock.Return(nil, nil).CreateMock.Return(user, nil),
				subscriptionRepo:        nil,
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).IncreaseMock.Return(errors.New("error")),
				postLogic:               nil,
				dailyBonus:              0,
				tokenExpiration:         time.Hour,
				secretKey:               "secretkey",
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
				user:     user,
				password: "fttuytuty",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "unsuccessful register: create",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetByLoginMock.Return(nil, nil).CreateMock.Return(user, errors.New("error")),
				subscriptionRepo:        nil,
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).IncreaseMock.Return(nil),
				postLogic:               nil,
				dailyBonus:              0,
				tokenExpiration:         time.Hour,
				secretKey:               "secretkey",
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
				user:     user,
				password: "fttuytuty",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Profile{
				userRepo:                tt.fields.userRepo,
				subscriptionRepo:        tt.fields.subscriptionRepo,
				balanceTransactionLogic: tt.fields.balanceTransactionLogic,
				postLogic:               tt.fields.postLogic,
				dailyBonus:              tt.fields.dailyBonus,
				tokenExpiration:         tt.fields.tokenExpiration,
				secretKey:               tt.fields.secretKey,
			}
			_, err := p.Register(tt.args.ctx, tt.args.user, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if got != tt.want {
			//	t.Errorf("Register() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestProfile_Enter(t *testing.T) {
	mc := minimock.NewController(t)
	password := "password"
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &models.User{
		UUID:        uuid.New(),
		Login:       "uehfkjsyg",
		Password:    string(hash),
		Balance:     0,
		Mail:        "mail@gmail.com",
		EnterTime:   time.Now().UTC(),
		Picture:     "jkjlkjlkj",
		Description: "fgfdydfj",
	}
	user2 := &models.User{
		UUID:        uuid.New(),
		Login:       "uehfkjsyg",
		Password:    string(hash),
		Balance:     0,
		Mail:        "mail@gmail.com",
		EnterTime:   time.Now().UTC().Add(-time.Hour * 25),
		Picture:     "jkjlkjlkj",
		Description: "fgfdydfj",
	}
	user3 := &models.User{
		UUID:        uuid.New(),
		Login:       "uehfkjsyg",
		Password:    string(hash),
		Balance:     0,
		Mail:        "mail@gmail.com",
		EnterTime:   time.Now().UTC().Add(-time.Hour * 25),
		Picture:     "jkjlkjlkj",
		Description: "fgfdydfj",
	}
	type fields struct {
		userRepo                interfaces.IUserRepository
		subscriptionRepo        interfaces.ISubscriptionRepository
		balanceTransactionLogic interfaces.IBalanceTransactionLogic
		postLogic               interfaces.IPostLogic
		dailyBonus              int
		tokenExpiration         time.Duration
		secretKey               string
	}
	type args struct {
		ctx      context.Context
		login    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "successful enter",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetByLoginMock.Return(user, nil).UpdateMock.Return(nil),
				subscriptionRepo:        nil,
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).IncreaseMock.Return(nil),
				postLogic:               mocks.NewIPostLogicMock(mc).PopularityCheckMock.Return(nil),
				dailyBonus:              5,
				tokenExpiration:         time.Hour,
				secretKey:               "secretkey",
			},
			args: args{
				ctx:      context.Background(),
				login:    "uehfkjsyg",
				password: password,
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "unsuccessful enter: wrong password",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetByLoginMock.Return(user, nil).UpdateMock.Return(nil),
				subscriptionRepo:        nil,
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).IncreaseMock.Return(nil),
				postLogic:               mocks.NewIPostLogicMock(mc).PopularityCheckMock.Return(nil),
				dailyBonus:              5,
				tokenExpiration:         time.Hour,
				secretKey:               "secretkey",
			},
			args: args{
				ctx:      context.Background(),
				login:    "uehfkjsyg",
				password: "r;glazyft",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "successful enter: with bonus",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetByLoginMock.Return(user2, nil).UpdateMock.Return(nil),
				subscriptionRepo:        nil,
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).IncreaseMock.Return(nil),
				postLogic:               mocks.NewIPostLogicMock(mc).PopularityCheckMock.Return(nil),
				dailyBonus:              5,
				tokenExpiration:         time.Hour,
				secretKey:               "secretkey",
			},
			args: args{
				ctx:      context.Background(),
				login:    "uehfkjsyg",
				password: password,
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "unsuccessful enter: with bonus",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetByLoginMock.Return(user3, nil).UpdateMock.Return(nil),
				subscriptionRepo:        nil,
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).IncreaseMock.Return(errors.New("error")),
				postLogic:               mocks.NewIPostLogicMock(mc).PopularityCheckMock.Return(nil),
				dailyBonus:              5,
				tokenExpiration:         time.Hour,
				secretKey:               "secretkey",
			},
			args: args{
				ctx:      context.Background(),
				login:    "uehfkjsyg",
				password: password,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "unsuccessful enter: get by login",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetByLoginMock.Return(user, errors.New("error")).UpdateMock.Return(nil),
				subscriptionRepo:        nil,
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).IncreaseMock.Return(nil),
				postLogic:               mocks.NewIPostLogicMock(mc).PopularityCheckMock.Return(nil),
				dailyBonus:              5,
				tokenExpiration:         time.Hour,
				secretKey:               "secretkey",
			},
			args: args{
				ctx:      context.Background(),
				login:    "uehfkjsyg",
				password: password,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "unsuccessful enter: popularity check",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetByLoginMock.Return(user, nil).UpdateMock.Return(nil),
				subscriptionRepo:        nil,
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).IncreaseMock.Return(nil),
				postLogic:               mocks.NewIPostLogicMock(mc).PopularityCheckMock.Return(errors.New("error")),
				dailyBonus:              5,
				tokenExpiration:         time.Hour,
				secretKey:               "secretkey",
			},
			args: args{
				ctx:      context.Background(),
				login:    "uehfkjsyg",
				password: password,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "unsuccessful enter: update",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetByLoginMock.Return(user, nil).UpdateMock.Return(errors.New("error")),
				subscriptionRepo:        nil,
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).IncreaseMock.Return(nil),
				postLogic:               mocks.NewIPostLogicMock(mc).PopularityCheckMock.Return(nil),
				dailyBonus:              5,
				tokenExpiration:         time.Hour,
				secretKey:               "secretkey",
			},
			args: args{
				ctx:      context.Background(),
				login:    "uehfkjsyg",
				password: password,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Profile{
				userRepo:                tt.fields.userRepo,
				subscriptionRepo:        tt.fields.subscriptionRepo,
				balanceTransactionLogic: tt.fields.balanceTransactionLogic,
				postLogic:               tt.fields.postLogic,
				dailyBonus:              tt.fields.dailyBonus,
				tokenExpiration:         tt.fields.tokenExpiration,
				secretKey:               tt.fields.secretKey,
			}
			_, err := p.Login(tt.args.ctx, tt.args.login, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProfile_Get(t *testing.T) {
	mc := minimock.NewController(t)
	user := &models.User{
		UUID:        uuid.New(),
		Login:       "uehfkjsyg",
		Balance:     0,
		Mail:        "mail@gmail.com",
		EnterTime:   time.Now().UTC(),
		Picture:     "jkjlkjlkj",
		Description: "fgfdydfj",
	}
	type fields struct {
		userRepo                interfaces.IUserRepository
		subscriptionRepo        interfaces.ISubscriptionRepository
		balanceTransactionLogic interfaces.IBalanceTransactionLogic
		postLogic               interfaces.IPostLogic
		dailyBonus              int
		tokenExpiration         time.Duration
		secretKey               string
	}
	type args struct {
		ctx   context.Context
		login string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "successful get",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetByLoginMock.Return(user, nil),
				subscriptionRepo:        nil,
				balanceTransactionLogic: nil,
				postLogic:               nil,
				dailyBonus:              0,
				tokenExpiration:         0,
				secretKey:               "",
			},
			args: args{
				ctx:   context.Background(),
				login: "rgrtgrrg",
			},
			want:    user,
			wantErr: false,
		},
		{
			name: "unsuccessful get: get by login",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetByLoginMock.Return(user, errors.New("error")),
				subscriptionRepo:        nil,
				balanceTransactionLogic: nil,
				postLogic:               nil,
				dailyBonus:              0,
				tokenExpiration:         0,
				secretKey:               "",
			},
			args: args{
				ctx:   context.Background(),
				login: "rgrtgrrg",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Profile{
				userRepo:                tt.fields.userRepo,
				subscriptionRepo:        tt.fields.subscriptionRepo,
				balanceTransactionLogic: tt.fields.balanceTransactionLogic,
				postLogic:               tt.fields.postLogic,
				dailyBonus:              tt.fields.dailyBonus,
				tokenExpiration:         tt.fields.tokenExpiration,
				secretKey:               tt.fields.secretKey,
			}
			got, err := p.Get(tt.args.ctx, tt.args.login)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProfile_Delete(t *testing.T) {
	mc := minimock.NewController(t)
	type fields struct {
		userRepo                interfaces.IUserRepository
		subscriptionRepo        interfaces.ISubscriptionRepository
		balanceTransactionLogic interfaces.IBalanceTransactionLogic
		postLogic               interfaces.IPostLogic
		dailyBonus              int
		tokenExpiration         time.Duration
		secretKey               string
	}
	type args struct {
		ctx   context.Context
		login string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "successful delete",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).DeleteMock.Return(nil),
				subscriptionRepo:        nil,
				balanceTransactionLogic: nil,
				postLogic:               nil,
				dailyBonus:              0,
				tokenExpiration:         0,
				secretKey:               "",
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
				login: "t3gtyghht",
			},
			wantErr: false,
		},
		{
			name: "unsuccessful delete: autodelete",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).DeleteMock.Return(nil),
				subscriptionRepo:        nil,
				balanceTransactionLogic: nil,
				postLogic:               nil,
				dailyBonus:              0,
				tokenExpiration:         0,
				secretKey:               "",
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
				login: "uehfkjs",
			},
			wantErr: true,
		},
		{
			name: "unsuccessful delete: delete",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).DeleteMock.Return(errors.New("error")),
				subscriptionRepo:        nil,
				balanceTransactionLogic: nil,
				postLogic:               nil,
				dailyBonus:              0,
				tokenExpiration:         0,
				secretKey:               "",
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
				login: "t3gtyghht",
			},
			wantErr: true,
		},
		{
			name: "unsuccessful delete: not an admin",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).DeleteMock.Return(nil),
				subscriptionRepo:        nil,
				balanceTransactionLogic: nil,
				postLogic:               nil,
				dailyBonus:              0,
				tokenExpiration:         0,
				secretKey:               "",
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
					IsAdmin:     false,
				}),
				login: "t3gtyghht",
			},
			wantErr: true,
		},
		{
			name: "unsuccessful delete: empty context",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).DeleteMock.Return(nil),
				subscriptionRepo:        nil,
				balanceTransactionLogic: nil,
				postLogic:               nil,
				dailyBonus:              0,
				tokenExpiration:         0,
				secretKey:               "",
			},
			args: args{
				ctx:   context.Background(),
				login: "t3gtyghht",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Profile{
				userRepo:                tt.fields.userRepo,
				subscriptionRepo:        tt.fields.subscriptionRepo,
				balanceTransactionLogic: tt.fields.balanceTransactionLogic,
				postLogic:               tt.fields.postLogic,
				dailyBonus:              tt.fields.dailyBonus,
				tokenExpiration:         tt.fields.tokenExpiration,
				secretKey:               tt.fields.secretKey,
			}
			if err := p.Delete(tt.args.ctx, tt.args.login); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProfile_GetAll(t *testing.T) {
	mc := minimock.NewController(t)
	users := []*models.User{{
		UUID:        uuid.New(),
		Login:       "uehfkjsyg",
		Balance:     0,
		Mail:        "mail@gmail.com",
		EnterTime:   time.Now().UTC(),
		Picture:     "jkjlkjlkj",
		Description: "fgfdydfj",
	}}
	type fields struct {
		userRepo                interfaces.IUserRepository
		subscriptionRepo        interfaces.ISubscriptionRepository
		balanceTransactionLogic interfaces.IBalanceTransactionLogic
		postLogic               interfaces.IPostLogic
		dailyBonus              int
		tokenExpiration         time.Duration
		secretKey               string
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
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetAllMock.Return(users, nil),
				subscriptionRepo:        nil,
				balanceTransactionLogic: nil,
				postLogic:               nil,
				dailyBonus:              0,
				tokenExpiration:         0,
				secretKey:               "",
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
			name: "unsuccessful get all: get all",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetAllMock.Return(users, errors.New("error")),
				subscriptionRepo:        nil,
				balanceTransactionLogic: nil,
				postLogic:               nil,
				dailyBonus:              0,
				tokenExpiration:         0,
				secretKey:               "",
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
			name: "unsuccessful get all: not an admin",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetAllMock.Return(users, nil),
				subscriptionRepo:        nil,
				balanceTransactionLogic: nil,
				postLogic:               nil,
				dailyBonus:              0,
				tokenExpiration:         0,
				secretKey:               "",
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
					IsAdmin:     false,
				}),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful get all: empty context",
			fields: fields{
				userRepo:                mocks.NewIUserRepositoryMock(mc).GetAllMock.Return(users, nil),
				subscriptionRepo:        nil,
				balanceTransactionLogic: nil,
				postLogic:               nil,
				dailyBonus:              0,
				tokenExpiration:         0,
				secretKey:               "",
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Profile{
				userRepo:                tt.fields.userRepo,
				subscriptionRepo:        tt.fields.subscriptionRepo,
				balanceTransactionLogic: tt.fields.balanceTransactionLogic,
				postLogic:               tt.fields.postLogic,
				dailyBonus:              tt.fields.dailyBonus,
				tokenExpiration:         tt.fields.tokenExpiration,
				secretKey:               tt.fields.secretKey,
			}
			got, err := p.GetAll(tt.args.ctx, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}
