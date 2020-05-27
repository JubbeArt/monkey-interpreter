package parser

import (
	"../ast"
	"../tokens"
)

func (pars *Parser) listExpression() ast.Expression {
	expression := ast.ListExpression{
		Token:  pars.currentToken,
		Values: pars.commaList(tokens.R_BRACKET),
	}

	return expression
}

func (pars *Parser) commaList(endToken tokens.TokenType) []ast.Expression {
	var args []ast.Expression

	if pars.nextTokenIf(endToken) {
		return args
	}

	pars.nextToken()
	args = append(args, pars.parseExpression(LOWEST))

	for pars.peekToken.Type == tokens.COMMA {
		pars.nextToken()
		pars.nextToken()
		args = append(args, pars.parseExpression(LOWEST))
	}

	if !pars.nextTokenIf(endToken) {
		pars.addError("expected %q at the end of list", endToken)
		return nil
	}

	return args
}
