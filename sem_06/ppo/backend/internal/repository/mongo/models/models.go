package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID        uuid.UUID `bson:"_id"`
	Login       string    `bson:"login"`
	Picture     string    `bson:"picture"`
	Description string    `bson:"description"`
	Password    string    `bson:"password"`
	Balance     int       `bson:"balance"`
	Mail        string    `bson:"mail"`
	EnterTime   time.Time `bson:"enter_time"`
	IsAdmin     bool      `bson:"is_admin"`
}

type Subscription struct {
	UUID     uuid.UUID `bson:"_id"`
	ReaderID uuid.UUID `bson:"reader_id"`
	WriterID uuid.UUID `bson:"writer_id"`
}

type BalanceTransaction struct {
	UUID   uuid.UUID `bson:"_id"`
	Reason string    `bson:"reason"`
	Time   time.Time `bson:"time"`
	Amount int       `bson:"amount"`
	UserID uuid.UUID `bson:"user_id"`
}

type Limit struct {
	UUID  uuid.UUID `bson:"_id"`
	Value int       `bson:"value"`
	Bonus int       `bson:"bonus"`
}

type Reaction struct {
	UUID           uuid.UUID `bson:"_id"`
	ReactionTypeID uuid.UUID `bson:"reaction_type_id"`
	ReactorID      uuid.UUID `bson:"reactor_id"`
	PostID         uuid.UUID `bson:"post_id"`
}

type Comment struct {
	UUID          uuid.UUID `bson:"_id"`
	Content       string    `bson:"content"`
	PublicTime    time.Time `bson:"public_time"`
	CommentatorID uuid.UUID `bson:"commentator_id"`
	PostID        uuid.UUID `bson:"post_id"`
}

type ReactionType struct {
	UUID uuid.UUID `bson:"_id"`
	Icon string    `bson:"icon"`
}

type Post struct {
	UUID       uuid.UUID `bson:"_id"`
	Content    string    `bson:"content"`
	PublicTime time.Time `bson:"public_time"`
	Perms      string    `bson:"perms"`
	WriterID   uuid.UUID `bson:"writer_id"`
	LimitID    uuid.UUID `bson:"limit_id"`
}
