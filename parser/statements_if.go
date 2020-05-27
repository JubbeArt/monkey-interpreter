package parser

import (
	"../ast"
	"../tokens"
)

func (pars *Parser) ifStatement() ast.Statement {
	stmt := ast.IfStatement{}

	pars.nextToken() // if -> expression

	stmt.Conditions = append(stmt.Conditions, pars.parseExpression(LOWEST))

	if !pars.nextTokenIf(tokens.THEN) {
		pars.addError("expected \"then\" after if condition")
		return nil
	}

	pars.nextToken() // then -> stmts

	stmt.Consequences = append(stmt.Consequences, pars.ifBlockStatements())

	for pars.currentToken.Type == tokens.ELSEIF {
		pars.nextToken()
		stmt.Conditions = append(stmt.Conditions, pars.parseExpression(LOWEST))

		if !pars.nextTokenIf(tokens.THEN) {
			pars.addError("expected \"then\" after elseif condition")
			return nil
		}

		pars.nextToken() // then
		stmt.Consequences = append(stmt.Consequences, pars.ifBlockStatements())
	}

	if pars.currentToken.Type == tokens.ELSE {
		pars.nextToken() // consume else
		stmt.Conditions = append(stmt.Conditions, ast.BooleanExpression{Value: true})
		stmt.Consequences = append(stmt.Consequences, pars.statements())
	}

	return stmt
}

func (pars *Parser) statements() ast.BlockStatement {
	stmts := ast.BlockStatement{
		Statements: []ast.Statement{},
		Token:      pars.currentToken,
	}

	for pars.currentToken.Type != tokens.END && pars.currentToken.Type != tokens.EOF {
		stmts.Statements = append(stmts.Statements, pars.parseStatement())
		pars.nextToken()
	}

	if pars.currentToken.Type != tokens.END {
		pars.addError("expected \"end\" at the end of ...")
		return stmts
	}

	return stmts
}

func (pars *Parser) ifBlockStatements() ast.BlockStatement {
	stmts := ast.BlockStatement{
		Statements: []ast.Statement{},
		Token:      pars.currentToken,
	}

	for pars.currentToken.Type != tokens.ELSEIF &&
		pars.currentToken.Type != tokens.ELSE &&
		pars.currentToken.Type != tokens.END &&
		pars.currentToken.Type != tokens.EOF {
		stmts.Statements = append(stmts.Statements, pars.parseStatement())
		pars.nextToken()
	}

	return stmts
}
