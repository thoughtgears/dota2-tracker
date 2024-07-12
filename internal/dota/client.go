package dota

import (
	"net/http"
	"time"
)

type ClientOpts struct {
	APIKey string
}

type Client struct {
	client  *http.Client
	baseUrl string
	apiKey  string
}

func NewClient(opts ...ClientOpts) *Client {
	var o ClientOpts
	if len(opts) > 0 {
		o = opts[0]
	}

	return &Client{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseUrl: "https://api.opendota.com/api",
		apiKey:  o.APIKey,
	}
}
