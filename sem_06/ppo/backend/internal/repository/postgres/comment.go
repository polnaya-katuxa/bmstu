package postgres

import (
	"context"
	"errors"
	"fmt"

	logicModels "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	repoModels "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/repository/postgres/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCR(db *gorm.DB) CommentRepository {
	return CommentRepository{db: db}
}

func (p CommentRepository) getUser(ctx context.Context, uuid uuid.UUID) (*logicModels.User, error) {
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

func (c CommentRepository) GetTotal(ctx context.Context, postID uuid.UUID) (int, error) {
	num := 0
	res := c.db.WithContext(ctx).Table("comments").Select("count(*)").Where("post_id = ?", postID).Find(&num)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) || (res.RowsAffected == 0 && res.Error == nil) {
		return 0, nil
	}
	if res.Error != nil {
		return 0, toErrorResponse(fmt.Errorf("select: %w", res.Error))
	}

	return num, nil
}

func (c CommentRepository) GetAll(ctx context.Context, postID uuid.UUID, pg *logicModels.Paginator) ([]*logicModels.Comment, error) {
	var commDB []*repoModels.Comment

	res := c.db.WithContext(ctx).Table("comments").Where("post_id = ?", postID).Order(
		"public_time desc")
	if pg != nil {
		res = res.Limit(pg.Num).Offset(pg.Num * (pg.Page - 1))
	}
	res = res.Find(&commDB)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) || (res.RowsAffected == 0 && res.Error == nil) {
		return nil, nil
	}
	if res.Error != nil {
		return nil, toErrorResponse(fmt.Errorf("select: %w", res.Error))
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
	var commDB *repoModels.Comment

	res := c.db.WithContext(ctx).Table("comments").Where("uuid = ?", commID).Take(&commDB)
	if res.Error != nil {
		return nil, toErrorResponse(fmt.Errorf("select: %w", res.Error))
	}

	u, err := c.getUser(ctx, commDB.CommentatorID)
	if err != nil {
		return nil, toErrorResponse(fmt.Errorf("get user: %w", err))
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
	id := uuid.New()
	commDB := &repoModels.Comment{
		UUID:          id,
		Content:       comm.Content,
		PublicTime:    comm.PublicTime,
		CommentatorID: comm.Commentator.UUID,
		PostID:        comm.PostID,
	}

	res := c.db.WithContext(ctx).Table("comments").Create(commDB)
	if res.Error != nil {
		return nil, toErrorResponse(fmt.Errorf("insert: %w", res.Error))
	}

	commLogic, err := c.GetByID(ctx, id)
	if err != nil {
		return nil, toErrorResponse(fmt.Errorf("get by id: %w", res.Error))
	}

	return commLogic, nil
}

func (c CommentRepository) Delete(ctx context.Context, commID uuid.UUID) error {
	res := c.db.WithContext(ctx).Table("comments").Where("uuid = ?", commID).Delete(&repoModels.Comment{})
	if res.Error != nil {
		return toErrorResponse(fmt.Errorf("delete: %w", res.Error))
	}

	return nil
}
