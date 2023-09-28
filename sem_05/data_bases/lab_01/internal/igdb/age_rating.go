package igdb

import (
	"db_gen/internal/igdb/query"
	"encoding/json"
	"errors"
)

const (
	rating3         = "EC"
	rating6         = "E"
	rating10        = "E10+"
	rating13        = "T"
	rating17        = "M"
	rating18        = "AO"
	ratingPending   = "RP"
	ratingPending17 = "RP 17+"
)

var ratingMap = map[int]string{
	1:  rating3,
	2:  rating6,
	3:  rating13,
	4:  rating17,
	5:  rating18,
	6:  ratingPending,
	7:  rating3,
	8:  rating6,
	9:  rating10,
	10: rating13,
	11: rating17,
	12: rating18,
	13: rating3,
	14: rating10,
	15: rating13,
	16: rating17,
	17: rating18,
	18: rating3,
	19: rating6,
	20: rating13,
	21: rating18,
	22: rating3,
	23: rating10,
	24: rating13,
	25: rating18,
	26: ratingPending,
	27: rating3,
	28: rating10,
	29: rating13,
	30: rating13,
	31: rating17,
	32: rating18,
	33: rating3,
	34: rating13,
	35: ratingPending17,
	36: rating17,
	37: rating18,
	38: rating18,
}

const (
	ageRatingURL = "https://api.igdb.com/v4/age_ratings"
)

type ageResponseItem struct {
	ID     int `json:"id"`
	Rating int `json:"rating"`
}

type ageResponse []ageResponseItem

func (c *Client) ageRating(ids []int) (map[int]string, error) {
	q := query.New().
		Fields("rating").
		Where(query.Equal("id", query.IDsToString(ids))).Limit(len(ids))

	req, err := c.createRequest(ageRatingURL, q)
	if err != nil {
		return nil, err
	}

	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var respParsed ageResponse
	err = json.NewDecoder(resp.Body).Decode(&respParsed)
	if err != nil {
		return nil, err
	}

	if len(respParsed) == 0 {
		return nil, errors.New("age rating not found")
	}

	ageRatingMap := make(map[int]string, len(respParsed))
	for _, g := range respParsed {
		ageRatingMap[g.ID] = ratingMap[g.Rating]
	}

	return ageRatingMap, nil
}
