package lexer

import (
	"testing"

	"../tokens"
)

func TestNextToken(t *testing.T) {
	input := `
		let five : 5;
		let ten : 10;
		let add : fn(x, y) {
			x + y;
			};
		let result : add(five, ten);
	`

	tests := []tokens.Token{
		{tokens.LET, ""},
		{tokens.IDENTIFIER, "five"},
		{tokens.ASSIGN, ""},
		{tokens.INT, "5"},
		{tokens.SEMICOLON, ""},
		{tokens.LET, ""},
		{tokens.IDENTIFIER, "ten"},
		{tokens.ASSIGN, ""},
		{tokens.INT, "10"},
		{tokens.SEMICOLON, ""},
		{tokens.LET, ""},
		{tokens.IDENTIFIER, "add"},
		{tokens.ASSIGN, ""},
		{tokens.FUNCTION, ""},
		{tokens.LEFT_PAREN, ""},
		{tokens.IDENTIFIER, "x"},
		{tokens.COMMA, ""},
		{tokens.IDENTIFIER, "y"},
		{tokens.RIGHT_PAREN, ""},
		{tokens.LEFT_BRACE, ""},
		{tokens.IDENTIFIER, "x"},
		{tokens.PLUS, ""},
		{tokens.IDENTIFIER, "y"},
		{tokens.SEMICOLON, ""},
		{tokens.RIGHT_BRACE, ""},
		{tokens.SEMICOLON, ""},
		{tokens.LET, ""},
		{tokens.IDENTIFIER, "result"},
		{tokens.ASSIGN, ""},
		{tokens.IDENTIFIER, "add"},
		{tokens.LEFT_PAREN, ""},
		{tokens.IDENTIFIER, "five"},
		{tokens.COMMA, ""},
		{tokens.IDENTIFIER, "ten"},
		{tokens.RIGHT_PAREN, ""},
		{tokens.SEMICOLON, ""},
		{tokens.EOF, ""},
	}

	lexer := New(input)

	for i, expected := range tests {
		actual := lexer.NextToken()

		if actual.Type != expected.Type {
			t.Fatalf("tests[%d] failed - wrong token type, expected %q, got %q", i, actual.Type, expected.Type)
		}
		if actual.Value != expected.Value {
			t.Fatalf("tests[%d] failed - wrong token literal, expected %q, got %q", i, actual.Value, expected.Value)
		}
	}
}
