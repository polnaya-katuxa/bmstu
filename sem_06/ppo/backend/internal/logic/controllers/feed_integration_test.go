package controllers

import (
	"context"
	"testing"
	"time"

	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/containers"
	userContext "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/context"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/interfaces"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/repository/postgres"
	"github.com/google/uuid"
)

func neededPosts(p1 []*models.Post, p2 []*models.Post) bool {
	if len(p1) != len(p2) {
		return false
	}

	for i := range p1 {
		if p1[i].Content != p2[i].Content {
			return false
		}

		if p1[i].UUID.String() != p2[i].UUID.String() {
			return false
		}
	}

	return true
}

func TestIntegrationFeed_View(t *testing.T) {
	dbContainer, db, err := containers.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = dbContainer.Terminate(context.Background())
	}()

	subRepo := postgres.NewSR(db)
	postRepo := postgres.NewPR(db)

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
			name: "successful view feed: got some subs",
			fields: fields{
				subscriptionRepo: subRepo,
				postRepo:         postRepo,
				span:             time.Hour,
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID: uuid.MustParse("a52b8aea-d751-4933-91bb-691132e3b760"),
				}),
			},
			want: []*models.Post{
				{
					UUID:    uuid.MustParse("1557a6be-1008-412a-88d9-1f06630d028c"),
					Content: "1",
				},
				{
					UUID:    uuid.MustParse("21bf7ace-965b-4679-86b8-93a89cba0094"),
					Content: "2",
				},
				{
					UUID:    uuid.MustParse("eea07263-3444-40d0-adc4-345ad7728298"),
					Content: "3",
				},
			},
			wantErr: false,
		},
		{
			name: "successful view feed: no subs",
			fields: fields{
				subscriptionRepo: subRepo,
				postRepo:         postRepo,
				span:             time.Hour,
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID: uuid.New(),
				}),
			},
			want: []*models.Post{
				{
					UUID:    uuid.MustParse("21bf7ace-965b-4679-86b8-93a89cba0094"),
					Content: "2",
				},
				{
					UUID:    uuid.MustParse("eea07263-3444-40d0-adc4-345ad7728298"),
					Content: "3",
				},
			},
			wantErr: false,
		},
		{
			name: "successful view feed: all subs",
			fields: fields{
				subscriptionRepo: subRepo,
				postRepo:         postRepo,
				span:             time.Hour,
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID: uuid.MustParse("b8b87f28-3cbb-425c-ac4d-52015710d61b"),
				}),
			},
			want: []*models.Post{
				{
					UUID:    uuid.MustParse("1557a6be-1008-412a-88d9-1f06630d028c"),
					Content: "1",
				},
				{
					UUID:    uuid.MustParse("21bf7ace-965b-4679-86b8-93a89cba0094"),
					Content: "2",
				},
				{
					UUID:    uuid.MustParse("eea07263-3444-40d0-adc4-345ad7728298"),
					Content: "3",
				},
				{
					UUID:    uuid.MustParse("c11add9b-207e-4b41-a964-2662ed3cae27"),
					Content: "4",
				},
			},
			wantErr: false,
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
			if !neededPosts(got, tt.want) {
				t.Errorf("View() got = %v, want %v", got, tt.want)
			}
		})
	}
}
