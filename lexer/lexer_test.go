package lexer

import (
	"testing"

	"../tokens"
)

func TestNextToken(t *testing.T) {
	input := `
		hello_world 123 123.4 "" "yes" 
		= 
		+-*/
		+=-=*=/=
		==!=<<=>>=
		, .
		([{}])
		not and or loop break continue func if then elseif
		else return end true false nil
		# hello world + 5
		"åäö" "öäå"
		a.b
		c[d]
		`
	tests := []struct {
		expectedType    tokens.TokenType
		expectedLiteral string
	}{
		{tokens.IDENT, "hello_world"},
		{tokens.NUMBER, "123"},
		{tokens.NUMBER, "123.4"},
		{tokens.STRING, ""},
		{tokens.STRING, "yes"},
		{tokens.ASSIGN, ""},
		{tokens.ADD, ""},
		{tokens.SUB, ""},
		{tokens.MUL, ""},
		{tokens.DIV, ""},
		{tokens.ADD_ASSIGN, ""},
		{tokens.SUB_ASSIGN, ""},
		{tokens.MUL_ASSIGN, ""},
		{tokens.DIV_ASSIGN, ""},
		{tokens.EQ, ""},
		{tokens.NOT_EQ, ""},
		{tokens.LESS, ""},
		{tokens.LESS_EQ, ""},
		{tokens.GREATER, ""},
		{tokens.GREATER_EQ, ""},
		{tokens.COMMA, ""},
		{tokens.DOT, ""},
		{tokens.L_PAREN, ""},
		{tokens.L_BRACKET, ""},
		{tokens.L_BRACE, ""},
		{tokens.R_BRACE, ""},
		{tokens.R_BRACKET, ""},
		{tokens.R_PAREN, ""},
		{tokens.NOT, ""},
		{tokens.AND, ""},
		{tokens.OR, ""},
		{tokens.LOOP, ""},
		{tokens.BREAK, ""},
		{tokens.CONTINUE, ""},
		{tokens.FUNC, ""},
		{tokens.IF, ""},
		{tokens.THEN, ""},
		{tokens.ELSEIF, ""},
		{tokens.ELSE, ""},
		{tokens.RETURN, ""},
		{tokens.END, ""},
		{tokens.TRUE, ""},
		{tokens.FALSE, ""},
		{tokens.NIL, ""},
		{tokens.STRING, "åäö"},
		{tokens.STRING, "öäå"},
		{tokens.IDENT, "a"},
		{tokens.DOT, ""},
		{tokens.IDENT, "b"},
		{tokens.IDENT, "c"},
		{tokens.L_BRACKET, ""},
		{tokens.IDENT, "d"},
		{tokens.R_BRACKET, ""},
		{tokens.EOF, ""},
	}
	l := New(input)

	for i, token := range tests {
		tok := l.NextToken()
		if tok.Type != token.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, token.expectedType, tok.Type)
		}
		if tok.Literal != token.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, token.expectedLiteral, tok.Literal)
		}
	}
}
