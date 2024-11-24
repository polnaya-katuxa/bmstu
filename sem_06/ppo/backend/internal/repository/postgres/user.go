package postgres

import (
	"context"
	"errors"
	"fmt"

	errors2 "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/errors"
	logicModels "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	repoModels "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/repository/postgres/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUR(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) GetByLogin(ctx context.Context, login string) (*logicModels.User, error) {
	user := repoModels.User{}

	res := u.db.WithContext(ctx).Table("users").Where("login = ?", login).Take(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, toErrorResponse(fmt.Errorf("select: %w", errors2.ErrNotFound))
	}
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

func (u *UserRepository) GetByID(ctx context.Context, uuid uuid.UUID) (*logicModels.User, error) {
	user := repoModels.User{}

	res := u.db.WithContext(ctx).Table("users").Where("uuid = ?", uuid).Take(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, toErrorResponse(fmt.Errorf("select: %w", errors2.ErrNotFound))
	}
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

func (u *UserRepository) Delete(ctx context.Context, login string) error {
	res := u.db.WithContext(ctx).Table("users").Where("login = ?", login).Delete(&repoModels.User{})
	if res.Error != nil {
		return toErrorResponse(fmt.Errorf("delete: %w", res.Error))
	}

	return nil
}

func (u *UserRepository) Update(ctx context.Context, user *logicModels.User) error {
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

	res := u.db.WithContext(ctx).Table("users").Save(userDB)
	if res.Error != nil {
		return toErrorResponse(fmt.Errorf("update: %w", res.Error))
	}

	return nil
}

func (u *UserRepository) GetTotal(ctx context.Context) (int, error) {
	num := 0
	res := u.db.WithContext(ctx).Table("users").Select("count(*)").Find(&num)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) || (res.RowsAffected == 0 && res.Error == nil) {
		return 0, nil
	}
	if res.Error != nil {
		return 0, toErrorResponse(fmt.Errorf("select: %w", res.Error))
	}

	return num, nil
}

func (u *UserRepository) GetAll(ctx context.Context, pg *logicModels.Paginator) ([]*logicModels.User, error) {
	var usersDB []*repoModels.User

	res := u.db.WithContext(ctx).Table("users").Order(
		"uuid desc")
	if pg != nil {
		res = res.Limit(pg.Num).Offset(pg.Num * (pg.Page - 1))
	}
	res = res.Find(&usersDB)
	if res.Error != nil {
		return nil, toErrorResponse(fmt.Errorf("select: %w", res.Error))
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

	res := u.db.WithContext(ctx).Table("users").Create(userDB)
	if res.Error != nil {
		return nil, toErrorResponse(fmt.Errorf("insert: %w", res.Error))
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
