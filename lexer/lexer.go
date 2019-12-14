package lexer

import (
	"unicode"

	"../tokens"
)

const EOF = 0

type Lexer struct {
	input string
	pos   int

	line int
	col  int

	char rune
}

func New(input string) *Lexer {
	lexer := Lexer{input: input, pos: -1}
	return &lexer
}

func (lexer *Lexer) NextToken() tokens.Token {
	lexer.readChar()

	for unicode.IsSpace(lexer.char) {
		lexer.readChar()
	}

	switch lexer.char {
	case '=':
		if lexer.peek() == '=' {
			lexer.readChar()
			return lexer.token(tokens.EQ)
		}
		return lexer.token(tokens.ASSIGN)
	case '+':
		if lexer.peek() == '=' {
			lexer.readChar()
			return lexer.token(tokens.ADD_ASSIGN)
		}
		return lexer.token(tokens.ADD)
	case '-':
		if lexer.peek() == '=' {
			lexer.readChar()
			return lexer.token(tokens.SUB_ASSIGN)
		}
		return lexer.token(tokens.SUB)
	case '*':
		if lexer.peek() == '=' {
			lexer.readChar()
			return lexer.token(tokens.MUL_ASSIGN)
		}
		return lexer.token(tokens.MUL)
	case '/':
		if lexer.peek() == '=' {
			lexer.readChar()
			return lexer.token(tokens.DIV_ASSIGN)
		}
		return lexer.token(tokens.DIV)
	case '!':
		if lexer.peek() == '=' {
			lexer.readChar()
			return lexer.token(tokens.NOT_EQ)
		}
	case '<':
		if lexer.peek() == '=' {
			lexer.readChar()
			return lexer.token(tokens.LESS_EQ)
		}
		return lexer.token(tokens.LESS)
	case '>':
		if lexer.peek() == '=' {
			lexer.readChar()
			return lexer.token(tokens.GREATER_EQ)
		}
		return lexer.token(tokens.GREATER)
	case ',':
		return lexer.token(tokens.COMMA)
	case '(':
		return lexer.token(tokens.L_PAREN)
	case ')':
		return lexer.token(tokens.R_PAREN)
	case '{':
		return lexer.token(tokens.L_BRACE)
	case '}':
		return lexer.token(tokens.R_BRACE)
	case '[':
		return lexer.token(tokens.L_BRACKET)
	case ']':
		return lexer.token(tokens.R_BRACKET)
	case '.':
		return lexer.token(tokens.DOT)
	case '"':
		text := lexer.getText()
		return lexer.tokenLiteral(tokens.STRING, text)
	case EOF:
		return lexer.token(tokens.EOF)
	}

	// check for numbers
	if unicode.IsDigit(lexer.char) {
		// TODO: RETURN ERROR
		number := lexer.getDigits()

		if lexer.peek() == '.' {
			lexer.readChar() // char is now "."
			lexer.readChar() // char is first digit (TODO: CHECK)
			number += "." + lexer.getDigits()
		}

		return lexer.tokenLiteral(tokens.NUMBER, number)
	}

	// check for keywords/identifiers
	if unicode.IsLetter(lexer.char) {
		// TODO: ADD ERROR
		identifier := lexer.getIdentifier()
		token := tokens.LookupIdentifier(identifier)

		if token == tokens.IDENT {
			return lexer.tokenLiteral(token, identifier)
		} else {
			return lexer.token(token)
		}
	}

	return lexer.tokenLiteral(tokens.ILLEGAL, string(lexer.char))
}

// TODO: utf8
func (lexer *Lexer) readChar() {
	lexer.pos += 1
	lexer.col += 1

	if lexer.pos >= len(lexer.input) {
		lexer.char = EOF
	} else {
		lexer.char = rune(lexer.input[lexer.pos])

		if lexer.char == '\n' {
			lexer.col = 0
			lexer.line += 1
		}
	}
}

func (lexer *Lexer) peek() rune {
	if lexer.pos+1 >= len(lexer.input) {
		return EOF
	}

	return rune(lexer.input[lexer.pos+1])
}

func (lexer *Lexer) getText() string {
	start := lexer.pos + 1
	lexer.readChar()

	for lexer.char != '"' {
		lexer.readChar()
	}

	return lexer.input[start:lexer.pos]
}

func (lexer *Lexer) getIdentifier() string {
	start := lexer.pos

	for unicode.IsLetter(lexer.peek()) || unicode.IsDigit(lexer.peek()) || lexer.peek() == '_' {
		lexer.readChar()
	}

	return lexer.input[start : lexer.pos+1]
}

func (lexer *Lexer) getDigits() string {
	start := lexer.pos

	for unicode.IsDigit(lexer.peek()) {
		lexer.readChar()
	}

	return lexer.input[start : lexer.pos+1]
}

func (lexer *Lexer) token(tokenType tokens.TokenType) tokens.Token {
	return tokens.Token{
		Type:    tokenType,
		Literal: "",
		Pos:     tokens.Pos{Line: lexer.line, Col: lexer.col},
	}
}

func (lexer *Lexer) tokenLiteral(tokenType tokens.TokenType, literal string) tokens.Token {
	return tokens.Token{
		Type:    tokenType,
		Literal: literal,
		Pos:     tokens.Pos{Line: lexer.line, Col: lexer.col},
	}
}
