package mongo

import (
	"context"
	"fmt"

	logicModels "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	repoModels "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/repository/mongo/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReactionRepository struct {
	db *mongo.Database
}

func NewRR(db *mongo.Database) ReactionRepository {
	return ReactionRepository{db: db}
}

func (r ReactionRepository) Reacted(ctx context.Context, uuid1 uuid.UUID, uuid2 uuid.UUID) (bool, error) {
	coll1 := r.db.Collection("reactions")
	react := repoModels.Reaction{}

	filter := bson.D{
		{
			"$and",
			bson.A{
				bson.D{{"reactor_id", bson.D{{"$eq", uuid1}}}},
				bson.D{{"post_id", bson.D{{"$eq", uuid2}}}},
			},
		},
	}

	err := coll1.FindOne(ctx, filter).Decode(&react)
	if err == mongo.ErrNoDocuments {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("select: %w", err)
	}

	return true, nil
}

func (r ReactionRepository) GetAll(ctx context.Context, uuid uuid.UUID) ([]*logicModels.Reaction, error) {
	coll1 := r.db.Collection("reactions")
	var reactsDB []*repoModels.Reaction

	cur, err := coll1.Find(ctx, bson.D{{"post_id", uuid}})
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	if err = cur.All(ctx, &reactsDB); err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	coll2 := r.db.Collection("reaction_types")
	var reactTypesDB []*repoModels.ReactionType

	cur, err = coll2.Find(ctx, bson.D{})
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	if err = cur.All(ctx, &reactTypesDB); err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	reactsLogic := make([]*logicModels.Reaction, 0, len(reactsDB))
	for _, r := range reactsDB {
		var icon string
		for _, rt := range reactTypesDB {
			if r.UUID == rt.UUID {
				icon = rt.Icon
			}
		}
		reactsLogic = append(reactsLogic, &logicModels.Reaction{
			UUID:      r.UUID,
			Icon:      icon,
			TypeID:    r.ReactionTypeID,
			ReactorID: r.ReactorID,
		})
	}

	return reactsLogic, nil
}

func (r ReactionRepository) Create(ctx context.Context, reaction *logicModels.Reaction, id uuid.UUID) error {
	coll := r.db.Collection("reactions")
	react := &repoModels.Reaction{
		UUID:           uuid.New(),
		ReactionTypeID: reaction.TypeID,
		ReactorID:      reaction.ReactorID,
		PostID:         id,
	}

	_, err := coll.InsertOne(ctx, react)
	if err != nil {
		return fmt.Errorf("insert: %w", err)
	}

	return nil
}

func (r ReactionRepository) Delete(ctx context.Context, uuid1 uuid.UUID, uuid2 uuid.UUID) error {
	coll := r.db.Collection("reactions")

	filter := bson.D{
		{
			"$and",
			bson.A{
				bson.D{{"reactor_id", bson.D{{"$eq", uuid1}}}},
				bson.D{{"post_id", bson.D{{"$eq", uuid2}}}},
			},
		},
	}

	_, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}
