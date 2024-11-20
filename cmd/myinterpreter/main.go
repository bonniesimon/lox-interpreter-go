package main

import (
	"fmt"
	"os"
	"strings"
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
			} else if char == LEFT_BRACE {
				fmt.Println("LEFT_BRACE { null")
			} else if char == RIGHT_BRACE {
				fmt.Println("RIGHT_BRACE } null")
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
