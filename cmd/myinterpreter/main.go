package main

import (
	"fmt"
	"os"
	"strings"
)

// Depreciated in favour of Scanner struct
func scannerDepreciated(fileContents string) ([]Token, bool) {
	var tokens []Token
	hasError := false

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
				// I should maybe move Scanner to a struct and then have methods for the Scanner struct
				// It would be much easier since we have the state inside the struct itself
				// This passes the tests, but I need a better approach to this.
				if charCount+1 < len(lines[lineCount]) && lines[lineCount][charCount+1] == '=' {
					AddToken(EQUAL_EQUAL, "==", "", &tokens)
					charCount++
				} else {
					AddToken(EQUAL, "=", "", &tokens)
				}
			} else if char == '>' {
				AddToken(GREATER, ">", "", &tokens)
			} else if char == '<' {
				AddToken(LESS, "<", "", &tokens)
			} else {
				hasError = true
				reportError(lineCount+1, "Unexpected character: "+string(char))
			}
		}
	}

	AddToken("EOF", "EOF", "", &tokens)

	return tokens, hasError
}

// Depreciated: moved to Scanner class
func reportError(lineNumber int, errorMessage string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error: %s\n", lineNumber, errorMessage)
}

func getCommand() (string, string) {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]
	filename := os.Args[2]

	return command, filename
}

func main() {
	command, filename := getCommand()

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
	}

	if len(fileContents) <= 0 {
		fmt.Println("EOF  null")
		return
	}

	scanner := NewScanner(string(fileContents))
	scanner.Lex()

	for _, token := range scanner.tokens {
		token.Debug()
		fmt.Println()
	}

	if scanner.hasError {
		os.Exit(65)
	}

}
