package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: go run . <input file> <output file>")
		os.Exit(1)
	}

	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "lecture :", err)
		os.Exit(1)
	}

	result := processText(string(input))
	if err := os.WriteFile(os.Args[2], []byte(result), 0644); err != nil {
		fmt.Fprintln(os.Stderr, "écriture :", err)
		os.Exit(1)
	}
}
