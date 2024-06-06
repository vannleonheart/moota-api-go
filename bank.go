package moota

import (
	"fmt"
	"github.com/vannleonheart/goutil"
)

func (c *Client) GetListOfBank(page, perPage uint8) (*GetListOfBankResponse, error) {
	token, err := c.getToken()

	if err != nil {
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

	if _, err = goutil.SendHttpGet(endpoint, queryString, &requestHeader, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) BankAccounts(page, perPage uint8) (*BankAccountsResponse, error) {
	token, err := c.getToken()

	if err != nil {
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

	if _, err = goutil.SendHttpGet(endpoint, queryString, &requestHeader, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CreateBankAccount(data CreateBankAccountRequest) (*CreateBankAccountResponse, error) {
	token, err := c.getToken()

	if err != nil {
		return nil, err
	}

	requestHeader := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", *token),
	}

	endpoint := fmt.Sprintf("%s%s", c.Config.BaseUrl, URLCreateBank)

	var result CreateBankAccountResponse

	if _, err = goutil.SendHttpPost(endpoint, data, &requestHeader, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteBankAccount(bankId string) error {
	token, err := c.getToken()

	if err != nil {
		return err
	}

	requestHeader := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", *token),
	}

	endpoint := fmt.Sprintf("%s%s/%s/destroy", c.Config.BaseUrl, URLBankAccounts, bankId)

	if _, err = goutil.SendHttpPost(endpoint, nil, &requestHeader, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) RequestEWalletOTP(bankId string) error {
	token, err := c.getToken()

	if err != nil {
		return err
	}

	requestHeader := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", *token),
	}

	endpoint := fmt.Sprintf("%s%s/%s", c.Config.BaseUrl, URLRequestOTP, bankId)

	if _, err = goutil.SendHttpPost(endpoint, nil, &requestHeader, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) VerifyEWalletOTP(bankId, otp string) error {
	token, err := c.getToken()

	if err != nil {
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

	if _, err = goutil.SendHttpPost(endpoint, requestBody, &requestHeader, nil); err != nil {
		return err
	}

	return nil
}
