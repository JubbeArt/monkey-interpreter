package ast

import (
	"bytes"
)

type Node interface {
	String() string
	//StartPos() tokens.Pos
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

	for i, stmt := range program.Statements {
		if i == len(program.Statements)-1 {
			out.WriteString(stmt.String() + "\n")
		} else {
			out.WriteString(stmt.String() + "\n\n")
		}
	}

	return out.String()
}
