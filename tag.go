package moota

import (
	"fmt"
	"github.com/vannleonheart/goutil"
)

func (c *Client) CreateTag(tag string) (*CreateTagResponse, error) {
	token, err := c.getToken()

	if err != nil {
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

	if _, err = goutil.SendHttpPost(endpoint, requestBody, &requestHeader, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateTag(idTag string, tag string) error {
	token, err := c.getToken()

	if err != nil {
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

	if _, err = goutil.SendHttpPut(endpoint, requestBody, &requestHeader, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteTag(idTag string) error {
	token, err := c.getToken()

	if err != nil {
		return err
	}

	requestHeader := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", *token),
	}

	endpoint := fmt.Sprintf("%s%s/%s", c.Config.BaseUrl, URLCreateTag, idTag)

	if _, err = goutil.SendHttpDelete(endpoint, &requestHeader, nil); err != nil {
		return err
	}

	return nil
}
