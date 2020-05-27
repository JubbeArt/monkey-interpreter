package parser

import (
	"../ast"
	"../tokens"
)

func (pars *Parser) functionExpression() ast.Expression {
	expression := ast.FunctionExpression{}

	if !pars.nextTokenIf(tokens.L_PAREN) {
		pars.addError("expected \"(\" after function parameters")
		return nil
	}

	expression.Parameters = pars.functionParameters()
	pars.nextToken() // )
	expression.Body = pars.statements()

	if pars.currentToken.Type != tokens.END {
		pars.addError("expected \"end\" at the end of function")
	}

	return expression
}

func (pars *Parser) functionParameters() []string {
	var parameters []string

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
	return ast.CallExpression{
		Function:  function,
		Token:     pars.currentToken,
		Arguments: pars.commaList(tokens.R_PAREN),
	}
}
