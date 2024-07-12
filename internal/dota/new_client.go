package dota

import (
	"net/http"
	"time"
)

type Client struct {
	client  *http.Client
	baseUrl string
}

type ClientOpts struct {
	Timeout time.Duration
	APIKey  string
}

func NewClient(clientOpts ...ClientOpts) *Client {
	var opts ClientOpts
	if len(clientOpts) > 0 {
		opts = clientOpts[0]
	}

	// Set default timeout to 10 seconds
	if opts.Timeout == 0 {
		opts.Timeout = 10 * time.Second
	}

	return &Client{
		client: &http.Client{
			Timeout: opts.Timeout,
		},
		baseUrl: "https://api.opendota.com/api",
	}
}
