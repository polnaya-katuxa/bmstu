package tech

import (
	"fmt"
	"time"
)

type DB struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
}

type Mongo struct {
	DB `mapstructure:",squash"`
}

type Postgres struct {
	DB `mapstructure:",squash"`
}

type Logger struct {
	Encoding string
	Level    string
	File     string
}

type Config struct {
	PG    Postgres
	Mongo Mongo

	DB string

	Cost       int
	TokenExp   time.Duration
	DailyBonus int
	SecretKey  string
	Span       time.Duration

	Logger Logger
}

func (d *Postgres) toDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", d.Host,
		d.Port, d.User, d.Password, d.DBName)
}

//nolint:unused
func (d *Mongo) toDSN() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", d.User, d.Password, d.Host,
		d.Port, d.DBName)
}
