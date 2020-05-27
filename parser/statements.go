package parser

import (
	"../ast"
	"../tokens"
)

func (pars *Parser) parseStatement() ast.Statement {
	switch pars.currentToken.Type {
	case tokens.IDENT:
		switch pars.peekToken.Type {
		case tokens.ASSIGN:
			return pars.assignmentStatement()
		case tokens.ADD_ASSIGN, tokens.SUB_ASSIGN, tokens.MUL_ASSIGN, tokens.DIV_ASSIGN:
			return pars.shorthandAssignmentStatements()
		}
	case tokens.RETURN:
		return pars.returnStatement()
	case tokens.IF:
		return pars.ifStatement()
	case tokens.LOOP:
		return pars.loopStatement()
	}

	return nil
	//return pars.expressionStatement()
}

//func (pars *Parser) expressionStatement() ast.ExpressionStatement {
//	token := pars.currentToken
//	expr := pars.parseExpression(LOWEST)
//
//	return ast.ExpressionStatement{
//		Expression: expr,
//		Token:      token,
//	}
//}

func (pars *Parser) assignmentStatement() ast.AssignmentStatement {
	stmt := ast.AssignmentStatement{
		Name: pars.currentToken.Literal,
	}

	pars.nextToken() // =
	pars.nextToken() // start expression
	stmt.Value = pars.parseExpression(LOWEST)
	return stmt
}

func (pars *Parser) shorthandAssignmentStatements() ast.ShorthandAssignmentStatement {
	stmt := ast.ShorthandAssignmentStatement{
		Name: pars.currentToken.Literal,
	}

	pars.nextToken() // +=
	stmt.Operator = pars.currentToken.Type
	pars.nextToken() // start expression
	stmt.Value = pars.parseExpression(LOWEST)
	return stmt
}

func (pars *Parser) returnStatement() ast.ReturnStatement {
	pars.nextToken()

	if pars.currentToken.Type == tokens.END ||
		pars.currentToken.Type == tokens.ELSEIF ||
		pars.currentToken.Type == tokens.ELSE {

		return ast.ReturnStatement{}
	}

	return ast.ReturnStatement{
		Value: pars.parseExpression(LOWEST),
	}
}

func (pars *Parser) loopStatement() ast.LoopStatement {
	statement := ast.LoopStatement{}
	pars.nextToken() // loop -> stmts
	statement.Body = pars.statements()
	return statement
}
