package main

import (
	"fmt"
	"os"
	"strings"
)

func scanner(fileContents string) []Token {
	var tokens []Token

	lines := strings.Split(fileContents, "\n")

	for lineCount := 0; lineCount < len(lines); lineCount++ {
		for charCount := 0; charCount < len(lines[lineCount]); charCount++ {
			char := lines[lineCount][charCount]

			if char == '(' {
				AddToken(LEFT_PAREN, "(", "", &tokens)
			} else if char == ')' {
				AddToken(RIGHT_PAREN, ")", "", &tokens)
			} else if char == '{' {
				AddToken(LEFT_BRACE, "{", "", &tokens)
			} else if char == '}' {
				AddToken(RIGHT_BRACE, "}", "", &tokens)
			} else if char == ',' {
				AddToken(COMMA, ",", "", &tokens)
			} else if char == '.' {
				AddToken(DOT, ".", "", &tokens)
			} else if char == '-' {
				AddToken(MINUS, "-", "", &tokens)
			} else if char == '+' {
				AddToken(PLUS, "+", "", &tokens)
			} else if char == ';' {
				AddToken(SEMICOLON, ";", "", &tokens)
			} else if char == '/' {
				AddToken(SLASH, "/", "", &tokens)
			} else if char == '*' {
				AddToken(STAR, "*", "", &tokens)
			} else if char == '!' {
				AddToken(BANG, "!", "", &tokens)
			} else if char == '=' {
				AddToken(EQUAL, "=", "", &tokens)
			} else if char == '>' {
				AddToken(GREATER, ">", "", &tokens)
			} else if char == '<' {
				AddToken(LESS, "<", "", &tokens)
			}
		}
	}

	AddToken("EOF", "EOF", "", &tokens)

	return tokens
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
		tokens := scanner(string(fileContents))

		for _, token := range tokens {
			token.Debug()
			fmt.Println()
		}
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}
}
