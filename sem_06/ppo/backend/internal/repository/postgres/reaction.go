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

type ReactionRepository struct {
	db *gorm.DB
}

func NewRR(db *gorm.DB) ReactionRepository {
	return ReactionRepository{db: db}
}

func (r ReactionRepository) Reacted(ctx context.Context, uuid1 uuid.UUID, uuid2 uuid.UUID, uuid3 uuid.UUID) (bool, error) {
	react := repoModels.Reaction{}

	res := r.db.WithContext(ctx).Table("reactions").Where("reactor_id = ? AND post_id = ? AND reaction_type_id = ?", uuid1, uuid2, uuid3).Take(&react)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) || (res.RowsAffected == 0 && res.Error == nil) {
		return false, nil
	}
	if res.Error != nil {
		return false, toErrorResponse(fmt.Errorf("select: %w", res.Error))
	}

	return true, nil
}

func (r ReactionRepository) GetAll(ctx context.Context, uuid uuid.UUID) ([]*logicModels.Reaction, error) {
	var reactsDB []*repoModels.Reaction

	res := r.db.WithContext(ctx).Table("reactions").Where("post_id = ?", uuid).Find(&reactsDB)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) || (res.RowsAffected == 0 && res.Error == nil) {
		return nil, nil
	}
	if res.Error != nil {
		return nil, toErrorResponse(fmt.Errorf("select: %w", res.Error))
	}

	var reactTypesDB []*repoModels.ReactionType
	res = r.db.WithContext(ctx).Table("reaction_types").Find(&reactTypesDB)
	if res.Error != nil {
		return nil, toErrorResponse(fmt.Errorf("select: %w", res.Error))
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
	react := &repoModels.Reaction{
		UUID:           uuid.New(),
		ReactionTypeID: reaction.TypeID,
		ReactorID:      reaction.ReactorID,
		PostID:         id,
	}

	res := r.db.WithContext(ctx).Table("reactions").Create(react)
	if res.Error != nil {
		return toErrorResponse(fmt.Errorf("insert: %w", res.Error))
	}

	return nil
}

func (r ReactionRepository) Delete(ctx context.Context, uuid1 uuid.UUID, uuid2 uuid.UUID, uuid3 uuid.UUID) error {
	res := r.db.WithContext(ctx).Table("reactions").Where("reactor_id = ? AND post_id = ? AND reaction_type_id = ?", uuid1, uuid2, uuid3).Delete(&repoModels.Reaction{})
	if res.Error != nil {
		return toErrorResponse(fmt.Errorf("delete: %w", res.Error))
	}

	return nil
}
