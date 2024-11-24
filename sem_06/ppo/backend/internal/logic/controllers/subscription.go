package controllers

import (
	"context"
	"fmt"

	mycontext "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/context"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/interfaces"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	"github.com/google/uuid"
)

type Subscription struct {
	userRepo         interfaces.IUserRepository
	subscriptionRepo interfaces.ISubscriptionRepository

	balanceTransactionLogic interfaces.IBalanceTransactionLogic

	cost int
}

func NewSL(u interfaces.IUserRepository, bt interfaces.IBalanceTransactionLogic, s interfaces.ISubscriptionRepository, c int) *Subscription {
	return &Subscription{
		userRepo:                u,
		balanceTransactionLogic: bt,
		subscriptionRepo:        s,
		cost:                    c,
	}
}

func (s *Subscription) Subscribe(ctx context.Context, writeID uuid.UUID) error {
	mycontext.LoggerFromContext(ctx).Infow("started changing subscription", "user-writer", writeID)

	reader, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user-reader", "user-writer", writeID, "error", err)
		return fmt.Errorf("user from context: %w", err)
	}

	flag, err := s.IsSubscribed(ctx, reader.UUID, writeID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot check if is subscribed", "user-reader", reader.UUID, "user-writer", writeID, "error", err)
		return fmt.Errorf("is subscribed: %w", err)
	}

	if flag {
		err = s.unsubscribe(ctx, reader.UUID, writeID)
		if err != nil {
			mycontext.LoggerFromContext(ctx).Warnw("cannot unsubscribe", "user-reader", reader.UUID, "user-writer", writeID, "error", err)
			return fmt.Errorf("unsubscribe: %w", err)
		}

		mycontext.LoggerFromContext(ctx).Infow("successfully unsubscribed", "user-reader", reader.UUID, "user-writer", writeID)

		return nil
	}

	err = s.balanceTransactionLogic.Decrease(ctx, s.cost, "paid subscription")
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot decrease reader balance on subscription", "user-reader", reader.UUID, "user-writer", writeID, "error", err)
		return fmt.Errorf("decrease: %w", err)
	}

	err = s.balanceTransactionLogic.IncreaseSub(ctx, s.cost, writeID)
	if err != nil {
		_ = s.balanceTransactionLogic.IncreaseSub(ctx, s.cost, reader.UUID)
		mycontext.LoggerFromContext(ctx).Warnw("cannot increase writer balance on subscription", "user-reader", reader.UUID, "user-writer", writeID, "error", err)
		return fmt.Errorf("increase: %w", err)
	}

	err = s.subscriptionRepo.Create(ctx, reader.UUID, writeID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Errorw("cannot create subscription", "user-reader", reader.UUID, "user-writer", writeID, "error", err)
		return fmt.Errorf("create: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("successfully subscribed", "user-reader", reader.UUID, "user-writer", writeID)

	return nil
}

func (s *Subscription) IsSubscribed(ctx context.Context, readID uuid.UUID, writeID uuid.UUID) (bool, error) {
	mycontext.LoggerFromContext(ctx).Infow("started checking if subscribed", "user-reader", readID, "user-writer", writeID)

	sub, err := s.subscriptionRepo.Get(ctx, readID, writeID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get subscription", "user-reader", readID, "user-writer", writeID, "error", err)
		return false, fmt.Errorf("get all: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("successfully checked subscription", "user-reader", readID, "user-writer", writeID)

	if sub == nil {
		return false, nil
	}

	return true, nil
}

func (s *Subscription) GetSubscriptions(ctx context.Context) ([]*models.User, error) {
	mycontext.LoggerFromContext(ctx).Infow("started getting subscriptions")

	reader, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to view subscriptions", "error", err)
		return nil, fmt.Errorf("user from context: %w", err)
	}

	subs, err := s.subscriptionRepo.GetAll(ctx, reader.UUID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get all subscriptions", "user-reader", reader.UUID, "error", err)
		return nil, fmt.Errorf("get all: %w", err)
	}

	users := make([]*models.User, 0, len(subs))
	for _, sub := range subs {
		user, err := s.userRepo.GetByID(ctx, sub.WriterID)
		if err != nil {
			mycontext.LoggerFromContext(ctx).Warnw("cannot get user's sub by id", "user-reader", reader.UUID, "user-writer", sub.WriterID, "error", err)
			return nil, fmt.Errorf("get: %w", err)
		}

		users = append(users, user)
	}

	mycontext.LoggerFromContext(ctx).Infow("successfully got all subscriptions", "user-reader", reader.UUID)

	return users, nil
}

func (s *Subscription) unsubscribe(ctx context.Context, readID uuid.UUID, writeID uuid.UUID) error {
	mycontext.LoggerFromContext(ctx).Infow("started unsubscribing", "user-reader", readID, "user-writer", writeID)

	err := s.subscriptionRepo.Delete(ctx, readID, writeID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot unsubscribe", "user-reader", readID, "user-writer", writeID, "error", err)
		return fmt.Errorf("delete: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("successfully unsubscribed", "user-reader", readID, "user-writer", writeID)

	return nil
}
