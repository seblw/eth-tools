package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	paths := make([]string, 0)

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		paths = append(paths, sc.Text())
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	remappings := make(map[string]string)

	for _, path := range paths {
		elems := strings.Split(path, "/")
		for i := range elems {
			if strings.HasPrefix(elems[i], "@") {
				if _, ok := remappings[elems[i]]; !ok {
					remappings[elems[i]] = strings.Join(elems[:i+1], "/")
				}
			}
		}
	}

	for k, v := range remappings {
		fmt.Fprintf(os.Stdout, "%s=%s\n", k, v)
	}
}
