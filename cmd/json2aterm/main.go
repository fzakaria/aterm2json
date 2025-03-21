package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/nix-community/go-nix/pkg/derivation"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <input.json>\n", os.Args[0])
		os.Exit(1)
	}

	inputFile := os.Args[1]

	var data *os.File
	var err error

	if inputFile == "-" {
		data = os.Stdin
	} else {
		data, err = os.Open(inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			os.Exit(1)
		}
		defer data.Close()
	}

	var drv derivation.Derivation
	if err := json.NewDecoder(data).Decode(&drv); err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding JSON: %v\n", err)
		os.Exit(1)
	}

	if err := drv.WriteDerivation(os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing ATerm: %v\n", err)
		os.Exit(1)
	}
}
