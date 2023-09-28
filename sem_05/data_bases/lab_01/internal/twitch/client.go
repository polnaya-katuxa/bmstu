package twitch

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const (
	twitchAuthURL = "https://id.twitch.tv/oauth2/token"
)

type Client struct {
	c *http.Client

	clientID  string
	secretKey string

	tokenStore expiringValue
}

func New(id string, key string) *Client {
	c := Client{
		c:         &http.Client{},
		clientID:  id,
		secretKey: key,
	}

	return &c
}

type twitchResponse struct {
	AccessToken string `json:"access_token"`
	Expire      int    `json:"expires_in"`
}

type expiringValue struct {
	value  string
	expire time.Time
}

func (e *expiringValue) set(val string, ttl int) {
	e.value = val
	e.expire = time.Now().Add(time.Duration(ttl) * time.Second)
}

func (e *expiringValue) get() (string, bool) {
	if time.Now().After(e.expire) {
		return "", false
	}

	return e.value, true
}

func (c *Client) GetClientID() string {
	return c.clientID
}

func (c *Client) GetAccessToken() (string, error) {
	key, ok := c.tokenStore.get()

	if ok {
		return key, nil
	}

	u, err := c.formURL()
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return "", err
	}

	resp, err := c.c.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var respParsed twitchResponse
	err = json.NewDecoder(resp.Body).Decode(&respParsed)
	if err != nil {
		return "", err
	}

	c.tokenStore.set(respParsed.AccessToken, respParsed.Expire)

	return respParsed.AccessToken, nil
}

func (c *Client) formURL() (string, error) {
	u, err := url.Parse(twitchAuthURL)

	if err != nil {
		return "", err
	}

	v := url.Values{}
	v.Add("client_id", c.clientID)
	v.Add("client_secret", c.secretKey)
	v.Add("grant_type", "client_credentials")

	u.RawQuery = v.Encode()

	return u.String(), nil
}
