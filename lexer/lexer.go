package lexer

import (
	"unicode"
	"unicode/utf8"

	"../tokens"
)

const EOF = 0

type Lexer struct {
	input string

	line int
	col  int

	current    rune
	currentPos int
	lastLen    int
}

func New(input string) *Lexer {
	lexer := Lexer{input: input, currentPos: 0}
	return &lexer
}

func (lexer *Lexer) NextToken() tokens.Token {
	lexer.readChar()

	for unicode.IsSpace(lexer.current) {
		lexer.readChar()
	}

	if lexer.current == '#' {
		lexer.readLine()
		return lexer.NextToken()
	}

	switch lexer.current {
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
		str := lexer.getString()
		return lexer.tokenLiteral(tokens.STRING, str)
	case EOF:
		return lexer.token(tokens.EOF)
	}

	// check for numbers
	if unicode.IsDigit(lexer.current) {
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
	if isLetter(lexer.current) {
		identifier := lexer.getIdentifier()
		token := tokens.LookupIdentifier(identifier)

		if token == tokens.IDENT {
			return lexer.tokenLiteral(token, identifier)
		} else {
			return lexer.token(token)
		}
	}

	return lexer.tokenLiteral(tokens.ILLEGAL, string(lexer.current))
}

func (lexer *Lexer) readChar() {
	lexer.currentPos += lexer.lastLen

	if lexer.currentPos >= len(lexer.input) {
		lexer.current = EOF
		return
	}

	current, currentLen := utf8.DecodeRuneInString(lexer.input[lexer.currentPos:])
	lexer.current = current
	lexer.lastLen = currentLen

	if lexer.current == '\n' {
		lexer.col = 0
		lexer.line += 1
	}
}

func (lexer *Lexer) peek() rune {
	if lexer.currentPos+lexer.lastLen > len(lexer.input) {
		return EOF
	}

	peek, _ := utf8.DecodeRuneInString(lexer.input[lexer.currentPos+lexer.lastLen:])
	return peek
}

func (lexer *Lexer) getString() string {
	lexer.readChar() // consume "
	start := lexer.currentPos

	for lexer.current != '"' {
		lexer.readChar()
	}

	return lexer.input[start:lexer.currentPos]
}

func (lexer *Lexer) getIdentifier() string {
	start := lexer.currentPos

	for isLetter(lexer.peek()) {
		lexer.readChar()
	}

	return lexer.input[start : lexer.currentPos+lexer.lastLen]
}

func (lexer *Lexer) getDigits() string {
	start := lexer.currentPos

	for isDigit(lexer.peek()) {
		lexer.readChar()
	}

	return lexer.input[start : lexer.currentPos+lexer.lastLen]
}

func (lexer *Lexer) readLine() {
	for peek := lexer.peek(); peek != '\n' && peek != EOF; peek = lexer.peek() {
		lexer.readChar()
	}
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

func isLetter(char rune) bool {
	return 'a' <= char && char <= 'z' || char == '_'
}

func isDigit(char rune) bool {
	return '0' <= char && char <= '9'
}
