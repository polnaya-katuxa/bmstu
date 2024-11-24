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

func TestPost_Publish(t *testing.T) {
	mc := minimock.NewController(t)

	post := &models.Post{
		UUID:       uuid.New(),
		Content:    "erysryth",
		PublicTime: time.Now().UTC(),
		Perms:      models.Free,
		Reactions:  nil,
		Author:     nil,
		NextLimit:  models.Limit{},
	}

	limits := []*models.Limit{{
		UUID:  uuid.New(),
		Value: 0,
		Bonus: 2,
	}, {
		UUID:  uuid.New(),
		Value: 2,
		Bonus: 5,
	}}
	type fields struct {
		postRepo                interfaces.IPostRepository
		reactionRepo            interfaces.IReactionRepository
		balanceTransactionLogic interfaces.IBalanceTransactionLogic
		subscriptionLogic       interfaces.ISubscriptionLogic
	}
	type args struct {
		ctx  context.Context
		post *models.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Post
		wantErr bool
	}{
		{
			name: "successful publish",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).CreateMock.Return(post, nil).GetSortedLimitsMock.Return(limits, nil),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
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
				post: post,
			},
			want:    post,
			wantErr: false,
		},
		{
			name: "unsuccessful publish: empty context",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).CreateMock.Return(post, nil).GetSortedLimitsMock.Return(limits, nil),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx:  context.Background(),
				post: post,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful publish: get limits",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).CreateMock.Return(post, nil).GetSortedLimitsMock.Return(limits, errors.New("error")),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
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
				post: post,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "successful publish: create",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).CreateMock.Return(nil, errors.New("error")).GetSortedLimitsMock.Return(limits, nil),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
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
				post: post,
			},
			want:    nil,
			wantErr: true,
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
			got, err := p.Publish(tt.args.ctx, tt.args.post)
			if (err != nil) != tt.wantErr {
				t.Errorf("Publish() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Publish() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPost_getReactions(t *testing.T) {
	mc := minimock.NewController(t)
	r := []*models.Reaction{{
		UUID:      uuid.New(),
		Icon:      "fdgrega",
		TypeID:    uuid.New(),
		ReactorID: uuid.New(),
	}}

	type fields struct {
		postRepo                interfaces.IPostRepository
		reactionRepo            interfaces.IReactionRepository
		balanceTransactionLogic interfaces.IBalanceTransactionLogic
		subscriptionLogic       interfaces.ISubscriptionLogic
	}
	type args struct {
		ctx  context.Context
		post *models.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Reaction
		wantErr bool
	}{
		{
			name: "sucessful get",
			fields: fields{
				postRepo:                nil,
				reactionRepo:            mocks.NewIReactionRepositoryMock(mc).GetAllMock.Return(r, nil),
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
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
				post: &models.Post{
					UUID:       uuid.New(),
					Content:    "erysryth",
					PublicTime: time.Now().UTC(),
					Perms:      models.Free,
					Reactions:  nil,
					Author:     nil,
					NextLimit:  models.Limit{},
				},
			},
			want:    r,
			wantErr: false,
		},
		{
			name: "unsucessful get",
			fields: fields{
				postRepo:                nil,
				reactionRepo:            mocks.NewIReactionRepositoryMock(mc).GetAllMock.Return(r, errors.New("error")),
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
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
				post: &models.Post{
					UUID:       uuid.New(),
					Content:    "erysryth",
					PublicTime: time.Now().UTC(),
					Perms:      models.Free,
					Reactions:  nil,
					Author:     nil,
					NextLimit:  models.Limit{},
				},
			},
			want:    nil,
			wantErr: true,
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
			got, err := p.getReactions(tt.args.ctx, tt.args.post)
			if (err != nil) != tt.wantErr {
				t.Errorf("getReactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getReactions() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPost_GetAll(t *testing.T) {
	mc := minimock.NewController(t)
	posts := []*models.Post{{
		UUID:       uuid.New(),
		Content:    "erysryth",
		PublicTime: time.Now().UTC(),
		Perms:      models.Free,
		Reactions:  nil,
		Author: &models.User{
			UUID:        uuid.New(),
			Login:       "erwgwrg",
			Balance:     10,
			Mail:        "ewgrfewr",
			EnterTime:   time.Now().UTC(),
			Picture:     "ert425t",
			Description: "34tr243t",
		},
		NextLimit: models.Limit{},
	}}
	posts2 := []*models.Post{{
		UUID:       uuid.New(),
		Content:    "erysryth",
		PublicTime: time.Now().UTC(),
		Perms:      models.Paid,
		Reactions:  nil,
		Author: &models.User{
			UUID:        uuid.New(),
			Login:       "erwgwrg",
			Balance:     10,
			Mail:        "ewgrfewr",
			EnterTime:   time.Now().UTC(),
			Picture:     "ert425t",
			Description: "34tr243t",
		},
		NextLimit: models.Limit{},
	}}

	type fields struct {
		postRepo                interfaces.IPostRepository
		reactionRepo            interfaces.IReactionRepository
		balanceTransactionLogic interfaces.IBalanceTransactionLogic
		subscriptionLogic       interfaces.ISubscriptionLogic
	}
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Post
		wantErr bool
	}{
		{
			name: "successful get all",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetAllMock.Return(posts, nil),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
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
				userID: uuid.New(),
			},
			want:    posts,
			wantErr: false,
		},
		{
			name: "successful get all own posts",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetAllMock.Return(posts, nil),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:        posts[0].Author.UUID,
					Login:       "uehfkjs",
					Balance:     0,
					Mail:        "mail@gmail.com",
					EnterTime:   time.Now().UTC(),
					Picture:     "jkjlkjlkj",
					Description: "fgfdydfj",
				}),
				userID: posts[0].Author.UUID,
			},
			want:    posts,
			wantErr: false,
		},
		{
			name: "unsuccessful get all posts: empty context",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetAllMock.Return(posts, nil),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx:    context.Background(),
				userID: posts[0].Author.UUID,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful get all posts",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetAllMock.Return(nil, errors.New("error")),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
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
				userID: posts[0].Author.UUID,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful get all paid posts: subscribe",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetAllMock.Return(posts2, nil),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       mocks.NewISubscriptionLogicMock(mc).IsSubscribedMock.Return(false, errors.New("error")),
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
				userID: uuid.New(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful get all paid posts: unsubscribed",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetAllMock.Return(posts2, nil),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       mocks.NewISubscriptionLogicMock(mc).IsSubscribedMock.Return(false, nil),
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
				userID: uuid.New(),
			},
			want:    []*models.Post{},
			wantErr: false,
		},
		{
			name: "successful get all paid posts: subscribed",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetAllMock.Return(posts2, nil),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       mocks.NewISubscriptionLogicMock(mc).IsSubscribedMock.Return(true, nil),
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
				userID: uuid.New(),
			},
			want:    posts2,
			wantErr: false,
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
			got, err := p.GetAll(tt.args.ctx, tt.args.userID, nil)
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

func TestPost_View(t *testing.T) {
	id := uuid.New()
	type fields struct {
		postRepo                interfaces.IPostRepository
		reactionRepo            interfaces.IReactionRepository
		balanceTransactionLogic interfaces.IBalanceTransactionLogic
		subscriptionLogic       interfaces.ISubscriptionLogic
	}
	type args struct {
		ctx    context.Context
		post   *models.Post
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "unsuccessful view own post",
			fields: fields{
				postRepo:                nil,
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
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
				post: &models.Post{
					UUID:       uuid.New(),
					Content:    "erysryth",
					PublicTime: time.Now().UTC(),
					Perms:      models.Free,
					Reactions:  nil,
					Author: &models.User{
						UUID:        id,
						Login:       "rtgrt",
						Balance:     0,
						Mail:        "3tg3t5",
						EnterTime:   time.Now().UTC(),
						Picture:     "rtgt",
						Description: "rgr2t",
					},
					NextLimit: models.Limit{},
				},
				userID: id,
			},
			want:    false,
			wantErr: false,
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
			got, err := p.View(tt.args.ctx, tt.args.post, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("View() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("View() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPost_Delete(t *testing.T) {
	mc := minimock.NewController(t)
	id := uuid.New()
	post := &models.Post{
		UUID:       uuid.New(),
		Content:    "erysryth",
		PublicTime: time.Now().UTC(),
		Perms:      models.Free,
		Reactions:  nil,
		Author: &models.User{
			UUID:        id,
			Login:       "erwgwrg",
			Balance:     10,
			Mail:        "ewgrfewr",
			EnterTime:   time.Now().UTC(),
			Picture:     "ert425t",
			Description: "34tr243t",
		},
		NextLimit: models.Limit{},
	}
	type fields struct {
		postRepo                interfaces.IPostRepository
		reactionRepo            interfaces.IReactionRepository
		balanceTransactionLogic interfaces.IBalanceTransactionLogic
		subscriptionLogic       interfaces.ISubscriptionLogic
	}
	type args struct {
		ctx    context.Context
		postID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "successful delete post",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).DeleteMock.Return(nil).GetMock.Return(post, nil),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:        id,
					Login:       "uehfkjs",
					Balance:     0,
					Mail:        "mail@gmail.com",
					EnterTime:   time.Now().UTC(),
					Picture:     "jkjlkjlkj",
					Description: "fgfdydfj",
				}),
				postID: uuid.New(),
			},
			wantErr: false,
		},
		{
			name: "unsuccessful delete post: not author",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).DeleteMock.Return(nil).GetMock.Return(post, nil),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
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
				postID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "unsuccessful delete post: empty context",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).DeleteMock.Return(nil).GetMock.Return(post, nil),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx:    context.Background(),
				postID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "unsuccessful delete post",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).DeleteMock.Return(errors.New("error")).GetMock.Return(post, nil),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:        id,
					Login:       "uehfkjs",
					Balance:     0,
					Mail:        "mail@gmail.com",
					EnterTime:   time.Now().UTC(),
					Picture:     "jkjlkjlkj",
					Description: "fgfdydfj",
				}),
				postID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "unsuccessful delete post: get",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).DeleteMock.Return(nil).GetMock.Return(nil, errors.New("error")),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:        id,
					Login:       "uehfkjs",
					Balance:     0,
					Mail:        "mail@gmail.com",
					EnterTime:   time.Now().UTC(),
					Picture:     "jkjlkjlkj",
					Description: "fgfdydfj",
				}),
				postID: uuid.New(),
			},
			wantErr: true,
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
			if err := p.Delete(tt.args.ctx, tt.args.postID); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPost_React(t *testing.T) {
	mc := minimock.NewController(t)
	id := uuid.New()
	post := &models.Post{
		UUID:       uuid.New(),
		Content:    "erysryth",
		PublicTime: time.Now().UTC(),
		Perms:      models.Free,
		Reactions: []*models.Reaction{{
			UUID:      uuid.New(),
			Icon:      "rtgtwt",
			TypeID:    uuid.New(),
			ReactorID: uuid.New(),
		}},
		Author: &models.User{
			UUID:        id,
			Login:       "erwgwrg",
			Balance:     10,
			Mail:        "ewgrfewr",
			EnterTime:   time.Now().UTC(),
			Picture:     "ert425t",
			Description: "34tr243t",
		},
		NextLimit: models.Limit{},
	}
	post1 := &models.Post{
		UUID:       uuid.New(),
		Content:    "erysryth",
		PublicTime: time.Now().UTC(),
		Perms:      models.Paid,
		Reactions: []*models.Reaction{{
			UUID:      uuid.New(),
			Icon:      "rtgtwt",
			TypeID:    uuid.New(),
			ReactorID: uuid.New(),
		}},
		Author: &models.User{
			UUID:        id,
			Login:       "erwgwrg",
			Balance:     10,
			Mail:        "ewgrfewr",
			EnterTime:   time.Now().UTC(),
			Picture:     "ert425t",
			Description: "34tr243t",
		},
		NextLimit: models.Limit{},
	}
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
	}{
		{
			name: "unsuccessful react: empty context",
			fields: fields{
				postRepo:                nil,
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx:    context.Background(),
				postID: uuid.New(),
				typeID: uuid.New(),
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "unsuccessful react: get",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, errors.New("error")),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:        id,
					Login:       "uehfkjs",
					Balance:     0,
					Mail:        "mail@gmail.com",
					EnterTime:   time.Now().UTC(),
					Picture:     "jkjlkjlkj",
					Description: "fgfdydfj",
				}),
				postID: uuid.New(),
				typeID: uuid.New(),
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "unsuccessful react: author",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, nil),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:        id,
					Login:       "uehfkjs",
					Balance:     0,
					Mail:        "mail@gmail.com",
					EnterTime:   time.Now().UTC(),
					Picture:     "jkjlkjlkj",
					Description: "fgfdydfj",
				}),
				postID: uuid.New(),
				typeID: uuid.New(),
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "unsuccessful react: paid post",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post1, nil),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       mocks.NewISubscriptionLogicMock(mc).IsSubscribedMock.Return(false, nil),
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
				postID: uuid.New(),
				typeID: uuid.New(),
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "unsuccessful react: reacted",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, nil),
				reactionRepo:            mocks.NewIReactionRepositoryMock(mc).ReactedMock.Return(false, errors.New("error")),
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
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
				postID: uuid.New(),
				typeID: uuid.New(),
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "unsuccessful react: already reacted delete",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, nil),
				reactionRepo:            mocks.NewIReactionRepositoryMock(mc).ReactedMock.Return(true, nil).DeleteMock.Return(errors.New("error")),
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
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
				postID: uuid.New(),
				typeID: uuid.New(),
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "unsuccessful react: react create",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, nil),
				reactionRepo:            mocks.NewIReactionRepositoryMock(mc).ReactedMock.Return(false, nil).CreateMock.Return(errors.New("error")),
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
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
				postID: uuid.New(),
				typeID: uuid.New(),
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "successful react: react create",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, nil),
				reactionRepo:            mocks.NewIReactionRepositoryMock(mc).ReactedMock.Return(false, nil).CreateMock.Return(nil),
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
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
				postID: uuid.New(),
				typeID: uuid.New(),
			},
			want:    true,
			wantErr: false,
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
		})
	}
}

func TestPost_ChangePerms(t *testing.T) {
	mc := minimock.NewController(t)
	id := uuid.New()
	post := &models.Post{
		UUID:       uuid.New(),
		Content:    "erysryth",
		PublicTime: time.Now().UTC(),
		Perms:      models.Free,
		Reactions: []*models.Reaction{{
			UUID:      uuid.New(),
			Icon:      "rtgtwt",
			TypeID:    uuid.New(),
			ReactorID: uuid.New(),
		}},
		Author: &models.User{
			UUID:        id,
			Login:       "erwgwrg",
			Balance:     10,
			Mail:        "ewgrfewr",
			EnterTime:   time.Now().UTC(),
			Picture:     "ert425t",
			Description: "34tr243t",
		},
		NextLimit: models.Limit{},
	}
	type fields struct {
		postRepo                interfaces.IPostRepository
		reactionRepo            interfaces.IReactionRepository
		balanceTransactionLogic interfaces.IBalanceTransactionLogic
		subscriptionLogic       interfaces.ISubscriptionLogic
	}
	type args struct {
		ctx    context.Context
		postID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "successful change perms",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, nil).UpdateMock.Return(nil),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:        id,
					Login:       "uehfkjs",
					Balance:     0,
					Mail:        "mail@gmail.com",
					EnterTime:   time.Now().UTC(),
					Picture:     "jkjlkjlkj",
					Description: "fgfdydfj",
				}),
				postID: uuid.New(),
			},
			wantErr: false,
		},
		{
			name: "unsuccessful change perms: empty context",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, nil).UpdateMock.Return(nil),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx:    context.Background(),
				postID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "unsuccessful change perms: get",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, errors.New("error")).UpdateMock.Return(nil),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:        id,
					Login:       "uehfkjs",
					Balance:     0,
					Mail:        "mail@gmail.com",
					EnterTime:   time.Now().UTC(),
					Picture:     "jkjlkjlkj",
					Description: "fgfdydfj",
				}),
				postID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "unsuccessful change perms: not author",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, nil).UpdateMock.Return(nil),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
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
				postID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "unsuccessful change perms: update",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, nil).UpdateMock.Return(errors.New("error")),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:        id,
					Login:       "uehfkjs",
					Balance:     0,
					Mail:        "mail@gmail.com",
					EnterTime:   time.Now().UTC(),
					Picture:     "jkjlkjlkj",
					Description: "fgfdydfj",
				}),
				postID: uuid.New(),
			},
			wantErr: true,
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
			if err := p.ChangePerms(tt.args.ctx, tt.args.postID); (err != nil) != tt.wantErr {
				t.Errorf("ChangePerms() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func genPosts_PopularityCheck(id uuid.UUID) []*models.Post {
	return []*models.Post{
		{
			UUID:       uuid.New(),
			Content:    "erysryth",
			PublicTime: time.Now().UTC(),
			Perms:      models.Free,
			Reactions: []*models.Reaction{
				{
					UUID:      uuid.New(),
					Icon:      "khlkjh",
					TypeID:    uuid.New(),
					ReactorID: uuid.New(),
				},
			},
			Author: &models.User{
				UUID:        uuid.New(),
				Login:       "erwgwrg",
				Balance:     10,
				Mail:        "ewgrfewr",
				EnterTime:   time.Now().UTC(),
				Picture:     "ert425t",
				Description: "34tr243t",
			},
			NextLimit: models.Limit{
				UUID:  id,
				Value: 0,
				Bonus: 2,
			},
		},
	}
}

func TestPost_PopularityCheck(t *testing.T) {
	limits := []*models.Limit{{
		UUID:  uuid.New(),
		Value: 0,
		Bonus: 2,
	}, {
		UUID:  uuid.New(),
		Value: 2,
		Bonus: 5,
	}}
	mc := minimock.NewController(t)
	id := uuid.New()
	type fields struct {
		postRepo                interfaces.IPostRepository
		reactionRepo            interfaces.IReactionRepository
		balanceTransactionLogic interfaces.IBalanceTransactionLogic
		subscriptionLogic       interfaces.ISubscriptionLogic
	}
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "successful popularity check",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetAllMock.Return(genPosts_PopularityCheck(id), nil).UpdateMock.Return(nil).GetSortedLimitsMock.Return(limits, nil),
				reactionRepo:            nil,
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).IncreaseMock.Return(nil),
				subscriptionLogic:       nil,
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
				userID: uuid.New(),
			},
			wantErr: false,
		},
		{
			name: "unsuccessful popularity check: get limits",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetAllMock.Return(genPosts_PopularityCheck(id), nil).UpdateMock.Return(nil).GetSortedLimitsMock.Return(limits, errors.New("error")),
				reactionRepo:            nil,
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).IncreaseMock.Return(nil),
				subscriptionLogic:       nil,
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
				userID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "unsuccessful popularity check: get all",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetAllMock.Return(nil, errors.New("error")).UpdateMock.Return(nil).GetSortedLimitsMock.Return(limits, nil),
				reactionRepo:            nil,
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).IncreaseMock.Return(nil),
				subscriptionLogic:       nil,
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
				userID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "unsuccessful popularity check: update",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetAllMock.Return(genPosts_PopularityCheck(id), nil).UpdateMock.Return(errors.New("error")).GetSortedLimitsMock.Return(limits, nil),
				reactionRepo:            nil,
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).IncreaseMock.Return(nil),
				subscriptionLogic:       nil,
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
				userID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "unsuccessful popularity check: transaction",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetAllMock.Return(genPosts_PopularityCheck(id), nil).UpdateMock.Return(nil).GetSortedLimitsMock.Return(limits, nil),
				reactionRepo:            nil,
				balanceTransactionLogic: mocks.NewIBalanceTransactionLogicMock(mc).IncreaseMock.Return(errors.New("error")),
				subscriptionLogic:       nil,
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
				userID: uuid.New(),
			},
			wantErr: true,
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
			if err := p.PopularityCheck(tt.args.ctx, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("PopularityCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPost_GetReactionTypes(t *testing.T) {
	mc := minimock.NewController(t)

	rts := []*models.ReactionType{{
		UUID: uuid.New(),
		Icon: "1",
	}, {
		UUID: uuid.New(),
		Icon: "2",
	}}

	type fields struct {
		postRepo                interfaces.IPostRepository
		reactionRepo            interfaces.IReactionRepository
		balanceTransactionLogic interfaces.IBalanceTransactionLogic
		subscriptionLogic       interfaces.ISubscriptionLogic
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.ReactionType
		wantErr bool
	}{
		{
			name: "successfull get reaction types",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetReactionTypesMock.Return(rts, nil),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{ctx: userContext.UserToContext(context.Background(), &models.User{
				UUID:        uuid.New(),
				Login:       "uehfkjs",
				Balance:     0,
				Mail:        "mail@gmail.com",
				EnterTime:   time.Now().UTC(),
				Picture:     "jkjlkjlkj",
				Description: "fgfdydfj",
			})},
			want:    rts,
			wantErr: false,
		},
		{
			name: "unsuccessfull get reaction types: get",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetReactionTypesMock.Return(nil, errors.New("error")),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{ctx: userContext.UserToContext(context.Background(), &models.User{
				UUID:        uuid.New(),
				Login:       "uehfkjs",
				Balance:     0,
				Mail:        "mail@gmail.com",
				EnterTime:   time.Now().UTC(),
				Picture:     "jkjlkjlkj",
				Description: "fgfdydfj",
			})},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessfull get reaction types: empty context",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetReactionTypesMock.Return(nil, errors.New("error")),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args:    args{ctx: context.Background()},
			want:    nil,
			wantErr: true,
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
			got, err := p.GetReactionTypes(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetReactionTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReactionTypes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPost_Get(t *testing.T) {
	mc := minimock.NewController(t)
	post := &models.Post{
		UUID:       uuid.New(),
		Content:    "erysryth",
		PublicTime: time.Now().UTC(),
		Perms:      models.Free,
		Reactions:  nil,
		Author: &models.User{
			UUID:        uuid.New(),
			Login:       "erwgwrg",
			Balance:     10,
			Mail:        "ewgrfewr",
			EnterTime:   time.Now().UTC(),
			Picture:     "ert425t",
			Description: "34tr243t",
		},
		NextLimit: models.Limit{},
	}
	post2 := &models.Post{
		UUID:       uuid.New(),
		Content:    "erysryth",
		PublicTime: time.Now().UTC(),
		Perms:      models.Paid,
		Reactions:  nil,
		Author: &models.User{
			UUID:        uuid.New(),
			Login:       "erwgwrg",
			Balance:     10,
			Mail:        "ewgrfewr",
			EnterTime:   time.Now().UTC(),
			Picture:     "ert425t",
			Description: "34tr243t",
		},
		NextLimit: models.Limit{},
	}

	type fields struct {
		postRepo                interfaces.IPostRepository
		reactionRepo            interfaces.IReactionRepository
		balanceTransactionLogic interfaces.IBalanceTransactionLogic
		subscriptionLogic       interfaces.ISubscriptionLogic
	}
	type args struct {
		ctx    context.Context
		postID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Post
		wantErr bool
	}{
		{
			name: "successful get post",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, nil),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
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
				postID: post.UUID,
			},
			want:    post,
			wantErr: false,
		},
		{
			name: "successful get post: own",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, nil),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:        post.Author.UUID,
					Login:       "uehfkjs",
					Balance:     0,
					Mail:        "mail@gmail.com",
					EnterTime:   time.Now().UTC(),
					Picture:     "jkjlkjlkj",
					Description: "fgfdydfj",
				}),
				postID: post.UUID,
			},
			want:    post,
			wantErr: false,
		},
		{
			name: "unsuccessful get post: user from context",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, nil),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx:    context.Background(),
				postID: post.UUID,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful get post: get",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(nil, errors.New("error")),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
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
				postID: post.UUID,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful get post: view",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post2, nil),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       mocks.NewISubscriptionLogicMock(mc).IsSubscribedMock.Return(false, nil),
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
				postID: post2.UUID,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful get post: view error",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post2, nil),
				reactionRepo:            nil,
				balanceTransactionLogic: nil,
				subscriptionLogic:       mocks.NewISubscriptionLogicMock(mc).IsSubscribedMock.Return(false, errors.New("error")),
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
				postID: post2.UUID,
			},
			want:    nil,
			wantErr: true,
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
			got, err := p.Get(tt.args.ctx, tt.args.postID)
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

func TestPost_Comment(t *testing.T) {
	mc := minimock.NewController(t)

	user := &models.User{
		UUID:        uuid.New(),
		Login:       "uehfkjs",
		Balance:     0,
		Mail:        "mail@gmail.com",
		EnterTime:   time.Now().UTC(),
		Picture:     "jkjlkjlkj",
		Description: "fgfdydfj",
	}

	post := &models.Post{
		UUID:       uuid.New(),
		Content:    "erysryth",
		PublicTime: time.Now().UTC(),
		Perms:      models.Free,
		Reactions:  nil,
		Author: &models.User{
			UUID:        uuid.New(),
			Login:       "erwgwrg",
			Balance:     10,
			Mail:        "ewgrfewr",
			EnterTime:   time.Now().UTC(),
			Picture:     "ert425t",
			Description: "34tr243t",
		},
		NextLimit: models.Limit{},
	}

	post1 := &models.Post{
		UUID:       uuid.New(),
		Content:    "erysryth",
		PublicTime: time.Now().UTC(),
		Perms:      models.Paid,
		Reactions:  nil,
		Author: &models.User{
			UUID:        uuid.New(),
			Login:       "erwgwrg",
			Balance:     10,
			Mail:        "ewgrfewr",
			EnterTime:   time.Now().UTC(),
			Picture:     "ert425t",
			Description: "34tr243t",
		},
		NextLimit: models.Limit{},
	}

	comm := &models.Comment{
		UUID:        uuid.New(),
		Content:     "like",
		PublicTime:  time.Now(),
		Commentator: user,
		PostID:      post.UUID,
	}

	ctx := userContext.UserToContext(context.Background(), user)

	type fields struct {
		postRepo                interfaces.IPostRepository
		reactionRepo            interfaces.IReactionRepository
		commentRepo             interfaces.ICommentRepository
		balanceTransactionLogic interfaces.IBalanceTransactionLogic
		subscriptionLogic       interfaces.ISubscriptionLogic
	}
	type args struct {
		ctx  context.Context
		comm *models.Comment
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Comment
		wantErr bool
	}{
		{
			name: "successful comment",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, nil),
				reactionRepo:            nil,
				commentRepo:             mocks.NewICommentRepositoryMock(mc).CreateMock.Return(comm, nil),
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx:  ctx,
				comm: comm,
			},
			want:    comm,
			wantErr: false,
		},
		{
			name: "unsuccessful comment: empty context",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, nil),
				reactionRepo:            nil,
				commentRepo:             mocks.NewICommentRepositoryMock(mc).CreateMock.Return(comm, nil),
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx:  context.Background(),
				comm: comm,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful comment: get post",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, errors.New("error")),
				reactionRepo:            nil,
				commentRepo:             mocks.NewICommentRepositoryMock(mc).CreateMock.Return(comm, nil),
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx:  ctx,
				comm: comm,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful comment: view",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post1, nil),
				reactionRepo:            nil,
				commentRepo:             mocks.NewICommentRepositoryMock(mc).CreateMock.Return(comm, nil),
				balanceTransactionLogic: nil,
				subscriptionLogic:       mocks.NewISubscriptionLogicMock(mc).IsSubscribedMock.Return(false, nil),
			},
			args: args{
				ctx:  ctx,
				comm: comm,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful comment: create",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, nil),
				reactionRepo:            nil,
				commentRepo:             mocks.NewICommentRepositoryMock(mc).CreateMock.Return(comm, errors.New("error")),
				balanceTransactionLogic: nil,
				subscriptionLogic:       mocks.NewISubscriptionLogicMock(mc).IsSubscribedMock.Return(false, nil),
			},
			args: args{
				ctx:  ctx,
				comm: comm,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Post{
				postRepo:                tt.fields.postRepo,
				reactionRepo:            tt.fields.reactionRepo,
				commentRepo:             tt.fields.commentRepo,
				balanceTransactionLogic: tt.fields.balanceTransactionLogic,
				subscriptionLogic:       tt.fields.subscriptionLogic,
			}
			got, err := p.Comment(tt.args.ctx, tt.args.comm)
			if (err != nil) != tt.wantErr {
				t.Errorf("Comment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllComments() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPost_Uncomment(t *testing.T) {
	mc := minimock.NewController(t)

	user := &models.User{
		UUID:        uuid.New(),
		Login:       "uehfkjs",
		Balance:     0,
		Mail:        "mail@gmail.com",
		EnterTime:   time.Now().UTC(),
		Picture:     "jkjlkjlkj",
		Description: "fgfdydfj",
	}

	post := &models.Post{
		UUID:       uuid.New(),
		Content:    "erysryth",
		PublicTime: time.Now().UTC(),
		Perms:      models.Free,
		Reactions:  nil,
		Author: &models.User{
			UUID:        uuid.New(),
			Login:       "erwgwrg",
			Balance:     10,
			Mail:        "ewgrfewr",
			EnterTime:   time.Now().UTC(),
			Picture:     "ert425t",
			Description: "34tr243t",
		},
		NextLimit: models.Limit{},
	}

	comm := &models.Comment{
		UUID:        uuid.New(),
		Content:     "like",
		PublicTime:  time.Now(),
		Commentator: user,
		PostID:      post.UUID,
	}

	ctx := userContext.UserToContext(context.Background(), user)

	type fields struct {
		postRepo                interfaces.IPostRepository
		reactionRepo            interfaces.IReactionRepository
		commentRepo             interfaces.ICommentRepository
		balanceTransactionLogic interfaces.IBalanceTransactionLogic
		subscriptionLogic       interfaces.ISubscriptionLogic
	}
	type args struct {
		ctx    context.Context
		commID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "successful uncomment",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, nil),
				reactionRepo:            nil,
				commentRepo:             mocks.NewICommentRepositoryMock(mc).GetByIDMock.Return(comm, nil).DeleteMock.Return(nil),
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx:    ctx,
				commID: comm.UUID,
			},
			wantErr: false,
		},
		{
			name: "unsuccessful uncomment: get by id",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, nil),
				reactionRepo:            nil,
				commentRepo:             mocks.NewICommentRepositoryMock(mc).GetByIDMock.Return(comm, errors.New("error")).DeleteMock.Return(nil),
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx:    ctx,
				commID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "unsuccessful uncomment: get post",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, errors.New("error")),
				reactionRepo:            nil,
				commentRepo:             mocks.NewICommentRepositoryMock(mc).GetByIDMock.Return(comm, nil).DeleteMock.Return(nil),
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx:    ctx,
				commID: comm.UUID,
			},
			wantErr: true,
		},
		{
			name: "unsuccessful uncomment: perms",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, nil),
				reactionRepo:            nil,
				commentRepo:             mocks.NewICommentRepositoryMock(mc).GetByIDMock.Return(comm, nil).DeleteMock.Return(nil),
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx: userContext.UserToContext(context.Background(), &models.User{
					UUID:        uuid.New(),
					Login:       "uehfgkjs",
					Balance:     0,
					Mail:        "mail@gmail.com",
					EnterTime:   time.Now().UTC(),
					Picture:     "jkjlkjlkj",
					Description: "fgfdydfj",
				}),
				commID: comm.UUID,
			},
			wantErr: true,
		},
		{
			name: "unsuccessful uncomment: delete",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, nil),
				reactionRepo:            nil,
				commentRepo:             mocks.NewICommentRepositoryMock(mc).GetByIDMock.Return(comm, nil).DeleteMock.Return(errors.New("error")),
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx:    ctx,
				commID: comm.UUID,
			},
			wantErr: true,
		},
		{
			name: "unsuccessful uncomment: empty context",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, nil),
				reactionRepo:            nil,
				commentRepo:             mocks.NewICommentRepositoryMock(mc).GetByIDMock.Return(comm, nil).DeleteMock.Return(nil),
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx:    context.Background(),
				commID: comm.UUID,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Post{
				postRepo:                tt.fields.postRepo,
				reactionRepo:            tt.fields.reactionRepo,
				commentRepo:             tt.fields.commentRepo,
				balanceTransactionLogic: tt.fields.balanceTransactionLogic,
				subscriptionLogic:       tt.fields.subscriptionLogic,
			}
			if err := p.Uncomment(tt.args.ctx, tt.args.commID); (err != nil) != tt.wantErr {
				t.Errorf("Uncomment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPost_GetAllComments(t *testing.T) {
	mc := minimock.NewController(t)

	user := &models.User{
		UUID:        uuid.New(),
		Login:       "uehfkjs",
		Balance:     0,
		Mail:        "mail@gmail.com",
		EnterTime:   time.Now().UTC(),
		Picture:     "jkjlkjlkj",
		Description: "fgfdydfj",
	}

	post := &models.Post{
		UUID:       uuid.New(),
		Content:    "erysryth",
		PublicTime: time.Now().UTC(),
		Perms:      models.Free,
		Reactions:  nil,
		Author: &models.User{
			UUID:        uuid.New(),
			Login:       "erwgwrg",
			Balance:     10,
			Mail:        "ewgrfewr",
			EnterTime:   time.Now().UTC(),
			Picture:     "ert425t",
			Description: "34tr243t",
		},
		NextLimit: models.Limit{},
	}

	post1 := &models.Post{
		UUID:       uuid.New(),
		Content:    "erysryth",
		PublicTime: time.Now().UTC(),
		Perms:      models.Paid,
		Reactions:  nil,
		Author: &models.User{
			UUID:        uuid.New(),
			Login:       "erwgwrg",
			Balance:     10,
			Mail:        "ewgrfewr",
			EnterTime:   time.Now().UTC(),
			Picture:     "ert425t",
			Description: "34tr243t",
		},
		NextLimit: models.Limit{},
	}

	ctx := userContext.UserToContext(context.Background(), user)

	comms := []*models.Comment{
		{
			UUID:        uuid.New(),
			Content:     "like",
			PublicTime:  time.Now(),
			Commentator: user,
			PostID:      post.UUID,
		},
		{
			UUID:        uuid.New(),
			Content:     "dislike",
			PublicTime:  time.Now(),
			Commentator: user,
			PostID:      post.UUID,
		},
	}

	type fields struct {
		postRepo                interfaces.IPostRepository
		reactionRepo            interfaces.IReactionRepository
		commentRepo             interfaces.ICommentRepository
		balanceTransactionLogic interfaces.IBalanceTransactionLogic
		subscriptionLogic       interfaces.ISubscriptionLogic
	}
	type args struct {
		ctx    context.Context
		postID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Comment
		wantErr bool
	}{
		{
			name: "successful get all comments",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, nil),
				reactionRepo:            nil,
				commentRepo:             mocks.NewICommentRepositoryMock(mc).GetAllMock.Return(comms, nil),
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx:    ctx,
				postID: post.UUID,
			},
			want:    comms,
			wantErr: false,
		},
		{
			name: "unsuccessful get all comments: get post",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, errors.New("error")),
				reactionRepo:            nil,
				commentRepo:             mocks.NewICommentRepositoryMock(mc).GetAllMock.Return(comms, nil),
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx:    ctx,
				postID: post.UUID,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful get all comments: view",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post1, nil),
				reactionRepo:            nil,
				commentRepo:             mocks.NewICommentRepositoryMock(mc).GetAllMock.Return(comms, nil),
				balanceTransactionLogic: nil,
				subscriptionLogic:       mocks.NewISubscriptionLogicMock(mc).IsSubscribedMock.Return(false, nil),
			},
			args: args{
				ctx:    ctx,
				postID: post1.UUID,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful get all comments: get all",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, nil),
				reactionRepo:            nil,
				commentRepo:             mocks.NewICommentRepositoryMock(mc).GetAllMock.Return(comms, errors.New("error")),
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx:    ctx,
				postID: post.UUID,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful get all comments: empty context",
			fields: fields{
				postRepo:                mocks.NewIPostRepositoryMock(mc).GetMock.Return(post, nil),
				reactionRepo:            nil,
				commentRepo:             mocks.NewICommentRepositoryMock(mc).GetAllMock.Return(comms, nil),
				balanceTransactionLogic: nil,
				subscriptionLogic:       nil,
			},
			args: args{
				ctx:    context.Background(),
				postID: post.UUID,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Post{
				postRepo:                tt.fields.postRepo,
				reactionRepo:            tt.fields.reactionRepo,
				commentRepo:             tt.fields.commentRepo,
				balanceTransactionLogic: tt.fields.balanceTransactionLogic,
				subscriptionLogic:       tt.fields.subscriptionLogic,
			}
			got, err := p.GetAllComments(tt.args.ctx, tt.args.postID, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllComments() got = %v, want %v", got, tt.want)
			}
		})
	}
}
