package main

import (
	"fmt"
	"os"

	"github.com/fzakaria/aterm2json/converter"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <input.aterm>\n", os.Args[0])
		os.Exit(1)
	}

	inputFile := os.Args[1]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	jsonOutput, err := converter.ATermToJSON(string(data))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Conversion error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(jsonOutput)
}
