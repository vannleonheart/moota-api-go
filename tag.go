package moota

import (
	"fmt"
	"github.com/vannleonheart/goutil"
)

func (c *Client) CreateTag(tag string) (*CreateTagResponse, error) {
	token, err := c.getToken()

	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "CreateTag",
			"error":   err,
			"message": "error when get token",
		})

		return nil, err
	}

	requestHeader := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", *token),
	}

	requestBody := map[string]interface{}{
		"name": tag,
	}

	endpoint := fmt.Sprintf("%s%s", c.Config.BaseUrl, URLCreateTag)

	var result CreateTagResponse

	if raw, err := goutil.SendHttpPost(endpoint, requestBody, &requestHeader, &result); err != nil {
		c.log("error", map[string]interface{}{
			"method":  "CreateTag",
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
		"method":  "CreateTag",
		"url":     endpoint,
		"body":    requestBody,
		"headers": requestHeader,
		"result":  result,
	})

	return &result, nil
}

func (c *Client) UpdateTag(idTag string, tag string) error {
	token, err := c.getToken()

	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "UpdateTag",
			"error":   err,
			"message": "error when get token",
		})

		return err
	}

	requestHeader := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", *token),
	}

	requestBody := map[string]interface{}{
		"name": tag,
	}

	endpoint := fmt.Sprintf("%s%s/%s", c.Config.BaseUrl, URLCreateTag, idTag)

	raw, err := goutil.SendHttpPut(endpoint, requestBody, &requestHeader, nil)
	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "UpdateTag",
			"error":   err,
			"message": "error when send http put",
			"url":     endpoint,
			"body":    requestBody,
			"headers": requestHeader,
			"result":  raw,
		})

		return err
	}

	c.log("debug", map[string]interface{}{
		"method":  "UpdateTag",
		"url":     endpoint,
		"body":    requestBody,
		"headers": requestHeader,
		"result":  raw,
	})

	return nil
}

func (c *Client) DeleteTag(idTag string) error {
	token, err := c.getToken()

	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "DeleteTag",
			"error":   err,
			"message": "error when get token",
		})

		return err
	}

	requestHeader := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", *token),
	}

	endpoint := fmt.Sprintf("%s%s/%s", c.Config.BaseUrl, URLCreateTag, idTag)

	raw, err := goutil.SendHttpDelete(endpoint, &requestHeader, nil)
	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "DeleteTag",
			"error":   err,
			"message": "error when send http delete",
			"url":     endpoint,
			"headers": requestHeader,
			"result":  raw,
		})

		return err
	}

	c.log("debug", map[string]interface{}{
		"method":  "DeleteTag",
		"url":     endpoint,
		"headers": requestHeader,
		"result":  raw,
	})

	return nil
}
