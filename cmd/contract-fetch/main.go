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
		envApiKey     = "ETH_TOOLS_APIKEY"
		envOutDir     = "ETH_TOOLS_OUTDIR"
		defaultOutDir = "./out/"
	)
	var usage = fmt.Sprintf(`Usage: contract-fetch <addr>\
Flags:
	%s [required]
	%s [default: %s]`, envApiKey, envOutDir, defaultOutDir)

	if len(os.Args) != 2 {
		fail(usage)
	}

	apiKey := os.Getenv(envApiKey)
	if apiKey == "" {
		fail("etherscan api key required")
	}
	outdir := os.Getenv(envOutDir)
	if outdir == "" {
		outdir = defaultOutDir
	}

	addr := os.Args[1]
	if apiKey == "" {
		fail(usage)
	}

	cli := etherscan.NewHTTPClient(apiKey)
	res, err := cli.GetSourceCode(context.Background(), addr)
	if err != nil {
		fail(err)
	}

	for _, result := range res.Result {
		for path, sourceCode := range parseAndFlatten(result) {
			if err := save(filepath.Join(outdir, path), sourceCode); err != nil {
				printerr(err)
			}
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
