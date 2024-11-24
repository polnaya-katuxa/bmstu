package postgres

import (
	"context"
	"fmt"

	logicModels "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	repoModels "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/repository/postgres/models"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type SubscriptionRepository struct {
	db *gorm.DB
}

func NewSR(db *gorm.DB) SubscriptionRepository {
	return SubscriptionRepository{db: db}
}

func (s SubscriptionRepository) Create(ctx context.Context, uuid1 uuid.UUID, uuid2 uuid.UUID) error {
	sub := &repoModels.Subscription{
		UUID:     uuid.New(),
		ReaderID: uuid1,
		WriterID: uuid2,
	}

	res := s.db.WithContext(ctx).Table("subscriptions").Create(sub)
	if res.Error != nil {
		return toErrorResponse(fmt.Errorf("insert: %w", res.Error))
	}

	return nil
}

func (s SubscriptionRepository) Delete(ctx context.Context, uuid1 uuid.UUID, uuid2 uuid.UUID) error {
	res := s.db.WithContext(ctx).Exec("delete from subscriptions where reader_id = ? AND writer_id = ?", uuid1, uuid2)
	if res.Error != nil {
		return toErrorResponse(fmt.Errorf("delete: %w", res.Error))
	}

	return nil
}

func (s SubscriptionRepository) GetAll(ctx context.Context, uuid uuid.UUID) ([]*logicModels.Subscription, error) {
	var subsDB []*repoModels.Subscription

	res := s.db.WithContext(ctx).Table("subscriptions").Where("reader_id = ?", uuid).Find(&subsDB)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) || (res.RowsAffected == 0 && res.Error == nil) {
		return nil, nil
	}
	if res.Error != nil {
		return nil, toErrorResponse(fmt.Errorf("select: %w", res.Error))
	}

	subsLogic := make([]*logicModels.Subscription, 0, len(subsDB))
	for _, s := range subsDB {
		subsLogic = append(subsLogic, &logicModels.Subscription{
			UUID:     s.UUID,
			ReaderID: s.ReaderID,
			WriterID: s.WriterID,
		})
	}

	return subsLogic, nil
}

func (s SubscriptionRepository) Get(ctx context.Context, uuid1 uuid.UUID, uuid2 uuid.UUID) (*logicModels.Subscription, error) {
	sub := repoModels.Subscription{}

	res := s.db.WithContext(ctx).Table("subscriptions").Where("reader_id = ? AND writer_id = ?", uuid1, uuid2).Take(&sub)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) || (res.RowsAffected == 0 && res.Error == nil) {
		return nil, nil
	}
	if res.Error != nil {
		return nil, toErrorResponse(fmt.Errorf("select: %w", res.Error))
	}

	resSub := logicModels.Subscription{
		UUID:     sub.UUID,
		ReaderID: sub.ReaderID,
		WriterID: sub.WriterID,
	}

	return &resSub, nil
}
