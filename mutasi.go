package moota

import (
	"fmt"
	"github.com/vannleonheart/goutil"
)

func (c *Client) RefreshMutasi(bankId string) error {
	token, err := c.getToken()

	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "RefreshMutasi",
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

	endpoint := fmt.Sprintf("%s%s/%s/refresh", c.Config.BaseUrl, URLBankAccounts, bankId)

	raw, err := goutil.SendHttpPost(endpoint, nil, &requestHeader, nil, c.httpClient)
	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "RefreshMutasi",
			"error":   err,
			"message": "error when send http post",
			"url":     endpoint,
			"headers": requestHeader,
			"result":  raw,
		})

		return err
	}

	c.log("debug", map[string]interface{}{
		"method":  "RefreshMutasi",
		"url":     endpoint,
		"headers": requestHeader,
		"result":  raw,
	})

	return nil
}

func (c *Client) Mutasi(transactiontype, bankId, start, end, tag *string, page, perPage *uint8) (*MutasiResponse, error) {
	token, err := c.getToken()

	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "Mutasi",
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

	querystring := map[string]string{}

	if transactiontype != nil {
		querystring["type"] = *transactiontype
	}

	if bankId != nil {
		querystring["bank"] = *bankId
	}

	if start != nil {
		querystring["start_date"] = *start
	}

	if end != nil {
		querystring["end_date"] = *end
	}

	if tag != nil {
		querystring["tag"] = *tag
	}

	if page != nil {
		querystring["page"] = fmt.Sprintf("%d", *page)
	}

	if perPage != nil {
		querystring["per_page"] = fmt.Sprintf("%d", *perPage)
	}

	endpoint := fmt.Sprintf("%s%s", c.Config.BaseUrl, URLMutasi)

	var result MutasiResponse

	if raw, err := goutil.SendHttpGet(endpoint, querystring, &requestHeader, &result, c.httpClient); err != nil {
		c.log("error", map[string]interface{}{
			"method":  "Mutasi",
			"error":   err,
			"message": "error when send http get",
			"url":     endpoint,
			"query":   querystring,
			"headers": requestHeader,
			"result":  raw,
		})

		return nil, err
	}

	c.log("debug", map[string]interface{}{
		"method":  "Mutasi",
		"url":     endpoint,
		"query":   querystring,
		"headers": requestHeader,
		"result":  result,
	})

	return &result, nil
}

// CreateDummyMutasi
// TODO: Error The route api/v2/mutation/store/xxx could not be found.
func (c *Client) CreateDummyMutasi(bankId string, data CreateMutasiRequest) error {
	token, err := c.getToken()

	if err != nil {
		return err
	}

	requestHeader := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", *token),
	}

	endpoint := fmt.Sprintf("%s%s/%s", c.Config.BaseUrl, URLCreateMutasi, bankId)

	if _, err = goutil.SendHttpPost(endpoint, data, &requestHeader, nil, c.httpClient); err != nil {
		return err
	}

	return nil
}

func (c *Client) NoteMutasi(idMutasi, note string) error {
	token, err := c.getToken()

	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "NoteMutasi",
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
		"note": note,
	}

	endpoint := fmt.Sprintf("%s%s/%s/note", c.Config.BaseUrl, URLMutasi, idMutasi)

	raw, err := goutil.SendHttpPost(endpoint, requestBody, &requestHeader, nil, c.httpClient)
	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "NoteMutasi",
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
		"method":  "NoteMutasi",
		"url":     endpoint,
		"body":    requestBody,
		"headers": requestHeader,
		"result":  raw,
	})

	return nil
}

func (c *Client) PushWebHookMutasi(idMutasi string) error {
	token, err := c.getToken()

	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "PushWebHookMutasi",
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

	endpoint := fmt.Sprintf("%s%s/%s/webhook", c.Config.BaseUrl, URLMutasi, idMutasi)

	raw, err := goutil.SendHttpPost(endpoint, nil, &requestHeader, nil, c.httpClient)
	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "PushWebHookMutasi",
			"error":   err,
			"message": "error when send http post",
			"url":     endpoint,
			"headers": requestHeader,
			"result":  raw,
		})

		return err
	}

	c.log("debug", map[string]interface{}{
		"method":  "DeleteMutasi",
		"url":     endpoint,
		"headers": requestHeader,
		"result":  raw,
	})

	return nil
}

func (c *Client) DeleteMutasi(idMutasi []string) error {
	token, err := c.getToken()

	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "DeleteMutasi",
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
		"mutations": idMutasi,
	}

	endpoint := fmt.Sprintf("%s%s", c.Config.BaseUrl, URLDeleteMutasi)

	raw, err := goutil.SendHttpPost(endpoint, requestBody, &requestHeader, nil, c.httpClient)
	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "DeleteMutasi",
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
		"method":  "DeleteMutasi",
		"url":     endpoint,
		"body":    requestBody,
		"headers": requestHeader,
		"result":  raw,
	})

	return nil
}

func (c *Client) CreateTagMutasi(idMutasi string, tags []string) error {
	token, err := c.getToken()

	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "CreateTagMutasi",
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
		"name": tags,
	}

	endpoint := fmt.Sprintf("%s%s/%s", c.Config.BaseUrl, URLTagMutasi, idMutasi)

	raw, err := goutil.SendHttpPost(endpoint, requestBody, &requestHeader, nil, c.httpClient)
	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "CreateTagMutasi",
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
		"method":  "CreateTagMutasi",
		"url":     endpoint,
		"body":    requestBody,
		"headers": requestHeader,
		"result":  raw,
	})

	return nil
}

func (c *Client) UpdateTagMutasi(idMutasi string, tags []string) error {
	token, err := c.getToken()

	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "UpdateTagMutasi",
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
		"name": tags,
	}

	endpoint := fmt.Sprintf("%s%s/%s", c.Config.BaseUrl, URLTagMutasi, idMutasi)

	raw, err := goutil.SendHttpPut(endpoint, requestBody, &requestHeader, nil, c.httpClient)
	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "UpdateTagMutasi",
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
		"method":  "UpdateTagMutasi",
		"url":     endpoint,
		"body":    requestBody,
		"headers": requestHeader,
		"result":  raw,
	})

	return nil
}

func (c *Client) SummaryMutasi(bankId, transactiontype, start, end *string) (*SummaryMutasiResponse, error) {
	token, err := c.getToken()

	if err != nil {
		c.log("error", map[string]interface{}{
			"method":  "SummaryMutasi",
			"error":   err,
			"message": "error when get token",
		})

		return nil, err
	}

	requestHeader := map[string]string{
		"Accept":        "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", *token),
	}

	querystring := map[string]string{}

	if bankId != nil {
		querystring["bank"] = *bankId
	}

	if transactiontype != nil {
		querystring["type"] = *transactiontype
	}

	if start != nil {
		querystring["start_date"] = *start
	}

	if end != nil {
		querystring["end_date"] = *end
	}

	endpoint := fmt.Sprintf("%s%s", c.Config.BaseUrl, URLSummaryMutasi)

	var result SummaryMutasiResponse

	if raw, err := goutil.SendHttpGet(endpoint, querystring, &requestHeader, &result, c.httpClient); err != nil {
		c.log("error", map[string]interface{}{
			"method":  "SummaryMutasi",
			"error":   err,
			"message": "error when send http get",
			"url":     endpoint,
			"query":   querystring,
			"headers": requestHeader,
			"result":  raw,
		})

		return nil, err
	}

	c.log("debug", map[string]interface{}{
		"method":  "SummaryMutasi",
		"url":     endpoint,
		"query":   querystring,
		"headers": requestHeader,
		"result":  result,
	})

	return &result, nil
}
