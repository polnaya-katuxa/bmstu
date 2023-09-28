package igdb

import (
	"db_gen/internal/igdb/query"
	"encoding/json"
	"math/rand"
	"time"
)

const (
	gamesURL = "https://api.igdb.com/v4/games"
)

type Game struct {
	Name            string
	Genre           string
	ReleaseDate     time.Time
	Company         string
	Country         string
	AgeRating       string
	Price           float32
	MultiplayerMode bool
}

type gamesResponseItem struct {
	Name              string `json:"name"`
	Genres            []int  `json:"genres"`
	FirstReleaseDate  int64  `json:"first_release_date"`
	InvolvedCompanies []int  `json:"involved_companies"`
	AgeRatings        []int  `json:"age_ratings"`
	MultiplayerModes  []int  `json:"multiplayer_modes"`
}

type gamesResponse []gamesResponseItem

func (c *Client) Games(limit int, dateStart time.Time, dateEnd time.Time) ([]Game, error) {
	gamesResp, err := c.games(limit, dateStart, dateEnd)
	if err != nil {
		return nil, err
	}

	genreIDs := make([]int, 0, len(gamesResp))
	ageRatingsIDs := make([]int, 0, len(gamesResp))
	involvedCompaniesIDs := make([]int, 0, len(gamesResp))
	for _, g := range gamesResp {
		if len(g.Genres) == 0 || len(g.AgeRatings) == 0 {
			continue
		}
		genreIDs = append(genreIDs, g.Genres[0])
		ageRatingsIDs = append(ageRatingsIDs, g.AgeRatings[0])
		involvedCompaniesIDs = append(involvedCompaniesIDs, g.InvolvedCompanies[0])
	}

	genresMap, err := c.genreName(genreIDs)
	if err != nil {
		return nil, err
	}
	ageRatingsMap, err := c.ageRating(ageRatingsIDs)
	if err != nil {
		return nil, err
	}
	companiesMap, err := c.companyCountry(involvedCompaniesIDs)
	if err != nil {
		return nil, err
	}

	games := make([]Game, 0, len(gamesResp))
	for _, g := range gamesResp {
		game := Game{
			Name:            g.Name,
			ReleaseDate:     time.Unix(g.FirstReleaseDate, 0),
			Price:           rand.Float32() * 5000,
			MultiplayerMode: len(g.MultiplayerModes) != 0,
		}

		if len(g.Genres) == 0 {
			continue
		}
		game.Genre = genresMap[g.Genres[0]]

		if len(g.InvolvedCompanies) == 0 {
			continue
		}
		game.Company = companiesMap[g.InvolvedCompanies[0]].Name
		game.Country = companiesMap[g.InvolvedCompanies[0]].Country

		if len(g.AgeRatings) == 0 {
			continue
		}
		game.AgeRating = ageRatingsMap[g.AgeRatings[0]]

		games = append(games, game)
	}

	return games, nil
}

func (c *Client) games(limit int, dateStart time.Time, dateEnd time.Time) (gamesResponse, error) {
	q := query.New().
		Fields("name", "genres", "first_release_date", "involved_companies", "age_ratings", "multiplayer_modes").
		Where(
			query.NotEqual("genres", nil),
			query.NotEqual("involved_companies", nil),
			query.NotEqual("age_ratings", nil),
			query.More("first_release_date", dateStart.Unix()),
			query.Less("first_release_date", dateEnd.Unix()),
		).Limit(limit)

	req, err := c.createRequest(gamesURL, q)
	if err != nil {
		return nil, err
	}

	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var respParsed gamesResponse
	err = json.NewDecoder(resp.Body).Decode(&respParsed)
	if err != nil {
		return nil, err
	}

	return respParsed, nil
}
