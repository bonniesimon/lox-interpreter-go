package main

import (
	"errors"
)

type Scanner struct {
	source  string
	char    byte
	current int
	lines   int
	tokens  []Token
}

func NewScanner(contents string) Scanner {
	return Scanner{contents, contents[0], 0, 0, []Token{}}
}

func (s Scanner) Lex() ([]Token, bool) {
	hasError := false

	for !s.isAtEnd() {
		if s.char == '\t' || s.char == '\r' || s.char == ' ' {
			// Skip
		} else if s.char == '(' {
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
			if s.peak() == '/' {
				s.skipLine()
				s.lines++
			} else {
				s.addToken(SLASH, "/", "")
			}
		} else if s.char == '*' {
			s.addToken(STAR, "*", "")
		} else if s.char == '!' {
			if s.peak() == '=' {
				s.addToken(BANG_EQUAL, "!=", "")
				s.advance()
			} else {
				s.addToken(BANG, "!", "")
			}
		} else if s.char == '=' {
			if s.peak() == '=' {
				s.addToken(EQUAL_EQUAL, "==", "")
				s.advance()
			} else {
				s.addToken(EQUAL, "=", "")
			}
		} else if s.char == '>' {
			if s.peak() == '=' {
				s.addToken(GREATER_EQUAL, ">=", "")
				s.advance()
			} else {
				s.addToken(GREATER, ">", "")
			}
		} else if s.char == '<' {
			if s.peak() == '=' {
				s.addToken(LESS_EQUAL, "<=", "")
				s.advance()
			} else {
				s.addToken(LESS, "<", "")
			}
		} else if s.char == '\n' {
			s.lines++
		} else {
			hasError = true
			reportError(s.lines+1, "Unexpected character: "+string(s.char))
		}

		s.advance()
	}

	s.addToken("EOF", "EOF", "")

	return s.tokens, hasError

}

func (s Scanner) peak() byte {
	if s.current+1 < len(s.source) {
		return s.source[s.current+1]
	}

	return '\x00'
}

func (s *Scanner) skipLine() {
	for s.char != '\n' {
		err := s.advance()

		if err != nil {
			break
		}
	}
}

func (s *Scanner) advance() error {
	s.current++

	if s.current < len(s.source) {
		s.char = s.source[s.current]

		return nil
	}

	return errors.New("EOF")
}

func (s Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) addToken(tokenType string, lexeme string, literal string) {
	newToken := Token{tokenType, lexeme, literal}

	s.tokens = append(s.tokens, newToken)
}
