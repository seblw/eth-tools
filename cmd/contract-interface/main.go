package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	b := strings.Builder{}
	m := strings.Builder{}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		text := sc.Text()
		if strings.HasPrefix(text, "//") {
			m.WriteString(text)
		} else {
			b.WriteString(text)
		}
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "scanner failed: %s\n", err)
	}

	abi := ABI{}
	err := json.Unmarshal([]byte(b.String()), &abi.Entries)
	if err != nil {
		fmt.Fprintf(os.Stderr, "json.Unmarshal failed: %s\n", err)
	}

	for _, e := range abi.Entries {
		switch e.Type {
		case "constructor":
			emitConstructor(e)
		case "function":
			emitFunction(e)
		case "event":
			emitEvent(e)
		}
	}
}

func emitConstructor(e Entry) {
	args := make([]string, 0)

	for _, in := range e.Inputs {
		args = append(args, fmt.Sprintf("%s %s", in.Type, in.Name))
	}
	fmt.Printf("constructor(%s) public\n", strings.Join(args, ", "))
}

func emitFunction(e Entry) {
	args := make([]string, 0, len(e.Inputs))
	rets := make([]string, 0, len(e.Outputs))

	for _, in := range e.Inputs {
		arg := ""
		if in.Name != "" {
			arg = fmt.Sprintf("%s %s", in.Type, in.Name)
		} else {
			arg = in.Type
		}
		args = append(args, arg)
	}
	for _, o := range e.Outputs {
		ret := ""
		if o.Name != "" {
			ret = fmt.Sprintf("%s %s", o.Type, o.Name)
		} else {
			ret = o.Type
		}
		rets = append(rets, ret)
	}
	part := fmt.Sprintf("function %s(%s) external %s", e.Name, strings.Join(args, ", "), e.StateMutability)
	if len(rets) > 0 {
		fmt.Printf("%s returns (%s)\n", part, strings.Join(rets, ","))
	} else {
		fmt.Printf("%s\n", part)
	}
}

func emitEvent(e Entry) {
	args := make([]string, 0)

	for _, in := range e.Inputs {
		arg := ""
		if in.Name != "" && in.Indexed {
			arg = fmt.Sprintf("%s indexed %s", in.Type, in.Name)
		} else if in.Name != "" {
			arg = fmt.Sprintf("%s %s", in.Type, in.Name)
		} else {
			arg = in.Type
		}
		args = append(args, arg)
	}
	fmt.Printf("event %s(%s)\n", e.Name, strings.Join(args, ", "))
}

type ABI struct {
	Entries []Entry
	Meta    Metadata
}

type Entry struct {
	Inputs          []Input  `json:"inputs"`
	StateMutability string   `json:"stateMutability,omitempty"`
	Type            string   `json:"type"`
	Anonymous       bool     `json:"anonymous,omitempty"`
	Name            string   `json:"name,omitempty"`
	Outputs         []Output `json:"outputs,omitempty"`
}

type Input struct {
	Indexed      bool   `json:"indexed"`
	InternalType string `json:"internalType"`
	Name         string `json:"name"`
	Type         string `json:"type"`
}
type Output struct {
	InternalType string `json:"internalType"`
	Name         string `json:"name"`
	Type         string `json:"type"`
}

type Metadata struct {
	Name, Compiler, Licence string
}
