package parser

import (
	"../ast"
	"../tokens"
)

const (
	_ int = iota
	LOWEST
	OR      // or
	AND     // and
	EQUALS  // ==
	SUM     // +
	PRODUCT // *
	PREFIX  // -X or not X
	CALL    // myFunction(X)
)

var precedences = map[tokens.TokenType]int{
	tokens.OR:         OR,
	tokens.AND:        AND,
	tokens.EQ:         EQUALS,
	tokens.NOT_EQ:     EQUALS,
	tokens.LESS_EQ:    EQUALS,
	tokens.GREATER_EQ: EQUALS,
	tokens.LESS:       EQUALS,
	tokens.GREATER:    EQUALS,
	tokens.ADD:        SUM,
	tokens.SUB:        SUM,
	tokens.MUL:        PRODUCT,
	tokens.DIV:        PRODUCT,
	tokens.L_PAREN:    CALL,
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
	pars.infixParseFuncs[tokens.NOT_EQ] = pars.infixExpression
	pars.infixParseFuncs[tokens.LESS] = pars.infixExpression
	pars.infixParseFuncs[tokens.LESS_EQ] = pars.infixExpression
	pars.infixParseFuncs[tokens.GREATER] = pars.infixExpression
	pars.infixParseFuncs[tokens.GREATER_EQ] = pars.infixExpression
	pars.infixParseFuncs[tokens.AND] = pars.infixExpression
	pars.infixParseFuncs[tokens.OR] = pars.infixExpression
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
