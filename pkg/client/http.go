package client

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
}

func NewClient(timeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) Get(url string) (*http.Response, error) {
	res, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	return res, nil
}
