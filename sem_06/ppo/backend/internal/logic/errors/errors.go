package errors

import (
	"errors"
	"fmt"
)

type InsufficientBalanceError struct {
	Want int
	Got  int
}

func (e *InsufficientBalanceError) Error() string {
	return fmt.Sprintf("insufficient balance: want %d, got %d", e.Want, e.Got)
}

func (e *InsufficientBalanceError) Is(err error) bool {
	if (e == nil) && (err == nil) {
		return true
	}

	var ibe *InsufficientBalanceError

	flag := errors.As(err, &ibe)
	if !flag {
		return false
	}

	if !(ibe.Want == e.Want && ibe.Got == e.Got) {
		return false
	}

	return true
}

var (
	ErrGet              = errors.New("cannot get user from context")
	ErrPaid             = errors.New("post is available only after subscription")
	ErrAuthor           = errors.New("not author of post")
	ErrReact            = errors.New("author of post")
	ErrUncomment        = errors.New("can delete only your comments or comments under your posts")
	ErrLogin            = errors.New("login must be at least 8 symbols")
	ErrPassword         = errors.New("password must be at least 8 symbols")
	ErrEmail            = errors.New("invalid email")
	ErrExists           = errors.New("user already exists")
	ErrPerms            = errors.New("not an admin")
	ErrNotFound         = errors.New("not found")
	ErrAutoDelete       = errors.New("cannot delete yourself")
	ErrPermissionDenied = errors.New("database permission denied")
)
