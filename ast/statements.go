package ast

import (
	"bytes"
	"fmt"

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
func (s AssignmentStatement) String() string {
	return fmt.Sprintf("%v = %v", s.Name, s.Value.String())
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
func (s ShorthandAssignmentStatement) String() string {
	return fmt.Sprintf("%v %v= %v", s.Name, s.Operator, s.Value.String())
}
func (s ShorthandAssignmentStatement) statementNode() {}

// -------------------------------------------
// ------------- IF STATEMENT ----------------
// -------------------------------------------
type IfStatement struct {
	IfCondition      Expression
	IfBlock          BlockStatement
	ElseIfConditions []Expression
	ElseIfBlocks     []BlockStatement
	ElseBlock        BlockStatement
	Token            tokens.Token
}

//func (e IfStatement) StartPos() tokens.Pos { return e.startPos }
func (e IfStatement) statementNode() {}
func (e IfStatement) String() string {
	var out bytes.Buffer

	out.WriteString(fmt.Sprintf("if %v then\n%v", e.IfCondition.String(), e.IfBlock.String()))

	for i, elseIf := range e.ElseIfConditions {
		out.WriteString(fmt.Sprintf("elseif %v then\n%v", elseIf, e.ElseIfBlocks[i]))
	}

	if len(e.ElseBlock.Statements) > 0 {
		out.WriteString(fmt.Sprintf("else \n%v", e.ElseBlock.String()))
	}

	out.WriteString("end")
	return out.String()
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
func (s ReturnStatement) String() string {
	if s.Value == nil {
		return fmt.Sprintf("return")
	} else {
		return fmt.Sprintf("return %v", s.Value.String())
	}
}

// -------------------------------------------
// -------------- FOR STATEMENT --------------
// -------------------------------------------
type ForStatement struct {
	ItemName string
	List     Expression
	Body     BlockStatement
	Token    tokens.Token
}

//func (s ForStatement) StartPos() tokens.Pos { return s.startPos }
func (s ForStatement) String() string {
	return fmt.Sprintf("for %v in %v do\n%vend", s.ItemName, s.List.String(), s.Body.String())
}
func (s ForStatement) statementNode() {}

// -------------------------------------------
// ------------ BLOCK STATEMENT --------------
// -------------------------------------------
type BlockStatement struct {
	Statements []Statement
	Token      tokens.Token
}

//func (e BlockStatement) StartPos() tokens.Pos { return e.startPos }
func (e BlockStatement) statementNode() {}
func (e BlockStatement) String() string {
	var out bytes.Buffer

	for _, stmt := range e.Statements {
		out.WriteString("  " + stmt.String() + "\n")
	}

	return out.String()
}
