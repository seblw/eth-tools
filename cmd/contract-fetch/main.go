package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/seblw/eth-tools/pkg/etherscan"
)

func main() {
	const (
		envApiKey      = "ETH_TOOLS_APIKEY"
		envBaseUrl     = "ETH_TOOLS_URL"
		defaultOutDir  = "./lib/"
		defaultBaseUrl = etherscan.BaseMainnet
	)
	var usage = fmt.Sprintf(`Usage: contract-fetch <addr> <outDir>

addr	- contract address (required)
outDir	- output directory (optional, default: %s)

Flags:
  %s (required)
  %s (optional, default: %s)`,
		defaultOutDir, envApiKey, envBaseUrl, defaultBaseUrl)

	if len(os.Args) != 2 && len(os.Args) != 3 {
		fail(usage)
	}

	apiKey := os.Getenv(envApiKey)
	if apiKey == "" {
		fail("etherscan api key required")
	}

	addr := os.Args[1]
	if apiKey == "" {
		fail(usage)
	}

	outdir := defaultOutDir
	if len(os.Args) == 3 {
		outdir = os.Args[2]
	}

	cli := etherscan.NewHTTPClient(apiKey, etherscan.WithBaseURL(defaultBaseUrl))
	res, err := cli.GetSourceCode(context.Background(), addr)
	if err != nil {
		fail(err)
	}

	for _, result := range res.Result {
		for path, sourceCode := range parseAndFlatten(result) {
			out := filepath.Join(outdir, result.ContractName, path)
			print(out)
			if err := save(out, sourceCode); err != nil {
				printerr(err)
			}
		}
	}
}

func printerr(v any) {
	fmt.Fprintf(os.Stdout, "%s\n", v)
}

func print(v any) {
	fmt.Fprintf(os.Stdout, "%s\n", v)
}

func fail(v any) {
	printerr(v)
	os.Exit(1)
}

func parseAndFlatten(res etherscan.SourceCodeResult) map[string]string {
	return flattenSources(parseSources(res))
}

// parseSources parse "sources" string from etherscan.SourceCodeResult to a map.
func parseSources(res etherscan.SourceCodeResult) map[string]map[string]string {
	type sourceCode struct {
		Sources map[string]map[string]string `json:"sources"`
	}

	sc := res.SourceCode
	// XXX: SourceCode content is sorrounded with {{ }} which fails json decoding.
	sc = strings.ReplaceAll(sc, "{{", "{")
	sc = strings.ReplaceAll(sc, "}}", "}")
	scc := sourceCode{}
	if err := json.NewDecoder(strings.NewReader(sc)).Decode(&scc); err != nil {
		fail(err)
	}
	return scc.Sources
}

// flattenSources flattens parsed map from etherscan.SourceCodeResult.sources to path => content map.
//
// map[filename]map["content"]source_code -> map[filename]source_code
func flattenSources(ss map[string]map[string]string) map[string]string {
	out := make(map[string]string)
	for k, v := range ss {
		out[k] = v["content"]
	}
	return out
}

// save saves content string at path.
// it creates parent directories if needed.
func save(path string, content string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	if err := os.WriteFile(path, []byte(content), os.ModePerm); err != nil {
		return err
	}
	return nil
}
