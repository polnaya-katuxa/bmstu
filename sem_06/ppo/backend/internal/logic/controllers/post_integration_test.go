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

func TestIntegrationPost_React(t *testing.T) {
	dbContainer, db, err := containers.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = dbContainer.Terminate(context.Background())
	}()

	postRepo := postgres.NewPR(db)
	subRepo := postgres.NewSR(db)
	reactRepo := postgres.NewRR(db)
	userRepo := postgres.NewUR(db)
	btRepo := postgres.NewBTR(db)
	btLog := NewBTL(userRepo, btRepo)
	subLog := NewSL(userRepo, btLog, subRepo, 10)

	type fields struct {
		postRepo                interfaces.IPostRepository
		reactionRepo            interfaces.IReactionRepository
		balanceTransactionLogic interfaces.IBalanceTransactionLogic
		subscriptionLogic       interfaces.ISubscriptionLogic
	}
	type args struct {
		ctx    context.Context
		postID uuid.UUID
		typeID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
		check   func(*Post) error
	}{
		{
			name: "successful react",
			fields: fields{
				postRepo:                postRepo,
				reactionRepo:            reactRepo,
				balanceTransactionLogic: btLog,
				subscriptionLogic:       subLog,
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID: uuid.MustParse("3f010aca-5008-4aa5-a1a3-a061a876783f"),
				}),
				postID: uuid.MustParse("21bf7ace-965b-4679-86b8-93a89cba0094"),
				typeID: uuid.MustParse("bde563b1-66a6-4d00-ac3f-4022be793c81"),
			},
			want:    true,
			wantErr: false,
			check: func(p *Post) error {
				rc, err := p.reactionRepo.Reacted(
					context.Background(),
					uuid.MustParse("3f010aca-5008-4aa5-a1a3-a061a876783f"),
					uuid.MustParse("21bf7ace-965b-4679-86b8-93a89cba0094"),
					uuid.MustParse("bde563b1-66a6-4d00-ac3f-4022be793c81"),
				)
				if err != nil {
					return fmt.Errorf("reacted: %w", err)
				}
				if !rc {
					return errors.New("not reacted")
				}
				return nil
			},
		},
		{
			name: "successful unreact",
			fields: fields{
				postRepo:                postRepo,
				reactionRepo:            reactRepo,
				balanceTransactionLogic: btLog,
				subscriptionLogic:       subLog,
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID: uuid.MustParse("3f010aca-5008-4aa5-a1a3-a061a876783f"),
				}),
				postID: uuid.MustParse("21bf7ace-965b-4679-86b8-93a89cba0094"),
				typeID: uuid.MustParse("bde563b1-66a6-4d00-ac3f-4022be793c81"),
			},
			want:    false,
			wantErr: false,
			check: func(p *Post) error {
				rc, err := p.reactionRepo.Reacted(
					context.Background(),
					uuid.MustParse("3f010aca-5008-4aa5-a1a3-a061a876783f"),
					uuid.MustParse("21bf7ace-965b-4679-86b8-93a89cba0094"),
					uuid.MustParse("bde563b1-66a6-4d00-ac3f-4022be793c81"),
				)
				if err != nil {
					return fmt.Errorf("reacted: %w", err)
				}
				if rc {
					return errors.New("reacted")
				}
				return nil
			},
		},
		{
			name: "unsuccessful react: author of post",
			fields: fields{
				postRepo:                postRepo,
				reactionRepo:            reactRepo,
				balanceTransactionLogic: btLog,
				subscriptionLogic:       subLog,
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID: uuid.MustParse("77dcd288-79b2-4655-9584-cc9b5329665d"),
				}),
				postID: uuid.MustParse("21bf7ace-965b-4679-86b8-93a89cba0094"),
				typeID: uuid.MustParse("bde563b1-66a6-4d00-ac3f-4022be793c81"),
			},
			want:    false,
			wantErr: true,
			check: func(p *Post) error {
				rc, err := p.reactionRepo.Reacted(
					context.Background(),
					uuid.MustParse("77dcd288-79b2-4655-9584-cc9b5329665d"),
					uuid.MustParse("21bf7ace-965b-4679-86b8-93a89cba0094"),
					uuid.MustParse("bde563b1-66a6-4d00-ac3f-4022be793c81"),
				)
				if err != nil {
					return fmt.Errorf("reacted: %w", err)
				}
				if rc {
					return errors.New("reacted")
				}
				return nil
			},
		},
		{
			name: "unsuccessful react: not exist post",
			fields: fields{
				postRepo:                postRepo,
				reactionRepo:            reactRepo,
				balanceTransactionLogic: btLog,
				subscriptionLogic:       subLog,
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID: uuid.MustParse("77dcd288-79b2-4655-9584-cc9b5329665d"),
				}),
				postID: uuid.MustParse("21bf7ace-965b-4679-86b8-93a89cba0095"),
				typeID: uuid.MustParse("bde563b1-66a6-4d00-ac3f-4022be793c81"),
			},
			want:    false,
			wantErr: true,
			check: func(p *Post) error {
				rc, err := p.reactionRepo.Reacted(
					context.Background(),
					uuid.MustParse("77dcd288-79b2-4655-9584-cc9b5329665d"),
					uuid.MustParse("21bf7ace-965b-4679-86b8-93a89cba0095"),
					uuid.MustParse("bde563b1-66a6-4d00-ac3f-4022be793c81"),
				)
				if err != nil {
					return fmt.Errorf("react: %w", err)
				}
				if rc {
					return errors.New("reacted")
				}
				return nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Post{
				postRepo:                tt.fields.postRepo,
				reactionRepo:            tt.fields.reactionRepo,
				balanceTransactionLogic: tt.fields.balanceTransactionLogic,
				subscriptionLogic:       tt.fields.subscriptionLogic,
			}
			got, err := p.React(tt.args.ctx, tt.args.postID, tt.args.typeID)
			if (err != nil) != tt.wantErr {
				t.Errorf("React() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("React() got = %v, want %v", got, tt.want)
			}
			if err := tt.check(p); err != nil {
				t.Errorf("React() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
