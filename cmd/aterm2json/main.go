package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/nix-community/go-nix/pkg/derivation"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <input.aterm>\n", os.Args[0])
		os.Exit(1)
	}

	inputFile := os.Args[1]

	data, err := os.Open(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	defer data.Close()

	drv, err := derivation.ReadDerivation(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing derivation: %v\n", err)
		os.Exit(1)
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(drv); err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
		os.Exit(1)
	}
}
