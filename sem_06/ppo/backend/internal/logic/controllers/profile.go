package controllers

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"time"

	mycontext "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/context"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/interfaces"
	errors2 "git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/errors"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/logic/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Profile struct {
	userRepo         interfaces.IUserRepository
	subscriptionRepo interfaces.ISubscriptionRepository

	balanceTransactionLogic interfaces.IBalanceTransactionLogic
	postLogic               interfaces.IPostLogic

	dailyBonus int

	tokenExpiration time.Duration
	secretKey       string
}

func NewPRL(u interfaces.IUserRepository, bt interfaces.IBalanceTransactionLogic, s interfaces.ISubscriptionRepository,
	p interfaces.IPostLogic, db int, te time.Duration, sk string,
) *Profile {
	return &Profile{
		userRepo:                u,
		subscriptionRepo:        s,
		balanceTransactionLogic: bt,
		postLogic:               p,
		dailyBonus:              db,
		tokenExpiration:         te,
		secretKey:               sk,
	}
}

func (p *Profile) AuthByToken(ctx context.Context, token string) (*models.User, error) {
	mycontext.LoggerFromContext(ctx).Infow("authentification started", "token", token)

	var claims jwt.RegisteredClaims
	_, err := jwt.ParseWithClaims(token, &claims, func(_ *jwt.Token) (interface{}, error) {
		return []byte(p.secretKey), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			mycontext.LoggerFromContext(ctx).Warnw("authentification failed: expired", "token", token, "error", err)
		} else {
			mycontext.LoggerFromContext(ctx).Errorw("authentification failed", "token", token, "error", err)
		}
		return nil, fmt.Errorf("authentification by token: %w", err)
	}

	id, err := uuid.Parse(claims.ID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Errorw("cannot parse token claims", "token", token, "error", err)
		return nil, fmt.Errorf("parse id: %w", err)
	}

	user, err := p.userRepo.GetByID(ctx, id)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Errorw("cannot get user by id", "token", token, "id", id, "error", err)
		return nil, fmt.Errorf("get: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("authentification succeeded", "user", user.Login)

	return user, nil
}

func (p *Profile) Register(ctx context.Context, user *models.User, password string) (string, error) {
	mycontext.LoggerFromContext(ctx).Infow("started registration", "user", user.Login)

	userOld, err := p.userRepo.GetByLogin(ctx, user.Login)
	if err != nil && !errors.Is(err, errors2.ErrNotFound) {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user by login", "user", user.Login, "error", err)
		return "", fmt.Errorf("get: %w", err)
	}

	if userOld != nil {
		mycontext.LoggerFromContext(ctx).Warnw("user already exists", "user", user.Login)
		return "", fmt.Errorf("already exists: %w", errors2.ErrExists)
	}

	if len(user.Login) < 4 {
		mycontext.LoggerFromContext(ctx).Errorw("incorrect login", "user", user.Login)
		return "", fmt.Errorf("login: %w", errors2.ErrLogin)
	}

	if len(password) < 8 {
		mycontext.LoggerFromContext(ctx).Errorw("incorrect password", "user", user.Login)
		return "", fmt.Errorf("password: %w", errors2.ErrPassword)
	}

	_, err = mail.ParseAddress(user.Mail)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Errorw("incorrect email", "user", user.Login, "error", err)
		return "", fmt.Errorf("email: %w", errors2.ErrEmail)
	}

	user.EnterTime = time.Now().UTC()

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.Password = string(hash)

	user, err = p.userRepo.Create(ctx, user)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Errorw("cannot create user", "user", user.Login, "error", err)
		return "", fmt.Errorf("register: %w", err)
	}

	ctx = mycontext.UserToContext(ctx, user)

	s := fmt.Sprintf("daily %s bonus", time.Now().UTC().Format(time.RFC822))
	err = p.balanceTransactionLogic.Increase(ctx, p.dailyBonus, s)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot add daily bonus", "user", user.Login, "error", err)
		return "", fmt.Errorf("increase: %w", err)
	}

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(p.tokenExpiration)),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ID:        user.UUID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString([]byte(p.secretKey))

	mycontext.LoggerFromContext(ctx).Infow("registration succeeded", "user", user.Login)

	return ss, nil
}

func areEqualDays(t1 time.Time, t2 time.Time) bool {
	d1, m1, y1 := t1.Date()
	d2, m2, y2 := t2.Date()

	if (d1 == d2) && (m1 == m2) && (y1 == y2) {
		return true
	}

	return false
}

func (p *Profile) Login(ctx context.Context, login string, password string) (string, error) {
	mycontext.LoggerFromContext(ctx).Infow("started login", "user", login)

	user, err := p.userRepo.GetByLogin(ctx, login)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user by login", "user", login, "error", err)
		return "", fmt.Errorf("get: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("login with wrong password", "user", login, "error", err)
		return "", fmt.Errorf("cmp hash & password: %w", err)
	}

	ctx = mycontext.UserToContext(ctx, user)

	if !areEqualDays(user.EnterTime, time.Now().UTC()) {
		s := fmt.Sprintf("daily %s bonus", time.Now().UTC().Format(time.RFC822))
		err = p.balanceTransactionLogic.Increase(ctx, p.dailyBonus, s)
		if err != nil {
			mycontext.LoggerFromContext(ctx).Warnw("cannot add daily bonus", "user", user.Login, "error", err)
			return "", fmt.Errorf("increase: %w", err)
		}
	}

	err = p.postLogic.PopularityCheck(ctx, user.UUID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot do popularity check", "user", user.Login, "error", err)
		return "", fmt.Errorf("popularity check: %w", err)
	}

	user.EnterTime = time.Now().UTC()

	err = p.userRepo.Update(ctx, user)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot update user", "user", user.Login, "error", err)
		return "", fmt.Errorf("update: %w", err)
	}

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(p.tokenExpiration)),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ID:        user.UUID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString([]byte(p.secretKey))

	mycontext.LoggerFromContext(ctx).Infow("successful login", "user", login)

	return ss, nil
}

func (p *Profile) Get(ctx context.Context, login string) (*models.User, error) {
	mycontext.LoggerFromContext(ctx).Infow("started getting user by login", "user", login)

	user, err := p.userRepo.GetByLogin(ctx, login)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user by login", "user", login, "error", err)
		return nil, fmt.Errorf("get: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("successfully got user by login", "user", user.Login)

	return user, nil
}

func (p *Profile) Delete(ctx context.Context, login string) error {
	mycontext.LoggerFromContext(ctx).Infow("started deleting user", "user to delete", login)

	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to delete another", "user to delete", login, "error", err)
		return fmt.Errorf("user from context: %w", err)
	}

	if !user.IsAdmin {
		mycontext.LoggerFromContext(ctx).Errorw("trying to delete user without admin perms", "user", user.Login, "deleted", login)
		return fmt.Errorf("delete: %w", errors2.ErrPerms)
	}

	if login == user.Login {
		mycontext.LoggerFromContext(ctx).Warnw("cannot delete yourself", "user", user.Login, "user to delete", login)
		return fmt.Errorf("delete: %w", errors2.ErrAutoDelete)
	}

	if err := p.userRepo.Delete(ctx, login); err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot delete user", "user", user.Login, "user to delete", login, "error", err)
		return fmt.Errorf("delete: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("admin deleted user", "user", user.Login, "deleted", login)

	return nil
}

func (p *Profile) GetTotal(ctx context.Context) (int, error) {
	_, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to get total number of posts by id", "error", err)
		return 0, fmt.Errorf("user from context: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("started getting total number of users")

	num, err := p.userRepo.GetTotal(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get total number of users", "error", err)
		return 0, fmt.Errorf("get total: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("successfully got total number of users")

	return num, nil
}

func (p *Profile) GetAll(ctx context.Context, pg *models.Paginator) ([]*models.User, error) {
	mycontext.LoggerFromContext(ctx).Infow("started getting all users")

	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to get all users", "error", err)
		return nil, fmt.Errorf("user from context: %w", err)
	}

	if !user.IsAdmin {
		mycontext.LoggerFromContext(ctx).Errorw("trying to get all users without admin perms", "user", user.Login)
		return nil, fmt.Errorf("check admin: %w", errors2.ErrPerms)
	}

	users, err := p.userRepo.GetAll(ctx, pg)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get all users", "user", user.Login, "error", err)
		return nil, fmt.Errorf("get all: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Infow("admin got all users", "user", user.Login)

	return users, nil
}
