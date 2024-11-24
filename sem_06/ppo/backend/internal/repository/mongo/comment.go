package mongo

import (
	"context"
	"fmt"

	myerrors "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/errors"
	logicModels "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	repoModels "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/repository/mongo/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CommentRepository struct {
	db *mongo.Database
}

func NewCR(db *mongo.Database) CommentRepository {
	return CommentRepository{db: db}
}

func (p CommentRepository) getUser(ctx context.Context, uuid uuid.UUID) (*logicModels.User, error) {
	coll := p.db.Collection("users")

	user := repoModels.User{}

	err := coll.FindOne(ctx, bson.D{{"_id", uuid}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("select: %w", myerrors.ErrNotFound)
	}
	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	resUser := logicModels.User{
		UUID:        user.UUID,
		Login:       user.Login,
		Password:    user.Password,
		Balance:     user.Balance,
		Mail:        user.Mail,
		EnterTime:   user.EnterTime,
		Picture:     user.Picture,
		Description: user.Description,
		IsAdmin:     user.IsAdmin,
	}

	return &resUser, nil
}

func (c CommentRepository) GetAll(ctx context.Context, postID uuid.UUID) ([]*logicModels.Comment, error) {
	coll := c.db.Collection("comments")

	var commDB []*repoModels.Comment

	cur, err := coll.Find(ctx, bson.D{{"post_id", postID}}, options.Find().SetSort(bson.D{{"public_time", -1}}))
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	if err = cur.All(ctx, &commDB); err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	commLogic := make([]*logicModels.Comment, 0, len(commDB))
	for _, r := range commDB {
		u, err := c.getUser(ctx, r.CommentatorID)
		if err != nil {
			return nil, fmt.Errorf("get user: %w", err)
		}

		commLogic = append(commLogic, &logicModels.Comment{
			UUID:        r.UUID,
			Content:     r.Content,
			PublicTime:  r.PublicTime,
			Commentator: u,
			PostID:      r.PostID,
		})
	}

	return commLogic, nil
}

func (c CommentRepository) GetByID(ctx context.Context, commID uuid.UUID) (*logicModels.Comment, error) {
	coll := c.db.Collection("comments")

	var commDB *repoModels.Comment

	err := coll.FindOne(ctx, bson.D{{"_id", commID}}).Decode(&commDB)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("select: %w", myerrors.ErrNotFound)
	}
	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	u, err := c.getUser(ctx, commDB.CommentatorID)
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}

	commLogic := &logicModels.Comment{
		UUID:        commDB.UUID,
		Content:     commDB.Content,
		PublicTime:  commDB.PublicTime,
		Commentator: u,
		PostID:      commDB.PostID,
	}

	return commLogic, nil
}

func (c CommentRepository) Create(ctx context.Context, comm *logicModels.Comment) (*logicModels.Comment, error) {
	coll := c.db.Collection("comments")

	id := uuid.New()
	commDB := &repoModels.Comment{
		UUID:          id,
		Content:       comm.Content,
		PublicTime:    comm.PublicTime,
		CommentatorID: comm.Commentator.UUID,
		PostID:        comm.PostID,
	}

	_, err := coll.InsertOne(ctx, commDB)
	if err != nil {
		return nil, fmt.Errorf("insert: %w", err)
	}

	commLogic, err := c.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get by id: %w", err)
	}

	return commLogic, nil
}

func (c CommentRepository) Delete(ctx context.Context, commID uuid.UUID) error {
	coll := c.db.Collection("comments")

	_, err := coll.DeleteOne(ctx, bson.D{{"_id", commID}})
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}
