package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	errors2 "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/errors"
	logicModels "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	repoModels "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/repository/postgres/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPR(db *gorm.DB) PostRepository {
	return PostRepository{db: db}
}

func (p PostRepository) getReactions(ctx context.Context, uuid uuid.UUID) ([]*logicModels.Reaction, error) {
	var reactsDB []*repoModels.Reaction

	res := p.db.WithContext(ctx).Table("reactions").Where("post_id = ?", uuid).Find(&reactsDB)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) || (res.RowsAffected == 0 && res.Error == nil) {
		return nil, nil
	}
	if res.Error != nil {
		return nil, toErrorResponse(fmt.Errorf("select: %w", res.Error))
	}

	var reactTypesDB []*repoModels.ReactionType
	res = p.db.WithContext(ctx).Table("reaction_types").Find(&reactTypesDB)
	if res.Error != nil {
		return nil, toErrorResponse(fmt.Errorf("select: %w", res.Error))
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
	user := repoModels.User{}

	res := p.db.WithContext(ctx).Table("users").Where("uuid = ?", uuid).Take(&user)
	if res.Error != nil {
		return nil, toErrorResponse(fmt.Errorf("select: %w", res.Error))
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
	limit := repoModels.Limit{}

	res := p.db.WithContext(ctx).Table("limits").Where("uuid = ?", uuid).Take(&limit)
	if res.Error != nil {
		return nil, toErrorResponse(fmt.Errorf("select: %w", res.Error))
	}

	resLimit := logicModels.Limit{
		UUID:  limit.UUID,
		Value: limit.Value,
		Bonus: limit.Bonus,
	}

	return &resLimit, nil
}

func (p PostRepository) GetTotalByUserID(ctx context.Context, userID uuid.UUID) (int, error) {
	num := 0
	res := p.db.WithContext(ctx).Table("posts").Select("count(*)").Where("writer_id = ?", userID).Find(&num)
	if res.Error != nil {
		return 0, toErrorResponse(fmt.Errorf("select: %w", res.Error))
	}

	return num, nil
}

func (p PostRepository) GetAll(ctx context.Context, uuid uuid.UUID, pg *logicModels.Paginator) ([]*logicModels.Post, error) {
	var postsDB []*repoModels.Post

	res := p.db.WithContext(ctx).Table("posts").Where("writer_id = ?", uuid).Order(
		"public_time desc")
	if pg != nil {
		res = res.Limit(pg.Num).Offset(pg.Num * (pg.Page - 1))
	}
	res = res.Find(&postsDB)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) || (res.RowsAffected == 0 && res.Error == nil) {
		return nil, nil
	}
	if res.Error != nil {
		return nil, toErrorResponse(fmt.Errorf("select: %w", res.Error))
	}

	postsLogic := make([]*logicModels.Post, 0, len(postsDB))
	for _, post := range postsDB {
		num := 0
		res = p.db.WithContext(ctx).Table("comments").Select("count(*)").Where("post_id = ?", post.UUID).Find(&num)
		if res.Error != nil {
			return nil, toErrorResponse(fmt.Errorf("select: %w", res.Error))
		}

		lPost, err := p.toPost(ctx, post)
		if err != nil {
			return nil, toErrorResponse(fmt.Errorf("to post: %w", err))
		}

		lPost.CommentNum = num

		postsLogic = append(postsLogic, lPost)
	}

	return postsLogic, nil
}

func (p PostRepository) Create(ctx context.Context, post *logicModels.Post) (*logicModels.Post, error) {
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

	res := p.db.WithContext(ctx).Table("posts").Create(postDB)
	if res.Error != nil {
		return nil, toErrorResponse(fmt.Errorf("insert: %w", res.Error))
	}

	created, err := p.Get(ctx, uuid)
	if err != nil {
		return nil, toErrorResponse(fmt.Errorf("get: %w", err))
	}

	return created, nil
}

func (p PostRepository) Update(ctx context.Context, post *logicModels.Post) error {
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

	res := p.db.WithContext(ctx).Table("posts").Save(postDB)
	if res.Error != nil {
		return toErrorResponse(fmt.Errorf("update: %w", res.Error))
	}

	return nil
}

func (p PostRepository) Delete(ctx context.Context, uuid uuid.UUID) error {
	res := p.db.WithContext(ctx).Table("posts").Where("uuid = ?", uuid).Delete(&repoModels.Post{})
	if res.Error != nil {
		return toErrorResponse(fmt.Errorf("delete: %w", res.Error))
	}

	return nil
}

func (p PostRepository) Get(ctx context.Context, uuid uuid.UUID) (*logicModels.Post, error) {
	post := repoModels.Post{}

	res := p.db.WithContext(ctx).Table("posts").Where("uuid = ?", uuid).Take(&post)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("select: %w", errors2.ErrNotFound)
	}
	if res.Error != nil {
		return nil, toErrorResponse(fmt.Errorf("select: %w", res.Error))
	}

	num := 0
	res = p.db.WithContext(ctx).Table("comments").Select("count(*)").Where("post_id = ?", uuid).Find(&num)
	if res.Error != nil {
		return nil, toErrorResponse(fmt.Errorf("select: %w", res.Error))
	}

	resPost, err := p.toPost(ctx, &post)
	if err != nil {
		return nil, toErrorResponse(fmt.Errorf("to post: %w", err))
	}

	resPost.CommentNum = num

	return resPost, nil
}

func (p PostRepository) GetTotalByIDAndSpan(ctx context.Context, uuids []uuid.UUID, time1 time.Time, time2 time.Time, isIn bool) (int, error) {
	num := 0
	var res *gorm.DB

	if isIn {
		res = p.db.WithContext(ctx).Table("posts").Where("public_time BETWEEN ? AND ? AND writer_id IN ?",
			time1, time2, uuids)
	} else {
		if len(uuids) == 0 {
			res = p.db.WithContext(ctx).Table("posts").Where("public_time BETWEEN ? AND ? AND perms = 'free'",
				time1, time2)
		} else {
			res = p.db.WithContext(ctx).Table("posts").Where("public_time BETWEEN ? AND ? AND writer_id "+
				"NOT IN ? AND perms = 'free'", time1, time2, uuids)
		}
	}

	res = res.Select("count(*)").Find(&num)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) || (res.RowsAffected == 0 && res.Error == nil) {
		return 0, nil
	}
	if res.Error != nil {
		return 0, toErrorResponse(fmt.Errorf("select: %w", res.Error))
	}

	return num, nil
}

func (p PostRepository) GetByIDAndSpan(ctx context.Context, uuids []uuid.UUID, time1 time.Time, time2 time.Time, isIn bool) ([]*logicModels.Post, error) {
	var postsDB []*repoModels.Post
	var res *gorm.DB

	if isIn {
		res = p.db.WithContext(ctx).Table("posts").Where("public_time BETWEEN ? AND ? AND writer_id IN ?",
			time1, time2, uuids)
	} else {
		if len(uuids) == 0 {
			res = p.db.WithContext(ctx).Table("posts").Where("public_time BETWEEN ? AND ? AND perms = 'free'",
				time1, time2)
		} else {
			res = p.db.WithContext(ctx).Table("posts").Where("public_time BETWEEN ? AND ? AND writer_id "+
				"NOT IN ? AND perms = 'free'", time1, time2, uuids)
		}
	}

	res = res.Order("public_time desc").Find(&postsDB)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) || (res.RowsAffected == 0 && res.Error == nil) {
		return nil, nil
	}
	if res.Error != nil {
		return nil, toErrorResponse(fmt.Errorf("select: %w", res.Error))
	}

	postsLogic := make([]*logicModels.Post, 0, len(postsDB))
	for _, post := range postsDB {
		num := 0
		res = p.db.WithContext(ctx).Table("comments").Select("count(*)").Where("post_id = ?", post.UUID).Find(&num)
		if res.Error != nil {
			return nil, toErrorResponse(fmt.Errorf("select: %w", res.Error))
		}

		lPost, err := p.toPost(ctx, post)
		if err != nil {
			return nil, toErrorResponse(fmt.Errorf("to post: %w", err))
		}

		lPost.CommentNum = num

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
		return nil, toErrorResponse(fmt.Errorf("get reactions: %w", err))
	}

	author, err := p.getUser(ctx, post.WriterID)
	if err != nil {
		return nil, toErrorResponse(fmt.Errorf("get author: %w", err))
	}

	limit, err := p.getLimit(ctx, post.LimitID)
	if err != nil {
		return nil, toErrorResponse(fmt.Errorf("get limit: %w", err))
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
	var limitsDB []*repoModels.Limit

	res := p.db.WithContext(ctx).Table("limits").Order("value asc").Find(&limitsDB)
	if res.Error != nil {
		return nil, toErrorResponse(fmt.Errorf("select: %w", res.Error))
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
	var rtDB []*repoModels.ReactionType

	res := p.db.WithContext(ctx).Table("reaction_types").Find(&rtDB)
	if res.Error != nil {
		return nil, toErrorResponse(fmt.Errorf("select: %w", res.Error))
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
