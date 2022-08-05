package etherscan

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Etherscan's contracts module.
// https://docs.etherscan.io/api-endpoints/contracts

// GetSourceCode fetches source code of contract at given address.
// https://docs.etherscan.io/api-endpoints/contracts#get-contract-source-code-for-verified-contract-source-codes
func (c *HTTPClient) GetSourceCode(ctx context.Context, address string) (*GetSourceCodeResponse, error) {
	url := fmt.Sprintf("%s/api?module=contract&action=getsourcecode&address=%s&apikey=%s", c.baseURL, address, c.apiKey)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected HTTP status: %s", res.Status)
	}
	jsonResp := GetSourceCodeResponse{}
	if err := json.NewDecoder(res.Body).Decode((&jsonResp)); err != nil {
		return nil, err
	}
	return &jsonResp, nil
}

type GetSourceCodeResponse struct {
	Status  string             `json:"status"`
	Message string             `json:"message"`
	Result  []SourceCodeResult `json:"result"`
}

type SourceCodeResult struct {
	SourceCode           string `json:"SourceCode"`
	Abi                  string `json:"ABI"`
	ContractName         string `json:"ContractName"`
	CompilerVersion      string `json:"CompilerVersion"`
	OptimizationUsed     string `json:"OptimizationUsed"`
	Runs                 string `json:"Runs"`
	ConstructorArguments string `json:"ConstructorArguments"`
	EVMVersion           string `json:"EVMVersion"`
	Library              string `json:"Library"`
	LicenseType          string `json:"LicenseType"`
	Proxy                string `json:"Proxy"`
	Implementation       string `json:"Implementation"`
	SwarmSource          string `json:"SwarmSource"`
}
