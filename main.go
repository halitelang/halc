package main

import (
	"fmt"
	"halc/lexer" // Use the correct import path for your lexer package
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: halc <file_path.hal>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	// Check if the file has a .hal extension
	if filepath.Ext(filePath) != ".hal" {
		fmt.Println("Warning: file does not have a .hal extension")
	}

	absPath, err := filepath.Abs(filePath)
	if err != nil {
		fmt.Printf("Error getting absolute path: %s\n", err)
		os.Exit(1)
	}

	content, err := os.ReadFile(absPath)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}

	l := lexer.NewLexer(string(content))
	for tok := l.NextToken(); tok.Type != lexer.EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}
}
