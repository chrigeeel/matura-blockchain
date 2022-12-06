package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/chrigeeel/matura-blockchain/core"
)

type SendTransactionResponse struct {
	Status string           `json:"status"`
	Data   core.Transaction `json:"data"`
}

func (c *Client) SendTransaction(transaction core.Transaction) (SendTransactionResponse, error) {
	if transaction.IsFieldEmpty() {
		return SendTransactionResponse{}, errors.New("transaction fields missing")
	}
	path := fmt.Sprintf("%s/transaction", c.url)

	payload, err := transaction.Marshal()
	if err != nil {
		return SendTransactionResponse{}, err
	}
	resp, err := c.httpClient.Post(path, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return SendTransactionResponse{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return SendTransactionResponse{}, err
	}

	if resp.StatusCode != 200 {
		return SendTransactionResponse{}, GetError(body)
	}

	var parsed SendTransactionResponse
	return parsed, json.Unmarshal(body, &parsed)
}
