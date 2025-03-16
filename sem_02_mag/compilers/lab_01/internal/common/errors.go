package common

import "errors"

var (
	ErrInvalidRegexp = errors.New("regexp is invalid")
	ErrRegexpNotSet  = errors.New("regexp not set")
	ErrUnclosedParen = errors.New("parentheses not closed")
)
