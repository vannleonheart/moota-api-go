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

	if raw, err := goutil.SendHttpPost(endpoint, requestBody, &requestHeader, &result); err != nil {
		c.log("error", map[string]interface{}{
			"method":  "GenerateToken",
			"error":   err,
			"message": "error when send http post",
			"url":     endpoint,
			"body":    requestBody,
			"headers": requestHeader,
			"result":  raw,
		})

		return nil, err
	}

	c.log("debug", map[string]interface{}{
		"method":  "GenerateToken",
		"url":     endpoint,
		"body":    requestBody,
		"headers": requestHeader,
		"result":  result,
	})

	return &result, nil
}

func (c *Client) DestroyToken() error {
	token, err := c.getToken()
	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "DestroyToken",
			"error":   err,
			"message": "error when get token",
		})

		return err
	}

	requestHeader := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", *token),
		"Content-Type":  "application/json",
		"Accept":        "application/json",
	}

	endpoint := fmt.Sprintf("%s%s", c.Config.BaseUrl, URLDestroyToken)

	raw, err := goutil.SendHttpPost(endpoint, nil, &requestHeader, nil)
	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "DestroyToken",
			"error":   err,
			"message": "error when send http post",
			"url":     endpoint,
			"headers": requestHeader,
			"result":  raw,
		})

		return err
	}

	c.log("debug", map[string]interface{}{
		"method":  "DestroyToken",
		"url":     endpoint,
		"headers": requestHeader,
		"result":  raw,
	})

	return nil
}
