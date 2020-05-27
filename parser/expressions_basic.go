package parser

import (
	"strconv"

	"../ast"
	"../tokens"
)

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
