package main

import (
	"fmt"
	"halc/lexer"
	"os"
	"path/filepath"
)

const (
	RedColor    = "\033[31m"
	YellowColor = "\033[33m"
	ResetColor  = "\033[0m"
)

func main() {
	fmt.Println("halc v0.0.1")
	if len(os.Args) != 2 {
		fmt.Printf("%sUsage: halc <file_path.hal>%s\n", RedColor, ResetColor)
		os.Exit(1)
	}

	filePath := os.Args[1]

	if filepath.Ext(filePath) != ".hal" {
		fmt.Printf("%s[WARNING] File does not have a .hal extension%s\n", YellowColor, ResetColor)
	}

	absPath, err := filepath.Abs(filePath)
	if err != nil {
		fmt.Printf("%s[ERROR] Error getting absolute path: %s%s\n", RedColor, err, ResetColor)
		os.Exit(1)
	}

	content, err := os.ReadFile(absPath)
	if err != nil {
		fmt.Printf("%s[ERROR] Error reading file: %s%s\n", RedColor, err, ResetColor)
		os.Exit(1)
	}

	//Lexing
	l := lexer.NewLexer(string(content))
	tokens := l.Lex()
	for _, token := range tokens {
		fmt.Println(token)
	}
}
