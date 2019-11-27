package parser

import (
	"strconv"

	"../ast"
	"../tokens"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // =
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or not X
	CALL        // myFunction(X)
)

var precedences = map[tokens.TokenType]int{
	tokens.EQUALITY: EQUALS,
	//tokens.NOT_EQ: EQUALS,
	tokens.LESS:       LESSGREATER,
	tokens.GREATER:    LESSGREATER,
	tokens.PLUS:       SUM,
	tokens.MINUS:      SUM,
	tokens.MULTIPLY:   PRODUCT,
	tokens.DIVIDE:     PRODUCT,
	tokens.LEFT_PAREN: CALL,
}

type prefixParseFunc func() ast.Expression
type infixParseFunc func(ast.Expression) ast.Expression

func (pars *Parser) installParseFuncs() {
	pars.prefixParseFuncs[tokens.IDENTIFIER] = pars.parseIdentifier
	pars.prefixParseFuncs[tokens.INT] = pars.parseInteger
	pars.prefixParseFuncs[tokens.TRUE] = pars.parseBoolean
	pars.prefixParseFuncs[tokens.FALSE] = pars.parseBoolean
	pars.prefixParseFuncs[tokens.NEGATION] = pars.parsePrefixExpression
	pars.prefixParseFuncs[tokens.MINUS] = pars.parsePrefixExpression
	pars.prefixParseFuncs[tokens.LEFT_PAREN] = pars.parseGroupedExpression
	pars.prefixParseFuncs[tokens.IF] = pars.parseIfExpression
	pars.prefixParseFuncs[tokens.FUNCTION] = pars.parseFunctionExpression

	pars.infixParseFuncs[tokens.PLUS] = pars.parseInfixExpression
	pars.infixParseFuncs[tokens.MINUS] = pars.parseInfixExpression
	pars.infixParseFuncs[tokens.MULTIPLY] = pars.parseInfixExpression
	pars.infixParseFuncs[tokens.DIVIDE] = pars.parseInfixExpression
	pars.infixParseFuncs[tokens.EQUALITY] = pars.parseInfixExpression
	pars.infixParseFuncs[tokens.GREATER] = pars.parseInfixExpression
	pars.infixParseFuncs[tokens.LESS] = pars.parseInfixExpression
	pars.infixParseFuncs[tokens.LEFT_PAREN] = pars.parseCallExpression
}

func (pars *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	statement := &ast.ExpressionStatement{}
	statement.Expression = pars.parseExpression(LOWEST)

	// optional semicolons (for e.g. repl)
	if pars.peekToken.Type == tokens.SEMICOLON {
		pars.nextToken()
	}

	return statement
}

func (pars *Parser) parseExpression(precedence int) ast.Expression {
	prefix := pars.prefixParseFuncs[pars.currentToken.Type]

	if prefix == nil {
		pars.addError("could not find prefix function for type %q", pars.currentToken.Type)
		return nil
	}

	leftExpression := prefix()

	// do rightside(s)
	for pars.peekToken.Type != tokens.SEMICOLON && precedence < pars.peekPrecedence() {
		infix := pars.infixParseFuncs[pars.peekToken.Type]

		if infix == nil {
			return leftExpression
		}
		pars.nextToken()
		leftExpression = infix(leftExpression)
	}

	return leftExpression
}

func (pars *Parser) parseIdentifier() ast.Expression {
	return &ast.IdentifierExpression{Name: pars.currentToken.Value}
}

func (pars *Parser) parseInteger() ast.Expression {
	value, err := strconv.ParseInt(pars.currentToken.Value, 10, 64)

	if err != nil {
		pars.addError("could not parse %q as an integer", pars.currentToken.Value)
		return nil
	}

	return &ast.IntegerExpression{Value: value}
}

func (pars *Parser) parseBoolean() ast.Expression {
	return &ast.BooleanExpression{Value: pars.currentToken.Type == tokens.TRUE}
}

func (pars *Parser) parseIfExpression() ast.Expression {
	expression := &ast.IfExpression{}

	if !pars.nextTokenIf(tokens.LEFT_PAREN) {
		pars.addError("expected \"(\" after if statement")
		return nil
	}

	pars.nextToken()
	expression.Condition = pars.parseExpression(LOWEST)

	if !pars.nextTokenIf(tokens.RIGHT_PAREN) {
		pars.addError("expected \")\" after if condition")
		return nil
	}

	if !pars.nextTokenIf(tokens.LEFT_BRACE) {
		pars.addError("expected \"{\" after if condition")
		return nil
	}

	expression.Consequence = pars.parseBlockStatements()

	if pars.nextTokenIf(tokens.ELSE) {
		if !pars.nextTokenIf(tokens.LEFT_BRACE) {
			pars.addError("expected \"{\" after else")
			return nil
		}

		expression.Alternative = pars.parseBlockStatements()
	}

	return expression
}

func (pars *Parser) parseBlockStatements() *ast.BlockStatement {
	blockStatement := &ast.BlockStatement{Statements: []ast.Statement{}}

	pars.nextToken()

	for pars.currentToken.Type != tokens.RIGHT_BRACE && pars.currentToken.Type != tokens.EOF {
		stmt := pars.parseStatement()

		if stmt != nil {
			blockStatement.Statements = append(blockStatement.Statements, stmt)
		}

		pars.nextToken()
	}

	return blockStatement
}

func (pars *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{Operator: pars.currentToken.Type}
	// consume prefix
	pars.nextToken()
	expression.RightSide = pars.parseExpression(PREFIX)
	return expression
}

func (pars *Parser) parseInfixExpression(leftSide ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Operator: pars.currentToken.Type,
		LeftSide: leftSide,
	}

	// get precedence for current operator and consume that operator
	precedence := pars.currentPrecedence()
	pars.nextToken()
	expression.RightSide = pars.parseExpression(precedence)

	return expression
}

func (pars *Parser) parseFunctionExpression() ast.Expression {
	expression := &ast.FunctionExpression{}

	if !pars.nextTokenIf(tokens.LEFT_PAREN) {
		pars.addError("expected \"(\" after function parameters")
		return nil
	}

	expression.Parameters = pars.parseFunctionParameters()

	if !pars.nextTokenIf(tokens.LEFT_BRACE) {
		pars.addError("expected \"{\" after function parameters")
		return nil
	}

	expression.Body = pars.parseBlockStatements()

	return expression
}

func (pars *Parser) parseFunctionParameters() []string {
	parameters := []string{}

	if pars.peekToken.Type == tokens.RIGHT_PAREN {
		pars.nextToken()
		return parameters
	}

	pars.nextToken()

	parameters = append(parameters, pars.currentToken.Value)

	for pars.peekToken.Type == tokens.COMMA {
		pars.nextToken() // ,
		pars.nextToken() // parameter
		parameters = append(parameters, pars.currentToken.Value)
	}

	if !pars.nextTokenIf(tokens.RIGHT_PAREN) {
		pars.addError("expected \")\" after function parameters")
		return nil
	}

	return parameters
}

func (pars *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	expression := &ast.CallExpression{Function: function}
	expression.Arguments = pars.parseCallArguments()
	return expression
}

func (pars *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}

	if pars.peekToken.Type == tokens.RIGHT_PAREN {
		pars.nextToken()
		return args
	}

	pars.nextToken()
	args = append(args, pars.parseExpression(LOWEST))

	for pars.peekToken.Type == tokens.COMMA {
		pars.nextToken()
		pars.nextToken()
		args = append(args, pars.parseExpression(LOWEST))
	}

	if !pars.nextTokenIf(tokens.RIGHT_PAREN) {
		pars.addError("expected \")\" at the end of function call")
		return nil
	}

	return args
}

func (pars *Parser) parseGroupedExpression() ast.Expression {
	pars.nextToken() // (
	expression := pars.parseExpression(LOWEST)

	if !pars.nextTokenIf(tokens.RIGHT_PAREN) {
		return nil
	}

	return expression
}

func (pars *Parser) peekPrecedence() int {
	if p, ok := precedences[pars.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (pars *Parser) currentPrecedence() int {
	if p, ok := precedences[pars.currentToken.Type]; ok {
		return p
	}
	return LOWEST
}
