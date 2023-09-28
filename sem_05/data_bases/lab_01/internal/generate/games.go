package generate

import (
	igdb2 "db_gen/internal/igdb"
	"db_gen/internal/twitch"
	"fmt"
	"time"
)

const (
	limit      = 200
	yearsDelta = 5
)

type Game struct {
	ID              int     `csv:"id"`
	Name            string  `csv:"name"`
	Genre           string  `csv:"genre"`
	ReleaseYear     int     `csv:"release_year"`
	Company         string  `csv:"company"`
	Country         string  `csv:"country"`
	AgeRating       string  `csv:"age_rating"`
	Price           float32 `csv:"price"`
	MultiplayerMode bool    `csv:"multiplayer_mode"`
}

func toGames(games []igdb2.Game) []Game {
	result := make([]Game, len(games))

	for i, game := range games {
		result[i] = Game{
			ID:              i + 1,
			Name:            game.Name,
			Genre:           game.Genre,
			ReleaseYear:     game.ReleaseDate.Year(),
			Company:         game.Company,
			Country:         game.Country,
			AgeRating:       game.AgeRating,
			Price:           game.Price,
			MultiplayerMode: game.MultiplayerMode,
		}
	}

	return result
}

func Games(n int) ([]Game, error) {
	twitchClient := twitch.New("u09k5fcj44m7w9cz1m6rr7fnsl1je8", "6ebcgpfye8r4j4llq0sc00t2l9y7so")
	igdbClient := igdb2.New(twitchClient)

	end := time.Now()
	start := end.AddDate(-yearsDelta, 0, 0)
	games := make([]igdb2.Game, 0, n)

	for len(games) < n {
		g, err := igdbClient.Games(limit, start, end)
		if err != nil {
			return nil, err
		}

		games = append(games, g...)

		fmt.Printf("games: %d/%d\r", len(games), n)

		end = start
		start = start.AddDate(-yearsDelta, 0, 0)
	}

	return toGames(games), nil
}
