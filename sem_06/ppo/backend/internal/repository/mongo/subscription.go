package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	logicModels "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	repoModels "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/repository/mongo/models"
	"github.com/google/uuid"
)

type SubscriptionRepository struct {
	db *mongo.Database
}

func NewSR(db *mongo.Database) SubscriptionRepository {
	return SubscriptionRepository{db: db}
}

func (s SubscriptionRepository) Create(ctx context.Context, uuid1 uuid.UUID, uuid2 uuid.UUID) error {
	coll := s.db.Collection("subscriptions")
	sub := &repoModels.Subscription{
		UUID:     uuid.New(),
		ReaderID: uuid1,
		WriterID: uuid2,
	}

	_, err := coll.InsertOne(ctx, sub)
	if err != nil {
		return fmt.Errorf("insert: %w", err)
	}

	return nil
}

func (s SubscriptionRepository) Delete(ctx context.Context, uuid1 uuid.UUID, uuid2 uuid.UUID) error {
	coll := s.db.Collection("subscriptions")

	filter := bson.D{
		{
			"$and",
			bson.A{
				bson.D{{"reader_id", bson.D{{"$eq", uuid1}}}},
				bson.D{{"writer_id", bson.D{{"$eq", uuid2}}}},
			},
		},
	}

	_, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}

func (s SubscriptionRepository) GetAll(ctx context.Context, uuid uuid.UUID) ([]*logicModels.Subscription, error) {
	coll1 := s.db.Collection("subscriptions")
	var subsDB []*repoModels.Subscription

	cur, err := coll1.Find(ctx, bson.D{{"reader_id", uuid}})
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	if err = cur.All(ctx, &subsDB); err != nil {
		return nil, fmt.Errorf("select: %w", err)
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
	coll1 := s.db.Collection("subscriptions")
	sub := repoModels.Subscription{}

	filter := bson.D{
		{
			"$and",
			bson.A{
				bson.D{{"reader_id", bson.D{{"$eq", uuid1}}}},
				bson.D{{"writer_id", bson.D{{"$eq", uuid2}}}},
			},
		},
	}

	err := coll1.FindOne(ctx, filter).Decode(&sub)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	resSub := logicModels.Subscription{
		UUID:     sub.UUID,
		ReaderID: sub.ReaderID,
		WriterID: sub.WriterID,
	}

	return &resSub, nil
}
