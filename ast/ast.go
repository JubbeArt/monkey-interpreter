package ast

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"../tokens"
)

type Node interface {
	String() string
	//Node()
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// -------------------------------------------
// ---------------- PROGRAM ------------------
// -------------------------------------------
type Program struct {
	Statements []Statement
}

func (program Program) String() string {
	var out bytes.Buffer

	for _, stmt := range program.Statements {
		out.WriteString(stmt.String())
	}

	return out.String()
}

// -------------------------------------------
// -------------- LET STATEMENT --------------
// -------------------------------------------
type LetStatement struct {
	Variable string
	Value    Expression
}

func (s LetStatement) String() string {
	if s.Value == nil {
		return fmt.Sprintf("let %v;", s.Variable)
	} else {
		return fmt.Sprintf("let %v : %v;", s.Variable, s.Value.String())

	}
}
func (s LetStatement) statementNode() {}

// -------------------------------------------
// ------------ RETURN STATEMENT -------------
// -------------------------------------------
type ReturnStatement struct {
	Value Expression
}

func (s ReturnStatement) String() string {
	if s.Value == nil {
		return fmt.Sprintf("return;")
	} else {
		return fmt.Sprintf("return %v;", s.Value.String())
	}
}
func (s ReturnStatement) statementNode() {}

// -------------------------------------------
// --------- EXPRESSION STATEMENT ------------
// -------------------------------------------
type ExpressionStatement struct {
	Expression Expression
}

func (s ExpressionStatement) statementNode() {}

func (s ExpressionStatement) String() string {
	return s.Expression.String() + ";"
}

// -------------------------------------------
// --------- IDENTIFIER EXPRESSION -----------
// -------------------------------------------
type IdentifierExpression struct {
	Name string
}

func (e IdentifierExpression) expressionNode() {}
func (e IdentifierExpression) String() string  { return e.Name }

// -------------------------------------------
// ---------- INTEGER EXPRESSION -------------
// -------------------------------------------
type IntegerExpression struct {
	Value int64
}

func (e IntegerExpression) expressionNode() {}
func (e IntegerExpression) String() string {
	return strconv.FormatInt(e.Value, 10)
}

// -------------------------------------------
// ---------- BOOLEAN EXPRESSION -------------
// -------------------------------------------
type BooleanExpression struct {
	Value bool
}

func (e BooleanExpression) expressionNode() {}
func (e BooleanExpression) String() string {
	return strconv.FormatBool(e.Value)
}

// -------------------------------------------
// ------------- IF EXPRESSION ---------------
// -------------------------------------------
type IfExpression struct {
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (e IfExpression) expressionNode() {}
func (e IfExpression) String() string {
	if e.Alternative == nil {
		return fmt.Sprintf("if (%v) {\n %v \n}", e.Condition.String(), e.Consequence.String())
	} else {
		return fmt.Sprintf("if (%v) {\n %v \n} else {\n %v \n}", e.Condition.String(), e.Consequence.String(), e.Alternative.String())
	}
}

// -------------------------------------------
// ---------- FUNCTION EXPRESSION ------------
// -------------------------------------------
type FunctionExpression struct {
	Parameters []string
	Body       *BlockStatement
}

func (e FunctionExpression) expressionNode() {}
func (e FunctionExpression) String() string {
	var out bytes.Buffer
	out.WriteString("fn (")
	out.WriteString(strings.Join(e.Parameters, ", "))
	out.WriteString(") {\n  " + e.Body.String() + "\n}")
	return out.String()
}

// -------------------------------------------
// ------------ CALL EXPRESSION --------------
// -------------------------------------------
type CallExpression struct {
	Function  Expression
	Arguments []Expression
}

func (e CallExpression) expressionNode() {}
func (e CallExpression) String() string {
	args := []string{}

	for _, arg := range e.Arguments {
		args = append(args, arg.String())
	}

	return fmt.Sprintf("%v(%v)", e.Function.String(), strings.Join(args, ", "))
}

// -------------------------------------------
// ------------ BLOCK STATEMENT --------------
// -------------------------------------------
type BlockStatement struct {
	Statements []Statement
}

func (e BlockStatement) statementNode() {}
func (e BlockStatement) String() string {
	var out bytes.Buffer

	for _, stmt := range e.Statements {
		out.WriteString(stmt.String())
	}

	return out.String()
}

// -------------------------------------------
// ----------- PREFIX EXPRESSION -------------
// -------------------------------------------
type PrefixExpression struct {
	Operator  tokens.TokenType
	RightSide Expression
}

func (e PrefixExpression) expressionNode() {}
func (e PrefixExpression) String() string {
	if e.Operator == tokens.MINUS {
		return fmt.Sprintf("-" + e.RightSide.String())
	} else if e.Operator == tokens.NEGATION {
		return fmt.Sprintf("not " + e.RightSide.String())
	} else {
		return fmt.Sprintf("unknown prefix " + e.RightSide.String())
	}
}

// -------------------------------------------
// ------------ INFIX EXPRESSION -------------
// -------------------------------------------
type InfixExpression struct {
	Operator  tokens.TokenType
	LeftSide  Expression
	RightSide Expression
}

func (e InfixExpression) expressionNode() {}
func (e InfixExpression) String() string {
	return fmt.Sprintf("(%v %v %v)", e.LeftSide.String(), e.Operator, e.RightSide.String())
}
