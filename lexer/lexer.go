package lexer

import (
	"unicode"

	"../tokens"
)

type Lexer struct {
	input    string
	position int
}

func New(input string) *Lexer {
	lexer := Lexer{input: input}
	return &lexer
}

func (lexer *Lexer) NextToken() tokens.Token {
	if lexer.position >= len(lexer.input) {
		return tokens.Token{Type: tokens.EOF}
	}

	char := lexer.input[lexer.position]
	lexer.position += 1

	// ignore whitespace
	if unicode.IsSpace(rune(char)) {
		return lexer.NextToken()
	}

	// check for tokens with length 1
	if token, ok := tokens.SimpleTokens[char]; ok {
		return tokens.Token{Type: token}
	}

	// check for numbers (ints for now)
	if unicode.IsDigit(rune(char)) {
		number := lexer.getInt()
		return tokens.Token{Type: tokens.INT, Value: number}
	}

	// check for keywords/identifiers
	if unicode.IsLetter(rune(char)) {
		identifier := lexer.getIdentifier()

		token := tokens.LookupIdentifier(identifier)

		if token == tokens.IDENTIFIER {
			return tokens.Token{Type: token, Value: identifier}
		} else {
			return tokens.Token{Type: token}
		}
	}

	return tokens.Token{Type: tokens.ILLEGAL}
}

func (lexer *Lexer) getIdentifier() string {
	start := lexer.position - 1

	for lexer.position < len(lexer.input) {
		char := lexer.input[lexer.position]

		if unicode.IsLetter(rune(char)) || unicode.IsDigit(rune(char)) || char == '_' {
			lexer.position += 1
		} else {
			break
		}
	}

	return lexer.input[start:lexer.position]
}

func (lexer *Lexer) getInt() string {
	start := lexer.position - 1

	for lexer.position < len(lexer.input) {
		if unicode.IsDigit(rune(lexer.input[lexer.position])) {
			lexer.position += 1
		} else {
			break
		}
	}

	return lexer.input[start:lexer.position]
}
