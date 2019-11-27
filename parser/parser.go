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

func (pars *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	for pars.currentToken.Type != tokens.EOF {
		statement := pars.parseStatement()

		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}

		pars.nextToken()
	}

	return program
}

func (pars *Parser) parseStatement() ast.Statement {
	switch pars.currentToken.Type {
	case tokens.LET:
		return pars.parseLetStatement()
	case tokens.RETURN:
		return pars.parseReturnStatement()
	default:
		//fmt.Println("doing exprStmt")
		return pars.parseExpressionStatement()
	}
}

func (pars *Parser) parseLetStatement() *ast.LetStatement {
	if !pars.nextTokenExpect(tokens.IDENTIFIER) {
		pars.addError("expected identifier after \"let\"")
		return nil
	}

	statement := &ast.LetStatement{
		Variable: pars.currentToken.Value,
	}

	if !pars.nextTokenExpect(tokens.ASSIGN) {
		pars.addError("expected an assignment operator (:) after \"let %v\"", statement.Variable)
		return nil
	}

	for pars.currentToken.Type != tokens.SEMICOLON {
		pars.nextToken()
	}

	return statement
}

func (pars *Parser) parseReturnStatement() *ast.ReturnStatement {
	statement := &ast.ReturnStatement{}

	for pars.currentToken.Type != tokens.SEMICOLON {
		pars.nextToken()
	}

	return statement
}

func (pars *Parser) nextTokenExpect(token tokens.TokenType) bool {
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

func (pars *Parser) Errors() []string {
	return pars.errors
}

func (pars *Parser) HasErrors() bool {
	return len(pars.errors) > 0
}
