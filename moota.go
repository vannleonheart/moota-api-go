package moota

import "fmt"

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

func (c *Client) getToken() (*string, error) {
	if c.token == nil {
		return nil, fmt.Errorf("token is not set")
	}

	if len(*c.token) == 0 {
		return nil, fmt.Errorf("token is empty")
	}

	return c.token, nil
}
