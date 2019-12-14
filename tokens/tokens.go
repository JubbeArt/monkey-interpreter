package tokens

import "fmt"

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
	Pos     Pos
}

type Pos struct {
	Line int
	Col  int
}

const (
	// special
	ILLEGAL TokenType = iota
	EOF
	COMMENT

	IDENT
	NUMBER
	STRING

	ASSIGN

	ADD
	SUB
	MUL
	DIV

	ADD_ASSIGN
	SUB_ASSIGN
	MUL_ASSIGN
	DIV_ASSIGN

	EQ
	NOT_EQ
	LESS
	LESS_EQ
	GREATER
	GREATER_EQ

	COMMA
	DOT

	L_PAREN
	L_BRACE
	L_BRACKET

	R_PAREN
	R_BRACE
	R_BRACKET

	keywords_begin

	NOT
	AND
	OR

	FOR
	IN
	DO
	BREAK
	CONTINUE
	FUNC
	IF
	THEN
	ELSE
	ELSEIF
	RETURN
	END

	TRUE
	FALSE
	NIL

	keywords_end
)

var tokenNames = [...]string{
	ILLEGAL: "illegal",
	EOF:     "eof",
	COMMENT: "comment",

	IDENT:  "ident",
	NUMBER: "number",
	STRING: "string",

	ASSIGN: "=",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	DIV: "/",

	ADD_ASSIGN: "+=",
	SUB_ASSIGN: "-=",
	MUL_ASSIGN: "*=",
	DIV_ASSIGN: "/=",

	EQ:         "=",
	NOT_EQ:     "!=",
	LESS:       "<",
	LESS_EQ:    "<=",
	GREATER:    ">",
	GREATER_EQ: ">=",

	COMMA: ",",
	DOT:   ".",

	L_PAREN:   "(",
	L_BRACE:   "{",
	L_BRACKET: "[",

	R_PAREN:   ")",
	R_BRACE:   "}",
	R_BRACKET: "]",

	NOT: "not",
	AND: "and",
	OR:  "or",

	FOR:      "for",
	IN:       "in",
	DO:       "do",
	BREAK:    "break",
	CONTINUE: "continue",
	FUNC:     "func",
	IF:       "if",
	THEN:     "then",
	ELSE:     "else",
	ELSEIF:   "elseif",
	RETURN:   "return",
	END:      "end",

	TRUE:  "true",
	FALSE: "false",
	NIL:   "nil",
}

func (tt TokenType) String() string {
	return tokenNames[tt]
}

func (t Token) String() string {
	if t.Literal != "" {
		return fmt.Sprintf("TOKEN(%v - %v)", tokenNames[int(t.Type)], t.Literal)
	}
	return fmt.Sprintf("TOKEN(%v)", tokenNames[int(t.Type)])
}

var keywords map[string]TokenType

func init() {
	keywords = make(map[string]TokenType)

	for i := keywords_begin + 1; i < keywords_end; i++ {
		keywords[tokenNames[i]] = i
	}
}

func LookupIdentifier(identifier string) TokenType {
	if token, ok := keywords[identifier]; ok {
		return token
	}

	return IDENT
}
