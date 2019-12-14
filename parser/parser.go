package parser

import (
	"fmt"

	"../ast"
	"../lexer"
	"../tokens"
)

type Parser struct {
	lex *lexer.Lexer

	currentToken tokens.Token
	peekToken    tokens.Token

	prefixParseFuncs map[tokens.TokenType]prefixParseFunc
	infixParseFuncs  map[tokens.TokenType]infixParseFunc

	errors []string
}

func New(lex *lexer.Lexer) *Parser {
	pars := Parser{
		lex:              lex,
		errors:           []string{},
		prefixParseFuncs: map[tokens.TokenType]prefixParseFunc{},
		infixParseFuncs:  map[tokens.TokenType]infixParseFunc{},
	}

	pars.installParseFuncs()
	pars.nextToken()
	pars.nextToken()

	return &pars
}

func (pars *Parser) ParseProgram() ast.Program {
	program := ast.Program{
		Statements: []ast.Statement{},
	}

	for pars.currentToken.Type != tokens.EOF {
		statement := pars.parseStatement()

		if pars.HasErrors() {
			break
		}

		program.Statements = append(program.Statements, statement)
		pars.nextToken()
	}

	return program
}

func (pars *Parser) nextTokenIf(token tokens.TokenType) bool {
	if pars.peekToken.Type != token {
		return false
	}

	pars.nextToken()
	return true
}

func (pars *Parser) nextToken() {
	pars.currentToken = pars.peekToken
	pars.peekToken = pars.lex.NextToken()
}

func (pars *Parser) addError(err string, args ...interface{}) {
	pars.errors = append(pars.errors, fmt.Sprintf(err, args...))
}

func (pars *Parser) PrintErrors() {
	if pars.HasErrors() {
		fmt.Println("Error while parsing:")
		for _, err := range pars.errors {
			fmt.Println("  ", err)
		}
	}
}

func (pars *Parser) HasErrors() bool {
	return len(pars.errors) > 0
}
