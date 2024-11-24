package controllers

import (
	"context"
	"fmt"
	"time"

	mycontext "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/context"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/interfaces"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	"github.com/google/uuid"
)

type Feed struct {
	subscriptionRepo interfaces.ISubscriptionRepository
	postRepo         interfaces.IPostRepository

	span time.Duration
}

func NewFL(s interfaces.ISubscriptionRepository, p interfaces.IPostRepository, span time.Duration) *Feed {
	return &Feed{
		subscriptionRepo: s,
		postRepo:         p,
		span:             span,
	}
}

func (f *Feed) View(ctx context.Context, pg *models.Paginator) ([]*models.Post, error) {
	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to view feed", "error", err)
		return nil, fmt.Errorf("user from context: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("started view feed", "user", user.Login)

	subs, err := f.subscriptionRepo.GetAll(ctx, user.UUID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get subs for feed view", "user", user.Login, "error", err)
		return nil, fmt.Errorf("user from context: %w", err)
	}

	subsID := make([]uuid.UUID, 0, len(subs))
	for _, s := range subs {
		subsID = append(subsID, s.WriterID)
	}

	subsPosts, err := f.postRepo.GetByIDAndSpan(ctx, subsID, time.Now().UTC().Add(-f.span), time.Now().UTC(), true)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get subs posts for feed view", "user", user.Login, "error", err)
		return nil, fmt.Errorf("get subs posts by span: %w", err)
	}

	notSubsPosts, err := f.postRepo.GetByIDAndSpan(ctx, subsID, time.Now().UTC().Add(-f.span), time.Now().UTC(), false)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get not subs free posts for feed view", "user", user.Login, "error", err)
		return nil, fmt.Errorf("get subs posts by span: %w", err)
	}

	posts := append([]*models.Post{}, subsPosts...)
	posts = append(posts, notSubsPosts...)

	if pg == nil {
		mycontext.LoggerFromContext(ctx).Infow("successfully view feed", "user", user.Login)
		return posts, nil
	}

	lim := pg.Num * (pg.Page - 1)
	if len(posts) <= lim {
		return nil, nil
	}

	res := make([]*models.Post, 0, pg.Num)
	for i := lim; i < len(posts) && i < lim+pg.Num; i++ {
		res = append(res, posts[i])
	}

	mycontext.LoggerFromContext(ctx).Infow("successfully view feed", "user", user.Login)

	return res, nil
}

func (f *Feed) GetTotalByIDAndSpan(ctx context.Context) (int, error) {
	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to get total number of posts by id and span", "error", err)
		return 0, fmt.Errorf("user from context: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("started getting total number of posts by id and span")

	subs, err := f.subscriptionRepo.GetAll(ctx, user.UUID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get subs for feed view", "user", user.Login, "error", err)
		return 0, fmt.Errorf("user from context: %w", err)
	}

	subsID := make([]uuid.UUID, 0, len(subs))
	for _, s := range subs {
		subsID = append(subsID, s.WriterID)
	}

	num1, err := f.postRepo.GetTotalByIDAndSpan(ctx, subsID, time.Now().UTC().Add(-f.span), time.Now().UTC(), true)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user's posts number", "user", user.Login, "error", err)
		return 0, fmt.Errorf("get total by id and span: %w", err)
	}

	num2, err := f.postRepo.GetTotalByIDAndSpan(ctx, subsID, time.Now().UTC().Add(-f.span), time.Now().UTC(), false)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user's posts number", "user", user.Login, "error", err)
		return 0, fmt.Errorf("get total by id and span: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("successfully got total number of posts by id and span", "user", user.Login)

	return num1 + num2, nil
}
