package rpc

import (
	"encoding/json"
	"errors"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func GetError(body []byte) error {
	var parsed ErrorResponse
	err := json.Unmarshal(body, &parsed)
	if err != nil {
		return err
	}

	return errors.New(parsed.Message)
}
