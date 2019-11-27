package tokens

import "fmt"

type TokenType string

type Token struct {
	Type  TokenType
	Value string
}

func (token Token) String() string {
	if token.Value == "" {
		return fmt.Sprintf("TOKEN [%s]", token.Type)
	} else {
		return fmt.Sprintf("TOKEN [%s] (%s)", token.Type, token.Value)
	}
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"
	REAL       = "REAL"

	ASSIGN   = ":"
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"

	EQUALITY = "="
	LESS     = "<"
	GREATER  = ">"

	COMMA     = ","
	SEMICOLON = ";"

	LEFT_PAREN    = "("
	RIGHT_PAREN   = ")"
	LEFT_BRACE    = "{"
	RIGHT_BRACE   = "}"
	LEFT_BRACKET  = "["
	RIGHT_BRACKET = "]"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	ELSEIF   = "ELSEIF"
	RETURN   = "RETURN"

	TRUE     = "TRUE"
	FALSE    = "FALSE"
	NEGATION = "NOT"
)

var (
	keyWords = map[string]TokenType{
		"let":    LET,
		"fn":     FUNCTION,
		"not":    NEGATION,
		"if":     IF,
		"else":   ELSE,
		"elseif": ELSEIF,
		"true":   TRUE,
		"false":  FALSE,
		"return": RETURN,
	}

	SimpleTokens = map[byte]TokenType{
		ASSIGN[0]:   ASSIGN,
		PLUS[0]:     PLUS,
		MINUS[0]:    MINUS,
		MULTIPLY[0]: MULTIPLY,
		DIVIDE[0]:   DIVIDE,

		EQUALITY[0]: EQUALITY,
		LESS[0]:     LESS,
		GREATER[0]:  GREATER,

		COMMA[0]:     COMMA,
		SEMICOLON[0]: SEMICOLON,

		LEFT_PAREN[0]:    LEFT_PAREN,
		RIGHT_PAREN[0]:   RIGHT_PAREN,
		LEFT_BRACE[0]:    LEFT_BRACE,
		RIGHT_BRACE[0]:   RIGHT_BRACE,
		LEFT_BRACKET[0]:  LEFT_BRACKET,
		RIGHT_BRACKET[0]: RIGHT_BRACKET,
	}
)

func LookupIdentifier(identifier string) TokenType {
	if token, ok := keyWords[identifier]; ok {
		return token
	}

	return IDENTIFIER
}
