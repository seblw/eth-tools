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

func NewHTTPClient(apikey string) *HTTPClient {
	return &HTTPClient{
		client:  http.DefaultClient,
		apiKey:  apikey,
		baseURL: BaseMainnet,
	}
}
