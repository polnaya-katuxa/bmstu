package igdb

import (
	"db_gen/internal/igdb/query"
	"encoding/json"
	"errors"
)

const (
	genresURL = "https://api.igdb.com/v4/genres"
)

type genreResponseItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type genresResponse []genreResponseItem

func (c *Client) genreName(ids []int) (map[int]string, error) {
	q := query.New().Fields("id, name").Where(query.Equal("id", query.IDsToString(ids))).Limit(len(ids))

	req, err := c.createRequest(genresURL, q)
	if err != nil {
		return nil, err
	}

	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var respParsed genresResponse
	err = json.NewDecoder(resp.Body).Decode(&respParsed)
	if err != nil {
		return nil, err
	}
	if len(respParsed) == 0 {
		return nil, errors.New("genre not found")
	}

	genresMap := make(map[int]string, len(respParsed))
	for _, g := range respParsed {
		genresMap[g.ID] = g.Name
	}

	return genresMap, nil
}
