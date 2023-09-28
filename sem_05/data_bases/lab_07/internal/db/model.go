package db

import (
	"time"
)

type Clients struct {
	ID          int `gorm:"primaryKey,autoIncrement"`
	Name        string
	Surname     string
	Patronymic  string
	BirthDate   time.Time
	Sex         string
	PhoneNumber string
	Login       string
}

type Attendance struct {
	TimeStart time.Time
	TimeEnd   time.Time
	IDClub    int
	IDClient  int
	Rating    int
	Price     float64
}

type MaxPriceByRating struct {
	Rating int
	Price  float64 //`gorm:"column:p"`
}

type LoyaltyPrograms struct {
	ID                 int
	Name               string
	Design             string
	CashbackPercent    int
	MinimumPurchaseSum float64
}

type FeedbackJSON struct {
	Staff      string
	Machines   string
	Atmosphere string
	Parking    string `json:"parking,omitempty"`
}

type Joined struct {
	Login     string
	TimeStart time.Time
	TimeEnd   time.Time
	Rating    int
}
