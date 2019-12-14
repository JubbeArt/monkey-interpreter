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
	tokens.EQ:      EQUALS,
	tokens.NOT_EQ:  EQUALS,
	tokens.LESS:    LESSGREATER,
	tokens.GREATER: LESSGREATER,
	tokens.ADD:     SUM,
	tokens.SUB:     SUM,
	tokens.MUL:     PRODUCT,
	tokens.DIV:     PRODUCT,
	tokens.L_PAREN: CALL,
}

type prefixParseFunc func() ast.Expression
type infixParseFunc func(ast.Expression) ast.Expression

func (pars *Parser) installParseFuncs() {
	pars.prefixParseFuncs[tokens.IDENT] = pars.identifier
	pars.prefixParseFuncs[tokens.NUMBER] = pars.number
	pars.prefixParseFuncs[tokens.TRUE] = pars.boolean
	pars.prefixParseFuncs[tokens.FALSE] = pars.boolean
	pars.prefixParseFuncs[tokens.STRING] = pars.text
	pars.prefixParseFuncs[tokens.NOT] = pars.prefixExpression
	pars.prefixParseFuncs[tokens.SUB] = pars.prefixExpression
	pars.prefixParseFuncs[tokens.L_PAREN] = pars.groupedExpression
	pars.prefixParseFuncs[tokens.FUNC] = pars.functionExpression
	pars.prefixParseFuncs[tokens.L_BRACKET] = pars.listExpression
	//pars.prefixParseFuncs[tokens.L_BRACE] = pars.recordExpresion

	pars.infixParseFuncs[tokens.ADD] = pars.infixExpression
	pars.infixParseFuncs[tokens.SUB] = pars.infixExpression
	pars.infixParseFuncs[tokens.MUL] = pars.infixExpression
	pars.infixParseFuncs[tokens.DIV] = pars.infixExpression
	pars.infixParseFuncs[tokens.EQ] = pars.infixExpression
	pars.infixParseFuncs[tokens.GREATER] = pars.infixExpression
	pars.infixParseFuncs[tokens.LESS] = pars.infixExpression
	pars.infixParseFuncs[tokens.L_PAREN] = pars.callExpression
	//pars.infixParseFuncs[tokens.DOT] = pars.dotExpression
}

func (pars *Parser) parseExpression(precedence int) ast.Expression {
	prefix := pars.prefixParseFuncs[pars.currentToken.Type]

	if prefix == nil {
		pars.addError("unexpected word: %q", pars.currentToken.Type)
		return nil
	}

	leftExpression := prefix()

	// do rightside(s)
	for precedence < pars.peekPrecedence() {
		infix := pars.infixParseFuncs[pars.peekToken.Type]

		if infix == nil {
			return leftExpression
		}
		pars.nextToken()
		leftExpression = infix(leftExpression)
	}

	return leftExpression
}

func (pars *Parser) identifier() ast.Expression {
	return ast.IdentifierExpression{
		Name:  pars.currentToken.Literal,
		Token: pars.currentToken,
	}
}

func (pars *Parser) number() ast.Expression {
	value, err := strconv.ParseFloat(pars.currentToken.Literal, 64)

	if err != nil {
		pars.addError("could not parse %q as an number", pars.currentToken.Literal)
		return nil
	}

	return ast.NumberExpression{
		Value: value,
		Token: pars.currentToken,
	}
}

func (pars *Parser) boolean() ast.Expression {
	return ast.BooleanExpression{
		Value: pars.currentToken.Type == tokens.TRUE,
		Token: pars.currentToken,
	}
}

func (pars *Parser) text() ast.Expression {
	return ast.TextExpression{
		Value: pars.currentToken.Literal,
		Token: pars.currentToken,
	}
}

func (pars *Parser) functionExpression() ast.Expression {
	expression := ast.FunctionExpression{}

	if !pars.nextTokenIf(tokens.L_PAREN) {
		pars.addError("expected \"(\" after function parameters")
		return nil
	}

	expression.Parameters = pars.functionParameters()
	pars.nextToken() // )
	expression.Body = pars.blockStatements()

	if pars.currentToken.Type != tokens.END {
		pars.addError("expected \"end\" at the end of function")
	}

	return expression
}

func (pars *Parser) listExpression() ast.Expression {
	expression := ast.ListExpression{
		Token:  pars.currentToken,
		Values: pars.listArguments(),
	}

	return expression
}

func (pars *Parser) listArguments() []ast.Expression {
	args := []ast.Expression{}

	if pars.nextTokenIf(tokens.R_BRACKET) {
		return args
	}

	pars.nextToken()
	args = append(args, pars.parseExpression(LOWEST))

	for pars.peekToken.Type == tokens.COMMA {
		pars.nextToken()
		pars.nextToken()
		args = append(args, pars.parseExpression(LOWEST))
	}

	if !pars.nextTokenIf(tokens.R_BRACKET) {
		pars.addError("expected \"]\" at the end of list")
		return nil
	}

	return args
}

func (pars *Parser) functionParameters() []string {
	parameters := []string{}

	if pars.peekToken.Type == tokens.R_PAREN {
		pars.nextToken()
		return parameters
	}

	pars.nextToken()

	parameters = append(parameters, pars.currentToken.Literal)

	for pars.peekToken.Type == tokens.COMMA {
		pars.nextToken() // ,
		pars.nextToken() // parameter
		parameters = append(parameters, pars.currentToken.Literal)
	}

	if !pars.nextTokenIf(tokens.R_PAREN) {
		pars.addError("expected \")\" after function parameters")
		return nil
	}

	return parameters
}

func (pars *Parser) callExpression(function ast.Expression) ast.Expression {
	expr := ast.CallExpression{
		Function:  function,
		Token:     pars.currentToken,
		Arguments: pars.callArguments(),
	}

	return expr
}

func (pars *Parser) callArguments() []ast.Expression {
	args := []ast.Expression{}

	if pars.nextTokenIf(tokens.R_PAREN) {
		return args
	}

	pars.nextToken()
	args = append(args, pars.parseExpression(LOWEST))

	for pars.peekToken.Type == tokens.COMMA {
		pars.nextToken()
		pars.nextToken()
		args = append(args, pars.parseExpression(LOWEST))
	}

	if !pars.nextTokenIf(tokens.R_PAREN) {
		pars.addError("expected \")\" at the end of function call")
		return nil
	}

	return args
}

//func (pars *Parser)

func (pars *Parser) groupedExpression() ast.Expression {
	pars.nextToken() // (
	expression := pars.parseExpression(LOWEST)

	if !pars.nextTokenIf(tokens.R_PAREN) {
		pars.addError("missing \")\" at the end of expression")
		return nil
	}

	return expression
}

func (pars *Parser) prefixExpression() ast.Expression {
	expression := ast.PrefixExpression{
		Operator: pars.currentToken.Type,
		Token:    pars.currentToken,
	}

	// consume prefix
	pars.nextToken()
	expression.RightSide = pars.parseExpression(PREFIX)
	return expression
}

func (pars *Parser) infixExpression(leftSide ast.Expression) ast.Expression {
	expression := ast.InfixExpression{
		Operator: pars.currentToken.Type,
		LeftSide: leftSide,
		Token:    pars.currentToken,
	}

	// get precedence for current operator and consume that operator
	precedence := pars.currentPrecedence()
	pars.nextToken()
	expression.RightSide = pars.parseExpression(precedence)

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
