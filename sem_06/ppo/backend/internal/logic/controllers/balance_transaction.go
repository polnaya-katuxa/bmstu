package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	mycontext "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/context"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/interfaces"
	errors2 "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/errors"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
)

type BalanceTransaction struct {
	userRepo               interfaces.IUserRepository
	balanceTransactionRepo interfaces.IBalanceTransactionRepository
}

func NewBTL(u interfaces.IUserRepository, bt interfaces.IBalanceTransactionRepository) *BalanceTransaction {
	return &BalanceTransaction{
		userRepo:               u,
		balanceTransactionRepo: bt,
	}
}

func (bt *BalanceTransaction) Increase(ctx context.Context, value int, reason string) error {
	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user for increase transaction", "value", value, "reason", reason, "error", err)
		return fmt.Errorf("user from context: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("started increase transaction", "user", user.Login, "value", value, "reason", reason)

	user.Balance += value

	err = bt.userRepo.Update(ctx, user)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot update user for increase transaction", "user", user.Login, "value", value, "reason", reason, "error", err)
		return fmt.Errorf("update: %w", err)
	}

	t := &models.BalanceTransaction{
		Reason: reason,
		Time:   time.Now().UTC(),
		Amount: value,
		UserID: user.UUID,
	}

	err = bt.balanceTransactionRepo.Create(ctx, t)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot create increase transaction", "user", user.Login, "value", t.Amount, "reason", t.Reason, "time", t.Time, "error", err)
		return fmt.Errorf("create: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("successful increase transaction", "user", user.Login, "value", t.Amount, "reason", t.Reason, "time", t.Time)

	return nil
}

func (bt *BalanceTransaction) Decrease(ctx context.Context, value int, reason string) error {
	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user for decrease transaction", "value", value, "reason", reason, "error", err)
		return fmt.Errorf("user from context: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("started decrease transaction", "user", user.Login, "value", value, "reason", reason)

	if user.Balance < value {
		mycontext.LoggerFromContext(ctx).Warnw("insufficient balance", "user", user.Login, "balance", user.Balance, "value", value, "reason", reason, "error", err)
		return fmt.Errorf("decrease: %w", &errors2.InsufficientBalanceError{
			Want: value,
			Got:  user.Balance,
		})
	}
	user.Balance -= value

	err = bt.userRepo.Update(ctx, user)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot update user for decrease transaction", "user", user.Login, "value", value, "reason", reason, "error", err)
		return fmt.Errorf("update: %w", err)
	}

	t := &models.BalanceTransaction{
		Reason: reason,
		Time:   time.Now().UTC(),
		Amount: -value,
		UserID: user.UUID,
	}

	err = bt.balanceTransactionRepo.Create(ctx, t)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot create decrease transaction", "user", user.Login, "value", t.Amount, "reason", t.Reason, "time", t.Time, "error", err)
		return fmt.Errorf("create: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("successful decrease transaction", "user", user.Login, "value", value)

	return nil
}

func (bt *BalanceTransaction) IncreaseSub(ctx context.Context, value int, userID uuid.UUID) error {
	mycontext.LoggerFromContext(ctx).Infow("started increase sub transaction", "to user", userID, "value", value)

	user, err := bt.userRepo.GetByID(ctx, userID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user by id for increase sub transaction", "to user", userID, "value", value, "error", err)
		return fmt.Errorf("update: %w", err)
	}

	user.Balance += value

	err = bt.userRepo.Update(ctx, user)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot update user for increase sub transaction", "to user", user.Login, "value", value, "error", err)
		return fmt.Errorf("update: %w", err)
	}

	t := &models.BalanceTransaction{
		Reason: "paid subscription: new subscriber",
		Time:   time.Now().UTC(),
		Amount: value,
		UserID: user.UUID,
	}

	err = bt.balanceTransactionRepo.Create(ctx, t)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot create increase sub transaction", "user", user.Login, "value", t.Amount, "reason", t.Reason, "time", t.Time, "error", err)
		return fmt.Errorf("create: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("successful increase sub transaction", "user", user.Login, "value", t.Amount, "reason", t.Reason, "time", t.Time)

	return nil
}
