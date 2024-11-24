package interfaces

import (
	"context"
	"time"

	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	"github.com/google/uuid"
)

type IUserRepository interface {
	Create(context.Context, *models.User) (*models.User, error)
	GetByLogin(context.Context, string) (*models.User, error)
	GetByID(context.Context, uuid.UUID) (*models.User, error)
	Delete(context.Context, string) error
	Update(context.Context, *models.User) error
	GetAll(context.Context, *models.Paginator) ([]*models.User, error)
	GetTotal(context.Context) (int, error)
}

type ISubscriptionRepository interface {
	Create(context.Context, uuid.UUID, uuid.UUID) error
	Delete(context.Context, uuid.UUID, uuid.UUID) error
	GetAll(context.Context, uuid.UUID) ([]*models.Subscription, error)
	Get(context.Context, uuid.UUID, uuid.UUID) (*models.Subscription, error)
}

type IReactionRepository interface {
	Reacted(context.Context, uuid.UUID, uuid.UUID, uuid.UUID) (bool, error)
	GetAll(context.Context, uuid.UUID) ([]*models.Reaction, error)
	Create(context.Context, *models.Reaction, uuid.UUID) error
	Delete(context.Context, uuid.UUID, uuid.UUID, uuid.UUID) error
}

type ICommentRepository interface {
	Create(context.Context, *models.Comment) (*models.Comment, error)
	Delete(context.Context, uuid.UUID) error
	GetAll(context.Context, uuid.UUID, *models.Paginator) ([]*models.Comment, error)
	GetByID(context.Context, uuid.UUID) (*models.Comment, error)
	GetTotal(context.Context, uuid.UUID) (int, error)
}

type IPostRepository interface {
	GetAll(context.Context, uuid.UUID, *models.Paginator) ([]*models.Post, error)
	Create(context.Context, *models.Post) (*models.Post, error)
	Update(context.Context, *models.Post) error
	Delete(context.Context, uuid.UUID) error
	Get(context.Context, uuid.UUID) (*models.Post, error)
	GetByIDAndSpan(context.Context, []uuid.UUID, time.Time, time.Time, bool) ([]*models.Post, error)
	GetSortedLimits(context.Context) ([]*models.Limit, error)
	GetReactionTypes(context.Context) ([]*models.ReactionType, error)
	GetTotalByUserID(ctx context.Context, userID uuid.UUID) (int, error)
	GetTotalByIDAndSpan(context.Context, []uuid.UUID, time.Time, time.Time, bool) (int, error)
}

type IBalanceTransactionRepository interface {
	Create(context.Context, *models.BalanceTransaction) error
}
