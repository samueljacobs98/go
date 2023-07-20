package main

import (
	"io"
	"log"
	"os"
)

func main() {
	args := os.Args
	// Validate input
	if len(args) != 2 {
		log.Fatal("Please provide a file name as an argument")
	}

	// Parse file name from command line arguments
	fileName := args[1]

	// Open the f
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file: %s, error: %v", fileName, err)
	}
	defer f.Close()

	// Copy the file content to stdout
	if _, err = io.Copy(os.Stdout, f); err != nil {
		log.Fatalf("Failed to read file: %s, error: %v", fileName, err)
	}
}
