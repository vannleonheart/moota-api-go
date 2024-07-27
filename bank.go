package moota

import (
	"fmt"
	"github.com/vannleonheart/goutil"
)

func (c *Client) GetListOfBank(page, perPage uint8) (*GetListOfBankResponse, error) {
	token, err := c.getToken()

	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "GetListOfBank",
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

	queryString := map[string]string{
		"page":     fmt.Sprintf("%d", page),
		"per_page": fmt.Sprintf("%d", perPage),
	}

	endpoint := fmt.Sprintf("%s%s", c.Config.BaseUrl, URLListOfBank)

	var result GetListOfBankResponse

	if raw, err := goutil.SendHttpGet(endpoint, queryString, &requestHeader, &result, c.httpClient); err != nil {
		c.log("error", map[string]interface{}{
			"method":  "GetListOfBank",
			"error":   err,
			"message": "error when send http get",
			"url":     endpoint,
			"query":   queryString,
			"headers": requestHeader,
			"result":  raw,
		})

		return nil, err
	}

	c.log("debug", map[string]interface{}{
		"method":  "GetListOfBank",
		"url":     endpoint,
		"query":   queryString,
		"headers": requestHeader,
		"result":  result,
	})

	return &result, nil
}

func (c *Client) BankAccounts(page, perPage uint8) (*BankAccountsResponse, error) {
	token, err := c.getToken()

	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "BankAccounts",
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

	queryString := map[string]string{
		"page":     fmt.Sprintf("%d", page),
		"per_page": fmt.Sprintf("%d", perPage),
	}

	endpoint := fmt.Sprintf("%s%s", c.Config.BaseUrl, URLBankAccounts)

	var result BankAccountsResponse

	if raw, err := goutil.SendHttpGet(endpoint, queryString, &requestHeader, &result, c.httpClient); err != nil {
		c.log("error", map[string]interface{}{
			"method":  "BankAccounts",
			"error":   err,
			"message": "error when send http get",
			"url":     endpoint,
			"query":   queryString,
			"headers": requestHeader,
			"result":  raw,
		})

		return nil, err
	}

	c.log("debug", map[string]interface{}{
		"method":  "BankAccounts",
		"url":     endpoint,
		"query":   queryString,
		"headers": requestHeader,
		"result":  result,
	})

	return &result, nil
}

func (c *Client) CreateBankAccount(data CreateBankAccountRequest) (*CreateBankAccountResponse, error) {
	token, err := c.getToken()

	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "CreateBankAccount",
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

	endpoint := fmt.Sprintf("%s%s", c.Config.BaseUrl, URLCreateBank)

	var result CreateBankAccountResponse

	if raw, err := goutil.SendHttpPost(endpoint, data, &requestHeader, &result, c.httpClient); err != nil {
		c.log("error", map[string]interface{}{
			"method":  "CreateBankAccount",
			"error":   err,
			"message": "error when send http post",
			"url":     endpoint,
			"body":    data,
			"headers": requestHeader,
			"result":  raw,
		})

		return nil, err
	}

	c.log("debug", map[string]interface{}{
		"method":  "CreateBankAccount",
		"url":     endpoint,
		"body":    data,
		"headers": requestHeader,
		"result":  result,
	})

	return &result, nil
}

func (c *Client) DeleteBankAccount(bankId string) error {
	token, err := c.getToken()

	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "DeleteBankAccount",
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

	endpoint := fmt.Sprintf("%s%s/%s/destroy", c.Config.BaseUrl, URLBankAccounts, bankId)

	raw, err := goutil.SendHttpPost(endpoint, nil, &requestHeader, nil, c.httpClient)
	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "DeleteBankAccount",
			"error":   err,
			"message": "error when send http post",
			"url":     endpoint,
			"headers": requestHeader,
			"result":  raw,
		})

		return err
	}

	c.log("debug", map[string]interface{}{
		"method":  "DeleteBankAccount",
		"url":     endpoint,
		"headers": requestHeader,
		"result":  raw,
	})

	return nil
}

func (c *Client) RequestEWalletOTP(bankId string) error {
	token, err := c.getToken()

	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "RequestEWalletOTP",
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

	endpoint := fmt.Sprintf("%s%s/%s", c.Config.BaseUrl, URLRequestOTP, bankId)

	raw, err := goutil.SendHttpPost(endpoint, nil, &requestHeader, nil, c.httpClient)
	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "RequestEWalletOTP",
			"error":   err,
			"message": "error when send http post",
			"url":     endpoint,
			"headers": requestHeader,
			"result":  raw,
		})

		return err
	}

	c.log("debug", map[string]interface{}{
		"method":  "RequestEWalletOTP",
		"url":     endpoint,
		"headers": requestHeader,
		"result":  raw,
	})

	return nil
}

func (c *Client) VerifyEWalletOTP(bankId, otp string) error {
	token, err := c.getToken()

	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "VerifyEWalletOTP",
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

	requestBody := map[string]string{
		"otp_code": otp,
	}

	endpoint := fmt.Sprintf("%s%s/%s", c.Config.BaseUrl, URLVerifyOTP, bankId)

	raw, err := goutil.SendHttpPost(endpoint, requestBody, &requestHeader, nil, c.httpClient)
	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "VerifyEWalletOTP",
			"error":   err,
			"message": "error when send http post",
			"url":     endpoint,
			"body":    requestBody,
			"headers": requestHeader,
			"result":  raw,
		})

		return err
	}

	c.log("debug", map[string]interface{}{
		"method":  "VerifyEWalletOTP",
		"url":     endpoint,
		"body":    requestBody,
		"headers": requestHeader,
		"result":  raw,
	})

	return nil
}
