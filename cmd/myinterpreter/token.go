package main

import "fmt"

const (
	// Single-character tokens
	LEFT_PAREN  string = "LEFT_PAREN"
	RIGHT_PAREN string = "RIGHT_PAREN"
	LEFT_BRACE  string = "LEFT_BRACE"
	RIGHT_BRACE string = "RIGHT_BRACE"
	COMMA       string = "COMMA"
	DOT         string = "DOT"
	MINUS       string = "MINUS"
	PLUS        string = "PLUS"
	SEMICOLON   string = "SEMICOLON"
	SLASH       string = "SLASH"
	STAR        string = "STAR"

	// One or two character tokens
	BANG          string = "BANG"
	BANG_EQUAL    string = "BANG_EQUAL"
	EQUAL         string = "EQUAL"
	EQUAL_EQUAL   string = "EQUAL_EQUAL"
	GREATER       string = "GREATER"
	GREATER_EQUAL string = "GREATER_EQUAL"
	LESS          string = "LESS"
	LESS_EQUAL    string = "LESS_EQUAL"

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

type Token struct {
	TokenType string
	Lexeme    string
	Literal   string
}

func AddToken(newToken Token, tokens *[]Token) {
	*tokens = append(*tokens, newToken)
}

func (t Token) Debug() {
	var literal string
	if t.Literal == "" {
		literal = "null"
	} else {
		literal = t.Literal
	}

	if t.TokenType == EOF {
		fmt.Printf("%s  %s", t.TokenType, literal)
		return
	}

	fmt.Printf("%s %s %s", t.TokenType, t.Lexeme, literal)
}
