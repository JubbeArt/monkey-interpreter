package ast

import (
	"fmt"
	"strings"

	"../tokens"
)

// -------------------------------------------
// --------- EXPRESSION STATEMENT ------------
// -------------------------------------------
//type ExpressionStatement struct {
//	Expression Expression
//	Token      tokens.Token
//}
//
////func (s ExpressionStatement) Token() tokens.Pos { return s.startPos }
//func (s ExpressionStatement) statementNode() {}
//func (s ExpressionStatement) String() string {
//	return s.Expression.String()
//}

// -------------------------------------------
// ----------- RECORD EXPRESSION -------------
// -------------------------------------------
type RecordExpression struct {
	Keys   []string
	Values []Expression
	Token  tokens.Token
}

//func (e RecordExpression) StartPos() tokens.Pos { return e.startPos }
func (e RecordExpression) expressionNode() {}
func (e RecordExpression) String(indent int) string {
	var pairs []string
	in := strings.Repeat(INDENT, indent)

	for i, key := range e.Keys {
		pairs = append(pairs, fmt.Sprintf("%v%v = %v", in, key, e.Values[i].String(indent+1)))
	}

	return fmt.Sprintf("{\n%v\n%v}", strings.Join(pairs, ",\n"), in)
}

// -------------------------------------------
// ---------- FUNCTION EXPRESSION ------------
// -------------------------------------------
type FunctionExpression struct {
	Parameters []string
	Body       BlockStatement
	Token      tokens.Token
}

//func (e FunctionExpression) StartPos() tokens.Pos { return e.startPos }
func (e FunctionExpression) expressionNode() {}
func (e FunctionExpression) String(indent int) string {
	var stmts []string
	in := strings.Repeat(INDENT, indent+1)

	for _, stmt := range e.Body.Statements {
		stmts = append(stmts, in+stmt.String(indent+1))
	}

	return fmt.Sprintf("func (%v)\n%v\n%vend",
		strings.Join(e.Parameters, ", "),
		strings.Join(stmts, "\n"),
		strings.Repeat(INDENT, indent))
}

// -------------------------------------------
// ------------ LIST EXPRESSION -------------
// -------------------------------------------
type ListExpression struct {
	Values []Expression
	Token  tokens.Token
}

//func (e ListExpression) StartPos() tokens.Pos { return e.startPos }
func (e ListExpression) expressionNode() {}
func (e ListExpression) String(indent int) string {
	if len(e.Values) == 0 {
		return "[]"
	}

	var values []string
	in := strings.Repeat(INDENT, indent+1)

	for _, value := range e.Values {
		values = append(values, in+value.String(indent+1))
	}

	return fmt.Sprintf("[\n%v\n%v]", strings.Join(values, ",\n"), strings.Repeat(INDENT, indent))
}

// -------------------------------------------
// ------------ CALL EXPRESSION --------------
// -------------------------------------------
type CallExpression struct {
	Function  Expression
	Arguments []Expression
	Token     tokens.Token
}

//func (e CallExpression) StartPos() tokens.Pos { return e.startPos }
func (e CallExpression) expressionNode() {}
func (e CallExpression) statementNode()  {}
func (e CallExpression) String(indent int) string {
	var args []string

	for _, arg := range e.Arguments {
		args = append(args, arg.String(indent))
	}

	return fmt.Sprintf("%v(%v)", e.Function.String(indent), strings.Join(args, ", "))
}

// -------------------------------------------
// ----------- PREFIX EXPRESSION -------------
// -------------------------------------------
type PrefixExpression struct {
	Operator  tokens.TokenType
	RightSide Expression
	Token     tokens.Token
}

//func (e PrefixExpression) StartPos() tokens.Pos { return e.startPos }
func (e PrefixExpression) expressionNode() {}
func (e PrefixExpression) String(indent int) string {
	if e.Operator == tokens.SUB {
		return "-" + e.RightSide.String(indent)
	}

	return fmt.Sprintf("%v %v", e.Operator, e.RightSide.String(indent))
}

// -------------------------------------------
// ------------ INFIX EXPRESSION -------------
// -------------------------------------------
type InfixExpression struct {
	Operator  tokens.TokenType
	LeftSide  Expression
	RightSide Expression
	Token     tokens.Token
}

//func (e InfixExpression) StartPos() tokens.Pos { return e.startPos }
func (e InfixExpression) expressionNode() {}
func (e InfixExpression) String(indent int) string {
	return fmt.Sprintf("(%v %v %v)", e.LeftSide.String(indent), e.Operator, e.RightSide.String(indent))
}
