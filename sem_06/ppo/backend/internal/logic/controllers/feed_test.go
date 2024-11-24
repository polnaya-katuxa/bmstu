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

func TestFeed_View(t *testing.T) {
	mc := minimock.NewController(t)
	id := uuid.New()
	id2 := uuid.New()
	s := []*models.Subscription{{
		UUID:     uuid.New(),
		ReaderID: id,
		WriterID: id2,
	}}
	c := userContext.UserToContext(context.Background(), &models.User{
		UUID:        id,
		Login:       "uehfkjs",
		Balance:     0,
		Mail:        "mail@gmail.com",
		EnterTime:   time.Now().UTC(),
		Picture:     "jkjlkjlkj",
		Description: "fgfdydfj",
	})
	posts1 := []*models.Post{
		{
			UUID:       uuid.New(),
			Content:    "erysryth",
			PublicTime: time.Now().UTC(),
			Perms:      models.Free,
			Reactions:  nil,
			Author: &models.User{
				UUID: id2,
			},
			NextLimit: models.Limit{},
		},
	}
	posts2 := []*models.Post{
		{
			UUID:       uuid.New(),
			Content:    "erysryth",
			PublicTime: time.Now().UTC(),
			Perms:      models.Free,
			Reactions:  nil,
			Author: &models.User{
				UUID: uuid.New(),
			},
			NextLimit: models.Limit{},
		},
		{
			UUID:       uuid.New(),
			Content:    "erysryth",
			PublicTime: time.Now().UTC(),
			Perms:      models.Free,
			Reactions:  nil,
			Author: &models.User{
				UUID: uuid.New(),
			},
			NextLimit: models.Limit{},
		},
	}

	type fields struct {
		subscriptionRepo interfaces.ISubscriptionRepository
		postRepo         interfaces.IPostRepository
		span             time.Duration
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Post
		wantErr bool
	}{
		{
			name: "successful view feed",
			fields: fields{
				subscriptionRepo: mocks.NewSubscriptionRepositoryMock(mc).GetAllMock.Return(s, nil),
				postRepo: mocks.NewIPostRepositoryMock(mc).GetByIDAndSpanMock.Set(func(_ context.Context, _ []uuid.UUID, _ time.Time, _ time.Time, b1 bool) (ppa1 []*models.Post, err error) {
					if b1 {
						return posts1, nil
					}
					return posts2, nil
				}),
				span: time.Hour,
			},
			args: args{
				ctx: c,
			},
			want:    append(posts1, posts2...),
			wantErr: false,
		},
		{
			name: "unsuccessful first get posts view feed",
			fields: fields{
				subscriptionRepo: mocks.NewSubscriptionRepositoryMock(mc).GetAllMock.Return(s, nil),
				postRepo: mocks.NewIPostRepositoryMock(mc).GetByIDAndSpanMock.Set(func(_ context.Context, _ []uuid.UUID, _ time.Time, _ time.Time, b1 bool) (ppa1 []*models.Post, err error) {
					if b1 {
						return posts1, errors.New("error")
					}
					return posts2, nil
				}),
				span: time.Hour,
			},
			args: args{
				ctx: c,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful second get posts view feed",
			fields: fields{
				subscriptionRepo: mocks.NewSubscriptionRepositoryMock(mc).GetAllMock.Return(s, nil),
				postRepo: mocks.NewIPostRepositoryMock(mc).GetByIDAndSpanMock.Set(func(_ context.Context, _ []uuid.UUID, _ time.Time, _ time.Time, b1 bool) (ppa1 []*models.Post, err error) {
					if b1 {
						return posts1, nil
					}
					return posts2, errors.New("error")
				}),
				span: time.Hour,
			},
			args: args{
				ctx: c,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful get subs view feed",
			fields: fields{
				subscriptionRepo: mocks.NewSubscriptionRepositoryMock(mc).GetAllMock.Return(s, errors.New("error")),
				postRepo: mocks.NewIPostRepositoryMock(mc).GetByIDAndSpanMock.Set(func(_ context.Context, _ []uuid.UUID, _ time.Time, _ time.Time, b1 bool) (ppa1 []*models.Post, err error) {
					if b1 {
						return posts1, nil
					}
					return posts2, nil
				}),
				span: time.Hour,
			},
			args: args{
				ctx: c,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful view feed: empty context",
			fields: fields{
				subscriptionRepo: mocks.NewSubscriptionRepositoryMock(mc).GetAllMock.Return(s, nil),
				postRepo: mocks.NewIPostRepositoryMock(mc).GetByIDAndSpanMock.Set(func(_ context.Context, _ []uuid.UUID, _ time.Time, _ time.Time, b1 bool) (ppa1 []*models.Post, err error) {
					if b1 {
						return posts1, nil
					}
					return posts2, nil
				}),
				span: time.Hour,
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
			f := &Feed{
				subscriptionRepo: tt.fields.subscriptionRepo,
				postRepo:         tt.fields.postRepo,
				span:             tt.fields.span,
			}
			got, err := f.View(tt.args.ctx, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("View() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("View() got = %v, want %v", got, tt.want)
			}
		})
	}
}
