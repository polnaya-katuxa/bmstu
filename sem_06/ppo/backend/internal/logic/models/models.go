package models

import (
	"time"

	"github.com/google/uuid"
)

type Perms int

const (
	Free Perms = iota
	Paid
)

func (p *Perms) String() string {
	if *p == Free {
		return "free"
	}

	return "paid"
}

type Paginator struct {
	Page int
	Num  int
}

type User struct {
	UUID        uuid.UUID
	Login       string
	Password    string
	Balance     int
	Mail        string
	EnterTime   time.Time
	Picture     string
	Description string
	IsAdmin     bool
}

type Subscription struct {
	UUID     uuid.UUID
	ReaderID uuid.UUID
	WriterID uuid.UUID
}

type BalanceTransaction struct {
	UUID   uuid.UUID
	Reason string
	Time   time.Time
	Amount int
	UserID uuid.UUID
}

type Limit struct {
	UUID  uuid.UUID
	Value int
	Bonus int
}

type Reaction struct {
	UUID      uuid.UUID
	Icon      string
	TypeID    uuid.UUID
	ReactorID uuid.UUID
}

type Comment struct {
	UUID        uuid.UUID
	Content     string
	PublicTime  time.Time
	Commentator *User
	PostID      uuid.UUID
}

type ReactionType struct {
	UUID uuid.UUID
	Icon string
}

type Post struct {
	UUID       uuid.UUID
	Content    string
	PublicTime time.Time
	Perms      Perms
	Reactions  []*Reaction
	CommentNum int
	Author     *User
	NextLimit  Limit
}
