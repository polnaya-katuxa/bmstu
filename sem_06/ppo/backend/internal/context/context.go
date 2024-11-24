package context

import (
	"context"
	"fmt"

	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/errors"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	"go.uber.org/zap"
)

type (
	keyUser struct{}
	keyLog  struct{}
)

func UserToContext(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, keyUser{}, user)
}

func UserFromContext(ctx context.Context) (*models.User, error) {
	user := ctx.Value(keyUser{})
	if user == nil {
		return nil, fmt.Errorf("get user: %w", errors.ErrGet)
	}

	return user.(*models.User), nil
}

func LoggerToContext(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, keyLog{}, logger)
}

func LoggerFromContext(ctx context.Context) *zap.SugaredLogger {
	v := ctx.Value(keyLog{})
	if v == nil {
		return zap.NewNop().Sugar()
	}
	return v.(*zap.SugaredLogger)
}
