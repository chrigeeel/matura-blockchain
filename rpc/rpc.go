package rpc

import "net/http"

type Client struct {
	url        string
	httpClient *http.Client
}

func NewClient(url string) *Client {
	return &Client{
		url:        url,
		httpClient: http.DefaultClient,
	}
}
