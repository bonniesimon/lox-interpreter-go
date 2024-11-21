package main

import (
	"fmt"
	"os"
	"strings"
)

var tokens []Token

func scanner(fileContents string) {
	lines := strings.Split(fileContents, "\n")

	for lineCount := 0; lineCount < len(lines); lineCount++ {
		for charCount := 0; charCount < len(lines[lineCount]); charCount++ {
			char := lines[lineCount][charCount]

			if char == '(' {
				AddToken(Token{"LEFT_PAREN", "(", ""}, &tokens)
			} else if char == ')' {
				AddToken(Token{"RIGHT_PAREN", ")", ""}, &tokens)
			} else if char == '{' {
				AddToken(Token{"LEFT_BRACE", "{", ""}, &tokens)
			} else if char == '}' {
				AddToken(Token{"RIGHT_BRACE", "}", ""}, &tokens)
			}
		}
	}

	AddToken(Token{"EOF", "EOF", ""}, &tokens)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	// Uncomment this block to pass the first stage

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(fileContents) > 0 {
		scanner(string(fileContents))
		for _, token := range tokens {
			token.Debug()
			fmt.Println()
		}
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}
}
