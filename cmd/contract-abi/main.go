package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/seblw/eth-tools/pkg/etherscan"
)

func main() {
	const (
		envApiKey      = "ETH_TOOLS_APIKEY"
		envBaseUrl     = "ETH_TOOLS_URL"
		defaultBaseUrl = etherscan.BaseMainnet
	)
	var usage = fmt.Sprintf(`Usage: contract-abi <addr>

addr	- contract address (required)

Flags:
  %s (required)
  %s (optional, default: %s)`,
		envApiKey, envBaseUrl, defaultBaseUrl)

	if len(os.Args) != 2 && len(os.Args) != 3 {
		fail(usage)
	}

	apiKey := os.Getenv(envApiKey)
	if apiKey == "" {
		fail("etherscan api key required")
	}

	baseUrl := os.Getenv(envBaseUrl)
	if baseUrl == "" {
		baseUrl = defaultBaseUrl
	}

	addr := os.Args[1]

	cli := etherscan.NewHTTPClient(apiKey, etherscan.WithBaseURL(baseUrl))
	res, err := unwrapProxies(context.Background(), cli, addr)
	if err != nil {
		fail(err)
	}

	for _, result := range res.Result {
		fmt.Fprintf(os.Stdout, "// ContractName: %s\n", result.ContractName)
		fmt.Fprintf(os.Stdout, "// CompilerVersion: %s\n", result.CompilerVersion)
		fmt.Fprintf(os.Stdout, "// LicenseType: %s\n", result.LicenseType)

		var out bytes.Buffer
		json.Indent(&out, []byte(result.Abi), "", "  ")
		out.WriteTo(os.Stdout)
	}
}

func unwrapProxies(ctx context.Context, cli etherscan.Client, addr string) (*etherscan.GetSourceCodeResponse, error) {
	resp, err := cli.GetSourceCode(ctx, addr)
	if err != nil {
		return nil, err
	}
	for {
		proxy := resp.Result[0].Proxy
		impl := resp.Result[0].Implementation
		if proxy == "0" {
			return resp, err
		}
		resp, err = cli.GetSourceCode(ctx, impl)
		if err != nil {
			return nil, err
		}
	}
}

func printerr(v any) {
	fmt.Fprintf(os.Stdout, "%s\n", v)
}

func fail(v any) {
	printerr(v)
	os.Exit(1)
}
