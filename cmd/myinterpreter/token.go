package main

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
