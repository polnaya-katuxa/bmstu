package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	errors2 "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/errors"
	logicModels "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	repoModels "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/repository/mongo/models"
	"github.com/google/uuid"
)

type UserRepository struct {
	db *mongo.Database
}

func NewUR(db *mongo.Database) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) GetByLogin(ctx context.Context, login string) (*logicModels.User, error) {
	coll := u.db.Collection("users")
	user := repoModels.User{}

	err := coll.FindOne(ctx, bson.D{{"login", login}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("delete: %w", errors2.ErrNotFound)
	}
	if err != nil {
		return nil, fmt.Errorf("delete: %w", err)
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

func (u *UserRepository) GetByID(ctx context.Context, uuid uuid.UUID) (*logicModels.User, error) {
	coll := u.db.Collection("users")
	user := repoModels.User{}

	err := coll.FindOne(ctx, bson.D{{"_id", uuid}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("select: %w", errors2.ErrNotFound)
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

func (u *UserRepository) Delete(ctx context.Context, login string) error {
	coll := u.db.Collection("users")

	user, err := u.GetByLogin(ctx, login)
	if err != nil {
		return fmt.Errorf("get by login: %w", err)
	}

	_, err = coll.DeleteOne(ctx, bson.D{{"login", login}})
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	coll1 := u.db.Collection("subscriptions")
	_, err = coll1.DeleteMany(ctx, bson.D{{"reader_id", user.UUID}})
	if err != nil {
		return fmt.Errorf("delete subs: %w", err)
	}

	coll2 := u.db.Collection("reactions")
	_, err = coll2.DeleteMany(ctx, bson.D{{"reactor_id", user.UUID}})
	if err != nil {
		return fmt.Errorf("delete reacts: %w", err)
	}

	coll3 := u.db.Collection("posts")
	_, err = coll3.DeleteMany(ctx, bson.D{{"writer_id", user.UUID}})
	if err != nil {
		return fmt.Errorf("delete posts: %w", err)
	}

	coll4 := u.db.Collection("comments")
	_, err = coll4.DeleteMany(ctx, bson.D{{"commentator_id", user.UUID}})
	if err != nil {
		return fmt.Errorf("delete comments: %w", err)
	}

	return nil
}

func (u *UserRepository) Update(ctx context.Context, user *logicModels.User) error {
	coll := u.db.Collection("users")

	userDB := &repoModels.User{
		UUID:        user.UUID,
		Login:       user.Login,
		Picture:     user.Picture,
		Description: user.Description,
		Password:    user.Password,
		Balance:     user.Balance,
		Mail:        user.Mail,
		EnterTime:   user.EnterTime,
		IsAdmin:     user.IsAdmin,
	}

	_, err := coll.ReplaceOne(ctx, bson.D{{"_id", user.UUID}}, userDB)
	if err != nil {
		return fmt.Errorf("update: %w", err)
	}

	return nil
}

func (u *UserRepository) GetAll(ctx context.Context) ([]*logicModels.User, error) {
	coll := u.db.Collection("users")
	var usersDB []*repoModels.User

	cur, err := coll.Find(ctx, bson.D{})
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("select: %w", errors2.ErrNotFound)
	}
	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	if err = cur.All(ctx, &usersDB); err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	usersLogic := make([]*logicModels.User, 0, len(usersDB))
	for _, u := range usersDB {
		usersLogic = append(usersLogic, &logicModels.User{
			UUID:        u.UUID,
			Login:       u.Login,
			Password:    u.Password,
			Balance:     u.Balance,
			Mail:        u.Mail,
			EnterTime:   u.EnterTime,
			Picture:     u.Picture,
			Description: u.Description,
			IsAdmin:     u.IsAdmin,
		})
	}

	return usersLogic, nil
}

func (u *UserRepository) Create(ctx context.Context, user *logicModels.User) (*logicModels.User, error) {
	coll := u.db.Collection("users")
	userDB := &repoModels.User{
		UUID:        uuid.New(),
		Login:       user.Login,
		Picture:     user.Picture,
		Description: user.Description,
		Password:    user.Password,
		Balance:     user.Balance,
		Mail:        user.Mail,
		EnterTime:   user.EnterTime,
		IsAdmin:     user.IsAdmin,
	}

	_, err := coll.InsertOne(ctx, userDB)
	if err != nil {
		return nil, fmt.Errorf("insert: %w", err)
	}

	resUser := logicModels.User{
		UUID:        userDB.UUID,
		Login:       userDB.Login,
		Password:    userDB.Password,
		Balance:     userDB.Balance,
		Mail:        userDB.Mail,
		EnterTime:   userDB.EnterTime,
		Picture:     userDB.Picture,
		Description: userDB.Description,
		IsAdmin:     userDB.IsAdmin,
	}

	return &resUser, nil
}
