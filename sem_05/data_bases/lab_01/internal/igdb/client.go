package igdb

import (
	"bytes"
	"db_gen/internal/igdb/query"
	"db_gen/internal/twitch"
	"fmt"
	"net/http"
)

type Client struct {
	c *http.Client

	cTwitch *twitch.Client
}

func New(cTwitch *twitch.Client) *Client {
	c := Client{
		c:       &http.Client{},
		cTwitch: cTwitch,
	}

	return &c
}

func (c *Client) createRequest(u string, q *query.Query) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, u, bytes.NewBufferString(q.String()))
	if err != nil {
		return nil, err
	}

	token, err := c.cTwitch.GetAccessToken()
	if err != nil {
		return nil, err
	}

	req.Header.Add("Client-ID", c.cTwitch.GetClientID())
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	return req, nil
}
