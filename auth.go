package moota

import (
	"fmt"
	"github.com/vannleonheart/goutil"
)

func (c *Client) GenerateToken(email, password string) (*GenerateTokenResponse, error) {
	requestBody := GenerateTokenRequest{
		Email:    email,
		Password: password,
		Scopes:   []string{"api"},
	}

	requestHeader := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	endpoint := fmt.Sprintf("%s%s", c.Config.BaseUrl, URLGenerateToken)

	var result GenerateTokenResponse

	if _, err := goutil.SendHttpPost(endpoint, requestBody, &requestHeader, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DestroyToken() error {
	token, err := c.getToken()

	if err != nil {
		return err
	}

	requestHeader := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", *token),
		"Content-Type":  "application/json",
		"Accept":        "application/json",
	}

	endpoint := fmt.Sprintf("%s%s", c.Config.BaseUrl, URLDestroyToken)

	if _, err = goutil.SendHttpPost(endpoint, nil, &requestHeader, nil); err != nil {
		return err
	}

	return nil
}
