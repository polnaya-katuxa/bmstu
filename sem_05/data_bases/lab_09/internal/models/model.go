package models

type Games struct {
	ID              int
	Name            string
	Genre           string
	ReleaseYear     int
	Company         string
	AgeRating       string
	Price           float64
	MultiplayerMode bool
}
