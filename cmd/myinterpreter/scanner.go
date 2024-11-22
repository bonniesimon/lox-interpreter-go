package main

import "strings"

type Scanner struct {
	content string
	cursor  int
	char    byte
	row     int
	col     int
	tokens  []Token
}

func NewScanner(contents string) Scanner {
	return Scanner{contents, 0, contents[0], 0, 0, []Token{}}
}

func (s Scanner) Lex() ([]Token, bool) {
	hasError := false

	lines := strings.Split(s.content, "\n")

	for ; s.row < len(lines); s.row++ {
		for s.col = 0; s.col < len(lines[s.row]); s.col++ {
			s.char = lines[s.row][s.col]

			if s.char == '(' {
				s.addToken(LEFT_PAREN, "(", "")
			} else if s.char == ')' {
				s.addToken(RIGHT_PAREN, ")", "")
			} else if s.char == '{' {
				s.addToken(LEFT_BRACE, "{", "")
			} else if s.char == '}' {
				s.addToken(RIGHT_BRACE, "}", "")
			} else if s.char == ',' {
				s.addToken(COMMA, ",", "")
			} else if s.char == '.' {
				s.addToken(DOT, ".", "")
			} else if s.char == '-' {
				s.addToken(MINUS, "-", "")
			} else if s.char == '+' {
				s.addToken(PLUS, "+", "")
			} else if s.char == ';' {
				s.addToken(SEMICOLON, ";", "")
			} else if s.char == '/' {
				s.addToken(SLASH, "/", "")
			} else if s.char == '*' {
				s.addToken(STAR, "*", "")
			} else if s.char == '!' {
				s.addToken(BANG, "!", "")
			} else if s.char == '=' {
				// I should maybe move Scanner to a struct and then have methods for the Scanner struct
				// It would be much easier since we have the state inside the struct itself
				// This passes the tests, but I need a better approach to this.
				if s.col+1 < len(lines[s.row]) && lines[s.row][s.col+1] == '=' {
					s.addToken(EQUAL_EQUAL, "==", "")
					s.col++
				} else {
					s.addToken(EQUAL, "=", "")
				}
			} else if s.char == '>' {
				s.addToken(GREATER, ">", "")
			} else if s.char == '<' {
				s.addToken(LESS, "<", "")
			} else {
				hasError = true
				reportError(s.row+1, "Unexpected character: "+string(s.char))
			}
		}
	}

	s.addToken("EOF", "EOF", "")

	return s.tokens, hasError

}

func (s *Scanner) addToken(tokenType string, lexeme string, literal string) {
	newToken := Token{tokenType, lexeme, literal}

	s.tokens = append(s.tokens, newToken)
}
