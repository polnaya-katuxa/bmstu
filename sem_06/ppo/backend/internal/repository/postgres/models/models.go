package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID        uuid.UUID `gorm:"primaryKey"`
	Login       string
	Picture     string
	Description string
	Password    string
	Balance     int
	Mail        string
	EnterTime   time.Time
	IsAdmin     bool
}

type Subscription struct {
	UUID     uuid.UUID `gorm:"primaryKey"`
	ReaderID uuid.UUID
	WriterID uuid.UUID
}

type BalanceTransaction struct {
	UUID   uuid.UUID `gorm:"primaryKey"`
	Reason string
	Time   time.Time
	Amount int
	UserID uuid.UUID
}

type Limit struct {
	UUID  uuid.UUID `gorm:"primaryKey"`
	Value int
	Bonus int
}

type Reaction struct {
	UUID           uuid.UUID `gorm:"primaryKey,autoIncrement"`
	ReactionTypeID uuid.UUID
	ReactorID      uuid.UUID
	PostID         uuid.UUID
}

type Comment struct {
	UUID          uuid.UUID `gorm:"primaryKey,autoIncrement"`
	Content       string
	PublicTime    time.Time
	CommentatorID uuid.UUID
	PostID        uuid.UUID
}

type ReactionType struct {
	UUID uuid.UUID `gorm:"primaryKey"`
	Icon string
}

type Post struct {
	UUID       uuid.UUID `gorm:"primaryKey"`
	Content    string
	PublicTime time.Time
	Perms      string
	WriterID   uuid.UUID
	LimitID    uuid.UUID
}
