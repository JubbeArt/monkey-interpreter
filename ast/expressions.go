package ast

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"../tokens"
)

// -------------------------------------------
// --------- EXPRESSION STATEMENT ------------
// -------------------------------------------
type ExpressionStatement struct {
	Expression Expression
	Token      tokens.Token
}

//func (s ExpressionStatement) Token() tokens.Pos { return s.startPos }
func (s ExpressionStatement) statementNode() {}
func (s ExpressionStatement) String() string {
	return s.Expression.String()
}

// -------------------------------------------
// --------- IDENTIFIER EXPRESSION -----------
// -------------------------------------------
type IdentifierExpression struct {
	Name  string
	Token tokens.Token
}

//func (e IdentifierExpression) StartPos() tokens.Pos { return e.startPos }
func (e IdentifierExpression) expressionNode() {}
func (e IdentifierExpression) String() string  { return e.Name }

// -------------------------------------------
// ----------- NUMBER EXPRESSION -------------
// -------------------------------------------
type NumberExpression struct {
	Value float64
	Token tokens.Token
}

//func (e NumberExpression) StartPos() tokens.Pos { return e.startPos }
func (e NumberExpression) expressionNode() {}
func (e NumberExpression) String() string {
	return strconv.FormatFloat(e.Value, 'f', -1, 64)
}

// -------------------------------------------
// ---------- BOOLEAN EXPRESSION -------------
// -------------------------------------------
type BooleanExpression struct {
	Value bool
	Token tokens.Token
}

//func (e BooleanExpression) StartPos() tokens.Pos { return e.startPos }
func (e BooleanExpression) expressionNode() {}
func (e BooleanExpression) String() string  { return strconv.FormatBool(e.Value) }

// -------------------------------------------
// ------------ TEXT EXPRESSION --------------
// -------------------------------------------
type TextExpression struct {
	Value string
	Token tokens.Token
}

//func (e TextExpression) StartPos() tokens.Pos { return e.startPos }
func (e TextExpression) expressionNode() {}
func (e TextExpression) String() string  { return fmt.Sprintf("%q", e.Value) }

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
func (e RecordExpression) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")

	for i, key := range e.Keys {
		out.WriteString(fmt.Sprintf("  %v = %v\n", key, e.Values[i].String()))
	}

	out.WriteString("}")
	return out.String()
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
func (e FunctionExpression) String() string {
	return fmt.Sprintf("func (%v)\n%vend", strings.Join(e.Parameters, ", "), e.Body.String())
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
func (e ListExpression) String() string {
	if len(e.Values) == 0 {
		return "[]"
	}

	var out bytes.Buffer
	out.WriteString("[\n")

	for _, value := range e.Values {
		out.WriteString(fmt.Sprintf("  %v,\n", value.String()))
	}

	out.WriteString("]")
	return out.String()
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
func (e CallExpression) String() string {
	args := []string{}

	for _, arg := range e.Arguments {
		args = append(args, arg.String())
	}

	return fmt.Sprintf("%v(%v)", e.Function.String(), strings.Join(args, ", "))
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
func (e PrefixExpression) String() string {
	if e.Operator == tokens.SUB {
		return fmt.Sprintf("-" + e.RightSide.String())
	} else if e.Operator == tokens.NOT {
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
	Token     tokens.Token
}

//func (e InfixExpression) StartPos() tokens.Pos { return e.startPos }
func (e InfixExpression) expressionNode() {}
func (e InfixExpression) String() string {
	return fmt.Sprintf("(%v %v %v)", e.LeftSide.String(), e.Operator, e.RightSide.String())
}
