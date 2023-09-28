package redis

import (
	"context"
	"encoding/json"
	redisLib "github.com/go-redis/redis/v9"
	"lab_09/internal/models"
	"time"
)

const (
	exp = 10 * time.Minute
)

type Client struct {
	c redisLib.UniversalClient
}

func New(c redisLib.UniversalClient) *Client {
	return &Client{
		c,
	}
}

func (c *Client) Get(key string) ([]models.Games, error) {
	data, err := c.c.Get(context.Background(), key).Bytes()
	if err != nil {
		return nil, err
	}

	var games []models.Games
	err = json.Unmarshal(data, &games)
	if err != nil {
		return nil, err
	}

	return games, nil
}

func (c *Client) Set(key string, g []models.Games) error {
	data, err := json.Marshal(g)
	if err != nil {
		return err
	}

	_, err = c.c.Set(context.Background(), key, data, exp).Result()
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Delete(key string) error {
	return c.c.Del(context.Background(), key).Err()
}
