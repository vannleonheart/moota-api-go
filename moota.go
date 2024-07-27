package moota

import (
	"fmt"
	"github.com/vannleonheart/goutil"
	"net/http"
	"time"
)

func New(config Config) *Client {
	return &Client{Config: config}
}

func (c *Client) SetToken(token string) {
	c.token = &token
}

func (c *Client) WithToken(token string) *Client {
	c.SetToken(token)

	return c
}

func (c *Client) SetHttpClient(cl *http.Client) {
	c.httpClient = cl
}

func (c *Client) WithHttpClient(cl *http.Client) *Client {
	c.SetHttpClient(cl)

	return c
}

func (c *Client) getToken() (*string, error) {
	if c.token == nil {
		return nil, fmt.Errorf("token is not set")
	}

	if len(*c.token) == 0 {
		return nil, fmt.Errorf("token is empty")
	}

	return c.token, nil
}

func (c *Client) getTimestamp() string {
	ts := time.Now().Format(DefaultTimestampFormat)

	if loc, err := time.LoadLocation(DefaultTimezone); err == nil {
		ts = time.Now().In(loc).Format(DefaultTimestampFormat)
	}

	return ts
}

func (c *Client) log(level string, data interface{}) {
	if c.Config.Log != nil && c.Config.Log.Enable {
		if c.Config.Log.Level == "error" && level != "error" {
			return
		}

		msg := map[string]interface{}{
			"timestamp": c.getTimestamp(),
			"level":     level,
			"data":      data,
		}

		_ = goutil.WriteJsonToFile(msg, c.Config.Log.Path, c.Config.Log.Filename, c.Config.Log.Extension, c.Config.Log.Rotation)
	}
}
