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
	case tokens.FOR:
		return pars.forStatement()
	case tokens.IF:
		return pars.ifStatement()
	}

	return pars.expressionStatement()
}

func (pars *Parser) expressionStatement() ast.ExpressionStatement {
	token := pars.currentToken
	expr := pars.parseExpression(LOWEST)

	return ast.ExpressionStatement{
		Expression: expr,
		Token:      token,
	}
}

func (pars *Parser) ifStatement() ast.Statement {
	expression := ast.IfStatement{}

	pars.nextToken() // step into expression
	expression.IfCondition = pars.parseExpression(LOWEST)

	if !pars.nextTokenIf(tokens.THEN) {
		pars.addError("expected \"then\" after if condition")
		return nil
	}

	pars.nextToken() // then
	expression.IfBlock = pars.blockStatements()

	for pars.currentToken.Type == tokens.ELSEIF {
		pars.nextToken()
		condition := pars.parseExpression(LOWEST)

		if !pars.nextTokenIf(tokens.THEN) {
			pars.addError("expected \"then\" after elseif condition")
			return nil
		}

		pars.nextToken() // then
		body := pars.blockStatements()
		expression.ElseIfConditions = append(expression.ElseIfConditions, condition)
		expression.ElseIfBlocks = append(expression.ElseIfBlocks, body)
	}

	if pars.currentToken.Type == tokens.ELSE {
		pars.nextToken() // consume else
		expression.ElseBlock = pars.blockStatements()
	}

	if pars.currentToken.Type != tokens.END {
		pars.addError("expected \"end\" after if statement")
	}

	return expression
}

func (pars *Parser) assignmentStatement() ast.AssignmentStatement {
	statement := ast.AssignmentStatement{
		Name: pars.currentToken.Literal,
	}

	pars.nextToken() // =
	pars.nextToken() // start expression
	statement.Value = pars.parseExpression(LOWEST)
	return statement
}

func (pars *Parser) shorthandAssignmentStatements() ast.ShorthandAssignmentStatement {
	return ast.ShorthandAssignmentStatement{}
}

func (pars *Parser) returnStatement() ast.ReturnStatement {
	pars.nextToken()

	return ast.ReturnStatement{
		Value: pars.parseExpression(LOWEST),
	}
}

func (pars *Parser) forStatement() ast.ForStatement {
	statement := ast.ForStatement{}

	if !pars.nextTokenIf(tokens.IDENT) {
		pars.addError("expected identifier after \"for\"")
		return statement
	}

	statement.ItemName = pars.currentToken.Literal

	if !pars.nextTokenIf(tokens.IN) {
		pars.addError("expected \"in\" after \"for\" %v", statement.ItemName)
		return statement
	}

	pars.nextToken() // step into expression
	statement.List = pars.parseExpression(LOWEST)

	if !pars.nextTokenIf(tokens.DO) {
		pars.addError("expected \"do\" before for loop body")
		return statement
	}

	pars.nextToken() // step into block
	statement.Body = pars.blockStatements()

	if pars.currentToken.Type != tokens.END {
		pars.addError("expected \"end\" at the end of for loop")
		return statement
	}

	return statement
}

func (pars *Parser) blockStatements() ast.BlockStatement {
	blockStatement := ast.BlockStatement{
		Statements: []ast.Statement{},
		Token:      pars.currentToken,
	}

	for pars.currentToken.Type != tokens.END &&
		pars.currentToken.Type != tokens.ELSE &&
		pars.currentToken.Type != tokens.ELSEIF &&
		pars.currentToken.Type != tokens.EOF {

		stmt := pars.parseStatement()

		if stmt != nil {
			blockStatement.Statements = append(blockStatement.Statements, stmt)
		}

		pars.nextToken()
	}

	// don't do nextToken() here! check in each instead to get right "end", "else" etc.
	return blockStatement
}
