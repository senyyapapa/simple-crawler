package client

import "net/http"

type Fetcher interface {
	Do(req *http.Request) (*http.Response, error)
}
