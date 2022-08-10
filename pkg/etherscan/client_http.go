package etherscan

import (
	"net/http"
)

const (
	BaseMainnet = "https://api.etherscan.io"
)

type HTTPClient struct {
	client  *http.Client
	apiKey  string
	baseURL string
}

type Option func(*HTTPClient)

func WithBaseURL(url string) Option {
	return func(c *HTTPClient) {
		c.baseURL = url
	}
}

func NewHTTPClient(apikey string, opts ...Option) *HTTPClient {
	c := &HTTPClient{
		client:  http.DefaultClient,
		apiKey:  apikey,
		baseURL: BaseMainnet,
	}
	for _, o := range opts {
		o(c)
	}
	return c
}
