package client

import (
	"net/http"
	"time"
)

type Client struct {
	client *http.Client
}

func New(timeout time.Duration) *Client {
	httpClient := &http.Client{
		Timeout: timeout,
	}

	return &Client{
		client: httpClient,
	}
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", "SimpleCrawler")
	return c.client.Do(req)
}
