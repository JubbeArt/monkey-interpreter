package ast

import (
	"fmt"
	"strconv"

	"../tokens"
)

// -------------------------------------------
// --------- IDENTIFIER EXPRESSION -----------
// -------------------------------------------
type IdentifierExpression struct {
	Name  string
	Token tokens.Token
}

//func (e IdentifierExpression) StartPos() tokens.Pos { return e.startPos }
func (e IdentifierExpression) expressionNode()   {}
func (e IdentifierExpression) String(int) string { return e.Name }

// -------------------------------------------
// ----------- NUMBER EXPRESSION -------------
// -------------------------------------------
type NumberExpression struct {
	Value float64
	Token tokens.Token
}

//func (e NumberExpression) StartPos() tokens.Pos { return e.startPos }
func (e NumberExpression) expressionNode() {}
func (e NumberExpression) String(int) string {
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
func (e BooleanExpression) expressionNode()   {}
func (e BooleanExpression) String(int) string { return strconv.FormatBool(e.Value) }

// -------------------------------------------
// ------------ TEXT EXPRESSION --------------
// -------------------------------------------
type TextExpression struct {
	Value string
	Token tokens.Token
}

//func (e TextExpression) StartPos() tokens.Pos { return e.startPos }
func (e TextExpression) expressionNode()   {}
func (e TextExpression) String(int) string { return fmt.Sprintf("%q", e.Value) }
