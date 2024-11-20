package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	// Single-character tokens
	LEFT_PAREN  byte = '('
	RIGHT_PAREN byte = ')'
	LEFT_BRACE  byte = '{'
	RIGHT_BRACE byte = '}'
	COMMA       byte = ','
	DOT         byte = '.'
	MINUS       byte = '-'
	PLUS        byte = '+'
	SEMICOLON   byte = ';'
	SLASH       byte = '/'
	STAR        byte = '*'

	// One or two character tokens
	BANG          string = "!"
	BANG_EQUAL    string = "!="
	EQUAL         string = "="
	EQUAL_EQUAL   string = "=="
	GREATER       string = ">"
	GREATER_EQUAL string = ">="
	LESS          string = "<"
	LESS_EQUAL    string = "<="

	// Literals
	IDENTIFIER string = "IDENTIFIER"
	STRING     string = "STRING"
	NUMBER     string = "NUMBER"

	// Keywords
	AND    string = "AND"
	CLASS  string = "CLASS"
	ELSE   string = "ELSE"
	FALSE  string = "FALSE"
	FUN    string = "FUN"
	FOR    string = "FOR"
	IF     string = "IF"
	NIL    string = "NIL"
	OR     string = "OR"
	PRINT  string = "PRINT"
	RETURN string = "RETURN"
	SUPER  string = "SUPER"
	THIS   string = "THIS"
	TRUE   string = "TRUE"
	VAR    string = "VAR"
	WHILE  string = "WHILE"
	EOF    string = "EOF"
)

func scanner(fileContents string) {
	lines := strings.Split(fileContents, "\n")

	for lineCount := 0; lineCount < len(lines); lineCount++ {
		for charCount := 0; charCount < len(lines[lineCount]); charCount++ {
			char := lines[lineCount][charCount]

			if char == LEFT_PAREN {
				fmt.Println("LEFT_PAREN ( null")
			} else if char == RIGHT_PAREN {
				fmt.Println("RIGHT_PAREN ) null")
			}
		}
	}

	fmt.Println("EOF  null")
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
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}
}
