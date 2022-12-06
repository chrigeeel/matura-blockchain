package rpc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/chrigeeel/matura-blockchain/core"
)

type GetBalanceResponse struct {
	Balance uint64 `json:"balance"`
}

func (c *Client) GetBalance(publicKey core.PublicKey) (GetBalanceResponse, error) {
	path := fmt.Sprintf("%s/balance/%s", c.url, publicKey.String())

	resp, err := c.httpClient.Get(path)
	if err != nil {
		return GetBalanceResponse{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GetBalanceResponse{}, err
	}

	if resp.StatusCode != 200 {
		return GetBalanceResponse{}, GetError(body)
	}

	var parsed GetBalanceResponse
	return parsed, json.Unmarshal(body, &parsed)
}
