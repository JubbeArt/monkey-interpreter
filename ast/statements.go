package ast

import (
	"fmt"
	"strings"

	"../tokens"
)

// -------------------------------------------
// --------- ASSIGNMENT STATEMENT ------------
// -------------------------------------------
type AssignmentStatement struct {
	Name  string
	Value Expression
	Token tokens.Token
}

//func (s AssignmentStatement) StartPos() tokens.Pos { return s.startPos }
func (s AssignmentStatement) String(indent int) string {
	return fmt.Sprintf("%v = %v", s.Name, s.Value.String(indent))
}
func (s AssignmentStatement) statementNode() {}

// -------------------------------------------
// ----- SHORTHAND ASSIGNMENT STATEMENT ------
// -------------------------------------------
type ShorthandAssignmentStatement struct {
	Name     string
	Value    Expression
	Operator tokens.TokenType
	Token    tokens.Token
}

//func (s AssignmentStatement) StartPos() tokens.Pos { return s.startPos }
func (s ShorthandAssignmentStatement) String(indent int) string {
	return fmt.Sprintf("%v %v %v", s.Name, s.Operator, s.Value.String(indent))
}
func (s ShorthandAssignmentStatement) statementNode() {}

// -------------------------------------------
// ------------- IF STATEMENT ----------------
// -------------------------------------------
type IfStatement struct {
	Conditions   []Expression
	Consequences []BlockStatement
	Token        tokens.Token
}

//func (e IfStatement) StartPos() tokens.Pos { return e.startPos }
func (e IfStatement) statementNode() {}
func (e IfStatement) String(indent int) string {
	var stmts []string

	for i, cond := range e.Conditions {
		stmts = append(stmts, fmt.Sprintf("if %v then\n%v", cond.String(indent), e.Consequences[i].String(indent+1)))
	}

	return strings.Join(stmts, "") + strings.Repeat(INDENT, indent) + "end"
}

// -------------------------------------------
// ------------ RETURN STATEMENT -------------
// -------------------------------------------
type ReturnStatement struct {
	Value Expression
	Token tokens.Token
}

//func (s ReturnStatement) StartPos() tokens.Pos { return s.startPos }
func (s ReturnStatement) statementNode() {}
func (s ReturnStatement) String(indent int) string {
	if s.Value == nil {
		return fmt.Sprintf("return")
	} else {
		return fmt.Sprintf("return %v", s.Value.String(indent))
	}
}

// -------------------------------------------
// ------------- LOOP STATEMENT --------------
// -------------------------------------------
type LoopStatement struct {
	Body  BlockStatement
	Token tokens.Token
}

//func (s LoopStatement) StartPos() tokens.Pos { return s.startPos }
func (s LoopStatement) String(indent int) string {
	return fmt.Sprintf("loop\n%v%vend", s.Body.String(indent+1), strings.Repeat(INDENT, indent))
}
func (s LoopStatement) statementNode() {}

// -------------------------------------------
// ------------ BLOCK STATEMENT --------------
// -------------------------------------------
type BlockStatement struct {
	Statements []Statement
	Token      tokens.Token
}

//func (e BlockStatement) StartPos() tokens.Pos { return e.startPos }
func (e BlockStatement) statementNode() {}
func (e BlockStatement) String(indent int) string {
	var stmts []string

	in := strings.Repeat(INDENT, indent)
	for _, stmt := range e.Statements {
		stmts = append(stmts, in+stmt.String(indent)+"\n")
	}

	return strings.Join(stmts, "")
}
