package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	myerrors "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/errors"
	logicModels "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	repoModels "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/repository/mongo/models"
	"github.com/google/uuid"
)

type PostRepository struct {
	db *mongo.Database
}

func NewPR(db *mongo.Database) PostRepository {
	return PostRepository{db: db}
}

func (p PostRepository) getReactions(ctx context.Context, postID uuid.UUID) ([]*logicModels.Reaction, error) {
	coll1 := p.db.Collection("reactions")
	var reactsDB []*repoModels.Reaction

	cur, err := coll1.Find(ctx, bson.D{{"post_id", postID}})
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	if err = cur.All(ctx, &reactsDB); err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	coll2 := p.db.Collection("reaction_types")
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
			if r.ReactionTypeID == rt.UUID {
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

func (p PostRepository) getUser(ctx context.Context, uuid uuid.UUID) (*logicModels.User, error) {
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

func (p PostRepository) getLimit(ctx context.Context, uuid uuid.UUID) (*logicModels.Limit, error) {
	coll := p.db.Collection("limits")
	limit := repoModels.Limit{}

	err := coll.FindOne(ctx, bson.D{{"_id", uuid}}).Decode(&limit)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("select: %w", myerrors.ErrNotFound)
	}
	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	resLimit := logicModels.Limit{
		UUID:  limit.UUID,
		Value: limit.Value,
		Bonus: limit.Bonus,
	}

	return &resLimit, nil
}

func (p PostRepository) GetAll(ctx context.Context, uuid uuid.UUID) ([]*logicModels.Post, error) {
	coll1 := p.db.Collection("posts")
	var postsDB []*repoModels.Post

	cur, err := coll1.Find(ctx, bson.D{{"writer_id", uuid}}, options.Find().SetSort(bson.D{{"public_time", -1}}))
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	if err = cur.All(ctx, &postsDB); err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	postsLogic := make([]*logicModels.Post, 0, len(postsDB))
	for _, post := range postsDB {
		coll2 := p.db.Collection("comments")

		num, err := coll2.CountDocuments(ctx, bson.D{{"post_id", post.UUID}})
		if err != nil {
			return nil, fmt.Errorf("count: %w", err)
		}

		lPost, err := p.toPost(ctx, post)
		if err != nil {
			return nil, fmt.Errorf("to post: %w", err)
		}

		lPost.CommentNum = int(num)

		postsLogic = append(postsLogic, lPost)
	}

	return postsLogic, nil
}

func (p PostRepository) Create(ctx context.Context, post *logicModels.Post) (*logicModels.Post, error) {
	coll := p.db.Collection("posts")

	uuid := uuid.New()
	postDB := &repoModels.Post{
		UUID:       uuid,
		Content:    post.Content,
		PublicTime: post.PublicTime,
		WriterID:   post.Author.UUID,
		LimitID:    post.NextLimit.UUID,
	}

	if post.Perms == logicModels.Free {
		postDB.Perms = "free"
	} else {
		postDB.Perms = "paid"
	}

	_, err := coll.InsertOne(ctx, postDB)
	if err != nil {
		return nil, fmt.Errorf("insert: %w", err)
	}

	created, err := p.Get(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	return created, nil
}

func (p PostRepository) Update(ctx context.Context, post *logicModels.Post) error {
	coll := p.db.Collection("posts")
	postDB := &repoModels.Post{
		UUID:       post.UUID,
		Content:    post.Content,
		PublicTime: post.PublicTime,
		WriterID:   post.Author.UUID,
		LimitID:    post.NextLimit.UUID,
	}

	if post.Perms == logicModels.Free {
		postDB.Perms = "free"
	} else {
		postDB.Perms = "paid"
	}

	_, err := coll.ReplaceOne(ctx, bson.D{{"_id", post.UUID}}, postDB)
	if err != nil {
		return fmt.Errorf("update: %w", err)
	}

	return nil
}

func (p PostRepository) Delete(ctx context.Context, uuid uuid.UUID) error {
	coll := p.db.Collection("posts")

	_, err := coll.DeleteOne(ctx, bson.D{{"_id", uuid}})
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	coll1 := p.db.Collection("reactions")
	_, err = coll1.DeleteMany(ctx, bson.D{{"post_id", uuid}})
	if err != nil {
		return fmt.Errorf("delete reacts: %w", err)
	}

	coll2 := p.db.Collection("comments")
	_, err = coll2.DeleteMany(ctx, bson.D{{"post_id", uuid}})
	if err != nil {
		return fmt.Errorf("delete comments: %w", err)
	}

	return nil
}

func (p PostRepository) Get(ctx context.Context, uuid uuid.UUID) (*logicModels.Post, error) {
	coll1 := p.db.Collection("posts")
	post := repoModels.Post{}

	err := coll1.FindOne(ctx, bson.D{{"_id", uuid}}).Decode(&post)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("select: %w", myerrors.ErrNotFound)
	}
	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	coll2 := p.db.Collection("comments")

	num, err := coll2.CountDocuments(ctx, bson.D{{"post_id", post.UUID}})
	if err != nil {
		return nil, fmt.Errorf("count: %w", err)
	}

	resPost, err := p.toPost(ctx, &post)
	if err != nil {
		return nil, fmt.Errorf("to post: %w", err)
	}

	resPost.CommentNum = int(num)

	return resPost, nil
}

func (p PostRepository) GetByIDAndSpan(ctx context.Context, uuids []uuid.UUID, time1 time.Time, time2 time.Time, isIn bool) ([]*logicModels.Post, error) {
	coll1 := p.db.Collection("posts")
	var postsDB []*repoModels.Post

	if isIn {
		filter := bson.D{
			{
				"$and",
				bson.A{
					bson.D{{"public_time", bson.D{{"$gte", time1}}}},
					bson.D{{"public_time", bson.D{{"$lte", time2}}}},
					bson.D{{"writer_id", bson.D{{"$in", uuids}}}},
				},
			},
		}

		cur, err := coll1.Find(ctx, filter, options.Find().SetSort(bson.D{{"public_time", -1}}))
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		if err != nil {
			return nil, fmt.Errorf("select: %w", err)
		}

		if err = cur.All(ctx, &postsDB); err != nil {
			return nil, fmt.Errorf("select: %w", err)
		}
	} else {
		if len(uuids) == 0 {
			filter := bson.D{
				{
					"$and",
					bson.A{
						bson.D{{"public_time", bson.D{{"$gte", time1}}}},
						bson.D{{"public_time", bson.D{{"$lte", time2}}}},
						bson.D{{"perms", "free"}},
					},
				},
			}

			cur, err := coll1.Find(ctx, filter, options.Find().SetSort(bson.D{{"public_time", -1}}))
			if err == mongo.ErrNoDocuments {
				return nil, nil
			}
			if err != nil {
				return nil, fmt.Errorf("select: %w", err)
			}

			if err = cur.All(ctx, &postsDB); err != nil {
				return nil, fmt.Errorf("select: %w", err)
			}
		} else {
			filter := bson.D{
				{
					"$and",
					bson.A{
						bson.D{{"public_time", bson.D{{"$gte", time1}}}},
						bson.D{{"public_time", bson.D{{"$lte", time2}}}},
						bson.D{{"writer_id", bson.D{{"$not", bson.D{{"$in", uuids}}}}}},
						bson.D{{"perms", "free"}},
					},
				},
			}

			cur, err := coll1.Find(ctx, filter, options.Find().SetSort(bson.D{{"public_time", -1}}))
			if err == mongo.ErrNoDocuments {
				return nil, nil
			}
			if err != nil {
				return nil, fmt.Errorf("select: %w", err)
			}

			if err = cur.All(ctx, &postsDB); err != nil {
				return nil, fmt.Errorf("select: %w", err)
			}
		}
	}

	postsLogic := make([]*logicModels.Post, 0, len(postsDB))
	for _, post := range postsDB {
		coll2 := p.db.Collection("comments")

		num, err := coll2.CountDocuments(ctx, bson.D{{"post_id", post.UUID}})
		if err != nil {
			return nil, fmt.Errorf("count: %w", err)
		}

		lPost, err := p.toPost(ctx, post)
		if err != nil {
			return nil, fmt.Errorf("to post: %w", err)
		}

		lPost.CommentNum = int(num)

		postsLogic = append(postsLogic, lPost)
	}

	return postsLogic, nil
}

func (p PostRepository) toPost(ctx context.Context, post *repoModels.Post) (*logicModels.Post, error) {
	var perm logicModels.Perms
	if post.Perms == "free" {
		perm = logicModels.Free
	} else {
		perm = logicModels.Paid
	}

	reacts, err := p.getReactions(ctx, post.UUID)
	if err != nil {
		return nil, fmt.Errorf("get reactions: %w", err)
	}

	author, err := p.getUser(ctx, post.WriterID)
	if err != nil {
		return nil, fmt.Errorf("get author: %w", err)
	}

	limit, err := p.getLimit(ctx, post.LimitID)
	if err != nil {
		return nil, fmt.Errorf("get limit: %w", err)
	}

	return &logicModels.Post{
		UUID:       post.UUID,
		Content:    post.Content,
		PublicTime: post.PublicTime,
		Perms:      perm,
		Reactions:  reacts,
		Author:     author,
		NextLimit:  *limit,
	}, nil
}

func (p PostRepository) GetSortedLimits(ctx context.Context) ([]*logicModels.Limit, error) {
	coll := p.db.Collection("limits")
	var limitsDB []*repoModels.Limit

	cur, err := coll.Find(ctx, bson.D{}, options.Find().SetSort(bson.D{{"value", 1}}))
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("select: %w", myerrors.ErrNotFound)
	}
	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	if err = cur.All(ctx, &limitsDB); err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	limitsLogic := make([]*logicModels.Limit, 0, len(limitsDB))
	for _, l := range limitsDB {
		lim := &logicModels.Limit{
			UUID:  l.UUID,
			Value: l.Value,
			Bonus: l.Bonus,
		}
		limitsLogic = append(limitsLogic, lim)
	}

	return limitsLogic, nil
}

func (p PostRepository) GetReactionTypes(ctx context.Context) ([]*logicModels.ReactionType, error) {
	coll := p.db.Collection("reaction_types")
	var rtDB []*repoModels.ReactionType

	cur, err := coll.Find(ctx, bson.D{})
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("select: %w", myerrors.ErrNotFound)
	}
	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	if err = cur.All(ctx, &rtDB); err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	rtLogic := make([]*logicModels.ReactionType, 0, len(rtDB))
	for _, rt := range rtDB {
		rt := &logicModels.ReactionType{
			UUID: rt.UUID,
			Icon: rt.Icon,
		}
		rtLogic = append(rtLogic, rt)
	}

	return rtLogic, nil
}
